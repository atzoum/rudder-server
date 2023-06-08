// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rudderlabs/rudder-server/jobsdb (interfaces: MultiTenantJobsDB)

// Package mocks_jobsdb is a generated GoMock package.
package mocks_jobsdb

import (
	context "context"
	json "encoding/json"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	jobsdb "github.com/rudderlabs/rudder-server/jobsdb"
)

// MockMultiTenantJobsDB is a mock of MultiTenantJobsDB interface.
type MockMultiTenantJobsDB struct {
	ctrl     *gomock.Controller
	recorder *MockMultiTenantJobsDBMockRecorder
}

// MockMultiTenantJobsDBMockRecorder is the mock recorder for MockMultiTenantJobsDB.
type MockMultiTenantJobsDBMockRecorder struct {
	mock *MockMultiTenantJobsDB
}

// NewMockMultiTenantJobsDB creates a new mock instance.
func NewMockMultiTenantJobsDB(ctrl *gomock.Controller) *MockMultiTenantJobsDB {
	mock := &MockMultiTenantJobsDB{ctrl: ctrl}
	mock.recorder = &MockMultiTenantJobsDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMultiTenantJobsDB) EXPECT() *MockMultiTenantJobsDBMockRecorder {
	return m.recorder
}

// DeleteExecuting mocks base method.
func (m *MockMultiTenantJobsDB) DeleteExecuting() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteExecuting")
}

// DeleteExecuting indicates an expected call of DeleteExecuting.
func (mr *MockMultiTenantJobsDBMockRecorder) DeleteExecuting() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExecuting", reflect.TypeOf((*MockMultiTenantJobsDB)(nil).DeleteExecuting))
}

// FailExecuting mocks base method.
func (m *MockMultiTenantJobsDB) FailExecuting() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FailExecuting")
}

// FailExecuting indicates an expected call of FailExecuting.
func (mr *MockMultiTenantJobsDBMockRecorder) FailExecuting() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FailExecuting", reflect.TypeOf((*MockMultiTenantJobsDB)(nil).FailExecuting))
}

// GetActiveWorkspaces mocks base method.
func (m *MockMultiTenantJobsDB) GetActiveWorkspaces(arg0 context.Context, arg1 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActiveWorkspaces", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActiveWorkspaces indicates an expected call of GetActiveWorkspaces.
func (mr *MockMultiTenantJobsDBMockRecorder) GetActiveWorkspaces(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActiveWorkspaces", reflect.TypeOf((*MockMultiTenantJobsDB)(nil).GetActiveWorkspaces), arg0, arg1)
}

// GetAllJobs mocks base method.
func (m *MockMultiTenantJobsDB) GetAllJobs(arg0 context.Context, arg1 map[string]int, arg2 jobsdb.GetQueryParamsT, arg3 int, arg4 jobsdb.MoreToken) (*jobsdb.GetAllJobsResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllJobs", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*jobsdb.GetAllJobsResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllJobs indicates an expected call of GetAllJobs.
func (mr *MockMultiTenantJobsDBMockRecorder) GetAllJobs(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllJobs", reflect.TypeOf((*MockMultiTenantJobsDB)(nil).GetAllJobs), arg0, arg1, arg2, arg3, arg4)
}

// GetDistinctParameterValues mocks base method.
func (m *MockMultiTenantJobsDB) GetDistinctParameterValues(arg0 context.Context, arg1 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDistinctParameterValues", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDistinctParameterValues indicates an expected call of GetDistinctParameterValues.
func (mr *MockMultiTenantJobsDBMockRecorder) GetDistinctParameterValues(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDistinctParameterValues", reflect.TypeOf((*MockMultiTenantJobsDB)(nil).GetDistinctParameterValues), arg0, arg1)
}

// GetJournalEntries mocks base method.
func (m *MockMultiTenantJobsDB) GetJournalEntries(arg0 string) []jobsdb.JournalEntryT {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJournalEntries", arg0)
	ret0, _ := ret[0].([]jobsdb.JournalEntryT)
	return ret0
}

// GetJournalEntries indicates an expected call of GetJournalEntries.
func (mr *MockMultiTenantJobsDBMockRecorder) GetJournalEntries(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJournalEntries", reflect.TypeOf((*MockMultiTenantJobsDB)(nil).GetJournalEntries), arg0)
}

// GetPileUpCounts mocks base method.
func (m *MockMultiTenantJobsDB) GetPileUpCounts(arg0 context.Context) (map[string]map[string]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPileUpCounts", arg0)
	ret0, _ := ret[0].(map[string]map[string]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPileUpCounts indicates an expected call of GetPileUpCounts.
func (mr *MockMultiTenantJobsDBMockRecorder) GetPileUpCounts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPileUpCounts", reflect.TypeOf((*MockMultiTenantJobsDB)(nil).GetPileUpCounts), arg0)
}

// JournalDeleteEntry mocks base method.
func (m *MockMultiTenantJobsDB) JournalDeleteEntry(arg0 int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "JournalDeleteEntry", arg0)
}

// JournalDeleteEntry indicates an expected call of JournalDeleteEntry.
func (mr *MockMultiTenantJobsDBMockRecorder) JournalDeleteEntry(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JournalDeleteEntry", reflect.TypeOf((*MockMultiTenantJobsDB)(nil).JournalDeleteEntry), arg0)
}

// JournalMarkStart mocks base method.
func (m *MockMultiTenantJobsDB) JournalMarkStart(arg0 string, arg1 json.RawMessage) int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "JournalMarkStart", arg0, arg1)
	ret0, _ := ret[0].(int64)
	return ret0
}

// JournalMarkStart indicates an expected call of JournalMarkStart.
func (mr *MockMultiTenantJobsDBMockRecorder) JournalMarkStart(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JournalMarkStart", reflect.TypeOf((*MockMultiTenantJobsDB)(nil).JournalMarkStart), arg0, arg1)
}

// UpdateJobStatus mocks base method.
func (m *MockMultiTenantJobsDB) UpdateJobStatus(arg0 context.Context, arg1 []*jobsdb.JobStatusT, arg2 []string, arg3 []jobsdb.ParameterFilterT) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateJobStatus", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateJobStatus indicates an expected call of UpdateJobStatus.
func (mr *MockMultiTenantJobsDBMockRecorder) UpdateJobStatus(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateJobStatus", reflect.TypeOf((*MockMultiTenantJobsDB)(nil).UpdateJobStatus), arg0, arg1, arg2, arg3)
}

// UpdateJobStatusInTx mocks base method.
func (m *MockMultiTenantJobsDB) UpdateJobStatusInTx(arg0 context.Context, arg1 jobsdb.UpdateSafeTx, arg2 []*jobsdb.JobStatusT, arg3 []string, arg4 []jobsdb.ParameterFilterT) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateJobStatusInTx", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateJobStatusInTx indicates an expected call of UpdateJobStatusInTx.
func (mr *MockMultiTenantJobsDBMockRecorder) UpdateJobStatusInTx(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateJobStatusInTx", reflect.TypeOf((*MockMultiTenantJobsDB)(nil).UpdateJobStatusInTx), arg0, arg1, arg2, arg3, arg4)
}

// WithUpdateSafeTx mocks base method.
func (m *MockMultiTenantJobsDB) WithUpdateSafeTx(arg0 context.Context, arg1 func(jobsdb.UpdateSafeTx) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithUpdateSafeTx", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WithUpdateSafeTx indicates an expected call of WithUpdateSafeTx.
func (mr *MockMultiTenantJobsDBMockRecorder) WithUpdateSafeTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithUpdateSafeTx", reflect.TypeOf((*MockMultiTenantJobsDB)(nil).WithUpdateSafeTx), arg0, arg1)
}
