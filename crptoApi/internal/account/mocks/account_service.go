// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// AccountService is an autogenerated mock type for the AccountService type
type AccountService struct {
	mock.Mock
}

type AccountService_Expecter struct {
	mock *mock.Mock
}

func (_m *AccountService) EXPECT() *AccountService_Expecter {
	return &AccountService_Expecter{mock: &_m.Mock}
}

// GetAccountBalance provides a mock function with no fields
func (_m *AccountService) GetAccountBalance() (float64, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAccountBalance")
	}

	var r0 float64
	var r1 error
	if rf, ok := ret.Get(0).(func() (float64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() float64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float64)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AccountService_GetAccountBalance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAccountBalance'
type AccountService_GetAccountBalance_Call struct {
	*mock.Call
}

// GetAccountBalance is a helper method to define mock.On call
func (_e *AccountService_Expecter) GetAccountBalance() *AccountService_GetAccountBalance_Call {
	return &AccountService_GetAccountBalance_Call{Call: _e.mock.On("GetAccountBalance")}
}

func (_c *AccountService_GetAccountBalance_Call) Run(run func()) *AccountService_GetAccountBalance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *AccountService_GetAccountBalance_Call) Return(_a0 float64, _a1 error) *AccountService_GetAccountBalance_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AccountService_GetAccountBalance_Call) RunAndReturn(run func() (float64, error)) *AccountService_GetAccountBalance_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateAccountBalance provides a mock function with given fields: amount
func (_m *AccountService) UpdateAccountBalance(amount float64) error {
	ret := _m.Called(amount)

	if len(ret) == 0 {
		panic("no return value specified for UpdateAccountBalance")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(float64) error); ok {
		r0 = rf(amount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AccountService_UpdateAccountBalance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateAccountBalance'
type AccountService_UpdateAccountBalance_Call struct {
	*mock.Call
}

// UpdateAccountBalance is a helper method to define mock.On call
//   - amount float64
func (_e *AccountService_Expecter) UpdateAccountBalance(amount interface{}) *AccountService_UpdateAccountBalance_Call {
	return &AccountService_UpdateAccountBalance_Call{Call: _e.mock.On("UpdateAccountBalance", amount)}
}

func (_c *AccountService_UpdateAccountBalance_Call) Run(run func(amount float64)) *AccountService_UpdateAccountBalance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(float64))
	})
	return _c
}

func (_c *AccountService_UpdateAccountBalance_Call) Return(_a0 error) *AccountService_UpdateAccountBalance_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AccountService_UpdateAccountBalance_Call) RunAndReturn(run func(float64) error) *AccountService_UpdateAccountBalance_Call {
	_c.Call.Return(run)
	return _c
}

// NewAccountService creates a new instance of AccountService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAccountService(t interface {
	mock.TestingT
	Cleanup(func())
}) *AccountService {
	mock := &AccountService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
