// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Services is an autogenerated mock type for the Services type
type Services struct {
	mock.Mock
}

// IsMutant provides a mock function with given fields: ctx, dna
func (_m *Services) IsMutant(ctx context.Context, dna []string) error {
	ret := _m.Called(ctx, dna)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []string) error); ok {
		r0 = rf(ctx, dna)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidateDna provides a mock function with given fields: ctx, dna
func (_m *Services) ValidateDna(ctx context.Context, dna []string) error {
	ret := _m.Called(ctx, dna)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []string) error); ok {
		r0 = rf(ctx, dna)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}