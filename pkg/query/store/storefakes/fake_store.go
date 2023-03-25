// Code generated by counterfeiter. DO NOT EDIT.
package storefakes

import (
	"context"
	"sync"

	"github.com/weaveworks/weave-gitops-enterprise/pkg/query/internal/models"
	"github.com/weaveworks/weave-gitops-enterprise/pkg/query/store"
)

type FakeStore struct {
	DeleteObjectsStub        func(context.Context, []models.Object) error
	deleteObjectsMutex       sync.RWMutex
	deleteObjectsArgsForCall []struct {
		arg1 context.Context
		arg2 []models.Object
	}
	deleteObjectsReturns struct {
		result1 error
	}
	deleteObjectsReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteRoleBindingsStub        func(context.Context, []models.RoleBinding) error
	deleteRoleBindingsMutex       sync.RWMutex
	deleteRoleBindingsArgsForCall []struct {
		arg1 context.Context
		arg2 []models.RoleBinding
	}
	deleteRoleBindingsReturns struct {
		result1 error
	}
	deleteRoleBindingsReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteRolesStub        func(context.Context, []models.Role) error
	deleteRolesMutex       sync.RWMutex
	deleteRolesArgsForCall []struct {
		arg1 context.Context
		arg2 []models.Role
	}
	deleteRolesReturns struct {
		result1 error
	}
	deleteRolesReturnsOnCall map[int]struct {
		result1 error
	}
	GetAccessRulesStub        func(context.Context) ([]models.AccessRule, error)
	getAccessRulesMutex       sync.RWMutex
	getAccessRulesArgsForCall []struct {
		arg1 context.Context
	}
	getAccessRulesReturns struct {
		result1 []models.AccessRule
		result2 error
	}
	getAccessRulesReturnsOnCall map[int]struct {
		result1 []models.AccessRule
		result2 error
	}
	GetObjectsStub        func(context.Context, store.Query, store.QueryOption) ([]models.Object, error)
	getObjectsMutex       sync.RWMutex
	getObjectsArgsForCall []struct {
		arg1 context.Context
		arg2 store.Query
		arg3 store.QueryOption
	}
	getObjectsReturns struct {
		result1 []models.Object
		result2 error
	}
	getObjectsReturnsOnCall map[int]struct {
		result1 []models.Object
		result2 error
	}
	StoreObjectsStub        func(context.Context, []models.Object) error
	storeObjectsMutex       sync.RWMutex
	storeObjectsArgsForCall []struct {
		arg1 context.Context
		arg2 []models.Object
	}
	storeObjectsReturns struct {
		result1 error
	}
	storeObjectsReturnsOnCall map[int]struct {
		result1 error
	}
	StoreRoleBindingsStub        func(context.Context, []models.RoleBinding) error
	storeRoleBindingsMutex       sync.RWMutex
	storeRoleBindingsArgsForCall []struct {
		arg1 context.Context
		arg2 []models.RoleBinding
	}
	storeRoleBindingsReturns struct {
		result1 error
	}
	storeRoleBindingsReturnsOnCall map[int]struct {
		result1 error
	}
	StoreRolesStub        func(context.Context, []models.Role) error
	storeRolesMutex       sync.RWMutex
	storeRolesArgsForCall []struct {
		arg1 context.Context
		arg2 []models.Role
	}
	storeRolesReturns struct {
		result1 error
	}
	storeRolesReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStore) DeleteObjects(arg1 context.Context, arg2 []models.Object) error {
	var arg2Copy []models.Object
	if arg2 != nil {
		arg2Copy = make([]models.Object, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.deleteObjectsMutex.Lock()
	ret, specificReturn := fake.deleteObjectsReturnsOnCall[len(fake.deleteObjectsArgsForCall)]
	fake.deleteObjectsArgsForCall = append(fake.deleteObjectsArgsForCall, struct {
		arg1 context.Context
		arg2 []models.Object
	}{arg1, arg2Copy})
	stub := fake.DeleteObjectsStub
	fakeReturns := fake.deleteObjectsReturns
	fake.recordInvocation("DeleteObjects", []interface{}{arg1, arg2Copy})
	fake.deleteObjectsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStore) DeleteObjectsCallCount() int {
	fake.deleteObjectsMutex.RLock()
	defer fake.deleteObjectsMutex.RUnlock()
	return len(fake.deleteObjectsArgsForCall)
}

func (fake *FakeStore) DeleteObjectsCalls(stub func(context.Context, []models.Object) error) {
	fake.deleteObjectsMutex.Lock()
	defer fake.deleteObjectsMutex.Unlock()
	fake.DeleteObjectsStub = stub
}

func (fake *FakeStore) DeleteObjectsArgsForCall(i int) (context.Context, []models.Object) {
	fake.deleteObjectsMutex.RLock()
	defer fake.deleteObjectsMutex.RUnlock()
	argsForCall := fake.deleteObjectsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStore) DeleteObjectsReturns(result1 error) {
	fake.deleteObjectsMutex.Lock()
	defer fake.deleteObjectsMutex.Unlock()
	fake.DeleteObjectsStub = nil
	fake.deleteObjectsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) DeleteObjectsReturnsOnCall(i int, result1 error) {
	fake.deleteObjectsMutex.Lock()
	defer fake.deleteObjectsMutex.Unlock()
	fake.DeleteObjectsStub = nil
	if fake.deleteObjectsReturnsOnCall == nil {
		fake.deleteObjectsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteObjectsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) DeleteRoleBindings(arg1 context.Context, arg2 []models.RoleBinding) error {
	var arg2Copy []models.RoleBinding
	if arg2 != nil {
		arg2Copy = make([]models.RoleBinding, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.deleteRoleBindingsMutex.Lock()
	ret, specificReturn := fake.deleteRoleBindingsReturnsOnCall[len(fake.deleteRoleBindingsArgsForCall)]
	fake.deleteRoleBindingsArgsForCall = append(fake.deleteRoleBindingsArgsForCall, struct {
		arg1 context.Context
		arg2 []models.RoleBinding
	}{arg1, arg2Copy})
	stub := fake.DeleteRoleBindingsStub
	fakeReturns := fake.deleteRoleBindingsReturns
	fake.recordInvocation("DeleteRoleBindings", []interface{}{arg1, arg2Copy})
	fake.deleteRoleBindingsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStore) DeleteRoleBindingsCallCount() int {
	fake.deleteRoleBindingsMutex.RLock()
	defer fake.deleteRoleBindingsMutex.RUnlock()
	return len(fake.deleteRoleBindingsArgsForCall)
}

func (fake *FakeStore) DeleteRoleBindingsCalls(stub func(context.Context, []models.RoleBinding) error) {
	fake.deleteRoleBindingsMutex.Lock()
	defer fake.deleteRoleBindingsMutex.Unlock()
	fake.DeleteRoleBindingsStub = stub
}

func (fake *FakeStore) DeleteRoleBindingsArgsForCall(i int) (context.Context, []models.RoleBinding) {
	fake.deleteRoleBindingsMutex.RLock()
	defer fake.deleteRoleBindingsMutex.RUnlock()
	argsForCall := fake.deleteRoleBindingsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStore) DeleteRoleBindingsReturns(result1 error) {
	fake.deleteRoleBindingsMutex.Lock()
	defer fake.deleteRoleBindingsMutex.Unlock()
	fake.DeleteRoleBindingsStub = nil
	fake.deleteRoleBindingsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) DeleteRoleBindingsReturnsOnCall(i int, result1 error) {
	fake.deleteRoleBindingsMutex.Lock()
	defer fake.deleteRoleBindingsMutex.Unlock()
	fake.DeleteRoleBindingsStub = nil
	if fake.deleteRoleBindingsReturnsOnCall == nil {
		fake.deleteRoleBindingsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteRoleBindingsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) DeleteRoles(arg1 context.Context, arg2 []models.Role) error {
	var arg2Copy []models.Role
	if arg2 != nil {
		arg2Copy = make([]models.Role, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.deleteRolesMutex.Lock()
	ret, specificReturn := fake.deleteRolesReturnsOnCall[len(fake.deleteRolesArgsForCall)]
	fake.deleteRolesArgsForCall = append(fake.deleteRolesArgsForCall, struct {
		arg1 context.Context
		arg2 []models.Role
	}{arg1, arg2Copy})
	stub := fake.DeleteRolesStub
	fakeReturns := fake.deleteRolesReturns
	fake.recordInvocation("DeleteRoles", []interface{}{arg1, arg2Copy})
	fake.deleteRolesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStore) DeleteRolesCallCount() int {
	fake.deleteRolesMutex.RLock()
	defer fake.deleteRolesMutex.RUnlock()
	return len(fake.deleteRolesArgsForCall)
}

func (fake *FakeStore) DeleteRolesCalls(stub func(context.Context, []models.Role) error) {
	fake.deleteRolesMutex.Lock()
	defer fake.deleteRolesMutex.Unlock()
	fake.DeleteRolesStub = stub
}

func (fake *FakeStore) DeleteRolesArgsForCall(i int) (context.Context, []models.Role) {
	fake.deleteRolesMutex.RLock()
	defer fake.deleteRolesMutex.RUnlock()
	argsForCall := fake.deleteRolesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStore) DeleteRolesReturns(result1 error) {
	fake.deleteRolesMutex.Lock()
	defer fake.deleteRolesMutex.Unlock()
	fake.DeleteRolesStub = nil
	fake.deleteRolesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) DeleteRolesReturnsOnCall(i int, result1 error) {
	fake.deleteRolesMutex.Lock()
	defer fake.deleteRolesMutex.Unlock()
	fake.DeleteRolesStub = nil
	if fake.deleteRolesReturnsOnCall == nil {
		fake.deleteRolesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteRolesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) GetAccessRules(arg1 context.Context) ([]models.AccessRule, error) {
	fake.getAccessRulesMutex.Lock()
	ret, specificReturn := fake.getAccessRulesReturnsOnCall[len(fake.getAccessRulesArgsForCall)]
	fake.getAccessRulesArgsForCall = append(fake.getAccessRulesArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.GetAccessRulesStub
	fakeReturns := fake.getAccessRulesReturns
	fake.recordInvocation("GetAccessRules", []interface{}{arg1})
	fake.getAccessRulesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStore) GetAccessRulesCallCount() int {
	fake.getAccessRulesMutex.RLock()
	defer fake.getAccessRulesMutex.RUnlock()
	return len(fake.getAccessRulesArgsForCall)
}

func (fake *FakeStore) GetAccessRulesCalls(stub func(context.Context) ([]models.AccessRule, error)) {
	fake.getAccessRulesMutex.Lock()
	defer fake.getAccessRulesMutex.Unlock()
	fake.GetAccessRulesStub = stub
}

func (fake *FakeStore) GetAccessRulesArgsForCall(i int) context.Context {
	fake.getAccessRulesMutex.RLock()
	defer fake.getAccessRulesMutex.RUnlock()
	argsForCall := fake.getAccessRulesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeStore) GetAccessRulesReturns(result1 []models.AccessRule, result2 error) {
	fake.getAccessRulesMutex.Lock()
	defer fake.getAccessRulesMutex.Unlock()
	fake.GetAccessRulesStub = nil
	fake.getAccessRulesReturns = struct {
		result1 []models.AccessRule
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) GetAccessRulesReturnsOnCall(i int, result1 []models.AccessRule, result2 error) {
	fake.getAccessRulesMutex.Lock()
	defer fake.getAccessRulesMutex.Unlock()
	fake.GetAccessRulesStub = nil
	if fake.getAccessRulesReturnsOnCall == nil {
		fake.getAccessRulesReturnsOnCall = make(map[int]struct {
			result1 []models.AccessRule
			result2 error
		})
	}
	fake.getAccessRulesReturnsOnCall[i] = struct {
		result1 []models.AccessRule
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) GetObjects(arg1 context.Context, arg2 store.Query, arg3 store.QueryOption) ([]models.Object, error) {
	fake.getObjectsMutex.Lock()
	ret, specificReturn := fake.getObjectsReturnsOnCall[len(fake.getObjectsArgsForCall)]
	fake.getObjectsArgsForCall = append(fake.getObjectsArgsForCall, struct {
		arg1 context.Context
		arg2 store.Query
		arg3 store.QueryOption
	}{arg1, arg2, arg3})
	stub := fake.GetObjectsStub
	fakeReturns := fake.getObjectsReturns
	fake.recordInvocation("GetObjects", []interface{}{arg1, arg2, arg3})
	fake.getObjectsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStore) GetObjectsCallCount() int {
	fake.getObjectsMutex.RLock()
	defer fake.getObjectsMutex.RUnlock()
	return len(fake.getObjectsArgsForCall)
}

func (fake *FakeStore) GetObjectsCalls(stub func(context.Context, store.Query, store.QueryOption) ([]models.Object, error)) {
	fake.getObjectsMutex.Lock()
	defer fake.getObjectsMutex.Unlock()
	fake.GetObjectsStub = stub
}

func (fake *FakeStore) GetObjectsArgsForCall(i int) (context.Context, store.Query, store.QueryOption) {
	fake.getObjectsMutex.RLock()
	defer fake.getObjectsMutex.RUnlock()
	argsForCall := fake.getObjectsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeStore) GetObjectsReturns(result1 []models.Object, result2 error) {
	fake.getObjectsMutex.Lock()
	defer fake.getObjectsMutex.Unlock()
	fake.GetObjectsStub = nil
	fake.getObjectsReturns = struct {
		result1 []models.Object
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) GetObjectsReturnsOnCall(i int, result1 []models.Object, result2 error) {
	fake.getObjectsMutex.Lock()
	defer fake.getObjectsMutex.Unlock()
	fake.GetObjectsStub = nil
	if fake.getObjectsReturnsOnCall == nil {
		fake.getObjectsReturnsOnCall = make(map[int]struct {
			result1 []models.Object
			result2 error
		})
	}
	fake.getObjectsReturnsOnCall[i] = struct {
		result1 []models.Object
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) StoreObjects(arg1 context.Context, arg2 []models.Object) error {
	var arg2Copy []models.Object
	if arg2 != nil {
		arg2Copy = make([]models.Object, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.storeObjectsMutex.Lock()
	ret, specificReturn := fake.storeObjectsReturnsOnCall[len(fake.storeObjectsArgsForCall)]
	fake.storeObjectsArgsForCall = append(fake.storeObjectsArgsForCall, struct {
		arg1 context.Context
		arg2 []models.Object
	}{arg1, arg2Copy})
	stub := fake.StoreObjectsStub
	fakeReturns := fake.storeObjectsReturns
	fake.recordInvocation("StoreObjects", []interface{}{arg1, arg2Copy})
	fake.storeObjectsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStore) StoreObjectsCallCount() int {
	fake.storeObjectsMutex.RLock()
	defer fake.storeObjectsMutex.RUnlock()
	return len(fake.storeObjectsArgsForCall)
}

func (fake *FakeStore) StoreObjectsCalls(stub func(context.Context, []models.Object) error) {
	fake.storeObjectsMutex.Lock()
	defer fake.storeObjectsMutex.Unlock()
	fake.StoreObjectsStub = stub
}

func (fake *FakeStore) StoreObjectsArgsForCall(i int) (context.Context, []models.Object) {
	fake.storeObjectsMutex.RLock()
	defer fake.storeObjectsMutex.RUnlock()
	argsForCall := fake.storeObjectsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStore) StoreObjectsReturns(result1 error) {
	fake.storeObjectsMutex.Lock()
	defer fake.storeObjectsMutex.Unlock()
	fake.StoreObjectsStub = nil
	fake.storeObjectsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) StoreObjectsReturnsOnCall(i int, result1 error) {
	fake.storeObjectsMutex.Lock()
	defer fake.storeObjectsMutex.Unlock()
	fake.StoreObjectsStub = nil
	if fake.storeObjectsReturnsOnCall == nil {
		fake.storeObjectsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.storeObjectsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) StoreRoleBindings(arg1 context.Context, arg2 []models.RoleBinding) error {
	var arg2Copy []models.RoleBinding
	if arg2 != nil {
		arg2Copy = make([]models.RoleBinding, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.storeRoleBindingsMutex.Lock()
	ret, specificReturn := fake.storeRoleBindingsReturnsOnCall[len(fake.storeRoleBindingsArgsForCall)]
	fake.storeRoleBindingsArgsForCall = append(fake.storeRoleBindingsArgsForCall, struct {
		arg1 context.Context
		arg2 []models.RoleBinding
	}{arg1, arg2Copy})
	stub := fake.StoreRoleBindingsStub
	fakeReturns := fake.storeRoleBindingsReturns
	fake.recordInvocation("StoreRoleBindings", []interface{}{arg1, arg2Copy})
	fake.storeRoleBindingsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStore) StoreRoleBindingsCallCount() int {
	fake.storeRoleBindingsMutex.RLock()
	defer fake.storeRoleBindingsMutex.RUnlock()
	return len(fake.storeRoleBindingsArgsForCall)
}

func (fake *FakeStore) StoreRoleBindingsCalls(stub func(context.Context, []models.RoleBinding) error) {
	fake.storeRoleBindingsMutex.Lock()
	defer fake.storeRoleBindingsMutex.Unlock()
	fake.StoreRoleBindingsStub = stub
}

func (fake *FakeStore) StoreRoleBindingsArgsForCall(i int) (context.Context, []models.RoleBinding) {
	fake.storeRoleBindingsMutex.RLock()
	defer fake.storeRoleBindingsMutex.RUnlock()
	argsForCall := fake.storeRoleBindingsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStore) StoreRoleBindingsReturns(result1 error) {
	fake.storeRoleBindingsMutex.Lock()
	defer fake.storeRoleBindingsMutex.Unlock()
	fake.StoreRoleBindingsStub = nil
	fake.storeRoleBindingsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) StoreRoleBindingsReturnsOnCall(i int, result1 error) {
	fake.storeRoleBindingsMutex.Lock()
	defer fake.storeRoleBindingsMutex.Unlock()
	fake.StoreRoleBindingsStub = nil
	if fake.storeRoleBindingsReturnsOnCall == nil {
		fake.storeRoleBindingsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.storeRoleBindingsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) StoreRoles(arg1 context.Context, arg2 []models.Role) error {
	var arg2Copy []models.Role
	if arg2 != nil {
		arg2Copy = make([]models.Role, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.storeRolesMutex.Lock()
	ret, specificReturn := fake.storeRolesReturnsOnCall[len(fake.storeRolesArgsForCall)]
	fake.storeRolesArgsForCall = append(fake.storeRolesArgsForCall, struct {
		arg1 context.Context
		arg2 []models.Role
	}{arg1, arg2Copy})
	stub := fake.StoreRolesStub
	fakeReturns := fake.storeRolesReturns
	fake.recordInvocation("StoreRoles", []interface{}{arg1, arg2Copy})
	fake.storeRolesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStore) StoreRolesCallCount() int {
	fake.storeRolesMutex.RLock()
	defer fake.storeRolesMutex.RUnlock()
	return len(fake.storeRolesArgsForCall)
}

func (fake *FakeStore) StoreRolesCalls(stub func(context.Context, []models.Role) error) {
	fake.storeRolesMutex.Lock()
	defer fake.storeRolesMutex.Unlock()
	fake.StoreRolesStub = stub
}

func (fake *FakeStore) StoreRolesArgsForCall(i int) (context.Context, []models.Role) {
	fake.storeRolesMutex.RLock()
	defer fake.storeRolesMutex.RUnlock()
	argsForCall := fake.storeRolesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStore) StoreRolesReturns(result1 error) {
	fake.storeRolesMutex.Lock()
	defer fake.storeRolesMutex.Unlock()
	fake.StoreRolesStub = nil
	fake.storeRolesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) StoreRolesReturnsOnCall(i int, result1 error) {
	fake.storeRolesMutex.Lock()
	defer fake.storeRolesMutex.Unlock()
	fake.StoreRolesStub = nil
	if fake.storeRolesReturnsOnCall == nil {
		fake.storeRolesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.storeRolesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deleteObjectsMutex.RLock()
	defer fake.deleteObjectsMutex.RUnlock()
	fake.deleteRoleBindingsMutex.RLock()
	defer fake.deleteRoleBindingsMutex.RUnlock()
	fake.deleteRolesMutex.RLock()
	defer fake.deleteRolesMutex.RUnlock()
	fake.getAccessRulesMutex.RLock()
	defer fake.getAccessRulesMutex.RUnlock()
	fake.getObjectsMutex.RLock()
	defer fake.getObjectsMutex.RUnlock()
	fake.storeObjectsMutex.RLock()
	defer fake.storeObjectsMutex.RUnlock()
	fake.storeRoleBindingsMutex.RLock()
	defer fake.storeRoleBindingsMutex.RUnlock()
	fake.storeRolesMutex.RLock()
	defer fake.storeRolesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStore) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ store.Store = new(FakeStore)
