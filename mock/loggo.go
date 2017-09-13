// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/go-figure/logr (interfaces: Sink)

// Package mock is a generated GoMock package.
package mock

import (
	logr "github.com/go-figure/logr"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockSink is a mock of Sink interface
type MockSink struct {
	ctrl     *gomock.Controller
	recorder *MockSinkMockRecorder
}

// MockSinkMockRecorder is the mock recorder for MockSink
type MockSinkMockRecorder struct {
	mock *MockSink
}

// NewMockSink creates a new mock instance
func NewMockSink(ctrl *gomock.Controller) *MockSink {
	mock := &MockSink{ctrl: ctrl}
	mock.recorder = &MockSinkMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSink) EXPECT() *MockSinkMockRecorder {
	return m.recorder
}

// Complete mocks base method
func (m *MockSink) Complete(arg0, arg1 string, arg2 logr.CompletionStatus, arg3 time.Duration, arg4 logr.KV) {
	m.ctrl.Call(m, "Complete", arg0, arg1, arg2, arg3, arg4)
}

// Complete indicates an expected call of Complete
func (mr *MockSinkMockRecorder) Complete(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Complete", reflect.TypeOf((*MockSink)(nil).Complete), arg0, arg1, arg2, arg3, arg4)
}

// Error mocks base method
func (m *MockSink) Error(arg0, arg1 string, arg2 error, arg3 logr.KV) {
	m.ctrl.Call(m, "Error", arg0, arg1, arg2, arg3)
}

// Error indicates an expected call of Error
func (mr *MockSinkMockRecorder) Error(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockSink)(nil).Error), arg0, arg1, arg2, arg3)
}

// Event mocks base method
func (m *MockSink) Event(arg0, arg1 string, arg2 logr.KV) {
	m.ctrl.Call(m, "Event", arg0, arg1, arg2)
}

// Event indicates an expected call of Event
func (mr *MockSinkMockRecorder) Event(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Event", reflect.TypeOf((*MockSink)(nil).Event), arg0, arg1, arg2)
}

// Gauge mocks base method
func (m *MockSink) Gauge(arg0, arg1 string, arg2 float64, arg3 logr.KV) {
	m.ctrl.Call(m, "Gauge", arg0, arg1, arg2, arg3)
}

// Gauge indicates an expected call of Gauge
func (mr *MockSinkMockRecorder) Gauge(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Gauge", reflect.TypeOf((*MockSink)(nil).Gauge), arg0, arg1, arg2, arg3)
}

// Timing mocks base method
func (m *MockSink) Timing(arg0, arg1 string, arg2 time.Duration, arg3 logr.KV) {
	m.ctrl.Call(m, "Timing", arg0, arg1, arg2, arg3)
}

// Timing indicates an expected call of Timing
func (mr *MockSinkMockRecorder) Timing(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Timing", reflect.TypeOf((*MockSink)(nil).Timing), arg0, arg1, arg2, arg3)
}
