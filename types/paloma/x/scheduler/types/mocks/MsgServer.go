// Code generated by mockery v2.11.0. DO NOT EDIT.

package mocks

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// MsgServer is an autogenerated mock type for the MsgServer type
type MsgServer struct {
	mock.Mock
}

// NewMsgServer creates a new instance of MsgServer. It also registers a cleanup function to assert the mocks expectations.
func NewMsgServer(t testing.TB) *MsgServer {
	mock := &MsgServer{}

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
