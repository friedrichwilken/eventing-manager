// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	context "context"

	handler "sigs.k8s.io/controller-runtime/pkg/handler"

	logr "github.com/go-logr/logr"

	mock "github.com/stretchr/testify/mock"

	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"

	reconcile "sigs.k8s.io/controller-runtime/pkg/reconcile"

	source "sigs.k8s.io/controller-runtime/pkg/source"
)

// Controller is an autogenerated mock type for the Controller type
type Controller struct {
	mock.Mock
}

type Controller_Expecter struct {
	mock *mock.Mock
}

func (_m *Controller) EXPECT() *Controller_Expecter {
	return &Controller_Expecter{mock: &_m.Mock}
}

// GetLogger provides a mock function with given fields:
func (_m *Controller) GetLogger() logr.Logger {
	ret := _m.Called()

	var r0 logr.Logger
	if rf, ok := ret.Get(0).(func() logr.Logger); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(logr.Logger)
	}

	return r0
}

// Controller_GetLogger_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLogger'
type Controller_GetLogger_Call struct {
	*mock.Call
}

// GetLogger is a helper method to define mock.On call
func (_e *Controller_Expecter) GetLogger() *Controller_GetLogger_Call {
	return &Controller_GetLogger_Call{Call: _e.mock.On("GetLogger")}
}

func (_c *Controller_GetLogger_Call) Run(run func()) *Controller_GetLogger_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Controller_GetLogger_Call) Return(_a0 logr.Logger) *Controller_GetLogger_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Controller_GetLogger_Call) RunAndReturn(run func() logr.Logger) *Controller_GetLogger_Call {
	_c.Call.Return(run)
	return _c
}

// Reconcile provides a mock function with given fields: _a0, _a1
func (_m *Controller) Reconcile(_a0 context.Context, _a1 reconcile.Request) (reconcile.Result, error) {
	ret := _m.Called(_a0, _a1)

	var r0 reconcile.Result
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, reconcile.Request) (reconcile.Result, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, reconcile.Request) reconcile.Result); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(reconcile.Result)
	}

	if rf, ok := ret.Get(1).(func(context.Context, reconcile.Request) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Controller_Reconcile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Reconcile'
type Controller_Reconcile_Call struct {
	*mock.Call
}

// Reconcile is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 reconcile.Request
func (_e *Controller_Expecter) Reconcile(_a0 interface{}, _a1 interface{}) *Controller_Reconcile_Call {
	return &Controller_Reconcile_Call{Call: _e.mock.On("Reconcile", _a0, _a1)}
}

func (_c *Controller_Reconcile_Call) Run(run func(_a0 context.Context, _a1 reconcile.Request)) *Controller_Reconcile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(reconcile.Request))
	})
	return _c
}

func (_c *Controller_Reconcile_Call) Return(_a0 reconcile.Result, _a1 error) *Controller_Reconcile_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Controller_Reconcile_Call) RunAndReturn(run func(context.Context, reconcile.Request) (reconcile.Result, error)) *Controller_Reconcile_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields: ctx
func (_m *Controller) Start(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Controller_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type Controller_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Controller_Expecter) Start(ctx interface{}) *Controller_Start_Call {
	return &Controller_Start_Call{Call: _e.mock.On("Start", ctx)}
}

func (_c *Controller_Start_Call) Run(run func(ctx context.Context)) *Controller_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Controller_Start_Call) Return(_a0 error) *Controller_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Controller_Start_Call) RunAndReturn(run func(context.Context) error) *Controller_Start_Call {
	_c.Call.Return(run)
	return _c
}

// Watch provides a mock function with given fields: src, eventhandler, predicates
func (_m *Controller) Watch(src source.Source, eventhandler handler.EventHandler, predicates ...predicate.Predicate) error {
	_va := make([]interface{}, len(predicates))
	for _i := range predicates {
		_va[_i] = predicates[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, src, eventhandler)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(source.Source, handler.EventHandler, ...predicate.Predicate) error); ok {
		r0 = rf(src, eventhandler, predicates...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Controller_Watch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Watch'
type Controller_Watch_Call struct {
	*mock.Call
}

// Watch is a helper method to define mock.On call
//   - src source.Source
//   - eventhandler handler.EventHandler
//   - predicates ...predicate.Predicate
func (_e *Controller_Expecter) Watch(src interface{}, eventhandler interface{}, predicates ...interface{}) *Controller_Watch_Call {
	return &Controller_Watch_Call{Call: _e.mock.On("Watch",
		append([]interface{}{src, eventhandler}, predicates...)...)}
}

func (_c *Controller_Watch_Call) Run(run func(src source.Source, eventhandler handler.EventHandler, predicates ...predicate.Predicate)) *Controller_Watch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]predicate.Predicate, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(predicate.Predicate)
			}
		}
		run(args[0].(source.Source), args[1].(handler.EventHandler), variadicArgs...)
	})
	return _c
}

func (_c *Controller_Watch_Call) Return(_a0 error) *Controller_Watch_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Controller_Watch_Call) RunAndReturn(run func(source.Source, handler.EventHandler, ...predicate.Predicate) error) *Controller_Watch_Call {
	_c.Call.Return(run)
	return _c
}

// NewController creates a new instance of Controller. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewController(t interface {
	mock.TestingT
	Cleanup(func())
}) *Controller {
	mock := &Controller{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}