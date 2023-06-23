// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package db is a generated GoMock package.
package db

import (
	reflect "reflect"

	model "github.com/blackhorseya/project-goapi/entity/domain/user/model"
	contextx "github.com/blackhorseya/project-goapi/pkg/contextx"
	gomock "github.com/golang/mock/gomock"
)

// MockReader is a mock of Reader interface.
type MockReader struct {
	ctrl     *gomock.Controller
	recorder *MockReaderMockRecorder
}

// MockReaderMockRecorder is the mock recorder for MockReader.
type MockReaderMockRecorder struct {
	mock *MockReader
}

// NewMockReader creates a new mock instance.
func NewMockReader(ctrl *gomock.Controller) *MockReader {
	mock := &MockReader{ctrl: ctrl}
	mock.recorder = &MockReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReader) EXPECT() *MockReaderMockRecorder {
	return m.recorder
}

// GetUserByName mocks base method.
func (m *MockReader) GetUserByName(ctx contextx.Contextx, name string) (*model.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByName", ctx, name)
	ret0, _ := ret[0].(*model.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByName indicates an expected call of GetUserByName.
func (mr *MockReaderMockRecorder) GetUserByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByName", reflect.TypeOf((*MockReader)(nil).GetUserByName), ctx, name)
}

// MockWriter is a mock of Writer interface.
type MockWriter struct {
	ctrl     *gomock.Controller
	recorder *MockWriterMockRecorder
}

// MockWriterMockRecorder is the mock recorder for MockWriter.
type MockWriterMockRecorder struct {
	mock *MockWriter
}

// NewMockWriter creates a new mock instance.
func NewMockWriter(ctrl *gomock.Controller) *MockWriter {
	mock := &MockWriter{ctrl: ctrl}
	mock.recorder = &MockWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWriter) EXPECT() *MockWriterMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockWriter) CreateUser(ctx contextx.Contextx, name, password string) (*model.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, name, password)
	ret0, _ := ret[0].(*model.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockWriterMockRecorder) CreateUser(ctx, name, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockWriter)(nil).CreateUser), ctx, name, password)
}

// UpdateUser mocks base method.
func (m *MockWriter) UpdateUser(ctx contextx.Contextx, user *model.Profile) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockWriterMockRecorder) UpdateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockWriter)(nil).UpdateUser), ctx, user)
}

// MockReaderWriter is a mock of ReaderWriter interface.
type MockReaderWriter struct {
	ctrl     *gomock.Controller
	recorder *MockReaderWriterMockRecorder
}

// MockReaderWriterMockRecorder is the mock recorder for MockReaderWriter.
type MockReaderWriterMockRecorder struct {
	mock *MockReaderWriter
}

// NewMockReaderWriter creates a new mock instance.
func NewMockReaderWriter(ctrl *gomock.Controller) *MockReaderWriter {
	mock := &MockReaderWriter{ctrl: ctrl}
	mock.recorder = &MockReaderWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReaderWriter) EXPECT() *MockReaderWriterMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockReaderWriter) CreateUser(ctx contextx.Contextx, name, password string) (*model.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, name, password)
	ret0, _ := ret[0].(*model.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockReaderWriterMockRecorder) CreateUser(ctx, name, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockReaderWriter)(nil).CreateUser), ctx, name, password)
}

// GetUserByName mocks base method.
func (m *MockReaderWriter) GetUserByName(ctx contextx.Contextx, name string) (*model.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByName", ctx, name)
	ret0, _ := ret[0].(*model.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByName indicates an expected call of GetUserByName.
func (mr *MockReaderWriterMockRecorder) GetUserByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByName", reflect.TypeOf((*MockReaderWriter)(nil).GetUserByName), ctx, name)
}

// UpdateUser mocks base method.
func (m *MockReaderWriter) UpdateUser(ctx contextx.Contextx, user *model.Profile) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockReaderWriterMockRecorder) UpdateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockReaderWriter)(nil).UpdateUser), ctx, user)
}
