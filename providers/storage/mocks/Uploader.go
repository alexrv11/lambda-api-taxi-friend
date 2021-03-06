// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	"github.com/alexrv11/lambda-api-taxi-friend/providers/domain"
	"github.com/stretchr/testify/mock"
)

// Uploader is an autogenerated mock type for the Uploader type
type Uploader struct {
	mock.Mock
}

// UploadFile provides a mock function with given fields: owner, contents
func (_m *Uploader) UploadFile(owner string, contents ...*domain.File) error {
	_va := make([]interface{}, len(contents))
	for _i := range contents {
		_va[_i] = contents[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, owner)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, ...*domain.File) error); ok {
		r0 = rf(owner, contents...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
