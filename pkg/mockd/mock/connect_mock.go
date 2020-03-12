// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/kt/connect/types.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	util "github.com/alibaba/kt-connect/pkg/kt/util"
	gomock "github.com/golang/mock/gomock"
)

// MockShadowInterface is a mock of ShadowInterface interface
type MockShadowInterface struct {
	ctrl     *gomock.Controller
	recorder *MockShadowInterfaceMockRecorder
}

// MockShadowInterfaceMockRecorder is the mock recorder for MockShadowInterface
type MockShadowInterfaceMockRecorder struct {
	mock *MockShadowInterface
}

// NewMockShadowInterface creates a new mock instance
func NewMockShadowInterface(ctrl *gomock.Controller) *MockShadowInterface {
	mock := &MockShadowInterface{ctrl: ctrl}
	mock.recorder = &MockShadowInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockShadowInterface) EXPECT() *MockShadowInterfaceMockRecorder {
	return m.recorder
}

// Inbound mocks base method
func (m *MockShadowInterface) Inbound(exposePort, podName, remoteIP string, credential *util.SSHCredential) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Inbound", exposePort, podName, remoteIP, credential)
	ret0, _ := ret[0].(error)
	return ret0
}

// Inbound indicates an expected call of Inbound
func (mr *MockShadowInterfaceMockRecorder) Inbound(exposePort, podName, remoteIP, credential interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Inbound", reflect.TypeOf((*MockShadowInterface)(nil).Inbound), exposePort, podName, remoteIP, credential)
}

// Outbound mocks base method
func (m *MockShadowInterface) Outbound(name, podIP string, credential *util.SSHCredential, cidrs []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Outbound", name, podIP, credential, cidrs)
	ret0, _ := ret[0].(error)
	return ret0
}

// Outbound indicates an expected call of Outbound
func (mr *MockShadowInterfaceMockRecorder) Outbound(name, podIP, credential, cidrs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Outbound", reflect.TypeOf((*MockShadowInterface)(nil).Outbound), name, podIP, credential, cidrs)
}