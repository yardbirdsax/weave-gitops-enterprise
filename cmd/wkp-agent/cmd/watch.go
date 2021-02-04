package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/weaveworks/wks/cmd/wkp-agent/internal/common"
	clusterclient "github.com/weaveworks/wks/pkg/cluster/client"
	clusterpoller "github.com/weaveworks/wks/pkg/cluster/poller"
	clusterwatcher "github.com/weaveworks/wks/pkg/cluster/watcher"
	"github.com/weaveworks/wks/pkg/messaging/handlers"
	"k8s.io/client-go/informers"
)

type watchCmdParamSet struct {
	ClusterInfoPollingInterval time.Duration
}

var watchCmdParams watchCmdParamSet

var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Watches for Kubernetes events and publishes them to NATS",
	Run:   run,
}

func init() {
	rootCmd.AddCommand(watchCmd)

	watchCmd.PersistentFlags().DurationVar(&watchCmdParams.ClusterInfoPollingInterval, "cluster-info-polling-interval", 10*time.Second, "Polling interval for ClusterInfo")
}

func run(cmd *cobra.Command, args []string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	k8sClient, err := clusterclient.GetClient(KubeconfigFile)
	if err != nil {
		log.Fatalf("Failed to create Kubernetes client: %s.", err.Error())
	}
	// The defaultResync duration is used as a default value for any handlers
	// added via AddEventHandler that don't specify a default value. This value
	// controls how often to re-list the resource that we want to be informed of.
	// Re-listing the resource results in re-fetching current items from the API
	// server and adding them to the store, resulting in Added events for new items,
	// Deleted events for removed items and Updated events for all existing items.
	// For our use case, if we set this value to anything other than 0, it will
	// fire OnUpdate calls for existing items in the store even if nothing has changed.
	// So we want to delay this as much as possible by setting it to 0.
	factory := informers.NewSharedInformerFactory(k8sClient, 0)

	client := common.CreateClient(ctx, NatsURL, Subject)

	log.Info("Agent starting watchers.")

	// Watch for Events
	notifier := handlers.NewEventNotifier("wkp-agent", client)
	events := clusterwatcher.NewWatcher(factory.Core().V1().Events().Informer(), notifier.Notify)
	go events.Run("Events", 1, ctx.Done())

	// Poll for ClusterInfo
	clusterInfoSender := handlers.NewClusterInfoSender("wkp-agent", client)
	clusterInfo := clusterpoller.NewClusterInfoPoller(k8sClient, watchCmdParams.ClusterInfoPollingInterval, clusterInfoSender)
	go clusterInfo.Run(ctx.Done())

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	cancel()
}
