// Code generated by MockGen. DO NOT EDIT.
// Source: logging.go

// Package logging_test is a generated GoMock package.
package logging_test

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSimpleLogger is a mock of SimpleLogger interface.
type MockSimpleLogger struct {
	ctrl     *gomock.Controller
	recorder *MockSimpleLoggerMockRecorder
}

// MockSimpleLoggerMockRecorder is the mock recorder for MockSimpleLogger.
type MockSimpleLoggerMockRecorder struct {
	mock *MockSimpleLogger
}

// NewMockSimpleLogger creates a new mock instance.
func NewMockSimpleLogger(ctrl *gomock.Controller) *MockSimpleLogger {
	mock := &MockSimpleLogger{ctrl: ctrl}
	mock.recorder = &MockSimpleLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSimpleLogger) EXPECT() *MockSimpleLoggerMockRecorder {
	return m.recorder
}

// Debugf mocks base method.
func (m *MockSimpleLogger) Debugf(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugf", varargs...)
}

// Debugf indicates an expected call of Debugf.
func (mr *MockSimpleLoggerMockRecorder) Debugf(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugf", reflect.TypeOf((*MockSimpleLogger)(nil).Debugf), varargs...)
}

// Errorf mocks base method.
func (m *MockSimpleLogger) Errorf(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf.
func (mr *MockSimpleLoggerMockRecorder) Errorf(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MockSimpleLogger)(nil).Errorf), varargs...)
}

// Fatalf mocks base method.
func (m *MockSimpleLogger) Fatalf(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Fatalf", varargs...)
}

// Fatalf indicates an expected call of Fatalf.
func (mr *MockSimpleLoggerMockRecorder) Fatalf(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatalf", reflect.TypeOf((*MockSimpleLogger)(nil).Fatalf), varargs...)
}

// Infof mocks base method.
func (m *MockSimpleLogger) Infof(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Infof", varargs...)
}

// Infof indicates an expected call of Infof.
func (mr *MockSimpleLoggerMockRecorder) Infof(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Infof", reflect.TypeOf((*MockSimpleLogger)(nil).Infof), varargs...)
}

// Warnf mocks base method.
func (m *MockSimpleLogger) Warnf(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warnf", varargs...)
}

// Warnf indicates an expected call of Warnf.
func (mr *MockSimpleLoggerMockRecorder) Warnf(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warnf", reflect.TypeOf((*MockSimpleLogger)(nil).Warnf), varargs...)
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

// EnableJsonLogFormat mocks base method.
func (m *Mockconfig) EnableJsonLogFormat() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableJsonLogFormat")
	ret0, _ := ret[0].(bool)
	return ret0
}

// EnableJsonLogFormat indicates an expected call of EnableJsonLogFormat.
func (mr *MockconfigMockRecorder) EnableJsonLogFormat() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableJsonLogFormat", reflect.TypeOf((*Mockconfig)(nil).EnableJsonLogFormat))
}

// LogLevel mocks base method.
func (m *Mockconfig) LogLevel() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LogLevel")
	ret0, _ := ret[0].(string)
	return ret0
}

// LogLevel indicates an expected call of LogLevel.
func (mr *MockconfigMockRecorder) LogLevel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogLevel", reflect.TypeOf((*Mockconfig)(nil).LogLevel))
}
