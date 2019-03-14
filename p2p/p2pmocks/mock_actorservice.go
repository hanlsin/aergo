// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aergoio/aergo/p2p/p2pcommon (interfaces: ActorService)

// Package mock_p2pcommon is a generated GoMock package.
package p2pmocks

import (
	reflect "reflect"
	time "time"

	actor "github.com/aergoio/aergo-actor/actor"
	types "github.com/aergoio/aergo/types"
	gomock "github.com/golang/mock/gomock"
)

// MockActorService is a mock of ActorService interface
type MockActorService struct {
	ctrl     *gomock.Controller
	recorder *MockActorServiceMockRecorder
}

// MockActorServiceMockRecorder is the mock recorder for MockActorService
type MockActorServiceMockRecorder struct {
	mock *MockActorService
}

// NewMockActorService creates a new mock instance
func NewMockActorService(ctrl *gomock.Controller) *MockActorService {
	mock := &MockActorService{ctrl: ctrl}
	mock.recorder = &MockActorServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockActorService) EXPECT() *MockActorServiceMockRecorder {
	return m.recorder
}

// CallRequest mocks base method
func (m *MockActorService) CallRequest(arg0 string, arg1 interface{}, arg2 time.Duration) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallRequest", arg0, arg1, arg2)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CallRequest indicates an expected call of CallRequest
func (mr *MockActorServiceMockRecorder) CallRequest(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallRequest", reflect.TypeOf((*MockActorService)(nil).CallRequest), arg0, arg1, arg2)
}

// CallRequestDefaultTimeout mocks base method
func (m *MockActorService) CallRequestDefaultTimeout(arg0 string, arg1 interface{}) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallRequestDefaultTimeout", arg0, arg1)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CallRequestDefaultTimeout indicates an expected call of CallRequestDefaultTimeout
func (mr *MockActorServiceMockRecorder) CallRequestDefaultTimeout(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallRequestDefaultTimeout", reflect.TypeOf((*MockActorService)(nil).CallRequestDefaultTimeout), arg0, arg1)
}

// FutureRequest mocks base method
func (m *MockActorService) FutureRequest(arg0 string, arg1 interface{}, arg2 time.Duration) *actor.Future {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FutureRequest", arg0, arg1, arg2)
	ret0, _ := ret[0].(*actor.Future)
	return ret0
}

// FutureRequest indicates an expected call of FutureRequest
func (mr *MockActorServiceMockRecorder) FutureRequest(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FutureRequest", reflect.TypeOf((*MockActorService)(nil).FutureRequest), arg0, arg1, arg2)
}

// FutureRequestDefaultTimeout mocks base method
func (m *MockActorService) FutureRequestDefaultTimeout(arg0 string, arg1 interface{}) *actor.Future {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FutureRequestDefaultTimeout", arg0, arg1)
	ret0, _ := ret[0].(*actor.Future)
	return ret0
}

// FutureRequestDefaultTimeout indicates an expected call of FutureRequestDefaultTimeout
func (mr *MockActorServiceMockRecorder) FutureRequestDefaultTimeout(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FutureRequestDefaultTimeout", reflect.TypeOf((*MockActorService)(nil).FutureRequestDefaultTimeout), arg0, arg1)
}

// GetChainAccessor mocks base method
func (m *MockActorService) GetChainAccessor() types.ChainAccessor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChainAccessor")
	ret0, _ := ret[0].(types.ChainAccessor)
	return ret0
}

// GetChainAccessor indicates an expected call of GetChainAccessor
func (mr *MockActorServiceMockRecorder) GetChainAccessor() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChainAccessor", reflect.TypeOf((*MockActorService)(nil).GetChainAccessor))
}

// SendRequest mocks base method
func (m *MockActorService) SendRequest(arg0 string, arg1 interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendRequest", arg0, arg1)
}

// SendRequest indicates an expected call of SendRequest
func (mr *MockActorServiceMockRecorder) SendRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendRequest", reflect.TypeOf((*MockActorService)(nil).SendRequest), arg0, arg1)
}

// TellRequest mocks base method
func (m *MockActorService) TellRequest(arg0 string, arg1 interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "TellRequest", arg0, arg1)
}

// TellRequest indicates an expected call of TellRequest
func (mr *MockActorServiceMockRecorder) TellRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TellRequest", reflect.TypeOf((*MockActorService)(nil).TellRequest), arg0, arg1)
}
