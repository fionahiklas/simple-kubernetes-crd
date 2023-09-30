// Code generated by MockGen. DO NOT EDIT.
// Source: handler.go

// Package treehandler_test is a generated GoMock package.
package treehandler_test

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockresponseWriter is a mock of responseWriter interface.
type MockresponseWriter struct {
	ctrl     *gomock.Controller
	recorder *MockresponseWriterMockRecorder
}

// MockresponseWriterMockRecorder is the mock recorder for MockresponseWriter.
type MockresponseWriterMockRecorder struct {
	mock *MockresponseWriter
}

// NewMockresponseWriter creates a new mock instance.
func NewMockresponseWriter(ctrl *gomock.Controller) *MockresponseWriter {
	mock := &MockresponseWriter{ctrl: ctrl}
	mock.recorder = &MockresponseWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockresponseWriter) EXPECT() *MockresponseWriterMockRecorder {
	return m.recorder
}

// Header mocks base method.
func (m *MockresponseWriter) Header() http.Header {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(http.Header)
	return ret0
}

// Header indicates an expected call of Header.
func (mr *MockresponseWriterMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockresponseWriter)(nil).Header))
}

// Write mocks base method.
func (m *MockresponseWriter) Write(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockresponseWriterMockRecorder) Write(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockresponseWriter)(nil).Write), arg0)
}

// WriteHeader mocks base method.
func (m *MockresponseWriter) WriteHeader(statusCode int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WriteHeader", statusCode)
}

// WriteHeader indicates an expected call of WriteHeader.
func (mr *MockresponseWriterMockRecorder) WriteHeader(statusCode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteHeader", reflect.TypeOf((*MockresponseWriter)(nil).WriteHeader), statusCode)
}

// Mocklogger is a mock of logger interface.
type Mocklogger struct {
	ctrl     *gomock.Controller
	recorder *MockloggerMockRecorder
}

// MockloggerMockRecorder is the mock recorder for Mocklogger.
type MockloggerMockRecorder struct {
	mock *Mocklogger
}

// NewMocklogger creates a new mock instance.
func NewMocklogger(ctrl *gomock.Controller) *Mocklogger {
	mock := &Mocklogger{ctrl: ctrl}
	mock.recorder = &MockloggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mocklogger) EXPECT() *MockloggerMockRecorder {
	return m.recorder
}

// Debugf mocks base method.
func (m *Mocklogger) Debugf(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugf", varargs...)
}

// Debugf indicates an expected call of Debugf.
func (mr *MockloggerMockRecorder) Debugf(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugf", reflect.TypeOf((*Mocklogger)(nil).Debugf), varargs...)
}

// Errorf mocks base method.
func (m *Mocklogger) Errorf(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf.
func (mr *MockloggerMockRecorder) Errorf(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*Mocklogger)(nil).Errorf), varargs...)
}

// Infof mocks base method.
func (m *Mocklogger) Infof(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Infof", varargs...)
}

// Infof indicates an expected call of Infof.
func (mr *MockloggerMockRecorder) Infof(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Infof", reflect.TypeOf((*Mocklogger)(nil).Infof), varargs...)
}

// Warnf mocks base method.
func (m *Mocklogger) Warnf(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warnf", varargs...)
}

// Warnf indicates an expected call of Warnf.
func (mr *MockloggerMockRecorder) Warnf(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warnf", reflect.TypeOf((*Mocklogger)(nil).Warnf), varargs...)
}

// Mockconfig is a mock of config interface.
type Mockconfig struct {
	ctrl     *gomock.Controller
	recorder *MockconfigMockRecorder
}

// MockconfigMockRecorder is the mock recorder for Mockconfig.
type MockconfigMockRecorder struct {
	mock *Mockconfig
}

// NewMockconfig creates a new mock instance.
func NewMockconfig(ctrl *gomock.Controller) *Mockconfig {
	mock := &Mockconfig{ctrl: ctrl}
	mock.recorder = &MockconfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockconfig) EXPECT() *MockconfigMockRecorder {
	return m.recorder
}

// EyesClosed mocks base method.
func (m *Mockconfig) EyesClosed() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EyesClosed")
	ret0, _ := ret[0].(bool)
	return ret0
}

// EyesClosed indicates an expected call of EyesClosed.
func (mr *MockconfigMockRecorder) EyesClosed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EyesClosed", reflect.TypeOf((*Mockconfig)(nil).EyesClosed))
}

// HowFarAway mocks base method.
func (m *Mockconfig) HowFarAway() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HowFarAway")
	ret0, _ := ret[0].(int)
	return ret0
}

// HowFarAway indicates an expected call of HowFarAway.
func (mr *MockconfigMockRecorder) HowFarAway() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HowFarAway", reflect.TypeOf((*Mockconfig)(nil).HowFarAway))
}

// TreeName mocks base method.
func (m *Mockconfig) TreeName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TreeName")
	ret0, _ := ret[0].(string)
	return ret0
}

// TreeName indicates an expected call of TreeName.
func (mr *MockconfigMockRecorder) TreeName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TreeName", reflect.TypeOf((*Mockconfig)(nil).TreeName))
}

// Try mocks base method.
func (m *Mockconfig) Try() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Try")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Try indicates an expected call of Try.
func (mr *MockconfigMockRecorder) Try() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Try", reflect.TypeOf((*Mockconfig)(nil).Try))
}
