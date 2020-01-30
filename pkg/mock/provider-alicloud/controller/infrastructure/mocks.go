// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/gardener-extension-provider-alicloud/pkg/controller/infrastructure (interfaces: TerraformChartOps)

// Package infrastructure is a generated GoMock package.
package infrastructure

import (
	v1alpha1 "github.com/gardener/gardener-extension-provider-alicloud/pkg/apis/alicloud/v1alpha1"
	infrastructure "github.com/gardener/gardener-extension-provider-alicloud/pkg/controller/infrastructure"
	v1alpha10 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockTerraformChartOps is a mock of TerraformChartOps interface
type MockTerraformChartOps struct {
	ctrl     *gomock.Controller
	recorder *MockTerraformChartOpsMockRecorder
}

// MockTerraformChartOpsMockRecorder is the mock recorder for MockTerraformChartOps
type MockTerraformChartOpsMockRecorder struct {
	mock *MockTerraformChartOps
}

// NewMockTerraformChartOps creates a new mock instance
func NewMockTerraformChartOps(ctrl *gomock.Controller) *MockTerraformChartOps {
	mock := &MockTerraformChartOps{ctrl: ctrl}
	mock.recorder = &MockTerraformChartOpsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTerraformChartOps) EXPECT() *MockTerraformChartOpsMockRecorder {
	return m.recorder
}

// ComputeChartValues mocks base method
func (m *MockTerraformChartOps) ComputeChartValues(arg0 *v1alpha10.Infrastructure, arg1 *v1alpha1.InfrastructureConfig, arg2 *infrastructure.InitializerValues) map[string]interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComputeChartValues", arg0, arg1, arg2)
	ret0, _ := ret[0].(map[string]interface{})
	return ret0
}

// ComputeChartValues indicates an expected call of ComputeChartValues
func (mr *MockTerraformChartOpsMockRecorder) ComputeChartValues(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComputeChartValues", reflect.TypeOf((*MockTerraformChartOps)(nil).ComputeChartValues), arg0, arg1, arg2)
}

// ComputeCreateVPCInitializerValues mocks base method
func (m *MockTerraformChartOps) ComputeCreateVPCInitializerValues(arg0 *v1alpha1.InfrastructureConfig, arg1 string) *infrastructure.InitializerValues {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComputeCreateVPCInitializerValues", arg0, arg1)
	ret0, _ := ret[0].(*infrastructure.InitializerValues)
	return ret0
}

// ComputeCreateVPCInitializerValues indicates an expected call of ComputeCreateVPCInitializerValues
func (mr *MockTerraformChartOpsMockRecorder) ComputeCreateVPCInitializerValues(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComputeCreateVPCInitializerValues", reflect.TypeOf((*MockTerraformChartOps)(nil).ComputeCreateVPCInitializerValues), arg0, arg1)
}

// ComputeUseVPCInitializerValues mocks base method
func (m *MockTerraformChartOps) ComputeUseVPCInitializerValues(arg0 *v1alpha1.InfrastructureConfig, arg1 *infrastructure.VPCInfo) *infrastructure.InitializerValues {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComputeUseVPCInitializerValues", arg0, arg1)
	ret0, _ := ret[0].(*infrastructure.InitializerValues)
	return ret0
}

// ComputeUseVPCInitializerValues indicates an expected call of ComputeUseVPCInitializerValues
func (mr *MockTerraformChartOpsMockRecorder) ComputeUseVPCInitializerValues(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComputeUseVPCInitializerValues", reflect.TypeOf((*MockTerraformChartOps)(nil).ComputeUseVPCInitializerValues), arg0, arg1)
}
