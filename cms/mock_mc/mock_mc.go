// Code generated by MockGen. DO NOT EDIT.
// Source: ./mc.go

// Package mock_mc is a generated GoMock package.
package mock_mc

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockConn is a mock of Conn interface
type MockConn struct {
	ctrl     *gomock.Controller
	recorder *MockConnMockRecorder
}

// MockConnMockRecorder is the mock recorder for MockConn
type MockConnMockRecorder struct {
	mock *MockConn
}

// NewMockConn creates a new mock instance
func NewMockConn(ctrl *gomock.Controller) *MockConn {
	mock := &MockConn{ctrl: ctrl}
	mock.recorder = &MockConnMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConn) EXPECT() *MockConnMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockConn) Get(key string) (string, uint32, uint64, error) {
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(uint32)
	ret2, _ := ret[2].(uint64)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// Get indicates an expected call of Get
func (mr *MockConnMockRecorder) Get(key interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockConn)(nil).Get), key)
}

// GAT mocks base method
func (m *MockConn) GAT(key string, exp uint32) (string, uint32, uint64, error) {
	ret := m.ctrl.Call(m, "GAT", key, exp)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(uint32)
	ret2, _ := ret[2].(uint64)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// GAT indicates an expected call of GAT
func (mr *MockConnMockRecorder) GAT(key, exp interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GAT", reflect.TypeOf((*MockConn)(nil).GAT), key, exp)
}

// Touch mocks base method
func (m *MockConn) Touch(key string, exp uint32) (uint64, error) {
	ret := m.ctrl.Call(m, "Touch", key, exp)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Touch indicates an expected call of Touch
func (mr *MockConnMockRecorder) Touch(key, exp interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Touch", reflect.TypeOf((*MockConn)(nil).Touch), key, exp)
}

// Set mocks base method
func (m *MockConn) Set(key, val string, flags, exp uint32, ocas uint64) (uint64, error) {
	ret := m.ctrl.Call(m, "Set", key, val, flags, exp, ocas)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Set indicates an expected call of Set
func (mr *MockConnMockRecorder) Set(key, val, flags, exp, ocas interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockConn)(nil).Set), key, val, flags, exp, ocas)
}

// Replace mocks base method
func (m *MockConn) Replace(key, val string, flags, exp uint32, ocas uint64) (uint64, error) {
	ret := m.ctrl.Call(m, "Replace", key, val, flags, exp, ocas)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Replace indicates an expected call of Replace
func (mr *MockConnMockRecorder) Replace(key, val, flags, exp, ocas interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Replace", reflect.TypeOf((*MockConn)(nil).Replace), key, val, flags, exp, ocas)
}

// Add mocks base method
func (m *MockConn) Add(key, val string, flags, exp uint32) (uint64, error) {
	ret := m.ctrl.Call(m, "Add", key, val, flags, exp)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add
func (mr *MockConnMockRecorder) Add(key, val, flags, exp interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockConn)(nil).Add), key, val, flags, exp)
}

// Incr mocks base method
func (m *MockConn) Incr(key string, delta, init uint64, exp uint32, ocas uint64) (uint64, uint64, error) {
	ret := m.ctrl.Call(m, "Incr", key, delta, init, exp, ocas)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(uint64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Incr indicates an expected call of Incr
func (mr *MockConnMockRecorder) Incr(key, delta, init, exp, ocas interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Incr", reflect.TypeOf((*MockConn)(nil).Incr), key, delta, init, exp, ocas)
}

// Decr mocks base method
func (m *MockConn) Decr(key string, delta, init uint64, exp uint32, ocas uint64) (uint64, uint64, error) {
	ret := m.ctrl.Call(m, "Decr", key, delta, init, exp, ocas)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(uint64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Decr indicates an expected call of Decr
func (mr *MockConnMockRecorder) Decr(key, delta, init, exp, ocas interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decr", reflect.TypeOf((*MockConn)(nil).Decr), key, delta, init, exp, ocas)
}

// Append mocks base method
func (m *MockConn) Append(key, val string, ocas uint64) (uint64, error) {
	ret := m.ctrl.Call(m, "Append", key, val, ocas)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Append indicates an expected call of Append
func (mr *MockConnMockRecorder) Append(key, val, ocas interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Append", reflect.TypeOf((*MockConn)(nil).Append), key, val, ocas)
}

// Prepend mocks base method
func (m *MockConn) Prepend(key, val string, ocas uint64) (uint64, error) {
	ret := m.ctrl.Call(m, "Prepend", key, val, ocas)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Prepend indicates an expected call of Prepend
func (mr *MockConnMockRecorder) Prepend(key, val, ocas interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prepend", reflect.TypeOf((*MockConn)(nil).Prepend), key, val, ocas)
}

// Del mocks base method
func (m *MockConn) Del(key string) error {
	ret := m.ctrl.Call(m, "Del", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Del indicates an expected call of Del
func (mr *MockConnMockRecorder) Del(key interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockConn)(nil).Del), key)
}

// DelCAS mocks base method
func (m *MockConn) DelCAS(key string, cas uint64) error {
	ret := m.ctrl.Call(m, "DelCAS", key, cas)
	ret0, _ := ret[0].(error)
	return ret0
}

// DelCAS indicates an expected call of DelCAS
func (mr *MockConnMockRecorder) DelCAS(key, cas interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelCAS", reflect.TypeOf((*MockConn)(nil).DelCAS), key, cas)
}

// Flush mocks base method
func (m *MockConn) Flush(when uint32) error {
	ret := m.ctrl.Call(m, "Flush", when)
	ret0, _ := ret[0].(error)
	return ret0
}

// Flush indicates an expected call of Flush
func (mr *MockConnMockRecorder) Flush(when interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockConn)(nil).Flush), when)
}

// NoOp mocks base method
func (m *MockConn) NoOp() error {
	ret := m.ctrl.Call(m, "NoOp")
	ret0, _ := ret[0].(error)
	return ret0
}

// NoOp indicates an expected call of NoOp
func (mr *MockConnMockRecorder) NoOp() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NoOp", reflect.TypeOf((*MockConn)(nil).NoOp))
}

// Version mocks base method
func (m *MockConn) Version() (string, error) {
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Version indicates an expected call of Version
func (mr *MockConnMockRecorder) Version() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockConn)(nil).Version))
}

// Quit mocks base method
func (m *MockConn) Quit() error {
	ret := m.ctrl.Call(m, "Quit")
	ret0, _ := ret[0].(error)
	return ret0
}

// Quit indicates an expected call of Quit
func (mr *MockConnMockRecorder) Quit() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Quit", reflect.TypeOf((*MockConn)(nil).Quit))
}

// StatsWithKey mocks base method
func (m *MockConn) StatsWithKey(key string) (map[string]string, error) {
	ret := m.ctrl.Call(m, "StatsWithKey", key)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StatsWithKey indicates an expected call of StatsWithKey
func (mr *MockConnMockRecorder) StatsWithKey(key interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatsWithKey", reflect.TypeOf((*MockConn)(nil).StatsWithKey), key)
}

// Stats mocks base method
func (m *MockConn) Stats() (map[string]string, error) {
	ret := m.ctrl.Call(m, "Stats")
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stats indicates an expected call of Stats
func (mr *MockConnMockRecorder) Stats() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stats", reflect.TypeOf((*MockConn)(nil).Stats))
}

// StatsReset mocks base method
func (m *MockConn) StatsReset() error {
	ret := m.ctrl.Call(m, "StatsReset")
	ret0, _ := ret[0].(error)
	return ret0
}

// StatsReset indicates an expected call of StatsReset
func (mr *MockConnMockRecorder) StatsReset() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatsReset", reflect.TypeOf((*MockConn)(nil).StatsReset))
}