// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/opensds/multi-cloud/backend/pkg/model"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *Repository) Close() {
	_m.Called()
}

// CreateBackend provides a mock function with given fields: ctx, backend
func (_m *Repository) CreateBackend(ctx context.Context, backend *model.Backend) (*model.Backend, error) {
	ret := _m.Called(ctx, backend)

	var r0 *model.Backend
	if rf, ok := ret.Get(0).(func(context.Context, *model.Backend) *model.Backend); ok {
		r0 = rf(ctx, backend)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Backend)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.Backend) error); ok {
		r1 = rf(ctx, backend)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTier provides a mock function with given fields: ctx, tier
func (_m *Repository) CreateTier(ctx context.Context, tier *model.Tier) (*model.Tier, error) {
	ret := _m.Called(ctx, tier)

	var r0 *model.Tier
	if rf, ok := ret.Get(0).(func(context.Context, *model.Tier) *model.Tier); ok {
		r0 = rf(ctx, tier)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Tier)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.Tier) error); ok {
		r1 = rf(ctx, tier)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteBackend provides a mock function with given fields: ctx, id
func (_m *Repository) DeleteBackend(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTier provides a mock function with given fields: ctx, id
func (_m *Repository) DeleteTier(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetBackend provides a mock function with given fields: ctx, id
func (_m *Repository) GetBackend(ctx context.Context, id string) (*model.Backend, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.Backend
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Backend); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Backend)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTier provides a mock function with given fields: ctx, id
func (_m *Repository) GetTier(ctx context.Context, id string) (*model.Tier, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.Tier
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Tier); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Tier)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListBackend provides a mock function with given fields: ctx, limit, offset, query
func (_m *Repository) ListBackend(ctx context.Context, limit int, offset int, query interface{}) ([]*model.Backend, error) {
	ret := _m.Called(ctx, limit, offset, query)

	var r0 []*model.Backend
	if rf, ok := ret.Get(0).(func(context.Context, int, int, interface{}) []*model.Backend); ok {
		r0 = rf(ctx, limit, offset, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Backend)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, int, interface{}) error); ok {
		r1 = rf(ctx, limit, offset, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTiers provides a mock function with given fields: ctx, limit, offset
func (_m *Repository) ListTiers(ctx context.Context, limit int, offset int) ([]*model.Tier, error) {
	ret := _m.Called(ctx, limit, offset)

	var r0 []*model.Tier
	if rf, ok := ret.Get(0).(func(context.Context, int, int) []*model.Tier); ok {
		r0 = rf(ctx, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Tier)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateBackend provides a mock function with given fields: ctx, backend
func (_m *Repository) UpdateBackend(ctx context.Context, backend *model.Backend) (*model.Backend, error) {
	ret := _m.Called(ctx, backend)

	var r0 *model.Backend
	if rf, ok := ret.Get(0).(func(context.Context, *model.Backend) *model.Backend); ok {
		r0 = rf(ctx, backend)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Backend)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.Backend) error); ok {
		r1 = rf(ctx, backend)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTier provides a mock function with given fields: ctx, tier
func (_m *Repository) UpdateTier(ctx context.Context, tier *model.Tier) (*model.Tier, error) {
	ret := _m.Called(ctx, tier)

	var r0 *model.Tier
	if rf, ok := ret.Get(0).(func(context.Context, *model.Tier) *model.Tier); ok {
		r0 = rf(ctx, tier)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Tier)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.Tier) error); ok {
		r1 = rf(ctx, tier)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
