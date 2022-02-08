// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"
	mutants "mutant-ms/models/mutants"

	mock "github.com/stretchr/testify/mock"
)

// Repositories is an autogenerated mock type for the Repositories type
type Repositories struct {
	mock.Mock
}

// GetAll provides a mock function with given fields: ctx
func (_m *Repositories) GetAll(ctx context.Context) ([]mutants.MutantDna, error) {
	ret := _m.Called(ctx)

	var r0 []mutants.MutantDna
	if rf, ok := ret.Get(0).(func(context.Context) []mutants.MutantDna); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]mutants.MutantDna)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, mutantDna
func (_m *Repositories) Save(ctx context.Context, mutantDna mutants.MutantDna) error {
	ret := _m.Called(ctx, mutantDna)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, mutants.MutantDna) error); ok {
		r0 = rf(ctx, mutantDna)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}