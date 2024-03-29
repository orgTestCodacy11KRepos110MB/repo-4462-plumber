// Code generated by counterfeiter. DO NOT EDIT.
package busfakes

import (
	"context"
	"sync"

	"github.com/batchcorp/plumber-schemas/build/go/protos/opts"
	"github.com/batchcorp/plumber/bus"
)

type FakeIBus struct {
	PublishCreateConnectionStub        func(context.Context, *opts.ConnectionOptions) error
	publishCreateConnectionMutex       sync.RWMutex
	publishCreateConnectionArgsForCall []struct {
		arg1 context.Context
		arg2 *opts.ConnectionOptions
	}
	publishCreateConnectionReturns struct {
		result1 error
	}
	publishCreateConnectionReturnsOnCall map[int]struct {
		result1 error
	}
	PublishCreateRelayStub        func(context.Context, *opts.RelayOptions) error
	publishCreateRelayMutex       sync.RWMutex
	publishCreateRelayArgsForCall []struct {
		arg1 context.Context
		arg2 *opts.RelayOptions
	}
	publishCreateRelayReturns struct {
		result1 error
	}
	publishCreateRelayReturnsOnCall map[int]struct {
		result1 error
	}
	PublishCreateTunnelStub        func(context.Context, *opts.TunnelOptions) error
	publishCreateTunnelMutex       sync.RWMutex
	publishCreateTunnelArgsForCall []struct {
		arg1 context.Context
		arg2 *opts.TunnelOptions
	}
	publishCreateTunnelReturns struct {
		result1 error
	}
	publishCreateTunnelReturnsOnCall map[int]struct {
		result1 error
	}
	PublishDeleteConnectionStub        func(context.Context, *opts.ConnectionOptions) error
	publishDeleteConnectionMutex       sync.RWMutex
	publishDeleteConnectionArgsForCall []struct {
		arg1 context.Context
		arg2 *opts.ConnectionOptions
	}
	publishDeleteConnectionReturns struct {
		result1 error
	}
	publishDeleteConnectionReturnsOnCall map[int]struct {
		result1 error
	}
	PublishDeleteRelayStub        func(context.Context, *opts.RelayOptions) error
	publishDeleteRelayMutex       sync.RWMutex
	publishDeleteRelayArgsForCall []struct {
		arg1 context.Context
		arg2 *opts.RelayOptions
	}
	publishDeleteRelayReturns struct {
		result1 error
	}
	publishDeleteRelayReturnsOnCall map[int]struct {
		result1 error
	}
	PublishDeleteTunnelStub        func(context.Context, *opts.TunnelOptions) error
	publishDeleteTunnelMutex       sync.RWMutex
	publishDeleteTunnelArgsForCall []struct {
		arg1 context.Context
		arg2 *opts.TunnelOptions
	}
	publishDeleteTunnelReturns struct {
		result1 error
	}
	publishDeleteTunnelReturnsOnCall map[int]struct {
		result1 error
	}
	PublishResumeRelayStub        func(context.Context, *opts.RelayOptions) error
	publishResumeRelayMutex       sync.RWMutex
	publishResumeRelayArgsForCall []struct {
		arg1 context.Context
		arg2 *opts.RelayOptions
	}
	publishResumeRelayReturns struct {
		result1 error
	}
	publishResumeRelayReturnsOnCall map[int]struct {
		result1 error
	}
	PublishResumeTunnelStub        func(context.Context, *opts.TunnelOptions) error
	publishResumeTunnelMutex       sync.RWMutex
	publishResumeTunnelArgsForCall []struct {
		arg1 context.Context
		arg2 *opts.TunnelOptions
	}
	publishResumeTunnelReturns struct {
		result1 error
	}
	publishResumeTunnelReturnsOnCall map[int]struct {
		result1 error
	}
	PublishStopRelayStub        func(context.Context, *opts.RelayOptions) error
	publishStopRelayMutex       sync.RWMutex
	publishStopRelayArgsForCall []struct {
		arg1 context.Context
		arg2 *opts.RelayOptions
	}
	publishStopRelayReturns struct {
		result1 error
	}
	publishStopRelayReturnsOnCall map[int]struct {
		result1 error
	}
	PublishStopTunnelStub        func(context.Context, *opts.TunnelOptions) error
	publishStopTunnelMutex       sync.RWMutex
	publishStopTunnelArgsForCall []struct {
		arg1 context.Context
		arg2 *opts.TunnelOptions
	}
	publishStopTunnelReturns struct {
		result1 error
	}
	publishStopTunnelReturnsOnCall map[int]struct {
		result1 error
	}
	PublishUpdateConnectionStub        func(context.Context, *opts.ConnectionOptions) error
	publishUpdateConnectionMutex       sync.RWMutex
	publishUpdateConnectionArgsForCall []struct {
		arg1 context.Context
		arg2 *opts.ConnectionOptions
	}
	publishUpdateConnectionReturns struct {
		result1 error
	}
	publishUpdateConnectionReturnsOnCall map[int]struct {
		result1 error
	}
	PublishUpdateRelayStub        func(context.Context, *opts.RelayOptions) error
	publishUpdateRelayMutex       sync.RWMutex
	publishUpdateRelayArgsForCall []struct {
		arg1 context.Context
		arg2 *opts.RelayOptions
	}
	publishUpdateRelayReturns struct {
		result1 error
	}
	publishUpdateRelayReturnsOnCall map[int]struct {
		result1 error
	}
	PublishUpdateTunnelStub        func(context.Context, *opts.TunnelOptions) error
	publishUpdateTunnelMutex       sync.RWMutex
	publishUpdateTunnelArgsForCall []struct {
		arg1 context.Context
		arg2 *opts.TunnelOptions
	}
	publishUpdateTunnelReturns struct {
		result1 error
	}
	publishUpdateTunnelReturnsOnCall map[int]struct {
		result1 error
	}
	StartStub        func(context.Context) error
	startMutex       sync.RWMutex
	startArgsForCall []struct {
		arg1 context.Context
	}
	startReturns struct {
		result1 error
	}
	startReturnsOnCall map[int]struct {
		result1 error
	}
	StopStub        func() error
	stopMutex       sync.RWMutex
	stopArgsForCall []struct {
	}
	stopReturns struct {
		result1 error
	}
	stopReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeIBus) PublishCreateConnection(arg1 context.Context, arg2 *opts.ConnectionOptions) error {
	fake.publishCreateConnectionMutex.Lock()
	ret, specificReturn := fake.publishCreateConnectionReturnsOnCall[len(fake.publishCreateConnectionArgsForCall)]
	fake.publishCreateConnectionArgsForCall = append(fake.publishCreateConnectionArgsForCall, struct {
		arg1 context.Context
		arg2 *opts.ConnectionOptions
	}{arg1, arg2})
	stub := fake.PublishCreateConnectionStub
	fakeReturns := fake.publishCreateConnectionReturns
	fake.recordInvocation("PublishCreateConnection", []interface{}{arg1, arg2})
	fake.publishCreateConnectionMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIBus) PublishCreateConnectionCallCount() int {
	fake.publishCreateConnectionMutex.RLock()
	defer fake.publishCreateConnectionMutex.RUnlock()
	return len(fake.publishCreateConnectionArgsForCall)
}

func (fake *FakeIBus) PublishCreateConnectionCalls(stub func(context.Context, *opts.ConnectionOptions) error) {
	fake.publishCreateConnectionMutex.Lock()
	defer fake.publishCreateConnectionMutex.Unlock()
	fake.PublishCreateConnectionStub = stub
}

func (fake *FakeIBus) PublishCreateConnectionArgsForCall(i int) (context.Context, *opts.ConnectionOptions) {
	fake.publishCreateConnectionMutex.RLock()
	defer fake.publishCreateConnectionMutex.RUnlock()
	argsForCall := fake.publishCreateConnectionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeIBus) PublishCreateConnectionReturns(result1 error) {
	fake.publishCreateConnectionMutex.Lock()
	defer fake.publishCreateConnectionMutex.Unlock()
	fake.PublishCreateConnectionStub = nil
	fake.publishCreateConnectionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishCreateConnectionReturnsOnCall(i int, result1 error) {
	fake.publishCreateConnectionMutex.Lock()
	defer fake.publishCreateConnectionMutex.Unlock()
	fake.PublishCreateConnectionStub = nil
	if fake.publishCreateConnectionReturnsOnCall == nil {
		fake.publishCreateConnectionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.publishCreateConnectionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishCreateRelay(arg1 context.Context, arg2 *opts.RelayOptions) error {
	fake.publishCreateRelayMutex.Lock()
	ret, specificReturn := fake.publishCreateRelayReturnsOnCall[len(fake.publishCreateRelayArgsForCall)]
	fake.publishCreateRelayArgsForCall = append(fake.publishCreateRelayArgsForCall, struct {
		arg1 context.Context
		arg2 *opts.RelayOptions
	}{arg1, arg2})
	stub := fake.PublishCreateRelayStub
	fakeReturns := fake.publishCreateRelayReturns
	fake.recordInvocation("PublishCreateRelay", []interface{}{arg1, arg2})
	fake.publishCreateRelayMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIBus) PublishCreateRelayCallCount() int {
	fake.publishCreateRelayMutex.RLock()
	defer fake.publishCreateRelayMutex.RUnlock()
	return len(fake.publishCreateRelayArgsForCall)
}

func (fake *FakeIBus) PublishCreateRelayCalls(stub func(context.Context, *opts.RelayOptions) error) {
	fake.publishCreateRelayMutex.Lock()
	defer fake.publishCreateRelayMutex.Unlock()
	fake.PublishCreateRelayStub = stub
}

func (fake *FakeIBus) PublishCreateRelayArgsForCall(i int) (context.Context, *opts.RelayOptions) {
	fake.publishCreateRelayMutex.RLock()
	defer fake.publishCreateRelayMutex.RUnlock()
	argsForCall := fake.publishCreateRelayArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeIBus) PublishCreateRelayReturns(result1 error) {
	fake.publishCreateRelayMutex.Lock()
	defer fake.publishCreateRelayMutex.Unlock()
	fake.PublishCreateRelayStub = nil
	fake.publishCreateRelayReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishCreateRelayReturnsOnCall(i int, result1 error) {
	fake.publishCreateRelayMutex.Lock()
	defer fake.publishCreateRelayMutex.Unlock()
	fake.PublishCreateRelayStub = nil
	if fake.publishCreateRelayReturnsOnCall == nil {
		fake.publishCreateRelayReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.publishCreateRelayReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishCreateTunnel(arg1 context.Context, arg2 *opts.TunnelOptions) error {
	fake.publishCreateTunnelMutex.Lock()
	ret, specificReturn := fake.publishCreateTunnelReturnsOnCall[len(fake.publishCreateTunnelArgsForCall)]
	fake.publishCreateTunnelArgsForCall = append(fake.publishCreateTunnelArgsForCall, struct {
		arg1 context.Context
		arg2 *opts.TunnelOptions
	}{arg1, arg2})
	stub := fake.PublishCreateTunnelStub
	fakeReturns := fake.publishCreateTunnelReturns
	fake.recordInvocation("PublishCreateTunnel", []interface{}{arg1, arg2})
	fake.publishCreateTunnelMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIBus) PublishCreateTunnelCallCount() int {
	fake.publishCreateTunnelMutex.RLock()
	defer fake.publishCreateTunnelMutex.RUnlock()
	return len(fake.publishCreateTunnelArgsForCall)
}

func (fake *FakeIBus) PublishCreateTunnelCalls(stub func(context.Context, *opts.TunnelOptions) error) {
	fake.publishCreateTunnelMutex.Lock()
	defer fake.publishCreateTunnelMutex.Unlock()
	fake.PublishCreateTunnelStub = stub
}

func (fake *FakeIBus) PublishCreateTunnelArgsForCall(i int) (context.Context, *opts.TunnelOptions) {
	fake.publishCreateTunnelMutex.RLock()
	defer fake.publishCreateTunnelMutex.RUnlock()
	argsForCall := fake.publishCreateTunnelArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeIBus) PublishCreateTunnelReturns(result1 error) {
	fake.publishCreateTunnelMutex.Lock()
	defer fake.publishCreateTunnelMutex.Unlock()
	fake.PublishCreateTunnelStub = nil
	fake.publishCreateTunnelReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishCreateTunnelReturnsOnCall(i int, result1 error) {
	fake.publishCreateTunnelMutex.Lock()
	defer fake.publishCreateTunnelMutex.Unlock()
	fake.PublishCreateTunnelStub = nil
	if fake.publishCreateTunnelReturnsOnCall == nil {
		fake.publishCreateTunnelReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.publishCreateTunnelReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishDeleteConnection(arg1 context.Context, arg2 *opts.ConnectionOptions) error {
	fake.publishDeleteConnectionMutex.Lock()
	ret, specificReturn := fake.publishDeleteConnectionReturnsOnCall[len(fake.publishDeleteConnectionArgsForCall)]
	fake.publishDeleteConnectionArgsForCall = append(fake.publishDeleteConnectionArgsForCall, struct {
		arg1 context.Context
		arg2 *opts.ConnectionOptions
	}{arg1, arg2})
	stub := fake.PublishDeleteConnectionStub
	fakeReturns := fake.publishDeleteConnectionReturns
	fake.recordInvocation("PublishDeleteConnection", []interface{}{arg1, arg2})
	fake.publishDeleteConnectionMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIBus) PublishDeleteConnectionCallCount() int {
	fake.publishDeleteConnectionMutex.RLock()
	defer fake.publishDeleteConnectionMutex.RUnlock()
	return len(fake.publishDeleteConnectionArgsForCall)
}

func (fake *FakeIBus) PublishDeleteConnectionCalls(stub func(context.Context, *opts.ConnectionOptions) error) {
	fake.publishDeleteConnectionMutex.Lock()
	defer fake.publishDeleteConnectionMutex.Unlock()
	fake.PublishDeleteConnectionStub = stub
}

func (fake *FakeIBus) PublishDeleteConnectionArgsForCall(i int) (context.Context, *opts.ConnectionOptions) {
	fake.publishDeleteConnectionMutex.RLock()
	defer fake.publishDeleteConnectionMutex.RUnlock()
	argsForCall := fake.publishDeleteConnectionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeIBus) PublishDeleteConnectionReturns(result1 error) {
	fake.publishDeleteConnectionMutex.Lock()
	defer fake.publishDeleteConnectionMutex.Unlock()
	fake.PublishDeleteConnectionStub = nil
	fake.publishDeleteConnectionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishDeleteConnectionReturnsOnCall(i int, result1 error) {
	fake.publishDeleteConnectionMutex.Lock()
	defer fake.publishDeleteConnectionMutex.Unlock()
	fake.PublishDeleteConnectionStub = nil
	if fake.publishDeleteConnectionReturnsOnCall == nil {
		fake.publishDeleteConnectionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.publishDeleteConnectionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishDeleteRelay(arg1 context.Context, arg2 *opts.RelayOptions) error {
	fake.publishDeleteRelayMutex.Lock()
	ret, specificReturn := fake.publishDeleteRelayReturnsOnCall[len(fake.publishDeleteRelayArgsForCall)]
	fake.publishDeleteRelayArgsForCall = append(fake.publishDeleteRelayArgsForCall, struct {
		arg1 context.Context
		arg2 *opts.RelayOptions
	}{arg1, arg2})
	stub := fake.PublishDeleteRelayStub
	fakeReturns := fake.publishDeleteRelayReturns
	fake.recordInvocation("PublishDeleteRelay", []interface{}{arg1, arg2})
	fake.publishDeleteRelayMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIBus) PublishDeleteRelayCallCount() int {
	fake.publishDeleteRelayMutex.RLock()
	defer fake.publishDeleteRelayMutex.RUnlock()
	return len(fake.publishDeleteRelayArgsForCall)
}

func (fake *FakeIBus) PublishDeleteRelayCalls(stub func(context.Context, *opts.RelayOptions) error) {
	fake.publishDeleteRelayMutex.Lock()
	defer fake.publishDeleteRelayMutex.Unlock()
	fake.PublishDeleteRelayStub = stub
}

func (fake *FakeIBus) PublishDeleteRelayArgsForCall(i int) (context.Context, *opts.RelayOptions) {
	fake.publishDeleteRelayMutex.RLock()
	defer fake.publishDeleteRelayMutex.RUnlock()
	argsForCall := fake.publishDeleteRelayArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeIBus) PublishDeleteRelayReturns(result1 error) {
	fake.publishDeleteRelayMutex.Lock()
	defer fake.publishDeleteRelayMutex.Unlock()
	fake.PublishDeleteRelayStub = nil
	fake.publishDeleteRelayReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishDeleteRelayReturnsOnCall(i int, result1 error) {
	fake.publishDeleteRelayMutex.Lock()
	defer fake.publishDeleteRelayMutex.Unlock()
	fake.PublishDeleteRelayStub = nil
	if fake.publishDeleteRelayReturnsOnCall == nil {
		fake.publishDeleteRelayReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.publishDeleteRelayReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishDeleteTunnel(arg1 context.Context, arg2 *opts.TunnelOptions) error {
	fake.publishDeleteTunnelMutex.Lock()
	ret, specificReturn := fake.publishDeleteTunnelReturnsOnCall[len(fake.publishDeleteTunnelArgsForCall)]
	fake.publishDeleteTunnelArgsForCall = append(fake.publishDeleteTunnelArgsForCall, struct {
		arg1 context.Context
		arg2 *opts.TunnelOptions
	}{arg1, arg2})
	stub := fake.PublishDeleteTunnelStub
	fakeReturns := fake.publishDeleteTunnelReturns
	fake.recordInvocation("PublishDeleteTunnel", []interface{}{arg1, arg2})
	fake.publishDeleteTunnelMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIBus) PublishDeleteTunnelCallCount() int {
	fake.publishDeleteTunnelMutex.RLock()
	defer fake.publishDeleteTunnelMutex.RUnlock()
	return len(fake.publishDeleteTunnelArgsForCall)
}

func (fake *FakeIBus) PublishDeleteTunnelCalls(stub func(context.Context, *opts.TunnelOptions) error) {
	fake.publishDeleteTunnelMutex.Lock()
	defer fake.publishDeleteTunnelMutex.Unlock()
	fake.PublishDeleteTunnelStub = stub
}

func (fake *FakeIBus) PublishDeleteTunnelArgsForCall(i int) (context.Context, *opts.TunnelOptions) {
	fake.publishDeleteTunnelMutex.RLock()
	defer fake.publishDeleteTunnelMutex.RUnlock()
	argsForCall := fake.publishDeleteTunnelArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeIBus) PublishDeleteTunnelReturns(result1 error) {
	fake.publishDeleteTunnelMutex.Lock()
	defer fake.publishDeleteTunnelMutex.Unlock()
	fake.PublishDeleteTunnelStub = nil
	fake.publishDeleteTunnelReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishDeleteTunnelReturnsOnCall(i int, result1 error) {
	fake.publishDeleteTunnelMutex.Lock()
	defer fake.publishDeleteTunnelMutex.Unlock()
	fake.PublishDeleteTunnelStub = nil
	if fake.publishDeleteTunnelReturnsOnCall == nil {
		fake.publishDeleteTunnelReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.publishDeleteTunnelReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishResumeRelay(arg1 context.Context, arg2 *opts.RelayOptions) error {
	fake.publishResumeRelayMutex.Lock()
	ret, specificReturn := fake.publishResumeRelayReturnsOnCall[len(fake.publishResumeRelayArgsForCall)]
	fake.publishResumeRelayArgsForCall = append(fake.publishResumeRelayArgsForCall, struct {
		arg1 context.Context
		arg2 *opts.RelayOptions
	}{arg1, arg2})
	stub := fake.PublishResumeRelayStub
	fakeReturns := fake.publishResumeRelayReturns
	fake.recordInvocation("PublishResumeRelay", []interface{}{arg1, arg2})
	fake.publishResumeRelayMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIBus) PublishResumeRelayCallCount() int {
	fake.publishResumeRelayMutex.RLock()
	defer fake.publishResumeRelayMutex.RUnlock()
	return len(fake.publishResumeRelayArgsForCall)
}

func (fake *FakeIBus) PublishResumeRelayCalls(stub func(context.Context, *opts.RelayOptions) error) {
	fake.publishResumeRelayMutex.Lock()
	defer fake.publishResumeRelayMutex.Unlock()
	fake.PublishResumeRelayStub = stub
}

func (fake *FakeIBus) PublishResumeRelayArgsForCall(i int) (context.Context, *opts.RelayOptions) {
	fake.publishResumeRelayMutex.RLock()
	defer fake.publishResumeRelayMutex.RUnlock()
	argsForCall := fake.publishResumeRelayArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeIBus) PublishResumeRelayReturns(result1 error) {
	fake.publishResumeRelayMutex.Lock()
	defer fake.publishResumeRelayMutex.Unlock()
	fake.PublishResumeRelayStub = nil
	fake.publishResumeRelayReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishResumeRelayReturnsOnCall(i int, result1 error) {
	fake.publishResumeRelayMutex.Lock()
	defer fake.publishResumeRelayMutex.Unlock()
	fake.PublishResumeRelayStub = nil
	if fake.publishResumeRelayReturnsOnCall == nil {
		fake.publishResumeRelayReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.publishResumeRelayReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishResumeTunnel(arg1 context.Context, arg2 *opts.TunnelOptions) error {
	fake.publishResumeTunnelMutex.Lock()
	ret, specificReturn := fake.publishResumeTunnelReturnsOnCall[len(fake.publishResumeTunnelArgsForCall)]
	fake.publishResumeTunnelArgsForCall = append(fake.publishResumeTunnelArgsForCall, struct {
		arg1 context.Context
		arg2 *opts.TunnelOptions
	}{arg1, arg2})
	stub := fake.PublishResumeTunnelStub
	fakeReturns := fake.publishResumeTunnelReturns
	fake.recordInvocation("PublishResumeTunnel", []interface{}{arg1, arg2})
	fake.publishResumeTunnelMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIBus) PublishResumeTunnelCallCount() int {
	fake.publishResumeTunnelMutex.RLock()
	defer fake.publishResumeTunnelMutex.RUnlock()
	return len(fake.publishResumeTunnelArgsForCall)
}

func (fake *FakeIBus) PublishResumeTunnelCalls(stub func(context.Context, *opts.TunnelOptions) error) {
	fake.publishResumeTunnelMutex.Lock()
	defer fake.publishResumeTunnelMutex.Unlock()
	fake.PublishResumeTunnelStub = stub
}

func (fake *FakeIBus) PublishResumeTunnelArgsForCall(i int) (context.Context, *opts.TunnelOptions) {
	fake.publishResumeTunnelMutex.RLock()
	defer fake.publishResumeTunnelMutex.RUnlock()
	argsForCall := fake.publishResumeTunnelArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeIBus) PublishResumeTunnelReturns(result1 error) {
	fake.publishResumeTunnelMutex.Lock()
	defer fake.publishResumeTunnelMutex.Unlock()
	fake.PublishResumeTunnelStub = nil
	fake.publishResumeTunnelReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishResumeTunnelReturnsOnCall(i int, result1 error) {
	fake.publishResumeTunnelMutex.Lock()
	defer fake.publishResumeTunnelMutex.Unlock()
	fake.PublishResumeTunnelStub = nil
	if fake.publishResumeTunnelReturnsOnCall == nil {
		fake.publishResumeTunnelReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.publishResumeTunnelReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishStopRelay(arg1 context.Context, arg2 *opts.RelayOptions) error {
	fake.publishStopRelayMutex.Lock()
	ret, specificReturn := fake.publishStopRelayReturnsOnCall[len(fake.publishStopRelayArgsForCall)]
	fake.publishStopRelayArgsForCall = append(fake.publishStopRelayArgsForCall, struct {
		arg1 context.Context
		arg2 *opts.RelayOptions
	}{arg1, arg2})
	stub := fake.PublishStopRelayStub
	fakeReturns := fake.publishStopRelayReturns
	fake.recordInvocation("PublishStopRelay", []interface{}{arg1, arg2})
	fake.publishStopRelayMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIBus) PublishStopRelayCallCount() int {
	fake.publishStopRelayMutex.RLock()
	defer fake.publishStopRelayMutex.RUnlock()
	return len(fake.publishStopRelayArgsForCall)
}

func (fake *FakeIBus) PublishStopRelayCalls(stub func(context.Context, *opts.RelayOptions) error) {
	fake.publishStopRelayMutex.Lock()
	defer fake.publishStopRelayMutex.Unlock()
	fake.PublishStopRelayStub = stub
}

func (fake *FakeIBus) PublishStopRelayArgsForCall(i int) (context.Context, *opts.RelayOptions) {
	fake.publishStopRelayMutex.RLock()
	defer fake.publishStopRelayMutex.RUnlock()
	argsForCall := fake.publishStopRelayArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeIBus) PublishStopRelayReturns(result1 error) {
	fake.publishStopRelayMutex.Lock()
	defer fake.publishStopRelayMutex.Unlock()
	fake.PublishStopRelayStub = nil
	fake.publishStopRelayReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishStopRelayReturnsOnCall(i int, result1 error) {
	fake.publishStopRelayMutex.Lock()
	defer fake.publishStopRelayMutex.Unlock()
	fake.PublishStopRelayStub = nil
	if fake.publishStopRelayReturnsOnCall == nil {
		fake.publishStopRelayReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.publishStopRelayReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishStopTunnel(arg1 context.Context, arg2 *opts.TunnelOptions) error {
	fake.publishStopTunnelMutex.Lock()
	ret, specificReturn := fake.publishStopTunnelReturnsOnCall[len(fake.publishStopTunnelArgsForCall)]
	fake.publishStopTunnelArgsForCall = append(fake.publishStopTunnelArgsForCall, struct {
		arg1 context.Context
		arg2 *opts.TunnelOptions
	}{arg1, arg2})
	stub := fake.PublishStopTunnelStub
	fakeReturns := fake.publishStopTunnelReturns
	fake.recordInvocation("PublishStopTunnel", []interface{}{arg1, arg2})
	fake.publishStopTunnelMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIBus) PublishStopTunnelCallCount() int {
	fake.publishStopTunnelMutex.RLock()
	defer fake.publishStopTunnelMutex.RUnlock()
	return len(fake.publishStopTunnelArgsForCall)
}

func (fake *FakeIBus) PublishStopTunnelCalls(stub func(context.Context, *opts.TunnelOptions) error) {
	fake.publishStopTunnelMutex.Lock()
	defer fake.publishStopTunnelMutex.Unlock()
	fake.PublishStopTunnelStub = stub
}

func (fake *FakeIBus) PublishStopTunnelArgsForCall(i int) (context.Context, *opts.TunnelOptions) {
	fake.publishStopTunnelMutex.RLock()
	defer fake.publishStopTunnelMutex.RUnlock()
	argsForCall := fake.publishStopTunnelArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeIBus) PublishStopTunnelReturns(result1 error) {
	fake.publishStopTunnelMutex.Lock()
	defer fake.publishStopTunnelMutex.Unlock()
	fake.PublishStopTunnelStub = nil
	fake.publishStopTunnelReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishStopTunnelReturnsOnCall(i int, result1 error) {
	fake.publishStopTunnelMutex.Lock()
	defer fake.publishStopTunnelMutex.Unlock()
	fake.PublishStopTunnelStub = nil
	if fake.publishStopTunnelReturnsOnCall == nil {
		fake.publishStopTunnelReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.publishStopTunnelReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishUpdateConnection(arg1 context.Context, arg2 *opts.ConnectionOptions) error {
	fake.publishUpdateConnectionMutex.Lock()
	ret, specificReturn := fake.publishUpdateConnectionReturnsOnCall[len(fake.publishUpdateConnectionArgsForCall)]
	fake.publishUpdateConnectionArgsForCall = append(fake.publishUpdateConnectionArgsForCall, struct {
		arg1 context.Context
		arg2 *opts.ConnectionOptions
	}{arg1, arg2})
	stub := fake.PublishUpdateConnectionStub
	fakeReturns := fake.publishUpdateConnectionReturns
	fake.recordInvocation("PublishUpdateConnection", []interface{}{arg1, arg2})
	fake.publishUpdateConnectionMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIBus) PublishUpdateConnectionCallCount() int {
	fake.publishUpdateConnectionMutex.RLock()
	defer fake.publishUpdateConnectionMutex.RUnlock()
	return len(fake.publishUpdateConnectionArgsForCall)
}

func (fake *FakeIBus) PublishUpdateConnectionCalls(stub func(context.Context, *opts.ConnectionOptions) error) {
	fake.publishUpdateConnectionMutex.Lock()
	defer fake.publishUpdateConnectionMutex.Unlock()
	fake.PublishUpdateConnectionStub = stub
}

func (fake *FakeIBus) PublishUpdateConnectionArgsForCall(i int) (context.Context, *opts.ConnectionOptions) {
	fake.publishUpdateConnectionMutex.RLock()
	defer fake.publishUpdateConnectionMutex.RUnlock()
	argsForCall := fake.publishUpdateConnectionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeIBus) PublishUpdateConnectionReturns(result1 error) {
	fake.publishUpdateConnectionMutex.Lock()
	defer fake.publishUpdateConnectionMutex.Unlock()
	fake.PublishUpdateConnectionStub = nil
	fake.publishUpdateConnectionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishUpdateConnectionReturnsOnCall(i int, result1 error) {
	fake.publishUpdateConnectionMutex.Lock()
	defer fake.publishUpdateConnectionMutex.Unlock()
	fake.PublishUpdateConnectionStub = nil
	if fake.publishUpdateConnectionReturnsOnCall == nil {
		fake.publishUpdateConnectionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.publishUpdateConnectionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishUpdateRelay(arg1 context.Context, arg2 *opts.RelayOptions) error {
	fake.publishUpdateRelayMutex.Lock()
	ret, specificReturn := fake.publishUpdateRelayReturnsOnCall[len(fake.publishUpdateRelayArgsForCall)]
	fake.publishUpdateRelayArgsForCall = append(fake.publishUpdateRelayArgsForCall, struct {
		arg1 context.Context
		arg2 *opts.RelayOptions
	}{arg1, arg2})
	stub := fake.PublishUpdateRelayStub
	fakeReturns := fake.publishUpdateRelayReturns
	fake.recordInvocation("PublishUpdateRelay", []interface{}{arg1, arg2})
	fake.publishUpdateRelayMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIBus) PublishUpdateRelayCallCount() int {
	fake.publishUpdateRelayMutex.RLock()
	defer fake.publishUpdateRelayMutex.RUnlock()
	return len(fake.publishUpdateRelayArgsForCall)
}

func (fake *FakeIBus) PublishUpdateRelayCalls(stub func(context.Context, *opts.RelayOptions) error) {
	fake.publishUpdateRelayMutex.Lock()
	defer fake.publishUpdateRelayMutex.Unlock()
	fake.PublishUpdateRelayStub = stub
}

func (fake *FakeIBus) PublishUpdateRelayArgsForCall(i int) (context.Context, *opts.RelayOptions) {
	fake.publishUpdateRelayMutex.RLock()
	defer fake.publishUpdateRelayMutex.RUnlock()
	argsForCall := fake.publishUpdateRelayArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeIBus) PublishUpdateRelayReturns(result1 error) {
	fake.publishUpdateRelayMutex.Lock()
	defer fake.publishUpdateRelayMutex.Unlock()
	fake.PublishUpdateRelayStub = nil
	fake.publishUpdateRelayReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishUpdateRelayReturnsOnCall(i int, result1 error) {
	fake.publishUpdateRelayMutex.Lock()
	defer fake.publishUpdateRelayMutex.Unlock()
	fake.PublishUpdateRelayStub = nil
	if fake.publishUpdateRelayReturnsOnCall == nil {
		fake.publishUpdateRelayReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.publishUpdateRelayReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishUpdateTunnel(arg1 context.Context, arg2 *opts.TunnelOptions) error {
	fake.publishUpdateTunnelMutex.Lock()
	ret, specificReturn := fake.publishUpdateTunnelReturnsOnCall[len(fake.publishUpdateTunnelArgsForCall)]
	fake.publishUpdateTunnelArgsForCall = append(fake.publishUpdateTunnelArgsForCall, struct {
		arg1 context.Context
		arg2 *opts.TunnelOptions
	}{arg1, arg2})
	stub := fake.PublishUpdateTunnelStub
	fakeReturns := fake.publishUpdateTunnelReturns
	fake.recordInvocation("PublishUpdateTunnel", []interface{}{arg1, arg2})
	fake.publishUpdateTunnelMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIBus) PublishUpdateTunnelCallCount() int {
	fake.publishUpdateTunnelMutex.RLock()
	defer fake.publishUpdateTunnelMutex.RUnlock()
	return len(fake.publishUpdateTunnelArgsForCall)
}

func (fake *FakeIBus) PublishUpdateTunnelCalls(stub func(context.Context, *opts.TunnelOptions) error) {
	fake.publishUpdateTunnelMutex.Lock()
	defer fake.publishUpdateTunnelMutex.Unlock()
	fake.PublishUpdateTunnelStub = stub
}

func (fake *FakeIBus) PublishUpdateTunnelArgsForCall(i int) (context.Context, *opts.TunnelOptions) {
	fake.publishUpdateTunnelMutex.RLock()
	defer fake.publishUpdateTunnelMutex.RUnlock()
	argsForCall := fake.publishUpdateTunnelArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeIBus) PublishUpdateTunnelReturns(result1 error) {
	fake.publishUpdateTunnelMutex.Lock()
	defer fake.publishUpdateTunnelMutex.Unlock()
	fake.PublishUpdateTunnelStub = nil
	fake.publishUpdateTunnelReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishUpdateTunnelReturnsOnCall(i int, result1 error) {
	fake.publishUpdateTunnelMutex.Lock()
	defer fake.publishUpdateTunnelMutex.Unlock()
	fake.PublishUpdateTunnelStub = nil
	if fake.publishUpdateTunnelReturnsOnCall == nil {
		fake.publishUpdateTunnelReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.publishUpdateTunnelReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) Start(arg1 context.Context) error {
	fake.startMutex.Lock()
	ret, specificReturn := fake.startReturnsOnCall[len(fake.startArgsForCall)]
	fake.startArgsForCall = append(fake.startArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.StartStub
	fakeReturns := fake.startReturns
	fake.recordInvocation("Start", []interface{}{arg1})
	fake.startMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIBus) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *FakeIBus) StartCalls(stub func(context.Context) error) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = stub
}

func (fake *FakeIBus) StartArgsForCall(i int) context.Context {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	argsForCall := fake.startArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeIBus) StartReturns(result1 error) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = nil
	fake.startReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) StartReturnsOnCall(i int, result1 error) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = nil
	if fake.startReturnsOnCall == nil {
		fake.startReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.startReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) Stop() error {
	fake.stopMutex.Lock()
	ret, specificReturn := fake.stopReturnsOnCall[len(fake.stopArgsForCall)]
	fake.stopArgsForCall = append(fake.stopArgsForCall, struct {
	}{})
	stub := fake.StopStub
	fakeReturns := fake.stopReturns
	fake.recordInvocation("Stop", []interface{}{})
	fake.stopMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIBus) StopCallCount() int {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return len(fake.stopArgsForCall)
}

func (fake *FakeIBus) StopCalls(stub func() error) {
	fake.stopMutex.Lock()
	defer fake.stopMutex.Unlock()
	fake.StopStub = stub
}

func (fake *FakeIBus) StopReturns(result1 error) {
	fake.stopMutex.Lock()
	defer fake.stopMutex.Unlock()
	fake.StopStub = nil
	fake.stopReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) StopReturnsOnCall(i int, result1 error) {
	fake.stopMutex.Lock()
	defer fake.stopMutex.Unlock()
	fake.StopStub = nil
	if fake.stopReturnsOnCall == nil {
		fake.stopReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.stopReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.publishCreateConnectionMutex.RLock()
	defer fake.publishCreateConnectionMutex.RUnlock()
	fake.publishCreateRelayMutex.RLock()
	defer fake.publishCreateRelayMutex.RUnlock()
	fake.publishCreateTunnelMutex.RLock()
	defer fake.publishCreateTunnelMutex.RUnlock()
	fake.publishDeleteConnectionMutex.RLock()
	defer fake.publishDeleteConnectionMutex.RUnlock()
	fake.publishDeleteRelayMutex.RLock()
	defer fake.publishDeleteRelayMutex.RUnlock()
	fake.publishDeleteTunnelMutex.RLock()
	defer fake.publishDeleteTunnelMutex.RUnlock()
	fake.publishResumeRelayMutex.RLock()
	defer fake.publishResumeRelayMutex.RUnlock()
	fake.publishResumeTunnelMutex.RLock()
	defer fake.publishResumeTunnelMutex.RUnlock()
	fake.publishStopRelayMutex.RLock()
	defer fake.publishStopRelayMutex.RUnlock()
	fake.publishStopTunnelMutex.RLock()
	defer fake.publishStopTunnelMutex.RUnlock()
	fake.publishUpdateConnectionMutex.RLock()
	defer fake.publishUpdateConnectionMutex.RUnlock()
	fake.publishUpdateRelayMutex.RLock()
	defer fake.publishUpdateRelayMutex.RUnlock()
	fake.publishUpdateTunnelMutex.RLock()
	defer fake.publishUpdateTunnelMutex.RUnlock()
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeIBus) recordInvocation(key string, args []interface{}) {
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

var _ bus.IBus = new(FakeIBus)
