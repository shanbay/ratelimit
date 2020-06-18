// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/envoyproxy/ratelimit/src/limiter (interfaces: RateLimitCache,TimeSource,JitterRandSource)

// Package mock_limiter is a generated GoMock package.
package mock_limiter

import (
	"github.com/envoyproxy/ratelimit/src/filter"

	context "context"
	v2 "github.com/envoyproxy/go-control-plane/envoy/service/ratelimit/v2"
	config "github.com/envoyproxy/ratelimit/src/config"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRateLimitCache is a mock of RateLimitCache interface
type MockRateLimitCache struct {
	ctrl     *gomock.Controller
	recorder *MockRateLimitCacheMockRecorder
}

// MockRateLimitCacheMockRecorder is the mock recorder for MockRateLimitCache
type MockRateLimitCacheMockRecorder struct {
	mock *MockRateLimitCache
}

// NewMockRateLimitCache creates a new mock instance
func NewMockRateLimitCache(ctrl *gomock.Controller) *MockRateLimitCache {
	mock := &MockRateLimitCache{ctrl: ctrl}
	mock.recorder = &MockRateLimitCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRateLimitCache) EXPECT() *MockRateLimitCacheMockRecorder {
	return m.recorder
}

// DoLimit mocks base method
func (m *MockRateLimitCache) DoLimit(arg0 context.Context, arg1 *v2.RateLimitRequest, arg2 []*config.RateLimit, arg3 bool, arg4, arg5 filter.Filter) []*v2.RateLimitResponse_DescriptorStatus {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoLimit", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].([]*v2.RateLimitResponse_DescriptorStatus)
	return ret0
}

// DoLimit indicates an expected call of DoLimit
func (mr *MockRateLimitCacheMockRecorder) DoLimit(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoLimit", reflect.TypeOf((*MockRateLimitCache)(nil).DoLimit), arg0, arg1, arg2, arg3, arg4, arg5)
}

// MockTimeSource is a mock of TimeSource interface
type MockTimeSource struct {
	ctrl     *gomock.Controller
	recorder *MockTimeSourceMockRecorder
}

// MockTimeSourceMockRecorder is the mock recorder for MockTimeSource
type MockTimeSourceMockRecorder struct {
	mock *MockTimeSource
}

// NewMockTimeSource creates a new mock instance
func NewMockTimeSource(ctrl *gomock.Controller) *MockTimeSource {
	mock := &MockTimeSource{ctrl: ctrl}
	mock.recorder = &MockTimeSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTimeSource) EXPECT() *MockTimeSourceMockRecorder {
	return m.recorder
}

// UnixNow mocks base method
func (m *MockTimeSource) UnixNow() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnixNow")
	ret0, _ := ret[0].(int64)
	return ret0
}

// UnixNow indicates an expected call of UnixNow
func (mr *MockTimeSourceMockRecorder) UnixNow() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnixNow", reflect.TypeOf((*MockTimeSource)(nil).UnixNow))
}

// MockJitterRandSource is a mock of JitterRandSource interface
type MockJitterRandSource struct {
	ctrl     *gomock.Controller
	recorder *MockJitterRandSourceMockRecorder
}

// MockJitterRandSourceMockRecorder is the mock recorder for MockJitterRandSource
type MockJitterRandSourceMockRecorder struct {
	mock *MockJitterRandSource
}

// NewMockJitterRandSource creates a new mock instance
func NewMockJitterRandSource(ctrl *gomock.Controller) *MockJitterRandSource {
	mock := &MockJitterRandSource{ctrl: ctrl}
	mock.recorder = &MockJitterRandSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockJitterRandSource) EXPECT() *MockJitterRandSourceMockRecorder {
	return m.recorder
}

// Int63 mocks base method
func (m *MockJitterRandSource) Int63() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Int63")
	ret0, _ := ret[0].(int64)
	return ret0
}

// Int63 indicates an expected call of Int63
func (mr *MockJitterRandSourceMockRecorder) Int63() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Int63", reflect.TypeOf((*MockJitterRandSource)(nil).Int63))
}

// Seed mocks base method
func (m *MockJitterRandSource) Seed(arg0 int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Seed", arg0)
}

// Seed indicates an expected call of Seed
func (mr *MockJitterRandSourceMockRecorder) Seed(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Seed", reflect.TypeOf((*MockJitterRandSource)(nil).Seed), arg0)
}
