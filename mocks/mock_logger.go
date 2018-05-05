// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/sirupsen/logrus (interfaces: FieldLogger)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	logrus "github.com/sirupsen/logrus"
	reflect "reflect"
)

// MockFieldLogger is a mock of FieldLogger interface
type MockFieldLogger struct {
	ctrl     *gomock.Controller
	recorder *MockFieldLoggerMockRecorder
}

// MockFieldLoggerMockRecorder is the mock recorder for MockFieldLogger
type MockFieldLoggerMockRecorder struct {
	mock *MockFieldLogger
}

// NewMockFieldLogger creates a new mock instance
func NewMockFieldLogger(ctrl *gomock.Controller) *MockFieldLogger {
	mock := &MockFieldLogger{ctrl: ctrl}
	mock.recorder = &MockFieldLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFieldLogger) EXPECT() *MockFieldLoggerMockRecorder {
	return m.recorder
}

// Debug mocks base method
func (m *MockFieldLogger) Debug(arg0 ...interface{}) {
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debug", varargs...)
}

// Debug indicates an expected call of Debug
func (mr *MockFieldLoggerMockRecorder) Debug(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debug", reflect.TypeOf((*MockFieldLogger)(nil).Debug), arg0...)
}

// Debugf mocks base method
func (m *MockFieldLogger) Debugf(arg0 string, arg1 ...interface{}) {
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugf", varargs...)
}

// Debugf indicates an expected call of Debugf
func (mr *MockFieldLoggerMockRecorder) Debugf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugf", reflect.TypeOf((*MockFieldLogger)(nil).Debugf), varargs...)
}

// Debugln mocks base method
func (m *MockFieldLogger) Debugln(arg0 ...interface{}) {
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugln", varargs...)
}

// Debugln indicates an expected call of Debugln
func (mr *MockFieldLoggerMockRecorder) Debugln(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugln", reflect.TypeOf((*MockFieldLogger)(nil).Debugln), arg0...)
}

// Error mocks base method
func (m *MockFieldLogger) Error(arg0 ...interface{}) {
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Error", varargs...)
}

// Error indicates an expected call of Error
func (mr *MockFieldLoggerMockRecorder) Error(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockFieldLogger)(nil).Error), arg0...)
}

// Errorf mocks base method
func (m *MockFieldLogger) Errorf(arg0 string, arg1 ...interface{}) {
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf
func (mr *MockFieldLoggerMockRecorder) Errorf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MockFieldLogger)(nil).Errorf), varargs...)
}

// Errorln mocks base method
func (m *MockFieldLogger) Errorln(arg0 ...interface{}) {
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorln", varargs...)
}

// Errorln indicates an expected call of Errorln
func (mr *MockFieldLoggerMockRecorder) Errorln(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorln", reflect.TypeOf((*MockFieldLogger)(nil).Errorln), arg0...)
}

// Fatal mocks base method
func (m *MockFieldLogger) Fatal(arg0 ...interface{}) {
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Fatal", varargs...)
}

// Fatal indicates an expected call of Fatal
func (mr *MockFieldLoggerMockRecorder) Fatal(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatal", reflect.TypeOf((*MockFieldLogger)(nil).Fatal), arg0...)
}

// Fatalf mocks base method
func (m *MockFieldLogger) Fatalf(arg0 string, arg1 ...interface{}) {
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Fatalf", varargs...)
}

// Fatalf indicates an expected call of Fatalf
func (mr *MockFieldLoggerMockRecorder) Fatalf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatalf", reflect.TypeOf((*MockFieldLogger)(nil).Fatalf), varargs...)
}

// Fatalln mocks base method
func (m *MockFieldLogger) Fatalln(arg0 ...interface{}) {
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Fatalln", varargs...)
}

// Fatalln indicates an expected call of Fatalln
func (mr *MockFieldLoggerMockRecorder) Fatalln(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatalln", reflect.TypeOf((*MockFieldLogger)(nil).Fatalln), arg0...)
}

// Info mocks base method
func (m *MockFieldLogger) Info(arg0 ...interface{}) {
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Info", varargs...)
}

// Info indicates an expected call of Info
func (mr *MockFieldLoggerMockRecorder) Info(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockFieldLogger)(nil).Info), arg0...)
}

// Infof mocks base method
func (m *MockFieldLogger) Infof(arg0 string, arg1 ...interface{}) {
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Infof", varargs...)
}

// Infof indicates an expected call of Infof
func (mr *MockFieldLoggerMockRecorder) Infof(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Infof", reflect.TypeOf((*MockFieldLogger)(nil).Infof), varargs...)
}

// Infoln mocks base method
func (m *MockFieldLogger) Infoln(arg0 ...interface{}) {
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Infoln", varargs...)
}

// Infoln indicates an expected call of Infoln
func (mr *MockFieldLoggerMockRecorder) Infoln(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Infoln", reflect.TypeOf((*MockFieldLogger)(nil).Infoln), arg0...)
}

// Panic mocks base method
func (m *MockFieldLogger) Panic(arg0 ...interface{}) {
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Panic", varargs...)
}

// Panic indicates an expected call of Panic
func (mr *MockFieldLoggerMockRecorder) Panic(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Panic", reflect.TypeOf((*MockFieldLogger)(nil).Panic), arg0...)
}

// Panicf mocks base method
func (m *MockFieldLogger) Panicf(arg0 string, arg1 ...interface{}) {
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Panicf", varargs...)
}

// Panicf indicates an expected call of Panicf
func (mr *MockFieldLoggerMockRecorder) Panicf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Panicf", reflect.TypeOf((*MockFieldLogger)(nil).Panicf), varargs...)
}

// Panicln mocks base method
func (m *MockFieldLogger) Panicln(arg0 ...interface{}) {
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Panicln", varargs...)
}

// Panicln indicates an expected call of Panicln
func (mr *MockFieldLoggerMockRecorder) Panicln(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Panicln", reflect.TypeOf((*MockFieldLogger)(nil).Panicln), arg0...)
}

// Print mocks base method
func (m *MockFieldLogger) Print(arg0 ...interface{}) {
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Print", varargs...)
}

// Print indicates an expected call of Print
func (mr *MockFieldLoggerMockRecorder) Print(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Print", reflect.TypeOf((*MockFieldLogger)(nil).Print), arg0...)
}

// Printf mocks base method
func (m *MockFieldLogger) Printf(arg0 string, arg1 ...interface{}) {
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Printf", varargs...)
}

// Printf indicates an expected call of Printf
func (mr *MockFieldLoggerMockRecorder) Printf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Printf", reflect.TypeOf((*MockFieldLogger)(nil).Printf), varargs...)
}

// Println mocks base method
func (m *MockFieldLogger) Println(arg0 ...interface{}) {
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Println", varargs...)
}

// Println indicates an expected call of Println
func (mr *MockFieldLoggerMockRecorder) Println(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Println", reflect.TypeOf((*MockFieldLogger)(nil).Println), arg0...)
}

// Warn mocks base method
func (m *MockFieldLogger) Warn(arg0 ...interface{}) {
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warn", varargs...)
}

// Warn indicates an expected call of Warn
func (mr *MockFieldLoggerMockRecorder) Warn(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warn", reflect.TypeOf((*MockFieldLogger)(nil).Warn), arg0...)
}

// Warnf mocks base method
func (m *MockFieldLogger) Warnf(arg0 string, arg1 ...interface{}) {
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warnf", varargs...)
}

// Warnf indicates an expected call of Warnf
func (mr *MockFieldLoggerMockRecorder) Warnf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warnf", reflect.TypeOf((*MockFieldLogger)(nil).Warnf), varargs...)
}

// Warning mocks base method
func (m *MockFieldLogger) Warning(arg0 ...interface{}) {
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warning", varargs...)
}

// Warning indicates an expected call of Warning
func (mr *MockFieldLoggerMockRecorder) Warning(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warning", reflect.TypeOf((*MockFieldLogger)(nil).Warning), arg0...)
}

// Warningf mocks base method
func (m *MockFieldLogger) Warningf(arg0 string, arg1 ...interface{}) {
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warningf", varargs...)
}

// Warningf indicates an expected call of Warningf
func (mr *MockFieldLoggerMockRecorder) Warningf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warningf", reflect.TypeOf((*MockFieldLogger)(nil).Warningf), varargs...)
}

// Warningln mocks base method
func (m *MockFieldLogger) Warningln(arg0 ...interface{}) {
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warningln", varargs...)
}

// Warningln indicates an expected call of Warningln
func (mr *MockFieldLoggerMockRecorder) Warningln(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warningln", reflect.TypeOf((*MockFieldLogger)(nil).Warningln), arg0...)
}

// Warnln mocks base method
func (m *MockFieldLogger) Warnln(arg0 ...interface{}) {
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warnln", varargs...)
}

// Warnln indicates an expected call of Warnln
func (mr *MockFieldLoggerMockRecorder) Warnln(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warnln", reflect.TypeOf((*MockFieldLogger)(nil).Warnln), arg0...)
}

// WithError mocks base method
func (m *MockFieldLogger) WithError(arg0 error) *logrus.Entry {
	ret := m.ctrl.Call(m, "WithError", arg0)
	ret0, _ := ret[0].(*logrus.Entry)
	return ret0
}

// WithError indicates an expected call of WithError
func (mr *MockFieldLoggerMockRecorder) WithError(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithError", reflect.TypeOf((*MockFieldLogger)(nil).WithError), arg0)
}

// WithField mocks base method
func (m *MockFieldLogger) WithField(arg0 string, arg1 interface{}) *logrus.Entry {
	ret := m.ctrl.Call(m, "WithField", arg0, arg1)
	ret0, _ := ret[0].(*logrus.Entry)
	return ret0
}

// WithField indicates an expected call of WithField
func (mr *MockFieldLoggerMockRecorder) WithField(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithField", reflect.TypeOf((*MockFieldLogger)(nil).WithField), arg0, arg1)
}

// WithFields mocks base method
func (m *MockFieldLogger) WithFields(arg0 logrus.Fields) *logrus.Entry {
	ret := m.ctrl.Call(m, "WithFields", arg0)
	ret0, _ := ret[0].(*logrus.Entry)
	return ret0
}

// WithFields indicates an expected call of WithFields
func (mr *MockFieldLoggerMockRecorder) WithFields(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithFields", reflect.TypeOf((*MockFieldLogger)(nil).WithFields), arg0)
}
