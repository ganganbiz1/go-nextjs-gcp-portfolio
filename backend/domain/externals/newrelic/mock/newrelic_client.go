// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/externals/newrelic (interfaces: IfNewrelicClient)
//
// Generated by this command:
//
//	mockgen -destination=../domain/externals/newrelic/mock/newrelic_client.go github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/externals/newrelic IfNewrelicClient
//

// Package mock_newrelic is a generated GoMock package.
package mock_newrelic

import (
	context "context"
	reflect "reflect"

	newrelic "github.com/newrelic/go-agent/v3/newrelic"
	gomock "go.uber.org/mock/gomock"
)

// MockIfNewrelicClient is a mock of IfNewrelicClient interface.
type MockIfNewrelicClient struct {
	ctrl     *gomock.Controller
	recorder *MockIfNewrelicClientMockRecorder
}

// MockIfNewrelicClientMockRecorder is the mock recorder for MockIfNewrelicClient.
type MockIfNewrelicClientMockRecorder struct {
	mock *MockIfNewrelicClient
}

// NewMockIfNewrelicClient creates a new mock instance.
func NewMockIfNewrelicClient(ctrl *gomock.Controller) *MockIfNewrelicClient {
	mock := &MockIfNewrelicClient{ctrl: ctrl}
	mock.recorder = &MockIfNewrelicClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIfNewrelicClient) EXPECT() *MockIfNewrelicClientMockRecorder {
	return m.recorder
}

// NewContext mocks base method.
func (m *MockIfNewrelicClient) NewContext(arg0 context.Context, arg1 *newrelic.Transaction) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NewContext", arg0, arg1)
}

// NewContext indicates an expected call of NewContext.
func (mr *MockIfNewrelicClientMockRecorder) NewContext(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewContext", reflect.TypeOf((*MockIfNewrelicClient)(nil).NewContext), arg0, arg1)
}

// StartTransaction mocks base method.
func (m *MockIfNewrelicClient) StartTransaction(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StartTransaction", arg0)
}

// StartTransaction indicates an expected call of StartTransaction.
func (mr *MockIfNewrelicClientMockRecorder) StartTransaction(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartTransaction", reflect.TypeOf((*MockIfNewrelicClient)(nil).StartTransaction), arg0)
}
