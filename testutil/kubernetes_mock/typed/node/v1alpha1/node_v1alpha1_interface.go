// Code generated by mockery v1.1.2. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	rest "k8s.io/client-go/rest"

	v1alpha1 "k8s.io/client-go/kubernetes/typed/node/v1alpha1"
)

// NodeV1alpha1Interface is an autogenerated mock type for the NodeV1alpha1Interface type
type NodeV1alpha1Interface struct {
	mock.Mock
}

// RESTClient provides a mock function with given fields:
func (_m *NodeV1alpha1Interface) RESTClient() rest.Interface {
	ret := _m.Called()

	var r0 rest.Interface
	if rf, ok := ret.Get(0).(func() rest.Interface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(rest.Interface)
		}
	}

	return r0
}

// RuntimeClasses provides a mock function with given fields:
func (_m *NodeV1alpha1Interface) RuntimeClasses() v1alpha1.RuntimeClassInterface {
	ret := _m.Called()

	var r0 v1alpha1.RuntimeClassInterface
	if rf, ok := ret.Get(0).(func() v1alpha1.RuntimeClassInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.RuntimeClassInterface)
		}
	}

	return r0
}