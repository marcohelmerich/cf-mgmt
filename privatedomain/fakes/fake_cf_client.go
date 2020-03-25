// This file was generated by counterfeiter
package fakes

import (
	"sync"

	go_cfclient "github.com/cloudfoundry-community/go-cfclient"
	"github.com/vmwarepivotallabs/cf-mgmt/privatedomain"
)

type FakeCFClient struct {
	ListDomainsStub        func() ([]go_cfclient.Domain, error)
	listDomainsMutex       sync.RWMutex
	listDomainsArgsForCall []struct{}
	listDomainsReturns     struct {
		result1 []go_cfclient.Domain
		result2 error
	}
	CreateDomainStub        func(name, orgGuid string) (*go_cfclient.Domain, error)
	createDomainMutex       sync.RWMutex
	createDomainArgsForCall []struct {
		name    string
		orgGuid string
	}
	createDomainReturns struct {
		result1 *go_cfclient.Domain
		result2 error
	}
	ShareOrgPrivateDomainStub        func(orgGUID, privateDomainGUID string) (*go_cfclient.Domain, error)
	shareOrgPrivateDomainMutex       sync.RWMutex
	shareOrgPrivateDomainArgsForCall []struct {
		orgGUID           string
		privateDomainGUID string
	}
	shareOrgPrivateDomainReturns struct {
		result1 *go_cfclient.Domain
		result2 error
	}
	ListOrgPrivateDomainsStub        func(orgGUID string) ([]go_cfclient.Domain, error)
	listOrgPrivateDomainsMutex       sync.RWMutex
	listOrgPrivateDomainsArgsForCall []struct {
		orgGUID string
	}
	listOrgPrivateDomainsReturns struct {
		result1 []go_cfclient.Domain
		result2 error
	}
	DeleteDomainStub        func(guid string) error
	deleteDomainMutex       sync.RWMutex
	deleteDomainArgsForCall []struct {
		guid string
	}
	deleteDomainReturns struct {
		result1 error
	}
	UnshareOrgPrivateDomainStub        func(orgGUID, privateDomainGUID string) error
	unshareOrgPrivateDomainMutex       sync.RWMutex
	unshareOrgPrivateDomainArgsForCall []struct {
		orgGUID           string
		privateDomainGUID string
	}
	unshareOrgPrivateDomainReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCFClient) ListDomains() ([]go_cfclient.Domain, error) {
	fake.listDomainsMutex.Lock()
	fake.listDomainsArgsForCall = append(fake.listDomainsArgsForCall, struct{}{})
	fake.recordInvocation("ListDomains", []interface{}{})
	fake.listDomainsMutex.Unlock()
	if fake.ListDomainsStub != nil {
		return fake.ListDomainsStub()
	} else {
		return fake.listDomainsReturns.result1, fake.listDomainsReturns.result2
	}
}

func (fake *FakeCFClient) ListDomainsCallCount() int {
	fake.listDomainsMutex.RLock()
	defer fake.listDomainsMutex.RUnlock()
	return len(fake.listDomainsArgsForCall)
}

func (fake *FakeCFClient) ListDomainsReturns(result1 []go_cfclient.Domain, result2 error) {
	fake.ListDomainsStub = nil
	fake.listDomainsReturns = struct {
		result1 []go_cfclient.Domain
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) CreateDomain(name string, orgGuid string) (*go_cfclient.Domain, error) {
	fake.createDomainMutex.Lock()
	fake.createDomainArgsForCall = append(fake.createDomainArgsForCall, struct {
		name    string
		orgGuid string
	}{name, orgGuid})
	fake.recordInvocation("CreateDomain", []interface{}{name, orgGuid})
	fake.createDomainMutex.Unlock()
	if fake.CreateDomainStub != nil {
		return fake.CreateDomainStub(name, orgGuid)
	} else {
		return fake.createDomainReturns.result1, fake.createDomainReturns.result2
	}
}

func (fake *FakeCFClient) CreateDomainCallCount() int {
	fake.createDomainMutex.RLock()
	defer fake.createDomainMutex.RUnlock()
	return len(fake.createDomainArgsForCall)
}

func (fake *FakeCFClient) CreateDomainArgsForCall(i int) (string, string) {
	fake.createDomainMutex.RLock()
	defer fake.createDomainMutex.RUnlock()
	return fake.createDomainArgsForCall[i].name, fake.createDomainArgsForCall[i].orgGuid
}

func (fake *FakeCFClient) CreateDomainReturns(result1 *go_cfclient.Domain, result2 error) {
	fake.CreateDomainStub = nil
	fake.createDomainReturns = struct {
		result1 *go_cfclient.Domain
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) ShareOrgPrivateDomain(orgGUID string, privateDomainGUID string) (*go_cfclient.Domain, error) {
	fake.shareOrgPrivateDomainMutex.Lock()
	fake.shareOrgPrivateDomainArgsForCall = append(fake.shareOrgPrivateDomainArgsForCall, struct {
		orgGUID           string
		privateDomainGUID string
	}{orgGUID, privateDomainGUID})
	fake.recordInvocation("ShareOrgPrivateDomain", []interface{}{orgGUID, privateDomainGUID})
	fake.shareOrgPrivateDomainMutex.Unlock()
	if fake.ShareOrgPrivateDomainStub != nil {
		return fake.ShareOrgPrivateDomainStub(orgGUID, privateDomainGUID)
	} else {
		return fake.shareOrgPrivateDomainReturns.result1, fake.shareOrgPrivateDomainReturns.result2
	}
}

func (fake *FakeCFClient) ShareOrgPrivateDomainCallCount() int {
	fake.shareOrgPrivateDomainMutex.RLock()
	defer fake.shareOrgPrivateDomainMutex.RUnlock()
	return len(fake.shareOrgPrivateDomainArgsForCall)
}

func (fake *FakeCFClient) ShareOrgPrivateDomainArgsForCall(i int) (string, string) {
	fake.shareOrgPrivateDomainMutex.RLock()
	defer fake.shareOrgPrivateDomainMutex.RUnlock()
	return fake.shareOrgPrivateDomainArgsForCall[i].orgGUID, fake.shareOrgPrivateDomainArgsForCall[i].privateDomainGUID
}

func (fake *FakeCFClient) ShareOrgPrivateDomainReturns(result1 *go_cfclient.Domain, result2 error) {
	fake.ShareOrgPrivateDomainStub = nil
	fake.shareOrgPrivateDomainReturns = struct {
		result1 *go_cfclient.Domain
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) ListOrgPrivateDomains(orgGUID string) ([]go_cfclient.Domain, error) {
	fake.listOrgPrivateDomainsMutex.Lock()
	fake.listOrgPrivateDomainsArgsForCall = append(fake.listOrgPrivateDomainsArgsForCall, struct {
		orgGUID string
	}{orgGUID})
	fake.recordInvocation("ListOrgPrivateDomains", []interface{}{orgGUID})
	fake.listOrgPrivateDomainsMutex.Unlock()
	if fake.ListOrgPrivateDomainsStub != nil {
		return fake.ListOrgPrivateDomainsStub(orgGUID)
	} else {
		return fake.listOrgPrivateDomainsReturns.result1, fake.listOrgPrivateDomainsReturns.result2
	}
}

func (fake *FakeCFClient) ListOrgPrivateDomainsCallCount() int {
	fake.listOrgPrivateDomainsMutex.RLock()
	defer fake.listOrgPrivateDomainsMutex.RUnlock()
	return len(fake.listOrgPrivateDomainsArgsForCall)
}

func (fake *FakeCFClient) ListOrgPrivateDomainsArgsForCall(i int) string {
	fake.listOrgPrivateDomainsMutex.RLock()
	defer fake.listOrgPrivateDomainsMutex.RUnlock()
	return fake.listOrgPrivateDomainsArgsForCall[i].orgGUID
}

func (fake *FakeCFClient) ListOrgPrivateDomainsReturns(result1 []go_cfclient.Domain, result2 error) {
	fake.ListOrgPrivateDomainsStub = nil
	fake.listOrgPrivateDomainsReturns = struct {
		result1 []go_cfclient.Domain
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) DeleteDomain(guid string) error {
	fake.deleteDomainMutex.Lock()
	fake.deleteDomainArgsForCall = append(fake.deleteDomainArgsForCall, struct {
		guid string
	}{guid})
	fake.recordInvocation("DeleteDomain", []interface{}{guid})
	fake.deleteDomainMutex.Unlock()
	if fake.DeleteDomainStub != nil {
		return fake.DeleteDomainStub(guid)
	} else {
		return fake.deleteDomainReturns.result1
	}
}

func (fake *FakeCFClient) DeleteDomainCallCount() int {
	fake.deleteDomainMutex.RLock()
	defer fake.deleteDomainMutex.RUnlock()
	return len(fake.deleteDomainArgsForCall)
}

func (fake *FakeCFClient) DeleteDomainArgsForCall(i int) string {
	fake.deleteDomainMutex.RLock()
	defer fake.deleteDomainMutex.RUnlock()
	return fake.deleteDomainArgsForCall[i].guid
}

func (fake *FakeCFClient) DeleteDomainReturns(result1 error) {
	fake.DeleteDomainStub = nil
	fake.deleteDomainReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCFClient) UnshareOrgPrivateDomain(orgGUID string, privateDomainGUID string) error {
	fake.unshareOrgPrivateDomainMutex.Lock()
	fake.unshareOrgPrivateDomainArgsForCall = append(fake.unshareOrgPrivateDomainArgsForCall, struct {
		orgGUID           string
		privateDomainGUID string
	}{orgGUID, privateDomainGUID})
	fake.recordInvocation("UnshareOrgPrivateDomain", []interface{}{orgGUID, privateDomainGUID})
	fake.unshareOrgPrivateDomainMutex.Unlock()
	if fake.UnshareOrgPrivateDomainStub != nil {
		return fake.UnshareOrgPrivateDomainStub(orgGUID, privateDomainGUID)
	} else {
		return fake.unshareOrgPrivateDomainReturns.result1
	}
}

func (fake *FakeCFClient) UnshareOrgPrivateDomainCallCount() int {
	fake.unshareOrgPrivateDomainMutex.RLock()
	defer fake.unshareOrgPrivateDomainMutex.RUnlock()
	return len(fake.unshareOrgPrivateDomainArgsForCall)
}

func (fake *FakeCFClient) UnshareOrgPrivateDomainArgsForCall(i int) (string, string) {
	fake.unshareOrgPrivateDomainMutex.RLock()
	defer fake.unshareOrgPrivateDomainMutex.RUnlock()
	return fake.unshareOrgPrivateDomainArgsForCall[i].orgGUID, fake.unshareOrgPrivateDomainArgsForCall[i].privateDomainGUID
}

func (fake *FakeCFClient) UnshareOrgPrivateDomainReturns(result1 error) {
	fake.UnshareOrgPrivateDomainStub = nil
	fake.unshareOrgPrivateDomainReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCFClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listDomainsMutex.RLock()
	defer fake.listDomainsMutex.RUnlock()
	fake.createDomainMutex.RLock()
	defer fake.createDomainMutex.RUnlock()
	fake.shareOrgPrivateDomainMutex.RLock()
	defer fake.shareOrgPrivateDomainMutex.RUnlock()
	fake.listOrgPrivateDomainsMutex.RLock()
	defer fake.listOrgPrivateDomainsMutex.RUnlock()
	fake.deleteDomainMutex.RLock()
	defer fake.deleteDomainMutex.RUnlock()
	fake.unshareOrgPrivateDomainMutex.RLock()
	defer fake.unshareOrgPrivateDomainMutex.RUnlock()
	return fake.invocations
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

var _ privatedomain.CFClient = new(FakeCFClient)
