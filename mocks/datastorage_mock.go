// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/q231950/alethea/datastorage (interfaces: DataStorage)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	model "github.com/q231950/alethea/model"
	reflect "reflect"
)

// MockDataStorage is a mock of DataStorage interface
type MockDataStorage struct {
	ctrl     *gomock.Controller
	recorder *MockDataStorageMockRecorder
}

// MockDataStorageMockRecorder is the mock recorder for MockDataStorage
type MockDataStorageMockRecorder struct {
	mock *MockDataStorage
}

// NewMockDataStorage creates a new mock instance
func NewMockDataStorage(ctrl *gomock.Controller) *MockDataStorage {
	mock := &MockDataStorage{ctrl: ctrl}
	mock.recorder = &MockDataStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDataStorage) EXPECT() *MockDataStorageMockRecorder {
	return m.recorder
}

// CreateIncidentsTable mocks base method
func (m *MockDataStorage) CreateIncidentsTable() {
	m.ctrl.Call(m, "CreateIncidentsTable")
}

// CreateIncidentsTable indicates an expected call of CreateIncidentsTable
func (mr *MockDataStorageMockRecorder) CreateIncidentsTable() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIncidentsTable", reflect.TypeOf((*MockDataStorage)(nil).CreateIncidentsTable))
}

// StoreIncident mocks base method
func (m *MockDataStorage) StoreIncident(arg0 model.Incident) error {
	ret := m.ctrl.Call(m, "StoreIncident", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreIncident indicates an expected call of StoreIncident
func (mr *MockDataStorageMockRecorder) StoreIncident(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreIncident", reflect.TypeOf((*MockDataStorage)(nil).StoreIncident), arg0)
}
