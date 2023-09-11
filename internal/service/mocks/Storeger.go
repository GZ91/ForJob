// Code generated by mockery v2.33.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Storeger is an autogenerated mock type for the Storeger type
type Storeger struct {
	mock.Mock
}

type Storeger_Expecter struct {
	mock *mock.Mock
}

func (_m *Storeger) EXPECT() *Storeger_Expecter {
	return &Storeger_Expecter{mock: &_m.Mock}
}

// AddURL provides a mock function with given fields: _a0, _a1
func (_m *Storeger) AddURL(_a0 context.Context, _a1 string) (string, error) {
	ret := _m.Called(_a0, _a1)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Storeger_AddURL_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddURL'
type Storeger_AddURL_Call struct {
	*mock.Call
}

// AddURL is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *Storeger_Expecter) AddURL(_a0 interface{}, _a1 interface{}) *Storeger_AddURL_Call {
	return &Storeger_AddURL_Call{Call: _e.mock.On("AddURL", _a0, _a1)}
}

func (_c *Storeger_AddURL_Call) Run(run func(_a0 context.Context, _a1 string)) *Storeger_AddURL_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Storeger_AddURL_Call) Return(_a0 string, _a1 error) *Storeger_AddURL_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Storeger_AddURL_Call) RunAndReturn(run func(context.Context, string) (string, error)) *Storeger_AddURL_Call {
	_c.Call.Return(run)
	return _c
}

// CheckToken provides a mock function with given fields: ctx, token
func (_m *Storeger) CheckToken(ctx context.Context, token string) (bool, error) {
	ret := _m.Called(ctx, token)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (bool, error)); ok {
		return rf(ctx, token)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, token)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Storeger_CheckToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckToken'
type Storeger_CheckToken_Call struct {
	*mock.Call
}

// CheckToken is a helper method to define mock.On call
//   - ctx context.Context
//   - token string
func (_e *Storeger_Expecter) CheckToken(ctx interface{}, token interface{}) *Storeger_CheckToken_Call {
	return &Storeger_CheckToken_Call{Call: _e.mock.On("CheckToken", ctx, token)}
}

func (_c *Storeger_CheckToken_Call) Run(run func(ctx context.Context, token string)) *Storeger_CheckToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Storeger_CheckToken_Call) Return(_a0 bool, _a1 error) *Storeger_CheckToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Storeger_CheckToken_Call) RunAndReturn(run func(context.Context, string) (bool, error)) *Storeger_CheckToken_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteLinkByLongLink provides a mock function with given fields: ctx, longLink, token
func (_m *Storeger) DeleteLinkByLongLink(ctx context.Context, longLink string, token string) error {
	ret := _m.Called(ctx, longLink, token)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, longLink, token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Storeger_DeleteLinkByLongLink_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteLinkByLongLink'
type Storeger_DeleteLinkByLongLink_Call struct {
	*mock.Call
}

// DeleteLinkByLongLink is a helper method to define mock.On call
//   - ctx context.Context
//   - longLink string
//   - token string
func (_e *Storeger_Expecter) DeleteLinkByLongLink(ctx interface{}, longLink interface{}, token interface{}) *Storeger_DeleteLinkByLongLink_Call {
	return &Storeger_DeleteLinkByLongLink_Call{Call: _e.mock.On("DeleteLinkByLongLink", ctx, longLink, token)}
}

func (_c *Storeger_DeleteLinkByLongLink_Call) Run(run func(ctx context.Context, longLink string, token string)) *Storeger_DeleteLinkByLongLink_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *Storeger_DeleteLinkByLongLink_Call) Return(_a0 error) *Storeger_DeleteLinkByLongLink_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Storeger_DeleteLinkByLongLink_Call) RunAndReturn(run func(context.Context, string, string) error) *Storeger_DeleteLinkByLongLink_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteLinkByShortLink provides a mock function with given fields: ctx, shortLink, token
func (_m *Storeger) DeleteLinkByShortLink(ctx context.Context, shortLink string, token string) error {
	ret := _m.Called(ctx, shortLink, token)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, shortLink, token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Storeger_DeleteLinkByShortLink_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteLinkByShortLink'
type Storeger_DeleteLinkByShortLink_Call struct {
	*mock.Call
}

// DeleteLinkByShortLink is a helper method to define mock.On call
//   - ctx context.Context
//   - shortLink string
//   - token string
func (_e *Storeger_Expecter) DeleteLinkByShortLink(ctx interface{}, shortLink interface{}, token interface{}) *Storeger_DeleteLinkByShortLink_Call {
	return &Storeger_DeleteLinkByShortLink_Call{Call: _e.mock.On("DeleteLinkByShortLink", ctx, shortLink, token)}
}

func (_c *Storeger_DeleteLinkByShortLink_Call) Run(run func(ctx context.Context, shortLink string, token string)) *Storeger_DeleteLinkByShortLink_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *Storeger_DeleteLinkByShortLink_Call) Return(_a0 error) *Storeger_DeleteLinkByShortLink_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Storeger_DeleteLinkByShortLink_Call) RunAndReturn(run func(context.Context, string, string) error) *Storeger_DeleteLinkByShortLink_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteToken provides a mock function with given fields: ctx, token
func (_m *Storeger) DeleteToken(ctx context.Context, token string) error {
	ret := _m.Called(ctx, token)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Storeger_DeleteToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteToken'
type Storeger_DeleteToken_Call struct {
	*mock.Call
}

// DeleteToken is a helper method to define mock.On call
//   - ctx context.Context
//   - token string
func (_e *Storeger_Expecter) DeleteToken(ctx interface{}, token interface{}) *Storeger_DeleteToken_Call {
	return &Storeger_DeleteToken_Call{Call: _e.mock.On("DeleteToken", ctx, token)}
}

func (_c *Storeger_DeleteToken_Call) Run(run func(ctx context.Context, token string)) *Storeger_DeleteToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Storeger_DeleteToken_Call) Return(_a0 error) *Storeger_DeleteToken_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Storeger_DeleteToken_Call) RunAndReturn(run func(context.Context, string) error) *Storeger_DeleteToken_Call {
	_c.Call.Return(run)
	return _c
}

// FindLongURL provides a mock function with given fields: _a0, _a1, _a2
func (_m *Storeger) FindLongURL(_a0 context.Context, _a1 string, _a2 string) (string, bool, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 string
	var r1 bool
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (string, bool, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) string); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) bool); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Get(1).(bool)
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, string) error); ok {
		r2 = rf(_a0, _a1, _a2)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Storeger_FindLongURL_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindLongURL'
type Storeger_FindLongURL_Call struct {
	*mock.Call
}

// FindLongURL is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
//   - _a2 string
func (_e *Storeger_Expecter) FindLongURL(_a0 interface{}, _a1 interface{}, _a2 interface{}) *Storeger_FindLongURL_Call {
	return &Storeger_FindLongURL_Call{Call: _e.mock.On("FindLongURL", _a0, _a1, _a2)}
}

func (_c *Storeger_FindLongURL_Call) Run(run func(_a0 context.Context, _a1 string, _a2 string)) *Storeger_FindLongURL_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *Storeger_FindLongURL_Call) Return(_a0 string, _a1 bool, _a2 error) *Storeger_FindLongURL_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *Storeger_FindLongURL_Call) RunAndReturn(run func(context.Context, string, string) (string, bool, error)) *Storeger_FindLongURL_Call {
	_c.Call.Return(run)
	return _c
}

// GetLinks provides a mock function with given fields: ctx, token
func (_m *Storeger) GetLinks(ctx context.Context, token string) (map[string]string, error) {
	ret := _m.Called(ctx, token)

	var r0 map[string]string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (map[string]string, error)); ok {
		return rf(ctx, token)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) map[string]string); ok {
		r0 = rf(ctx, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Storeger_GetLinks_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLinks'
type Storeger_GetLinks_Call struct {
	*mock.Call
}

// GetLinks is a helper method to define mock.On call
//   - ctx context.Context
//   - token string
func (_e *Storeger_Expecter) GetLinks(ctx interface{}, token interface{}) *Storeger_GetLinks_Call {
	return &Storeger_GetLinks_Call{Call: _e.mock.On("GetLinks", ctx, token)}
}

func (_c *Storeger_GetLinks_Call) Run(run func(ctx context.Context, token string)) *Storeger_GetLinks_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Storeger_GetLinks_Call) Return(_a0 map[string]string, _a1 error) *Storeger_GetLinks_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Storeger_GetLinks_Call) RunAndReturn(run func(context.Context, string) (map[string]string, error)) *Storeger_GetLinks_Call {
	_c.Call.Return(run)
	return _c
}

// GetServices provides a mock function with given fields: ctx, name
func (_m *Storeger) GetServices(ctx context.Context, name string) (map[string]string, error) {
	ret := _m.Called(ctx, name)

	var r0 map[string]string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (map[string]string, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) map[string]string); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Storeger_GetServices_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetServices'
type Storeger_GetServices_Call struct {
	*mock.Call
}

// GetServices is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *Storeger_Expecter) GetServices(ctx interface{}, name interface{}) *Storeger_GetServices_Call {
	return &Storeger_GetServices_Call{Call: _e.mock.On("GetServices", ctx, name)}
}

func (_c *Storeger_GetServices_Call) Run(run func(ctx context.Context, name string)) *Storeger_GetServices_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Storeger_GetServices_Call) Return(_a0 map[string]string, _a1 error) *Storeger_GetServices_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Storeger_GetServices_Call) RunAndReturn(run func(context.Context, string) (map[string]string, error)) *Storeger_GetServices_Call {
	_c.Call.Return(run)
	return _c
}

// GetTokens provides a mock function with given fields: ctx, namesServices
func (_m *Storeger) GetTokens(ctx context.Context, namesServices []string) (map[string]string, error) {
	ret := _m.Called(ctx, namesServices)

	var r0 map[string]string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []string) (map[string]string, error)); ok {
		return rf(ctx, namesServices)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string) map[string]string); ok {
		r0 = rf(ctx, namesServices)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(ctx, namesServices)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Storeger_GetTokens_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTokens'
type Storeger_GetTokens_Call struct {
	*mock.Call
}

// GetTokens is a helper method to define mock.On call
//   - ctx context.Context
//   - namesServices []string
func (_e *Storeger_Expecter) GetTokens(ctx interface{}, namesServices interface{}) *Storeger_GetTokens_Call {
	return &Storeger_GetTokens_Call{Call: _e.mock.On("GetTokens", ctx, namesServices)}
}

func (_c *Storeger_GetTokens_Call) Run(run func(ctx context.Context, namesServices []string)) *Storeger_GetTokens_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]string))
	})
	return _c
}

func (_c *Storeger_GetTokens_Call) Return(_a0 map[string]string, _a1 error) *Storeger_GetTokens_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Storeger_GetTokens_Call) RunAndReturn(run func(context.Context, []string) (map[string]string, error)) *Storeger_GetTokens_Call {
	_c.Call.Return(run)
	return _c
}

// GetURL provides a mock function with given fields: _a0, _a1
func (_m *Storeger) GetURL(_a0 context.Context, _a1 string) (string, bool, error) {
	ret := _m.Called(_a0, _a1)

	var r0 string
	var r1 bool
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, bool, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) bool); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Get(1).(bool)
	}

	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(_a0, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Storeger_GetURL_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetURL'
type Storeger_GetURL_Call struct {
	*mock.Call
}

// GetURL is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *Storeger_Expecter) GetURL(_a0 interface{}, _a1 interface{}) *Storeger_GetURL_Call {
	return &Storeger_GetURL_Call{Call: _e.mock.On("GetURL", _a0, _a1)}
}

func (_c *Storeger_GetURL_Call) Run(run func(_a0 context.Context, _a1 string)) *Storeger_GetURL_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Storeger_GetURL_Call) Return(_a0 string, _a1 bool, _a2 error) *Storeger_GetURL_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *Storeger_GetURL_Call) RunAndReturn(run func(context.Context, string) (string, bool, error)) *Storeger_GetURL_Call {
	_c.Call.Return(run)
	return _c
}

// Ping provides a mock function with given fields: _a0
func (_m *Storeger) Ping(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Storeger_Ping_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Ping'
type Storeger_Ping_Call struct {
	*mock.Call
}

// Ping is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *Storeger_Expecter) Ping(_a0 interface{}) *Storeger_Ping_Call {
	return &Storeger_Ping_Call{Call: _e.mock.On("Ping", _a0)}
}

func (_c *Storeger_Ping_Call) Run(run func(_a0 context.Context)) *Storeger_Ping_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Storeger_Ping_Call) Return(_a0 error) *Storeger_Ping_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Storeger_Ping_Call) RunAndReturn(run func(context.Context) error) *Storeger_Ping_Call {
	_c.Call.Return(run)
	return _c
}

// NewStoreger creates a new instance of Storeger. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStoreger(t interface {
	mock.TestingT
	Cleanup(func())
}) *Storeger {
	mock := &Storeger{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
