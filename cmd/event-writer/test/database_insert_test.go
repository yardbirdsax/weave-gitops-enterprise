package test

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/tj/assert"
	"github.com/weaveworks/wks/cmd/event-writer/converter"
	"github.com/weaveworks/wks/cmd/event-writer/database/models"
	"github.com/weaveworks/wks/cmd/event-writer/database/utils"
	"gorm.io/gorm"
)

var (
	clusterCount    = 25
	eventCount      = 100
	batchInsertSize = 1000
)

func TestDBInsertSuite(t *testing.T) {
	testDB, err := utils.Open("test.db")
	defer os.Remove("test.db")

	assert.NoError(t, err)
	err = utils.MigrateTables(testDB)
	assert.NoError(t, err)

	testSingleEventDBInsertion(t, testDB)
	testSingleEventDBSelect(t, testDB)
	testBatchEventDBInsertion(t, testDB)
	testDBQuerying(t, testDB)
}

func testSingleEventDBInsertion(t *testing.T, db *gorm.DB) {
	// Read the event from the sampleWarningEvent.json test file
	dbEvent, err := dbEventFromFile(t, "sampleWarningEvent.json")
	assert.NoError(t, err)

	// Insert the event in the table
	db.Create(&dbEvent)

	// Ensure the row count of the event table is 1
	var events []models.Event
	var count int64
	db.Model(&events).Count(&count)
	assert.Equal(t, int(count), 1)

	// Ensure that the event row contains the correct message
	var result models.Event
	db.First(&models.Event{}).First(&result)
	assert.Contains(t, dbEvent.Message, "synchronization of release 'grafana'")
	assert.Equal(t, dbEvent.Message, result.Message)
}

func testBatchEventDBInsertion(t *testing.T, db *gorm.DB) {
	var allEvents = []models.Event{}
	var allClusters = []models.Cluster{}

	for i := 0; i < clusterCount; i++ {
		clusterRow := models.Cluster{
			Name: RandomString(5),
		}
		allClusters = append(allClusters, clusterRow)
	}
	db.CreateInBatches(allClusters, clusterCount)

	// Read the event from the sampleWarningEvent.json test file
	dbEvent, err := dbEventFromFile(t, "sampleWarningEvent.json")
	assert.NoError(t, err)

	// Time the insertion of 250 batches of 10000 events
	start := time.Now()
	// Change the name and insert it to the array
	for _, clusterRow := range allClusters {
		for i := 0; i < eventCount; i++ {
			dbEvent.Name = RandomString(10)
			dbEvent.ClusterName = clusterRow.Name
			allEvents = append(allEvents, dbEvent)
		}
		// Insert the event in the table
		db.CreateInBatches(allEvents, batchInsertSize)
	}
	assert.Equal(t, len(allEvents), clusterCount*eventCount)
	elapsed1 := time.Since(start)

	// Delete all event rows
	start = time.Now()
	db.Exec("DELETE FROM events")
	elapsed2 := time.Since(start)

	// Time the insertion of 250 batches of 10000 events in a single insert
	start = time.Now()
	db.CreateInBatches(allEvents, batchInsertSize)
	assert.Equal(t, len(allEvents), clusterCount*eventCount)
	elapsed3 := time.Since(start)
	log.Printf("Insertion of %d batches of %d events took %s", clusterCount, eventCount, elapsed1)
	log.Printf("Batch deletion of %d events took %s", clusterCount*eventCount, elapsed2)
	log.Printf("Single batch insertion of %d events took %s", clusterCount*eventCount, elapsed3)
}

func testDBQuerying(t *testing.T, db *gorm.DB) {
	// Ensure the row count of the event table is 2500000
	var events []models.Event
	var count int64
	db.Model(&events).Count(&count)
	assert.Equal(t, int(count), clusterCount*eventCount)

	// Get all events for a single cluster
	// Get a random cluster name
	var result models.Event
	db.First(&result)
	assert.NotNil(t, result.ClusterName)

	// Get all events with that cluster name
	start := time.Now()
	db.Where("cluster_name == ?", result.ClusterName).Find(&events)
	elapsed := time.Since(start)
	assert.Equal(t, len(events), eventCount)
	log.Printf("Querying %d events by cluster name took %s", len(events), elapsed)
}

func testSingleEventDBSelect(t *testing.T, db *gorm.DB) {
	// Get the event with primary key (ID) 1
	var event models.Event
	db.First(&event, 1)
	assert.Contains(t, event.Message, "synchronization of release 'grafana'")

	// Return all event rows, should be 1
	result := db.Find(&event)
	assert.Equal(t, int(result.RowsAffected), 1)

	// Return all events from the wkp-grafana namespace
	db.Where("Namespace = ?", "wkp-grafana").Find(&event)
	assert.Equal(t, event.Namespace, "wkp-grafana")
	_, err := json.Marshal(event)
	assert.NoError(t, err)
}

// Read file at path and return bytes
func readTestFile(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Create a models.Event object from an Event exported to a JSON file
func dbEventFromFile(t *testing.T, path string) (models.Event, error) {
	data, err := readTestFile(path)
	assert.NoError(t, err)
	event, err := converter.DeserializeJSONToEvent(data)
	assert.NoError(t, err)

	dbEvent, err := converter.ConvertEvent(*event)
	assert.NoError(t, err)
	return dbEvent, nil
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RandomString(length int) string {
	return StringWithCharset(length, charset)
}