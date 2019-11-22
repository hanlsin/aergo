// Code generated by MockGen. DO NOT EDIT.
// Source: peerrole.go

// Package p2pmock is a generated GoMock package.
package p2pmock

import (
	p2pcommon "github.com/aergoio/aergo/p2p/p2pcommon"
	types "github.com/aergoio/aergo/types"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPeerRoleManager is a mock of PeerRoleManager interface
type MockPeerRoleManager struct {
	ctrl     *gomock.Controller
	recorder *MockPeerRoleManagerMockRecorder
}

// MockPeerRoleManagerMockRecorder is the mock recorder for MockPeerRoleManager
type MockPeerRoleManagerMockRecorder struct {
	mock *MockPeerRoleManager
}

// NewMockPeerRoleManager creates a new mock instance
func NewMockPeerRoleManager(ctrl *gomock.Controller) *MockPeerRoleManager {
	mock := &MockPeerRoleManager{ctrl: ctrl}
	mock.recorder = &MockPeerRoleManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPeerRoleManager) EXPECT() *MockPeerRoleManagerMockRecorder {
	return m.recorder
}

// UpdateBP mocks base method
func (m *MockPeerRoleManager) UpdateBP(toAdd, toRemove []types.PeerID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateBP", toAdd, toRemove)
}

// UpdateBP indicates an expected call of UpdateBP
func (mr *MockPeerRoleManagerMockRecorder) UpdateBP(toAdd, toRemove interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBP", reflect.TypeOf((*MockPeerRoleManager)(nil).UpdateBP), toAdd, toRemove)
}

// SelfRole mocks base method
func (m *MockPeerRoleManager) SelfRole() types.PeerRole {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelfRole")
	ret0, _ := ret[0].(types.PeerRole)
	return ret0
}

// SelfRole indicates an expected call of SelfRole
func (mr *MockPeerRoleManagerMockRecorder) SelfRole() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelfRole", reflect.TypeOf((*MockPeerRoleManager)(nil).SelfRole))
}

// GetRole mocks base method
func (m *MockPeerRoleManager) GetRole(pid types.PeerID) types.PeerRole {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRole", pid)
	ret0, _ := ret[0].(types.PeerRole)
	return ret0
}

// GetRole indicates an expected call of GetRole
func (mr *MockPeerRoleManagerMockRecorder) GetRole(pid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRole", reflect.TypeOf((*MockPeerRoleManager)(nil).GetRole), pid)
}

// FilterBPNoticeReceiver mocks base method
func (m *MockPeerRoleManager) FilterBPNoticeReceiver(block *types.Block, pm p2pcommon.PeerManager, targetZone p2pcommon.PeerZone) []p2pcommon.RemotePeer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilterBPNoticeReceiver", block, pm)
	ret0, _ := ret[0].([]p2pcommon.RemotePeer)
	return ret0
}

// FilterBPNoticeReceiver indicates an expected call of FilterBPNoticeReceiver
func (mr *MockPeerRoleManagerMockRecorder) FilterBPNoticeReceiver(block, pm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterBPNoticeReceiver", reflect.TypeOf((*MockPeerRoleManager)(nil).FilterBPNoticeReceiver), block, pm)
}

// FilterNewBlockNoticeReceiver mocks base method
func (m *MockPeerRoleManager) FilterNewBlockNoticeReceiver(block *types.Block, pm p2pcommon.PeerManager) []p2pcommon.RemotePeer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilterNewBlockNoticeReceiver", block, pm)
	ret0, _ := ret[0].([]p2pcommon.RemotePeer)
	return ret0
}

// FilterNewBlockNoticeReceiver indicates an expected call of FilterNewBlockNoticeReceiver
func (mr *MockPeerRoleManagerMockRecorder) FilterNewBlockNoticeReceiver(block, pm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterNewBlockNoticeReceiver", reflect.TypeOf((*MockPeerRoleManager)(nil).FilterNewBlockNoticeReceiver), block, pm)
}
