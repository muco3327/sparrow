// Code generated by mockery v2.11.0. DO NOT EDIT.

package mocks

import (
	context "context"

	chain "github.com/palomachain/sparrow/chain"

	mock "github.com/stretchr/testify/mock"

	paloma "github.com/palomachain/sparrow/chain/paloma"

	testing "testing"

	types "github.com/palomachain/sparrow/types/paloma/x/valset/types"
)

// PalomaClienter is an autogenerated mock type for the PalomaClienter type
type PalomaClienter struct {
	mock.Mock
}

// AddExternalChainInfo provides a mock function with given fields: ctx, chainInfos
func (_m *PalomaClienter) AddExternalChainInfo(ctx context.Context, chainInfos ...paloma.ChainInfoIn) error {
	_va := make([]interface{}, len(chainInfos))
	for _i := range chainInfos {
		_va[_i] = chainInfos[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ...paloma.ChainInfoIn) error); ok {
		r0 = rf(ctx, chainInfos...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BroadcastMessageSignatures provides a mock function with given fields: ctx, signatures
func (_m *PalomaClienter) BroadcastMessageSignatures(ctx context.Context, signatures ...paloma.BroadcastMessageSignatureIn) error {
	_va := make([]interface{}, len(signatures))
	for _i := range signatures {
		_va[_i] = signatures[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ...paloma.BroadcastMessageSignatureIn) error); ok {
		r0 = rf(ctx, signatures...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QueryMessagesForSigning provides a mock function with given fields: ctx, queueTypeName
func (_m *PalomaClienter) QueryMessagesForSigning(ctx context.Context, queueTypeName string) ([]chain.QueuedMessage, error) {
	ret := _m.Called(ctx, queueTypeName)

	var r0 []chain.QueuedMessage
	if rf, ok := ret.Get(0).(func(context.Context, string) []chain.QueuedMessage); ok {
		r0 = rf(ctx, queueTypeName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]chain.QueuedMessage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, queueTypeName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryMessagesInQueue provides a mock function with given fields: ctx, queueTypeName
func (_m *PalomaClienter) QueryMessagesInQueue(ctx context.Context, queueTypeName string) ([]chain.MessageWithSignatures, error) {
	ret := _m.Called(ctx, queueTypeName)

	var r0 []chain.MessageWithSignatures
	if rf, ok := ret.Get(0).(func(context.Context, string) []chain.MessageWithSignatures); ok {
		r0 = rf(ctx, queueTypeName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]chain.MessageWithSignatures)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, queueTypeName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryValidatorInfo provides a mock function with given fields: ctx
func (_m *PalomaClienter) QueryValidatorInfo(ctx context.Context) ([]*types.ExternalChainInfo, error) {
	ret := _m.Called(ctx)

	var r0 []*types.ExternalChainInfo
	if rf, ok := ret.Get(0).(func(context.Context) []*types.ExternalChainInfo); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.ExternalChainInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPalomaClienter creates a new instance of PalomaClienter. It also registers a cleanup function to assert the mocks expectations.
func NewPalomaClienter(t testing.TB) *PalomaClienter {
	mock := &PalomaClienter{}

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
