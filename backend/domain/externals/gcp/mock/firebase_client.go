// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/externals/gcp (interfaces: IfFirebaseClient)
//
// Generated by this command:
//
//	mockgen -destination=../domain/externals/gcp/mock/firebase_client.go github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/externals/gcp IfFirebaseClient
//

// Package mock_gcp is a generated GoMock package.
package mock_gcp

import (
	context "context"
	reflect "reflect"

	entity "github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockIfFirebaseClient is a mock of IfFirebaseClient interface.
type MockIfFirebaseClient struct {
	ctrl     *gomock.Controller
	recorder *MockIfFirebaseClientMockRecorder
}

// MockIfFirebaseClientMockRecorder is the mock recorder for MockIfFirebaseClient.
type MockIfFirebaseClientMockRecorder struct {
	mock *MockIfFirebaseClient
}

// NewMockIfFirebaseClient creates a new mock instance.
func NewMockIfFirebaseClient(ctrl *gomock.Controller) *MockIfFirebaseClient {
	mock := &MockIfFirebaseClient{ctrl: ctrl}
	mock.recorder = &MockIfFirebaseClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIfFirebaseClient) EXPECT() *MockIfFirebaseClientMockRecorder {
	return m.recorder
}

// Signup mocks base method.
func (m *MockIfFirebaseClient) Signup(arg0 context.Context, arg1, arg2 string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Signup", arg0, arg1, arg2)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Signup indicates an expected call of Signup.
func (mr *MockIfFirebaseClientMockRecorder) Signup(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Signup", reflect.TypeOf((*MockIfFirebaseClient)(nil).Signup), arg0, arg1, arg2)
}

// VerifyIDToken mocks base method.
func (m *MockIfFirebaseClient) VerifyIDToken(arg0 context.Context, arg1 string) (*entity.TokenInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyIDToken", arg0, arg1)
	ret0, _ := ret[0].(*entity.TokenInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyIDToken indicates an expected call of VerifyIDToken.
func (mr *MockIfFirebaseClientMockRecorder) VerifyIDToken(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyIDToken", reflect.TypeOf((*MockIfFirebaseClient)(nil).VerifyIDToken), arg0, arg1)
}
