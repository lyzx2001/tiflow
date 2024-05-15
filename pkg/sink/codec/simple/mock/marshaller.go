// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/sink/codec/simple/marshaller.go

// Package mock_simple is a generated GoMock package.
package mock_simple

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/pingcap/tiflow/cdc/model"
)

// Mockmarshaller is a mock of marshaller interface.
type Mockmarshaller struct {
	ctrl     *gomock.Controller
	recorder *MockmarshallerMockRecorder
}

// MockmarshallerMockRecorder is the mock recorder for Mockmarshaller.
type MockmarshallerMockRecorder struct {
	mock *Mockmarshaller
}

// NewMockmarshaller creates a new mock instance.
func NewMockmarshaller(ctrl *gomock.Controller) *Mockmarshaller {
	mock := &Mockmarshaller{ctrl: ctrl}
	mock.recorder = &MockmarshallerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockmarshaller) EXPECT() *MockmarshallerMockRecorder {
	return m.recorder
}

// MarshalCheckpoint mocks base method.
func (m *Mockmarshaller) MarshalCheckpoint(ts uint64) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarshalCheckpoint", ts)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalCheckpoint indicates an expected call of MarshalCheckpoint.
func (mr *MockmarshallerMockRecorder) MarshalCheckpoint(ts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalCheckpoint", reflect.TypeOf((*Mockmarshaller)(nil).MarshalCheckpoint), ts)
}

// MarshalDDLEvent mocks base method.
func (m *Mockmarshaller) MarshalDDLEvent(event *model.DDLEvent) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarshalDDLEvent", event)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalDDLEvent indicates an expected call of MarshalDDLEvent.
func (mr *MockmarshallerMockRecorder) MarshalDDLEvent(event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalDDLEvent", reflect.TypeOf((*Mockmarshaller)(nil).MarshalDDLEvent), event)
}

// MarshalRowChangedEvent mocks base method.
func (m *Mockmarshaller) MarshalRowChangedEvent(event *model.RowChangedEvent, handleKeyOnly bool, claimCheckFileName string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarshalRowChangedEvent", event, handleKeyOnly, claimCheckFileName)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalRowChangedEvent indicates an expected call of MarshalRowChangedEvent.
func (mr *MockmarshallerMockRecorder) MarshalRowChangedEvent(event, handleKeyOnly, claimCheckFileName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalRowChangedEvent", reflect.TypeOf((*Mockmarshaller)(nil).MarshalRowChangedEvent), event, handleKeyOnly, claimCheckFileName)
}

// Unmarshal mocks base method.
func (m *Mockmarshaller) Unmarshal(data []byte, v any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unmarshal", data, v)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unmarshal indicates an expected call of Unmarshal.
func (mr *MockmarshallerMockRecorder) Unmarshal(data, v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unmarshal", reflect.TypeOf((*Mockmarshaller)(nil).Unmarshal), data, v)
}