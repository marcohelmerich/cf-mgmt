// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	sync "sync"

	cfclient "github.com/cloudfoundry-community/go-cfclient"
	space "github.com/vmwarepivotallabs/cf-mgmt/space"
)

type FakeCFClient struct {
	CreateSpaceStub        func(cfclient.SpaceRequest) (cfclient.Space, error)
	createSpaceMutex       sync.RWMutex
	createSpaceArgsForCall []struct {
		arg1 cfclient.SpaceRequest
	}
	createSpaceReturns struct {
		result1 cfclient.Space
		result2 error
	}
	createSpaceReturnsOnCall map[int]struct {
		result1 cfclient.Space
		result2 error
	}
	DeleteSpaceStub        func(string, bool, bool) error
	deleteSpaceMutex       sync.RWMutex
	deleteSpaceArgsForCall []struct {
		arg1 string
		arg2 bool
		arg3 bool
	}
	deleteSpaceReturns struct {
		result1 error
	}
	deleteSpaceReturnsOnCall map[int]struct {
		result1 error
	}
	GetSpaceByGuidStub        func(string) (cfclient.Space, error)
	getSpaceByGuidMutex       sync.RWMutex
	getSpaceByGuidArgsForCall []struct {
		arg1 string
	}
	getSpaceByGuidReturns struct {
		result1 cfclient.Space
		result2 error
	}
	getSpaceByGuidReturnsOnCall map[int]struct {
		result1 cfclient.Space
		result2 error
	}
	ListOrgsStub        func() ([]cfclient.Org, error)
	listOrgsMutex       sync.RWMutex
	listOrgsArgsForCall []struct {
	}
	listOrgsReturns struct {
		result1 []cfclient.Org
		result2 error
	}
	listOrgsReturnsOnCall map[int]struct {
		result1 []cfclient.Org
		result2 error
	}
	ListSpacesStub        func() ([]cfclient.Space, error)
	listSpacesMutex       sync.RWMutex
	listSpacesArgsForCall []struct {
	}
	listSpacesReturns struct {
		result1 []cfclient.Space
		result2 error
	}
	listSpacesReturnsOnCall map[int]struct {
		result1 []cfclient.Space
		result2 error
	}
	RemoveSpaceMetadataStub        func(string) error
	removeSpaceMetadataMutex       sync.RWMutex
	removeSpaceMetadataArgsForCall []struct {
		arg1 string
	}
	removeSpaceMetadataReturns struct {
		result1 error
	}
	removeSpaceMetadataReturnsOnCall map[int]struct {
		result1 error
	}
	SpaceMetadataStub        func(string) (*cfclient.Metadata, error)
	spaceMetadataMutex       sync.RWMutex
	spaceMetadataArgsForCall []struct {
		arg1 string
	}
	spaceMetadataReturns struct {
		result1 *cfclient.Metadata
		result2 error
	}
	spaceMetadataReturnsOnCall map[int]struct {
		result1 *cfclient.Metadata
		result2 error
	}
	SupportsMetadataAPIStub        func() (bool, error)
	supportsMetadataAPIMutex       sync.RWMutex
	supportsMetadataAPIArgsForCall []struct {
	}
	supportsMetadataAPIReturns struct {
		result1 bool
		result2 error
	}
	supportsMetadataAPIReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	UpdateSpaceStub        func(string, cfclient.SpaceRequest) (cfclient.Space, error)
	updateSpaceMutex       sync.RWMutex
	updateSpaceArgsForCall []struct {
		arg1 string
		arg2 cfclient.SpaceRequest
	}
	updateSpaceReturns struct {
		result1 cfclient.Space
		result2 error
	}
	updateSpaceReturnsOnCall map[int]struct {
		result1 cfclient.Space
		result2 error
	}
	UpdateSpaceMetadataStub        func(string, cfclient.Metadata) error
	updateSpaceMetadataMutex       sync.RWMutex
	updateSpaceMetadataArgsForCall []struct {
		arg1 string
		arg2 cfclient.Metadata
	}
	updateSpaceMetadataReturns struct {
		result1 error
	}
	updateSpaceMetadataReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCFClient) CreateSpace(arg1 cfclient.SpaceRequest) (cfclient.Space, error) {
	fake.createSpaceMutex.Lock()
	ret, specificReturn := fake.createSpaceReturnsOnCall[len(fake.createSpaceArgsForCall)]
	fake.createSpaceArgsForCall = append(fake.createSpaceArgsForCall, struct {
		arg1 cfclient.SpaceRequest
	}{arg1})
	fake.recordInvocation("CreateSpace", []interface{}{arg1})
	fake.createSpaceMutex.Unlock()
	if fake.CreateSpaceStub != nil {
		return fake.CreateSpaceStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createSpaceReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCFClient) CreateSpaceCallCount() int {
	fake.createSpaceMutex.RLock()
	defer fake.createSpaceMutex.RUnlock()
	return len(fake.createSpaceArgsForCall)
}

func (fake *FakeCFClient) CreateSpaceCalls(stub func(cfclient.SpaceRequest) (cfclient.Space, error)) {
	fake.createSpaceMutex.Lock()
	defer fake.createSpaceMutex.Unlock()
	fake.CreateSpaceStub = stub
}

func (fake *FakeCFClient) CreateSpaceArgsForCall(i int) cfclient.SpaceRequest {
	fake.createSpaceMutex.RLock()
	defer fake.createSpaceMutex.RUnlock()
	argsForCall := fake.createSpaceArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCFClient) CreateSpaceReturns(result1 cfclient.Space, result2 error) {
	fake.createSpaceMutex.Lock()
	defer fake.createSpaceMutex.Unlock()
	fake.CreateSpaceStub = nil
	fake.createSpaceReturns = struct {
		result1 cfclient.Space
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) CreateSpaceReturnsOnCall(i int, result1 cfclient.Space, result2 error) {
	fake.createSpaceMutex.Lock()
	defer fake.createSpaceMutex.Unlock()
	fake.CreateSpaceStub = nil
	if fake.createSpaceReturnsOnCall == nil {
		fake.createSpaceReturnsOnCall = make(map[int]struct {
			result1 cfclient.Space
			result2 error
		})
	}
	fake.createSpaceReturnsOnCall[i] = struct {
		result1 cfclient.Space
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) DeleteSpace(arg1 string, arg2 bool, arg3 bool) error {
	fake.deleteSpaceMutex.Lock()
	ret, specificReturn := fake.deleteSpaceReturnsOnCall[len(fake.deleteSpaceArgsForCall)]
	fake.deleteSpaceArgsForCall = append(fake.deleteSpaceArgsForCall, struct {
		arg1 string
		arg2 bool
		arg3 bool
	}{arg1, arg2, arg3})
	fake.recordInvocation("DeleteSpace", []interface{}{arg1, arg2, arg3})
	fake.deleteSpaceMutex.Unlock()
	if fake.DeleteSpaceStub != nil {
		return fake.DeleteSpaceStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteSpaceReturns
	return fakeReturns.result1
}

func (fake *FakeCFClient) DeleteSpaceCallCount() int {
	fake.deleteSpaceMutex.RLock()
	defer fake.deleteSpaceMutex.RUnlock()
	return len(fake.deleteSpaceArgsForCall)
}

func (fake *FakeCFClient) DeleteSpaceCalls(stub func(string, bool, bool) error) {
	fake.deleteSpaceMutex.Lock()
	defer fake.deleteSpaceMutex.Unlock()
	fake.DeleteSpaceStub = stub
}

func (fake *FakeCFClient) DeleteSpaceArgsForCall(i int) (string, bool, bool) {
	fake.deleteSpaceMutex.RLock()
	defer fake.deleteSpaceMutex.RUnlock()
	argsForCall := fake.deleteSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCFClient) DeleteSpaceReturns(result1 error) {
	fake.deleteSpaceMutex.Lock()
	defer fake.deleteSpaceMutex.Unlock()
	fake.DeleteSpaceStub = nil
	fake.deleteSpaceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCFClient) DeleteSpaceReturnsOnCall(i int, result1 error) {
	fake.deleteSpaceMutex.Lock()
	defer fake.deleteSpaceMutex.Unlock()
	fake.DeleteSpaceStub = nil
	if fake.deleteSpaceReturnsOnCall == nil {
		fake.deleteSpaceReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteSpaceReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCFClient) GetSpaceByGuid(arg1 string) (cfclient.Space, error) {
	fake.getSpaceByGuidMutex.Lock()
	ret, specificReturn := fake.getSpaceByGuidReturnsOnCall[len(fake.getSpaceByGuidArgsForCall)]
	fake.getSpaceByGuidArgsForCall = append(fake.getSpaceByGuidArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetSpaceByGuid", []interface{}{arg1})
	fake.getSpaceByGuidMutex.Unlock()
	if fake.GetSpaceByGuidStub != nil {
		return fake.GetSpaceByGuidStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getSpaceByGuidReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCFClient) GetSpaceByGuidCallCount() int {
	fake.getSpaceByGuidMutex.RLock()
	defer fake.getSpaceByGuidMutex.RUnlock()
	return len(fake.getSpaceByGuidArgsForCall)
}

func (fake *FakeCFClient) GetSpaceByGuidCalls(stub func(string) (cfclient.Space, error)) {
	fake.getSpaceByGuidMutex.Lock()
	defer fake.getSpaceByGuidMutex.Unlock()
	fake.GetSpaceByGuidStub = stub
}

func (fake *FakeCFClient) GetSpaceByGuidArgsForCall(i int) string {
	fake.getSpaceByGuidMutex.RLock()
	defer fake.getSpaceByGuidMutex.RUnlock()
	argsForCall := fake.getSpaceByGuidArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCFClient) GetSpaceByGuidReturns(result1 cfclient.Space, result2 error) {
	fake.getSpaceByGuidMutex.Lock()
	defer fake.getSpaceByGuidMutex.Unlock()
	fake.GetSpaceByGuidStub = nil
	fake.getSpaceByGuidReturns = struct {
		result1 cfclient.Space
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) GetSpaceByGuidReturnsOnCall(i int, result1 cfclient.Space, result2 error) {
	fake.getSpaceByGuidMutex.Lock()
	defer fake.getSpaceByGuidMutex.Unlock()
	fake.GetSpaceByGuidStub = nil
	if fake.getSpaceByGuidReturnsOnCall == nil {
		fake.getSpaceByGuidReturnsOnCall = make(map[int]struct {
			result1 cfclient.Space
			result2 error
		})
	}
	fake.getSpaceByGuidReturnsOnCall[i] = struct {
		result1 cfclient.Space
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) ListOrgs() ([]cfclient.Org, error) {
	fake.listOrgsMutex.Lock()
	ret, specificReturn := fake.listOrgsReturnsOnCall[len(fake.listOrgsArgsForCall)]
	fake.listOrgsArgsForCall = append(fake.listOrgsArgsForCall, struct {
	}{})
	fake.recordInvocation("ListOrgs", []interface{}{})
	fake.listOrgsMutex.Unlock()
	if fake.ListOrgsStub != nil {
		return fake.ListOrgsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listOrgsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCFClient) ListOrgsCallCount() int {
	fake.listOrgsMutex.RLock()
	defer fake.listOrgsMutex.RUnlock()
	return len(fake.listOrgsArgsForCall)
}

func (fake *FakeCFClient) ListOrgsCalls(stub func() ([]cfclient.Org, error)) {
	fake.listOrgsMutex.Lock()
	defer fake.listOrgsMutex.Unlock()
	fake.ListOrgsStub = stub
}

func (fake *FakeCFClient) ListOrgsReturns(result1 []cfclient.Org, result2 error) {
	fake.listOrgsMutex.Lock()
	defer fake.listOrgsMutex.Unlock()
	fake.ListOrgsStub = nil
	fake.listOrgsReturns = struct {
		result1 []cfclient.Org
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) ListOrgsReturnsOnCall(i int, result1 []cfclient.Org, result2 error) {
	fake.listOrgsMutex.Lock()
	defer fake.listOrgsMutex.Unlock()
	fake.ListOrgsStub = nil
	if fake.listOrgsReturnsOnCall == nil {
		fake.listOrgsReturnsOnCall = make(map[int]struct {
			result1 []cfclient.Org
			result2 error
		})
	}
	fake.listOrgsReturnsOnCall[i] = struct {
		result1 []cfclient.Org
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) ListSpaces() ([]cfclient.Space, error) {
	fake.listSpacesMutex.Lock()
	ret, specificReturn := fake.listSpacesReturnsOnCall[len(fake.listSpacesArgsForCall)]
	fake.listSpacesArgsForCall = append(fake.listSpacesArgsForCall, struct {
	}{})
	fake.recordInvocation("ListSpaces", []interface{}{})
	fake.listSpacesMutex.Unlock()
	if fake.ListSpacesStub != nil {
		return fake.ListSpacesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listSpacesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCFClient) ListSpacesCallCount() int {
	fake.listSpacesMutex.RLock()
	defer fake.listSpacesMutex.RUnlock()
	return len(fake.listSpacesArgsForCall)
}

func (fake *FakeCFClient) ListSpacesCalls(stub func() ([]cfclient.Space, error)) {
	fake.listSpacesMutex.Lock()
	defer fake.listSpacesMutex.Unlock()
	fake.ListSpacesStub = stub
}

func (fake *FakeCFClient) ListSpacesReturns(result1 []cfclient.Space, result2 error) {
	fake.listSpacesMutex.Lock()
	defer fake.listSpacesMutex.Unlock()
	fake.ListSpacesStub = nil
	fake.listSpacesReturns = struct {
		result1 []cfclient.Space
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) ListSpacesReturnsOnCall(i int, result1 []cfclient.Space, result2 error) {
	fake.listSpacesMutex.Lock()
	defer fake.listSpacesMutex.Unlock()
	fake.ListSpacesStub = nil
	if fake.listSpacesReturnsOnCall == nil {
		fake.listSpacesReturnsOnCall = make(map[int]struct {
			result1 []cfclient.Space
			result2 error
		})
	}
	fake.listSpacesReturnsOnCall[i] = struct {
		result1 []cfclient.Space
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) RemoveSpaceMetadata(arg1 string) error {
	fake.removeSpaceMetadataMutex.Lock()
	ret, specificReturn := fake.removeSpaceMetadataReturnsOnCall[len(fake.removeSpaceMetadataArgsForCall)]
	fake.removeSpaceMetadataArgsForCall = append(fake.removeSpaceMetadataArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("RemoveSpaceMetadata", []interface{}{arg1})
	fake.removeSpaceMetadataMutex.Unlock()
	if fake.RemoveSpaceMetadataStub != nil {
		return fake.RemoveSpaceMetadataStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.removeSpaceMetadataReturns
	return fakeReturns.result1
}

func (fake *FakeCFClient) RemoveSpaceMetadataCallCount() int {
	fake.removeSpaceMetadataMutex.RLock()
	defer fake.removeSpaceMetadataMutex.RUnlock()
	return len(fake.removeSpaceMetadataArgsForCall)
}

func (fake *FakeCFClient) RemoveSpaceMetadataCalls(stub func(string) error) {
	fake.removeSpaceMetadataMutex.Lock()
	defer fake.removeSpaceMetadataMutex.Unlock()
	fake.RemoveSpaceMetadataStub = stub
}

func (fake *FakeCFClient) RemoveSpaceMetadataArgsForCall(i int) string {
	fake.removeSpaceMetadataMutex.RLock()
	defer fake.removeSpaceMetadataMutex.RUnlock()
	argsForCall := fake.removeSpaceMetadataArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCFClient) RemoveSpaceMetadataReturns(result1 error) {
	fake.removeSpaceMetadataMutex.Lock()
	defer fake.removeSpaceMetadataMutex.Unlock()
	fake.RemoveSpaceMetadataStub = nil
	fake.removeSpaceMetadataReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCFClient) RemoveSpaceMetadataReturnsOnCall(i int, result1 error) {
	fake.removeSpaceMetadataMutex.Lock()
	defer fake.removeSpaceMetadataMutex.Unlock()
	fake.RemoveSpaceMetadataStub = nil
	if fake.removeSpaceMetadataReturnsOnCall == nil {
		fake.removeSpaceMetadataReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeSpaceMetadataReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCFClient) SpaceMetadata(arg1 string) (*cfclient.Metadata, error) {
	fake.spaceMetadataMutex.Lock()
	ret, specificReturn := fake.spaceMetadataReturnsOnCall[len(fake.spaceMetadataArgsForCall)]
	fake.spaceMetadataArgsForCall = append(fake.spaceMetadataArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("SpaceMetadata", []interface{}{arg1})
	fake.spaceMetadataMutex.Unlock()
	if fake.SpaceMetadataStub != nil {
		return fake.SpaceMetadataStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.spaceMetadataReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCFClient) SpaceMetadataCallCount() int {
	fake.spaceMetadataMutex.RLock()
	defer fake.spaceMetadataMutex.RUnlock()
	return len(fake.spaceMetadataArgsForCall)
}

func (fake *FakeCFClient) SpaceMetadataCalls(stub func(string) (*cfclient.Metadata, error)) {
	fake.spaceMetadataMutex.Lock()
	defer fake.spaceMetadataMutex.Unlock()
	fake.SpaceMetadataStub = stub
}

func (fake *FakeCFClient) SpaceMetadataArgsForCall(i int) string {
	fake.spaceMetadataMutex.RLock()
	defer fake.spaceMetadataMutex.RUnlock()
	argsForCall := fake.spaceMetadataArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCFClient) SpaceMetadataReturns(result1 *cfclient.Metadata, result2 error) {
	fake.spaceMetadataMutex.Lock()
	defer fake.spaceMetadataMutex.Unlock()
	fake.SpaceMetadataStub = nil
	fake.spaceMetadataReturns = struct {
		result1 *cfclient.Metadata
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) SpaceMetadataReturnsOnCall(i int, result1 *cfclient.Metadata, result2 error) {
	fake.spaceMetadataMutex.Lock()
	defer fake.spaceMetadataMutex.Unlock()
	fake.SpaceMetadataStub = nil
	if fake.spaceMetadataReturnsOnCall == nil {
		fake.spaceMetadataReturnsOnCall = make(map[int]struct {
			result1 *cfclient.Metadata
			result2 error
		})
	}
	fake.spaceMetadataReturnsOnCall[i] = struct {
		result1 *cfclient.Metadata
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) SupportsMetadataAPI() (bool, error) {
	fake.supportsMetadataAPIMutex.Lock()
	ret, specificReturn := fake.supportsMetadataAPIReturnsOnCall[len(fake.supportsMetadataAPIArgsForCall)]
	fake.supportsMetadataAPIArgsForCall = append(fake.supportsMetadataAPIArgsForCall, struct {
	}{})
	fake.recordInvocation("SupportsMetadataAPI", []interface{}{})
	fake.supportsMetadataAPIMutex.Unlock()
	if fake.SupportsMetadataAPIStub != nil {
		return fake.SupportsMetadataAPIStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.supportsMetadataAPIReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCFClient) SupportsMetadataAPICallCount() int {
	fake.supportsMetadataAPIMutex.RLock()
	defer fake.supportsMetadataAPIMutex.RUnlock()
	return len(fake.supportsMetadataAPIArgsForCall)
}

func (fake *FakeCFClient) SupportsMetadataAPICalls(stub func() (bool, error)) {
	fake.supportsMetadataAPIMutex.Lock()
	defer fake.supportsMetadataAPIMutex.Unlock()
	fake.SupportsMetadataAPIStub = stub
}

func (fake *FakeCFClient) SupportsMetadataAPIReturns(result1 bool, result2 error) {
	fake.supportsMetadataAPIMutex.Lock()
	defer fake.supportsMetadataAPIMutex.Unlock()
	fake.SupportsMetadataAPIStub = nil
	fake.supportsMetadataAPIReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) SupportsMetadataAPIReturnsOnCall(i int, result1 bool, result2 error) {
	fake.supportsMetadataAPIMutex.Lock()
	defer fake.supportsMetadataAPIMutex.Unlock()
	fake.SupportsMetadataAPIStub = nil
	if fake.supportsMetadataAPIReturnsOnCall == nil {
		fake.supportsMetadataAPIReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.supportsMetadataAPIReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) UpdateSpace(arg1 string, arg2 cfclient.SpaceRequest) (cfclient.Space, error) {
	fake.updateSpaceMutex.Lock()
	ret, specificReturn := fake.updateSpaceReturnsOnCall[len(fake.updateSpaceArgsForCall)]
	fake.updateSpaceArgsForCall = append(fake.updateSpaceArgsForCall, struct {
		arg1 string
		arg2 cfclient.SpaceRequest
	}{arg1, arg2})
	fake.recordInvocation("UpdateSpace", []interface{}{arg1, arg2})
	fake.updateSpaceMutex.Unlock()
	if fake.UpdateSpaceStub != nil {
		return fake.UpdateSpaceStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.updateSpaceReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCFClient) UpdateSpaceCallCount() int {
	fake.updateSpaceMutex.RLock()
	defer fake.updateSpaceMutex.RUnlock()
	return len(fake.updateSpaceArgsForCall)
}

func (fake *FakeCFClient) UpdateSpaceCalls(stub func(string, cfclient.SpaceRequest) (cfclient.Space, error)) {
	fake.updateSpaceMutex.Lock()
	defer fake.updateSpaceMutex.Unlock()
	fake.UpdateSpaceStub = stub
}

func (fake *FakeCFClient) UpdateSpaceArgsForCall(i int) (string, cfclient.SpaceRequest) {
	fake.updateSpaceMutex.RLock()
	defer fake.updateSpaceMutex.RUnlock()
	argsForCall := fake.updateSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCFClient) UpdateSpaceReturns(result1 cfclient.Space, result2 error) {
	fake.updateSpaceMutex.Lock()
	defer fake.updateSpaceMutex.Unlock()
	fake.UpdateSpaceStub = nil
	fake.updateSpaceReturns = struct {
		result1 cfclient.Space
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) UpdateSpaceReturnsOnCall(i int, result1 cfclient.Space, result2 error) {
	fake.updateSpaceMutex.Lock()
	defer fake.updateSpaceMutex.Unlock()
	fake.UpdateSpaceStub = nil
	if fake.updateSpaceReturnsOnCall == nil {
		fake.updateSpaceReturnsOnCall = make(map[int]struct {
			result1 cfclient.Space
			result2 error
		})
	}
	fake.updateSpaceReturnsOnCall[i] = struct {
		result1 cfclient.Space
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) UpdateSpaceMetadata(arg1 string, arg2 cfclient.Metadata) error {
	fake.updateSpaceMetadataMutex.Lock()
	ret, specificReturn := fake.updateSpaceMetadataReturnsOnCall[len(fake.updateSpaceMetadataArgsForCall)]
	fake.updateSpaceMetadataArgsForCall = append(fake.updateSpaceMetadataArgsForCall, struct {
		arg1 string
		arg2 cfclient.Metadata
	}{arg1, arg2})
	fake.recordInvocation("UpdateSpaceMetadata", []interface{}{arg1, arg2})
	fake.updateSpaceMetadataMutex.Unlock()
	if fake.UpdateSpaceMetadataStub != nil {
		return fake.UpdateSpaceMetadataStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.updateSpaceMetadataReturns
	return fakeReturns.result1
}

func (fake *FakeCFClient) UpdateSpaceMetadataCallCount() int {
	fake.updateSpaceMetadataMutex.RLock()
	defer fake.updateSpaceMetadataMutex.RUnlock()
	return len(fake.updateSpaceMetadataArgsForCall)
}

func (fake *FakeCFClient) UpdateSpaceMetadataCalls(stub func(string, cfclient.Metadata) error) {
	fake.updateSpaceMetadataMutex.Lock()
	defer fake.updateSpaceMetadataMutex.Unlock()
	fake.UpdateSpaceMetadataStub = stub
}

func (fake *FakeCFClient) UpdateSpaceMetadataArgsForCall(i int) (string, cfclient.Metadata) {
	fake.updateSpaceMetadataMutex.RLock()
	defer fake.updateSpaceMetadataMutex.RUnlock()
	argsForCall := fake.updateSpaceMetadataArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCFClient) UpdateSpaceMetadataReturns(result1 error) {
	fake.updateSpaceMetadataMutex.Lock()
	defer fake.updateSpaceMetadataMutex.Unlock()
	fake.UpdateSpaceMetadataStub = nil
	fake.updateSpaceMetadataReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCFClient) UpdateSpaceMetadataReturnsOnCall(i int, result1 error) {
	fake.updateSpaceMetadataMutex.Lock()
	defer fake.updateSpaceMetadataMutex.Unlock()
	fake.UpdateSpaceMetadataStub = nil
	if fake.updateSpaceMetadataReturnsOnCall == nil {
		fake.updateSpaceMetadataReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateSpaceMetadataReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCFClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createSpaceMutex.RLock()
	defer fake.createSpaceMutex.RUnlock()
	fake.deleteSpaceMutex.RLock()
	defer fake.deleteSpaceMutex.RUnlock()
	fake.getSpaceByGuidMutex.RLock()
	defer fake.getSpaceByGuidMutex.RUnlock()
	fake.listOrgsMutex.RLock()
	defer fake.listOrgsMutex.RUnlock()
	fake.listSpacesMutex.RLock()
	defer fake.listSpacesMutex.RUnlock()
	fake.removeSpaceMetadataMutex.RLock()
	defer fake.removeSpaceMetadataMutex.RUnlock()
	fake.spaceMetadataMutex.RLock()
	defer fake.spaceMetadataMutex.RUnlock()
	fake.supportsMetadataAPIMutex.RLock()
	defer fake.supportsMetadataAPIMutex.RUnlock()
	fake.updateSpaceMutex.RLock()
	defer fake.updateSpaceMutex.RUnlock()
	fake.updateSpaceMetadataMutex.RLock()
	defer fake.updateSpaceMetadataMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCFClient) recordInvocation(key string, args []interface{}) {
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

var _ space.CFClient = new(FakeCFClient)
