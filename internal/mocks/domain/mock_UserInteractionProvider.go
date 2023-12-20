// SPDX-FileCopyrightText: 2023 INDITEX S.A
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by mockery v2.32.4. DO NOT EDIT.

package domain

import mock "github.com/stretchr/testify/mock"

// MockUserInteractionProvider is an autogenerated mock type for the UserInteractionProvider type
type MockUserInteractionProvider struct {
	mock.Mock
}

type MockUserInteractionProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUserInteractionProvider) EXPECT() *MockUserInteractionProvider_Expecter {
	return &MockUserInteractionProvider_Expecter{mock: &_m.Mock}
}

// AskUserForConfirmation provides a mock function with given fields: msg, defaultAnswer
func (_m *MockUserInteractionProvider) AskUserForConfirmation(msg string, defaultAnswer bool) (bool, error) {
	ret := _m.Called(msg, defaultAnswer)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string, bool) (bool, error)); ok {
		return rf(msg, defaultAnswer)
	}
	if rf, ok := ret.Get(0).(func(string, bool) bool); ok {
		r0 = rf(msg, defaultAnswer)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string, bool) error); ok {
		r1 = rf(msg, defaultAnswer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserInteractionProvider_AskUserForConfirmation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AskUserForConfirmation'
type MockUserInteractionProvider_AskUserForConfirmation_Call struct {
	*mock.Call
}

// AskUserForConfirmation is a helper method to define mock.On call
//   - msg string
//   - defaultAnswer bool
func (_e *MockUserInteractionProvider_Expecter) AskUserForConfirmation(msg interface{}, defaultAnswer interface{}) *MockUserInteractionProvider_AskUserForConfirmation_Call {
	return &MockUserInteractionProvider_AskUserForConfirmation_Call{Call: _e.mock.On("AskUserForConfirmation", msg, defaultAnswer)}
}

func (_c *MockUserInteractionProvider_AskUserForConfirmation_Call) Run(run func(msg string, defaultAnswer bool)) *MockUserInteractionProvider_AskUserForConfirmation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(bool))
	})
	return _c
}

func (_c *MockUserInteractionProvider_AskUserForConfirmation_Call) Return(answer bool, err error) *MockUserInteractionProvider_AskUserForConfirmation_Call {
	_c.Call.Return(answer, err)
	return _c
}

func (_c *MockUserInteractionProvider_AskUserForConfirmation_Call) RunAndReturn(run func(string, bool) (bool, error)) *MockUserInteractionProvider_AskUserForConfirmation_Call {
	_c.Call.Return(run)
	return _c
}

// SelectOrInput provides a mock function with given fields: name, validValues, variable, required
func (_m *MockUserInteractionProvider) SelectOrInput(name string, validValues []string, variable *string, required bool) error {
	ret := _m.Called(name, validValues, variable, required)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []string, *string, bool) error); ok {
		r0 = rf(name, validValues, variable, required)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUserInteractionProvider_SelectOrInput_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SelectOrInput'
type MockUserInteractionProvider_SelectOrInput_Call struct {
	*mock.Call
}

// SelectOrInput is a helper method to define mock.On call
//   - name string
//   - validValues []string
//   - variable *string
//   - required bool
func (_e *MockUserInteractionProvider_Expecter) SelectOrInput(name interface{}, validValues interface{}, variable interface{}, required interface{}) *MockUserInteractionProvider_SelectOrInput_Call {
	return &MockUserInteractionProvider_SelectOrInput_Call{Call: _e.mock.On("SelectOrInput", name, validValues, variable, required)}
}

func (_c *MockUserInteractionProvider_SelectOrInput_Call) Run(run func(name string, validValues []string, variable *string, required bool)) *MockUserInteractionProvider_SelectOrInput_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].([]string), args[2].(*string), args[3].(bool))
	})
	return _c
}

func (_c *MockUserInteractionProvider_SelectOrInput_Call) Return(_a0 error) *MockUserInteractionProvider_SelectOrInput_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUserInteractionProvider_SelectOrInput_Call) RunAndReturn(run func(string, []string, *string, bool) error) *MockUserInteractionProvider_SelectOrInput_Call {
	_c.Call.Return(run)
	return _c
}

// SelectOrInputPrompt provides a mock function with given fields: message, validValues, variable, required
func (_m *MockUserInteractionProvider) SelectOrInputPrompt(message string, validValues []string, variable *string, required bool) error {
	ret := _m.Called(message, validValues, variable, required)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []string, *string, bool) error); ok {
		r0 = rf(message, validValues, variable, required)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUserInteractionProvider_SelectOrInputPrompt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SelectOrInputPrompt'
type MockUserInteractionProvider_SelectOrInputPrompt_Call struct {
	*mock.Call
}

// SelectOrInputPrompt is a helper method to define mock.On call
//   - message string
//   - validValues []string
//   - variable *string
//   - required bool
func (_e *MockUserInteractionProvider_Expecter) SelectOrInputPrompt(message interface{}, validValues interface{}, variable interface{}, required interface{}) *MockUserInteractionProvider_SelectOrInputPrompt_Call {
	return &MockUserInteractionProvider_SelectOrInputPrompt_Call{Call: _e.mock.On("SelectOrInputPrompt", message, validValues, variable, required)}
}

func (_c *MockUserInteractionProvider_SelectOrInputPrompt_Call) Run(run func(message string, validValues []string, variable *string, required bool)) *MockUserInteractionProvider_SelectOrInputPrompt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].([]string), args[2].(*string), args[3].(bool))
	})
	return _c
}

func (_c *MockUserInteractionProvider_SelectOrInputPrompt_Call) Return(_a0 error) *MockUserInteractionProvider_SelectOrInputPrompt_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUserInteractionProvider_SelectOrInputPrompt_Call) RunAndReturn(run func(string, []string, *string, bool) error) *MockUserInteractionProvider_SelectOrInputPrompt_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUserInteractionProvider creates a new instance of MockUserInteractionProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUserInteractionProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUserInteractionProvider {
	mock := &MockUserInteractionProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
