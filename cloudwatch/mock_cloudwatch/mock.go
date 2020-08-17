// Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-cloudwatch-logs-for-fluent-bit/cloudwatch (interfaces: LogsClient)

// Package mock_cloudwatch is a generated GoMock package.
package mock_cloudwatch

import (
	reflect "reflect"

	cloudwatchlogs "github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	gomock "github.com/golang/mock/gomock"
)

// MockLogsClient is a mock of LogsClient interface
type MockLogsClient struct {
	ctrl     *gomock.Controller
	recorder *MockLogsClientMockRecorder
}

// MockLogsClientMockRecorder is the mock recorder for MockLogsClient
type MockLogsClientMockRecorder struct {
	mock *MockLogsClient
}

// NewMockLogsClient creates a new mock instance
func NewMockLogsClient(ctrl *gomock.Controller) *MockLogsClient {
	mock := &MockLogsClient{ctrl: ctrl}
	mock.recorder = &MockLogsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLogsClient) EXPECT() *MockLogsClientMockRecorder {
	return m.recorder
}

// CreateLogGroup mocks base method
func (m *MockLogsClient) CreateLogGroup(arg0 *cloudwatchlogs.CreateLogGroupInput) (*cloudwatchlogs.CreateLogGroupOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLogGroup", arg0)
	ret0, _ := ret[0].(*cloudwatchlogs.CreateLogGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLogGroup indicates an expected call of CreateLogGroup
func (mr *MockLogsClientMockRecorder) CreateLogGroup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLogGroup", reflect.TypeOf((*MockLogsClient)(nil).CreateLogGroup), arg0)
}

// CreateLogStream mocks base method
func (m *MockLogsClient) CreateLogStream(arg0 *cloudwatchlogs.CreateLogStreamInput) (*cloudwatchlogs.CreateLogStreamOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLogStream", arg0)
	ret0, _ := ret[0].(*cloudwatchlogs.CreateLogStreamOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLogStream indicates an expected call of CreateLogStream
func (mr *MockLogsClientMockRecorder) CreateLogStream(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLogStream", reflect.TypeOf((*MockLogsClient)(nil).CreateLogStream), arg0)
}

// DescribeLogStreams mocks base method
func (m *MockLogsClient) DescribeLogStreams(arg0 *cloudwatchlogs.DescribeLogStreamsInput) (*cloudwatchlogs.DescribeLogStreamsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeLogStreams", arg0)
	ret0, _ := ret[0].(*cloudwatchlogs.DescribeLogStreamsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeLogStreams indicates an expected call of DescribeLogStreams
func (mr *MockLogsClientMockRecorder) DescribeLogStreams(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeLogStreams", reflect.TypeOf((*MockLogsClient)(nil).DescribeLogStreams), arg0)
}

// PutLogEvents mocks base method
func (m *MockLogsClient) PutLogEvents(arg0 *cloudwatchlogs.PutLogEventsInput) (*cloudwatchlogs.PutLogEventsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutLogEvents", arg0)
	ret0, _ := ret[0].(*cloudwatchlogs.PutLogEventsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutLogEvents indicates an expected call of PutLogEvents
func (mr *MockLogsClientMockRecorder) PutLogEvents(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutLogEvents", reflect.TypeOf((*MockLogsClient)(nil).PutLogEvents), arg0)
}

// PutRetentionPolicy mocks base method
func (m *MockLogsClient) PutRetentionPolicy(arg0 *cloudwatchlogs.PutRetentionPolicyInput) (*cloudwatchlogs.PutRetentionPolicyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutRetentionPolicy", arg0)
	ret0, _ := ret[0].(*cloudwatchlogs.PutRetentionPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutRetentionPolicy indicates an expected call of PutRetentionPolicy
func (mr *MockLogsClientMockRecorder) PutRetentionPolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutRetentionPolicy", reflect.TypeOf((*MockLogsClient)(nil).PutRetentionPolicy), arg0)
}
