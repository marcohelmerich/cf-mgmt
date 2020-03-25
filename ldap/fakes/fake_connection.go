// This file was generated by counterfeiter
package fakes

import (
	"sync"

	ldapgo_ldap "github.com/go-ldap/ldap"
	"github.com/vmwarepivotallabs/cf-mgmt/ldap"
)

type FakeConnection struct {
	CloseStub         func()
	closeMutex        sync.RWMutex
	closeArgsForCall  []struct{}
	SearchStub        func(*ldapgo_ldap.SearchRequest) (*ldapgo_ldap.SearchResult, error)
	searchMutex       sync.RWMutex
	searchArgsForCall []struct {
		arg1 *ldapgo_ldap.SearchRequest
	}
	searchReturns struct {
		result1 *ldapgo_ldap.SearchResult
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeConnection) Close() {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct{}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		fake.CloseStub()
	}
}

func (fake *FakeConnection) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeConnection) Search(arg1 *ldapgo_ldap.SearchRequest) (*ldapgo_ldap.SearchResult, error) {
	fake.searchMutex.Lock()
	fake.searchArgsForCall = append(fake.searchArgsForCall, struct {
		arg1 *ldapgo_ldap.SearchRequest
	}{arg1})
	fake.recordInvocation("Search", []interface{}{arg1})
	fake.searchMutex.Unlock()
	if fake.SearchStub != nil {
		return fake.SearchStub(arg1)
	} else {
		return fake.searchReturns.result1, fake.searchReturns.result2
	}
}

func (fake *FakeConnection) SearchCallCount() int {
	fake.searchMutex.RLock()
	defer fake.searchMutex.RUnlock()
	return len(fake.searchArgsForCall)
}

func (fake *FakeConnection) SearchArgsForCall(i int) *ldapgo_ldap.SearchRequest {
	fake.searchMutex.RLock()
	defer fake.searchMutex.RUnlock()
	return fake.searchArgsForCall[i].arg1
}

func (fake *FakeConnection) SearchReturns(result1 *ldapgo_ldap.SearchResult, result2 error) {
	fake.SearchStub = nil
	fake.searchReturns = struct {
		result1 *ldapgo_ldap.SearchResult
		result2 error
	}{result1, result2}
}

func (fake *FakeConnection) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.searchMutex.RLock()
	defer fake.searchMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeConnection) recordInvocation(key string, args []interface{}) {
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

var _ ldap.Connection = new(FakeConnection)
