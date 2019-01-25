// Code generated by counterfeiter. DO NOT EDIT.
package fab3

import (
	sync "sync"

	fab3 "github.com/hyperledger/fabric-chaincode-evm/fab3"
	channel "github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

type MockChannelClient struct {
	ExecuteStub        func(channel.Request, ...channel.RequestOption) (channel.Response, error)
	executeMutex       sync.RWMutex
	executeArgsForCall []struct {
		arg1 channel.Request
		arg2 []channel.RequestOption
	}
	executeReturns struct {
		result1 channel.Response
		result2 error
	}
	executeReturnsOnCall map[int]struct {
		result1 channel.Response
		result2 error
	}
	QueryStub        func(channel.Request, ...channel.RequestOption) (channel.Response, error)
	queryMutex       sync.RWMutex
	queryArgsForCall []struct {
		arg1 channel.Request
		arg2 []channel.RequestOption
	}
	queryReturns struct {
		result1 channel.Response
		result2 error
	}
	queryReturnsOnCall map[int]struct {
		result1 channel.Response
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *MockChannelClient) Execute(arg1 channel.Request, arg2 ...channel.RequestOption) (channel.Response, error) {
	fake.executeMutex.Lock()
	ret, specificReturn := fake.executeReturnsOnCall[len(fake.executeArgsForCall)]
	fake.executeArgsForCall = append(fake.executeArgsForCall, struct {
		arg1 channel.Request
		arg2 []channel.RequestOption
	}{arg1, arg2})
	fake.recordInvocation("Execute", []interface{}{arg1, arg2})
	fake.executeMutex.Unlock()
	if fake.ExecuteStub != nil {
		return fake.ExecuteStub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.executeReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *MockChannelClient) ExecuteCallCount() int {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return len(fake.executeArgsForCall)
}

func (fake *MockChannelClient) ExecuteCalls(stub func(channel.Request, ...channel.RequestOption) (channel.Response, error)) {
	fake.executeMutex.Lock()
	defer fake.executeMutex.Unlock()
	fake.ExecuteStub = stub
}

func (fake *MockChannelClient) ExecuteArgsForCall(i int) (channel.Request, []channel.RequestOption) {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	argsForCall := fake.executeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *MockChannelClient) ExecuteReturns(result1 channel.Response, result2 error) {
	fake.executeMutex.Lock()
	defer fake.executeMutex.Unlock()
	fake.ExecuteStub = nil
	fake.executeReturns = struct {
		result1 channel.Response
		result2 error
	}{result1, result2}
}

func (fake *MockChannelClient) ExecuteReturnsOnCall(i int, result1 channel.Response, result2 error) {
	fake.executeMutex.Lock()
	defer fake.executeMutex.Unlock()
	fake.ExecuteStub = nil
	if fake.executeReturnsOnCall == nil {
		fake.executeReturnsOnCall = make(map[int]struct {
			result1 channel.Response
			result2 error
		})
	}
	fake.executeReturnsOnCall[i] = struct {
		result1 channel.Response
		result2 error
	}{result1, result2}
}

func (fake *MockChannelClient) Query(arg1 channel.Request, arg2 ...channel.RequestOption) (channel.Response, error) {
	fake.queryMutex.Lock()
	ret, specificReturn := fake.queryReturnsOnCall[len(fake.queryArgsForCall)]
	fake.queryArgsForCall = append(fake.queryArgsForCall, struct {
		arg1 channel.Request
		arg2 []channel.RequestOption
	}{arg1, arg2})
	fake.recordInvocation("Query", []interface{}{arg1, arg2})
	fake.queryMutex.Unlock()
	if fake.QueryStub != nil {
		return fake.QueryStub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.queryReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *MockChannelClient) QueryCallCount() int {
	fake.queryMutex.RLock()
	defer fake.queryMutex.RUnlock()
	return len(fake.queryArgsForCall)
}

func (fake *MockChannelClient) QueryCalls(stub func(channel.Request, ...channel.RequestOption) (channel.Response, error)) {
	fake.queryMutex.Lock()
	defer fake.queryMutex.Unlock()
	fake.QueryStub = stub
}

func (fake *MockChannelClient) QueryArgsForCall(i int) (channel.Request, []channel.RequestOption) {
	fake.queryMutex.RLock()
	defer fake.queryMutex.RUnlock()
	argsForCall := fake.queryArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *MockChannelClient) QueryReturns(result1 channel.Response, result2 error) {
	fake.queryMutex.Lock()
	defer fake.queryMutex.Unlock()
	fake.QueryStub = nil
	fake.queryReturns = struct {
		result1 channel.Response
		result2 error
	}{result1, result2}
}

func (fake *MockChannelClient) QueryReturnsOnCall(i int, result1 channel.Response, result2 error) {
	fake.queryMutex.Lock()
	defer fake.queryMutex.Unlock()
	fake.QueryStub = nil
	if fake.queryReturnsOnCall == nil {
		fake.queryReturnsOnCall = make(map[int]struct {
			result1 channel.Response
			result2 error
		})
	}
	fake.queryReturnsOnCall[i] = struct {
		result1 channel.Response
		result2 error
	}{result1, result2}
}

func (fake *MockChannelClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	fake.queryMutex.RLock()
	defer fake.queryMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *MockChannelClient) recordInvocation(key string, args []interface{}) {
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

var _ fab3.ChannelClient = new(MockChannelClient)