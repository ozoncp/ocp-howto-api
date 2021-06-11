// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ozoncp/ocp-howto-api/internal/repo (interfaces: Repo)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	howto "github.com/ozoncp/ocp-howto-api/internal/howto"
)

// MockRepo is a mock of Repo interface.
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo.
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance.
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// AddHowto mocks base method.
func (m *MockRepo) AddHowto(arg0 howto.Howto) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddHowto", arg0)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddHowto indicates an expected call of AddHowto.
func (mr *MockRepoMockRecorder) AddHowto(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddHowto", reflect.TypeOf((*MockRepo)(nil).AddHowto), arg0)
}

// AddHowtos mocks base method.
func (m *MockRepo) AddHowtos(arg0 []howto.Howto) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddHowtos", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddHowtos indicates an expected call of AddHowtos.
func (mr *MockRepoMockRecorder) AddHowtos(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddHowtos", reflect.TypeOf((*MockRepo)(nil).AddHowtos), arg0)
}

// DescribeHowto mocks base method.
func (m *MockRepo) DescribeHowto(arg0 uint64) (howto.Howto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeHowto", arg0)
	ret0, _ := ret[0].(howto.Howto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeHowto indicates an expected call of DescribeHowto.
func (mr *MockRepoMockRecorder) DescribeHowto(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeHowto", reflect.TypeOf((*MockRepo)(nil).DescribeHowto), arg0)
}

// ListHowtos mocks base method.
func (m *MockRepo) ListHowtos(arg0, arg1 uint64) ([]howto.Howto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListHowtos", arg0, arg1)
	ret0, _ := ret[0].([]howto.Howto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListHowtos indicates an expected call of ListHowtos.
func (mr *MockRepoMockRecorder) ListHowtos(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListHowtos", reflect.TypeOf((*MockRepo)(nil).ListHowtos), arg0, arg1)
}

// RemoveHowto mocks base method.
func (m *MockRepo) RemoveHowto(arg0 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveHowto", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveHowto indicates an expected call of RemoveHowto.
func (mr *MockRepoMockRecorder) RemoveHowto(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveHowto", reflect.TypeOf((*MockRepo)(nil).RemoveHowto), arg0)
}
