// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/pingcap/tiflow/engine/pkg/rpcutil (interfaces: FeatureChecker)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFeatureChecker is a mock of FeatureChecker interface.
type MockFeatureChecker struct {
	ctrl     *gomock.Controller
	recorder *MockFeatureCheckerMockRecorder
}

// MockFeatureCheckerMockRecorder is the mock recorder for MockFeatureChecker.
type MockFeatureCheckerMockRecorder struct {
	mock *MockFeatureChecker
}

// NewMockFeatureChecker creates a new mock instance.
func NewMockFeatureChecker(ctrl *gomock.Controller) *MockFeatureChecker {
	mock := &MockFeatureChecker{ctrl: ctrl}
	mock.recorder = &MockFeatureCheckerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFeatureChecker) EXPECT() *MockFeatureCheckerMockRecorder {
	return m.recorder
}

// Available mocks base method.
func (m *MockFeatureChecker) Available(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Available", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Available indicates an expected call of Available.
func (mr *MockFeatureCheckerMockRecorder) Available(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Available", reflect.TypeOf((*MockFeatureChecker)(nil).Available), arg0)
}