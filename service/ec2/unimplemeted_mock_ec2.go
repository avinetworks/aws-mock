/*
Copyright 2018 The Avi Networks.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package ec2


import aws "github.com/aws/aws-sdk-go/aws"
import ec2 "github.com/aws/aws-sdk-go/service/ec2"

import request "github.com/aws/aws-sdk-go/aws/request"


func (_m *EC2API) ReleaseAddressRequest(*ec2.ReleaseAddressInput) (r *request.Request, output *ec2.ReleaseAddressOutput) {
	return
}


// AcceptReservedInstancesExchangeQuote provides a mock function with given fields: _a0
func (_m *EC2API) AcceptReservedInstancesExchangeQuote(_a0 *ec2.AcceptReservedInstancesExchangeQuoteInput) (*ec2.AcceptReservedInstancesExchangeQuoteOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.AcceptReservedInstancesExchangeQuoteOutput
	if rf, ok := ret.Get(0).(func(*ec2.AcceptReservedInstancesExchangeQuoteInput) *ec2.AcceptReservedInstancesExchangeQuoteOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AcceptReservedInstancesExchangeQuoteOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.AcceptReservedInstancesExchangeQuoteInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AcceptReservedInstancesExchangeQuoteRequest provides a mock function with given fields: _a0
func (_m *EC2API) AcceptReservedInstancesExchangeQuoteRequest(_a0 *ec2.AcceptReservedInstancesExchangeQuoteInput) (*request.Request, *ec2.AcceptReservedInstancesExchangeQuoteOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.AcceptReservedInstancesExchangeQuoteInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.AcceptReservedInstancesExchangeQuoteOutput
	if rf, ok := ret.Get(1).(func(*ec2.AcceptReservedInstancesExchangeQuoteInput) *ec2.AcceptReservedInstancesExchangeQuoteOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.AcceptReservedInstancesExchangeQuoteOutput)
		}
	}

	return r0, r1
}

// AcceptReservedInstancesExchangeQuoteWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) AcceptReservedInstancesExchangeQuoteWithContext(_a0 aws.Context, _a1 *ec2.AcceptReservedInstancesExchangeQuoteInput, _a2 ...request.Option) (*ec2.AcceptReservedInstancesExchangeQuoteOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.AcceptReservedInstancesExchangeQuoteOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.AcceptReservedInstancesExchangeQuoteInput, ...request.Option) *ec2.AcceptReservedInstancesExchangeQuoteOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AcceptReservedInstancesExchangeQuoteOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.AcceptReservedInstancesExchangeQuoteInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AcceptVpcEndpointConnections provides a mock function with given fields: _a0
func (_m *EC2API) AcceptVpcEndpointConnections(_a0 *ec2.AcceptVpcEndpointConnectionsInput) (*ec2.AcceptVpcEndpointConnectionsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.AcceptVpcEndpointConnectionsOutput
	if rf, ok := ret.Get(0).(func(*ec2.AcceptVpcEndpointConnectionsInput) *ec2.AcceptVpcEndpointConnectionsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AcceptVpcEndpointConnectionsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.AcceptVpcEndpointConnectionsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AcceptVpcEndpointConnectionsRequest provides a mock function with given fields: _a0
func (_m *EC2API) AcceptVpcEndpointConnectionsRequest(_a0 *ec2.AcceptVpcEndpointConnectionsInput) (*request.Request, *ec2.AcceptVpcEndpointConnectionsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.AcceptVpcEndpointConnectionsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.AcceptVpcEndpointConnectionsOutput
	if rf, ok := ret.Get(1).(func(*ec2.AcceptVpcEndpointConnectionsInput) *ec2.AcceptVpcEndpointConnectionsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.AcceptVpcEndpointConnectionsOutput)
		}
	}

	return r0, r1
}

// AcceptVpcEndpointConnectionsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) AcceptVpcEndpointConnectionsWithContext(_a0 aws.Context, _a1 *ec2.AcceptVpcEndpointConnectionsInput, _a2 ...request.Option) (*ec2.AcceptVpcEndpointConnectionsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.AcceptVpcEndpointConnectionsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.AcceptVpcEndpointConnectionsInput, ...request.Option) *ec2.AcceptVpcEndpointConnectionsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AcceptVpcEndpointConnectionsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.AcceptVpcEndpointConnectionsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AcceptVpcPeeringConnection provides a mock function with given fields: _a0
func (_m *EC2API) AcceptVpcPeeringConnection(_a0 *ec2.AcceptVpcPeeringConnectionInput) (*ec2.AcceptVpcPeeringConnectionOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.AcceptVpcPeeringConnectionOutput
	if rf, ok := ret.Get(0).(func(*ec2.AcceptVpcPeeringConnectionInput) *ec2.AcceptVpcPeeringConnectionOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AcceptVpcPeeringConnectionOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.AcceptVpcPeeringConnectionInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AcceptVpcPeeringConnectionRequest provides a mock function with given fields: _a0
func (_m *EC2API) AcceptVpcPeeringConnectionRequest(_a0 *ec2.AcceptVpcPeeringConnectionInput) (*request.Request, *ec2.AcceptVpcPeeringConnectionOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.AcceptVpcPeeringConnectionInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.AcceptVpcPeeringConnectionOutput
	if rf, ok := ret.Get(1).(func(*ec2.AcceptVpcPeeringConnectionInput) *ec2.AcceptVpcPeeringConnectionOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.AcceptVpcPeeringConnectionOutput)
		}
	}

	return r0, r1
}

// AcceptVpcPeeringConnectionWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) AcceptVpcPeeringConnectionWithContext(_a0 aws.Context, _a1 *ec2.AcceptVpcPeeringConnectionInput, _a2 ...request.Option) (*ec2.AcceptVpcPeeringConnectionOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.AcceptVpcPeeringConnectionOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.AcceptVpcPeeringConnectionInput, ...request.Option) *ec2.AcceptVpcPeeringConnectionOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AcceptVpcPeeringConnectionOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.AcceptVpcPeeringConnectionInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AllocateAddressRequest provides a mock function with given fields: _a0
func (_m *EC2API) AllocateAddressRequest(_a0 *ec2.AllocateAddressInput) (*request.Request, *ec2.AllocateAddressOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.AllocateAddressInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.AllocateAddressOutput
	if rf, ok := ret.Get(1).(func(*ec2.AllocateAddressInput) *ec2.AllocateAddressOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.AllocateAddressOutput)
		}
	}

	return r0, r1
}

// AllocateAddressWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) AllocateAddressWithContext(_a0 aws.Context, _a1 *ec2.AllocateAddressInput, _a2 ...request.Option) (*ec2.AllocateAddressOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.AllocateAddressOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.AllocateAddressInput, ...request.Option) *ec2.AllocateAddressOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AllocateAddressOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.AllocateAddressInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AllocateHosts provides a mock function with given fields: _a0
func (_m *EC2API) AllocateHosts(_a0 *ec2.AllocateHostsInput) (*ec2.AllocateHostsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.AllocateHostsOutput
	if rf, ok := ret.Get(0).(func(*ec2.AllocateHostsInput) *ec2.AllocateHostsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AllocateHostsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.AllocateHostsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AllocateHostsRequest provides a mock function with given fields: _a0
func (_m *EC2API) AllocateHostsRequest(_a0 *ec2.AllocateHostsInput) (*request.Request, *ec2.AllocateHostsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.AllocateHostsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.AllocateHostsOutput
	if rf, ok := ret.Get(1).(func(*ec2.AllocateHostsInput) *ec2.AllocateHostsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.AllocateHostsOutput)
		}
	}

	return r0, r1
}

// AllocateHostsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) AllocateHostsWithContext(_a0 aws.Context, _a1 *ec2.AllocateHostsInput, _a2 ...request.Option) (*ec2.AllocateHostsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.AllocateHostsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.AllocateHostsInput, ...request.Option) *ec2.AllocateHostsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AllocateHostsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.AllocateHostsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AssignIpv6Addresses provides a mock function with given fields: _a0
func (_m *EC2API) AssignIpv6Addresses(_a0 *ec2.AssignIpv6AddressesInput) (*ec2.AssignIpv6AddressesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.AssignIpv6AddressesOutput
	if rf, ok := ret.Get(0).(func(*ec2.AssignIpv6AddressesInput) *ec2.AssignIpv6AddressesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AssignIpv6AddressesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.AssignIpv6AddressesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AssignIpv6AddressesRequest provides a mock function with given fields: _a0
func (_m *EC2API) AssignIpv6AddressesRequest(_a0 *ec2.AssignIpv6AddressesInput) (*request.Request, *ec2.AssignIpv6AddressesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.AssignIpv6AddressesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.AssignIpv6AddressesOutput
	if rf, ok := ret.Get(1).(func(*ec2.AssignIpv6AddressesInput) *ec2.AssignIpv6AddressesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.AssignIpv6AddressesOutput)
		}
	}

	return r0, r1
}

// AssignIpv6AddressesWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) AssignIpv6AddressesWithContext(_a0 aws.Context, _a1 *ec2.AssignIpv6AddressesInput, _a2 ...request.Option) (*ec2.AssignIpv6AddressesOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.AssignIpv6AddressesOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.AssignIpv6AddressesInput, ...request.Option) *ec2.AssignIpv6AddressesOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AssignIpv6AddressesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.AssignIpv6AddressesInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}



// AssignPrivateIpAddressesRequest provides a mock function with given fields: _a0
func (_m *EC2API) AssignPrivateIpAddressesRequest(_a0 *ec2.AssignPrivateIpAddressesInput) (*request.Request, *ec2.AssignPrivateIpAddressesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.AssignPrivateIpAddressesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.AssignPrivateIpAddressesOutput
	if rf, ok := ret.Get(1).(func(*ec2.AssignPrivateIpAddressesInput) *ec2.AssignPrivateIpAddressesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.AssignPrivateIpAddressesOutput)
		}
	}

	return r0, r1
}

// AssignPrivateIpAddressesWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) AssignPrivateIpAddressesWithContext(_a0 aws.Context, _a1 *ec2.AssignPrivateIpAddressesInput, _a2 ...request.Option) (*ec2.AssignPrivateIpAddressesOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.AssignPrivateIpAddressesOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.AssignPrivateIpAddressesInput, ...request.Option) *ec2.AssignPrivateIpAddressesOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AssignPrivateIpAddressesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.AssignPrivateIpAddressesInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}



// AssociateAddressRequest provides a mock function with given fields: _a0
func (_m *EC2API) AssociateAddressRequest(_a0 *ec2.AssociateAddressInput) (*request.Request, *ec2.AssociateAddressOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.AssociateAddressInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.AssociateAddressOutput
	if rf, ok := ret.Get(1).(func(*ec2.AssociateAddressInput) *ec2.AssociateAddressOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.AssociateAddressOutput)
		}
	}

	return r0, r1
}

// AssociateAddressWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) AssociateAddressWithContext(_a0 aws.Context, _a1 *ec2.AssociateAddressInput, _a2 ...request.Option) (*ec2.AssociateAddressOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.AssociateAddressOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.AssociateAddressInput, ...request.Option) *ec2.AssociateAddressOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AssociateAddressOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.AssociateAddressInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AssociateDhcpOptions provides a mock function with given fields: _a0
func (_m *EC2API) AssociateDhcpOptions(_a0 *ec2.AssociateDhcpOptionsInput) (*ec2.AssociateDhcpOptionsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.AssociateDhcpOptionsOutput
	if rf, ok := ret.Get(0).(func(*ec2.AssociateDhcpOptionsInput) *ec2.AssociateDhcpOptionsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AssociateDhcpOptionsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.AssociateDhcpOptionsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AssociateDhcpOptionsRequest provides a mock function with given fields: _a0
func (_m *EC2API) AssociateDhcpOptionsRequest(_a0 *ec2.AssociateDhcpOptionsInput) (*request.Request, *ec2.AssociateDhcpOptionsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.AssociateDhcpOptionsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.AssociateDhcpOptionsOutput
	if rf, ok := ret.Get(1).(func(*ec2.AssociateDhcpOptionsInput) *ec2.AssociateDhcpOptionsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.AssociateDhcpOptionsOutput)
		}
	}

	return r0, r1
}

// AssociateDhcpOptionsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) AssociateDhcpOptionsWithContext(_a0 aws.Context, _a1 *ec2.AssociateDhcpOptionsInput, _a2 ...request.Option) (*ec2.AssociateDhcpOptionsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.AssociateDhcpOptionsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.AssociateDhcpOptionsInput, ...request.Option) *ec2.AssociateDhcpOptionsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AssociateDhcpOptionsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.AssociateDhcpOptionsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AssociateIamInstanceProfile provides a mock function with given fields: _a0
func (_m *EC2API) AssociateIamInstanceProfile(_a0 *ec2.AssociateIamInstanceProfileInput) (*ec2.AssociateIamInstanceProfileOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.AssociateIamInstanceProfileOutput
	if rf, ok := ret.Get(0).(func(*ec2.AssociateIamInstanceProfileInput) *ec2.AssociateIamInstanceProfileOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AssociateIamInstanceProfileOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.AssociateIamInstanceProfileInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AssociateIamInstanceProfileRequest provides a mock function with given fields: _a0
func (_m *EC2API) AssociateIamInstanceProfileRequest(_a0 *ec2.AssociateIamInstanceProfileInput) (*request.Request, *ec2.AssociateIamInstanceProfileOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.AssociateIamInstanceProfileInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.AssociateIamInstanceProfileOutput
	if rf, ok := ret.Get(1).(func(*ec2.AssociateIamInstanceProfileInput) *ec2.AssociateIamInstanceProfileOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.AssociateIamInstanceProfileOutput)
		}
	}

	return r0, r1
}

// AssociateIamInstanceProfileWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) AssociateIamInstanceProfileWithContext(_a0 aws.Context, _a1 *ec2.AssociateIamInstanceProfileInput, _a2 ...request.Option) (*ec2.AssociateIamInstanceProfileOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.AssociateIamInstanceProfileOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.AssociateIamInstanceProfileInput, ...request.Option) *ec2.AssociateIamInstanceProfileOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AssociateIamInstanceProfileOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.AssociateIamInstanceProfileInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AssociateRouteTable provides a mock function with given fields: _a0
func (_m *EC2API) AssociateRouteTable(_a0 *ec2.AssociateRouteTableInput) (*ec2.AssociateRouteTableOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.AssociateRouteTableOutput
	if rf, ok := ret.Get(0).(func(*ec2.AssociateRouteTableInput) *ec2.AssociateRouteTableOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AssociateRouteTableOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.AssociateRouteTableInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AssociateRouteTableRequest provides a mock function with given fields: _a0
func (_m *EC2API) AssociateRouteTableRequest(_a0 *ec2.AssociateRouteTableInput) (*request.Request, *ec2.AssociateRouteTableOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.AssociateRouteTableInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.AssociateRouteTableOutput
	if rf, ok := ret.Get(1).(func(*ec2.AssociateRouteTableInput) *ec2.AssociateRouteTableOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.AssociateRouteTableOutput)
		}
	}

	return r0, r1
}

// AssociateRouteTableWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) AssociateRouteTableWithContext(_a0 aws.Context, _a1 *ec2.AssociateRouteTableInput, _a2 ...request.Option) (*ec2.AssociateRouteTableOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.AssociateRouteTableOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.AssociateRouteTableInput, ...request.Option) *ec2.AssociateRouteTableOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AssociateRouteTableOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.AssociateRouteTableInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AssociateSubnetCidrBlock provides a mock function with given fields: _a0
func (_m *EC2API) AssociateSubnetCidrBlock(_a0 *ec2.AssociateSubnetCidrBlockInput) (*ec2.AssociateSubnetCidrBlockOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.AssociateSubnetCidrBlockOutput
	if rf, ok := ret.Get(0).(func(*ec2.AssociateSubnetCidrBlockInput) *ec2.AssociateSubnetCidrBlockOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AssociateSubnetCidrBlockOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.AssociateSubnetCidrBlockInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AssociateSubnetCidrBlockRequest provides a mock function with given fields: _a0
func (_m *EC2API) AssociateSubnetCidrBlockRequest(_a0 *ec2.AssociateSubnetCidrBlockInput) (*request.Request, *ec2.AssociateSubnetCidrBlockOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.AssociateSubnetCidrBlockInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.AssociateSubnetCidrBlockOutput
	if rf, ok := ret.Get(1).(func(*ec2.AssociateSubnetCidrBlockInput) *ec2.AssociateSubnetCidrBlockOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.AssociateSubnetCidrBlockOutput)
		}
	}

	return r0, r1
}

// AssociateSubnetCidrBlockWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) AssociateSubnetCidrBlockWithContext(_a0 aws.Context, _a1 *ec2.AssociateSubnetCidrBlockInput, _a2 ...request.Option) (*ec2.AssociateSubnetCidrBlockOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.AssociateSubnetCidrBlockOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.AssociateSubnetCidrBlockInput, ...request.Option) *ec2.AssociateSubnetCidrBlockOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AssociateSubnetCidrBlockOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.AssociateSubnetCidrBlockInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AssociateVpcCidrBlock provides a mock function with given fields: _a0
func (_m *EC2API) AssociateVpcCidrBlock(_a0 *ec2.AssociateVpcCidrBlockInput) (*ec2.AssociateVpcCidrBlockOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.AssociateVpcCidrBlockOutput
	if rf, ok := ret.Get(0).(func(*ec2.AssociateVpcCidrBlockInput) *ec2.AssociateVpcCidrBlockOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AssociateVpcCidrBlockOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.AssociateVpcCidrBlockInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AssociateVpcCidrBlockRequest provides a mock function with given fields: _a0
func (_m *EC2API) AssociateVpcCidrBlockRequest(_a0 *ec2.AssociateVpcCidrBlockInput) (*request.Request, *ec2.AssociateVpcCidrBlockOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.AssociateVpcCidrBlockInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.AssociateVpcCidrBlockOutput
	if rf, ok := ret.Get(1).(func(*ec2.AssociateVpcCidrBlockInput) *ec2.AssociateVpcCidrBlockOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.AssociateVpcCidrBlockOutput)
		}
	}

	return r0, r1
}

// AssociateVpcCidrBlockWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) AssociateVpcCidrBlockWithContext(_a0 aws.Context, _a1 *ec2.AssociateVpcCidrBlockInput, _a2 ...request.Option) (*ec2.AssociateVpcCidrBlockOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.AssociateVpcCidrBlockOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.AssociateVpcCidrBlockInput, ...request.Option) *ec2.AssociateVpcCidrBlockOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AssociateVpcCidrBlockOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.AssociateVpcCidrBlockInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AttachClassicLinkVpc provides a mock function with given fields: _a0
func (_m *EC2API) AttachClassicLinkVpc(_a0 *ec2.AttachClassicLinkVpcInput) (*ec2.AttachClassicLinkVpcOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.AttachClassicLinkVpcOutput
	if rf, ok := ret.Get(0).(func(*ec2.AttachClassicLinkVpcInput) *ec2.AttachClassicLinkVpcOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AttachClassicLinkVpcOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.AttachClassicLinkVpcInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AttachClassicLinkVpcRequest provides a mock function with given fields: _a0
func (_m *EC2API) AttachClassicLinkVpcRequest(_a0 *ec2.AttachClassicLinkVpcInput) (*request.Request, *ec2.AttachClassicLinkVpcOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.AttachClassicLinkVpcInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.AttachClassicLinkVpcOutput
	if rf, ok := ret.Get(1).(func(*ec2.AttachClassicLinkVpcInput) *ec2.AttachClassicLinkVpcOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.AttachClassicLinkVpcOutput)
		}
	}

	return r0, r1
}

// AttachClassicLinkVpcWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) AttachClassicLinkVpcWithContext(_a0 aws.Context, _a1 *ec2.AttachClassicLinkVpcInput, _a2 ...request.Option) (*ec2.AttachClassicLinkVpcOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.AttachClassicLinkVpcOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.AttachClassicLinkVpcInput, ...request.Option) *ec2.AttachClassicLinkVpcOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AttachClassicLinkVpcOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.AttachClassicLinkVpcInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AttachInternetGateway provides a mock function with given fields: _a0
func (_m *EC2API) AttachInternetGateway(_a0 *ec2.AttachInternetGatewayInput) (*ec2.AttachInternetGatewayOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.AttachInternetGatewayOutput
	if rf, ok := ret.Get(0).(func(*ec2.AttachInternetGatewayInput) *ec2.AttachInternetGatewayOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AttachInternetGatewayOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.AttachInternetGatewayInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AttachInternetGatewayRequest provides a mock function with given fields: _a0
func (_m *EC2API) AttachInternetGatewayRequest(_a0 *ec2.AttachInternetGatewayInput) (*request.Request, *ec2.AttachInternetGatewayOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.AttachInternetGatewayInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.AttachInternetGatewayOutput
	if rf, ok := ret.Get(1).(func(*ec2.AttachInternetGatewayInput) *ec2.AttachInternetGatewayOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.AttachInternetGatewayOutput)
		}
	}

	return r0, r1
}

// AttachInternetGatewayWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) AttachInternetGatewayWithContext(_a0 aws.Context, _a1 *ec2.AttachInternetGatewayInput, _a2 ...request.Option) (*ec2.AttachInternetGatewayOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.AttachInternetGatewayOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.AttachInternetGatewayInput, ...request.Option) *ec2.AttachInternetGatewayOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AttachInternetGatewayOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.AttachInternetGatewayInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AttachNetworkInterface provides a mock function with given fields: _a0
func (_m *EC2API) AttachNetworkInterface(_a0 *ec2.AttachNetworkInterfaceInput) (*ec2.AttachNetworkInterfaceOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.AttachNetworkInterfaceOutput
	if rf, ok := ret.Get(0).(func(*ec2.AttachNetworkInterfaceInput) *ec2.AttachNetworkInterfaceOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AttachNetworkInterfaceOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.AttachNetworkInterfaceInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AttachNetworkInterfaceRequest provides a mock function with given fields: _a0
func (_m *EC2API) AttachNetworkInterfaceRequest(_a0 *ec2.AttachNetworkInterfaceInput) (*request.Request, *ec2.AttachNetworkInterfaceOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.AttachNetworkInterfaceInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.AttachNetworkInterfaceOutput
	if rf, ok := ret.Get(1).(func(*ec2.AttachNetworkInterfaceInput) *ec2.AttachNetworkInterfaceOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.AttachNetworkInterfaceOutput)
		}
	}

	return r0, r1
}

// AttachNetworkInterfaceWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) AttachNetworkInterfaceWithContext(_a0 aws.Context, _a1 *ec2.AttachNetworkInterfaceInput, _a2 ...request.Option) (*ec2.AttachNetworkInterfaceOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.AttachNetworkInterfaceOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.AttachNetworkInterfaceInput, ...request.Option) *ec2.AttachNetworkInterfaceOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AttachNetworkInterfaceOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.AttachNetworkInterfaceInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AttachVolume provides a mock function with given fields: _a0
func (_m *EC2API) AttachVolume(_a0 *ec2.AttachVolumeInput) (*ec2.VolumeAttachment, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.VolumeAttachment
	if rf, ok := ret.Get(0).(func(*ec2.AttachVolumeInput) *ec2.VolumeAttachment); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.VolumeAttachment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.AttachVolumeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AttachVolumeRequest provides a mock function with given fields: _a0
func (_m *EC2API) AttachVolumeRequest(_a0 *ec2.AttachVolumeInput) (*request.Request, *ec2.VolumeAttachment) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.AttachVolumeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.VolumeAttachment
	if rf, ok := ret.Get(1).(func(*ec2.AttachVolumeInput) *ec2.VolumeAttachment); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.VolumeAttachment)
		}
	}

	return r0, r1
}

// AttachVolumeWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) AttachVolumeWithContext(_a0 aws.Context, _a1 *ec2.AttachVolumeInput, _a2 ...request.Option) (*ec2.VolumeAttachment, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.VolumeAttachment
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.AttachVolumeInput, ...request.Option) *ec2.VolumeAttachment); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.VolumeAttachment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.AttachVolumeInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AttachVpnGateway provides a mock function with given fields: _a0
func (_m *EC2API) AttachVpnGateway(_a0 *ec2.AttachVpnGatewayInput) (*ec2.AttachVpnGatewayOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.AttachVpnGatewayOutput
	if rf, ok := ret.Get(0).(func(*ec2.AttachVpnGatewayInput) *ec2.AttachVpnGatewayOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AttachVpnGatewayOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.AttachVpnGatewayInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AttachVpnGatewayRequest provides a mock function with given fields: _a0
func (_m *EC2API) AttachVpnGatewayRequest(_a0 *ec2.AttachVpnGatewayInput) (*request.Request, *ec2.AttachVpnGatewayOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.AttachVpnGatewayInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.AttachVpnGatewayOutput
	if rf, ok := ret.Get(1).(func(*ec2.AttachVpnGatewayInput) *ec2.AttachVpnGatewayOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.AttachVpnGatewayOutput)
		}
	}

	return r0, r1
}

// AttachVpnGatewayWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) AttachVpnGatewayWithContext(_a0 aws.Context, _a1 *ec2.AttachVpnGatewayInput, _a2 ...request.Option) (*ec2.AttachVpnGatewayOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.AttachVpnGatewayOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.AttachVpnGatewayInput, ...request.Option) *ec2.AttachVpnGatewayOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AttachVpnGatewayOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.AttachVpnGatewayInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthorizeSecurityGroupEgress provides a mock function with given fields: _a0
func (_m *EC2API) AuthorizeSecurityGroupEgress(_a0 *ec2.AuthorizeSecurityGroupEgressInput) (*ec2.AuthorizeSecurityGroupEgressOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.AuthorizeSecurityGroupEgressOutput
	if rf, ok := ret.Get(0).(func(*ec2.AuthorizeSecurityGroupEgressInput) *ec2.AuthorizeSecurityGroupEgressOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AuthorizeSecurityGroupEgressOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.AuthorizeSecurityGroupEgressInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthorizeSecurityGroupEgressRequest provides a mock function with given fields: _a0
func (_m *EC2API) AuthorizeSecurityGroupEgressRequest(_a0 *ec2.AuthorizeSecurityGroupEgressInput) (*request.Request, *ec2.AuthorizeSecurityGroupEgressOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.AuthorizeSecurityGroupEgressInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.AuthorizeSecurityGroupEgressOutput
	if rf, ok := ret.Get(1).(func(*ec2.AuthorizeSecurityGroupEgressInput) *ec2.AuthorizeSecurityGroupEgressOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.AuthorizeSecurityGroupEgressOutput)
		}
	}

	return r0, r1
}

// AuthorizeSecurityGroupEgressWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) AuthorizeSecurityGroupEgressWithContext(_a0 aws.Context, _a1 *ec2.AuthorizeSecurityGroupEgressInput, _a2 ...request.Option) (*ec2.AuthorizeSecurityGroupEgressOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.AuthorizeSecurityGroupEgressOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.AuthorizeSecurityGroupEgressInput, ...request.Option) *ec2.AuthorizeSecurityGroupEgressOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AuthorizeSecurityGroupEgressOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.AuthorizeSecurityGroupEgressInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}



// AuthorizeSecurityGroupIngressRequest provides a mock function with given fields: _a0
func (_m *EC2API) AuthorizeSecurityGroupIngressRequest(_a0 *ec2.AuthorizeSecurityGroupIngressInput) (*request.Request, *ec2.AuthorizeSecurityGroupIngressOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.AuthorizeSecurityGroupIngressInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.AuthorizeSecurityGroupIngressOutput
	if rf, ok := ret.Get(1).(func(*ec2.AuthorizeSecurityGroupIngressInput) *ec2.AuthorizeSecurityGroupIngressOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.AuthorizeSecurityGroupIngressOutput)
		}
	}

	return r0, r1
}

// AuthorizeSecurityGroupIngressWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) AuthorizeSecurityGroupIngressWithContext(_a0 aws.Context, _a1 *ec2.AuthorizeSecurityGroupIngressInput, _a2 ...request.Option) (*ec2.AuthorizeSecurityGroupIngressOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.AuthorizeSecurityGroupIngressOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.AuthorizeSecurityGroupIngressInput, ...request.Option) *ec2.AuthorizeSecurityGroupIngressOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.AuthorizeSecurityGroupIngressOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.AuthorizeSecurityGroupIngressInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BundleInstance provides a mock function with given fields: _a0
func (_m *EC2API) BundleInstance(_a0 *ec2.BundleInstanceInput) (*ec2.BundleInstanceOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.BundleInstanceOutput
	if rf, ok := ret.Get(0).(func(*ec2.BundleInstanceInput) *ec2.BundleInstanceOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.BundleInstanceOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.BundleInstanceInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BundleInstanceRequest provides a mock function with given fields: _a0
func (_m *EC2API) BundleInstanceRequest(_a0 *ec2.BundleInstanceInput) (*request.Request, *ec2.BundleInstanceOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.BundleInstanceInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.BundleInstanceOutput
	if rf, ok := ret.Get(1).(func(*ec2.BundleInstanceInput) *ec2.BundleInstanceOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.BundleInstanceOutput)
		}
	}

	return r0, r1
}

// BundleInstanceWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) BundleInstanceWithContext(_a0 aws.Context, _a1 *ec2.BundleInstanceInput, _a2 ...request.Option) (*ec2.BundleInstanceOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.BundleInstanceOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.BundleInstanceInput, ...request.Option) *ec2.BundleInstanceOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.BundleInstanceOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.BundleInstanceInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CancelBundleTask provides a mock function with given fields: _a0
func (_m *EC2API) CancelBundleTask(_a0 *ec2.CancelBundleTaskInput) (*ec2.CancelBundleTaskOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CancelBundleTaskOutput
	if rf, ok := ret.Get(0).(func(*ec2.CancelBundleTaskInput) *ec2.CancelBundleTaskOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CancelBundleTaskOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CancelBundleTaskInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CancelBundleTaskRequest provides a mock function with given fields: _a0
func (_m *EC2API) CancelBundleTaskRequest(_a0 *ec2.CancelBundleTaskInput) (*request.Request, *ec2.CancelBundleTaskOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CancelBundleTaskInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CancelBundleTaskOutput
	if rf, ok := ret.Get(1).(func(*ec2.CancelBundleTaskInput) *ec2.CancelBundleTaskOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CancelBundleTaskOutput)
		}
	}

	return r0, r1
}

// CancelBundleTaskWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CancelBundleTaskWithContext(_a0 aws.Context, _a1 *ec2.CancelBundleTaskInput, _a2 ...request.Option) (*ec2.CancelBundleTaskOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.CancelBundleTaskOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CancelBundleTaskInput, ...request.Option) *ec2.CancelBundleTaskOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CancelBundleTaskOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CancelBundleTaskInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CancelConversionTask provides a mock function with given fields: _a0
func (_m *EC2API) CancelConversionTask(_a0 *ec2.CancelConversionTaskInput) (*ec2.CancelConversionTaskOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CancelConversionTaskOutput
	if rf, ok := ret.Get(0).(func(*ec2.CancelConversionTaskInput) *ec2.CancelConversionTaskOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CancelConversionTaskOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CancelConversionTaskInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CancelConversionTaskRequest provides a mock function with given fields: _a0
func (_m *EC2API) CancelConversionTaskRequest(_a0 *ec2.CancelConversionTaskInput) (*request.Request, *ec2.CancelConversionTaskOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CancelConversionTaskInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CancelConversionTaskOutput
	if rf, ok := ret.Get(1).(func(*ec2.CancelConversionTaskInput) *ec2.CancelConversionTaskOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CancelConversionTaskOutput)
		}
	}

	return r0, r1
}

// CancelConversionTaskWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CancelConversionTaskWithContext(_a0 aws.Context, _a1 *ec2.CancelConversionTaskInput, _a2 ...request.Option) (*ec2.CancelConversionTaskOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.CancelConversionTaskOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CancelConversionTaskInput, ...request.Option) *ec2.CancelConversionTaskOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CancelConversionTaskOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CancelConversionTaskInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CancelExportTask provides a mock function with given fields: _a0
func (_m *EC2API) CancelExportTask(_a0 *ec2.CancelExportTaskInput) (*ec2.CancelExportTaskOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CancelExportTaskOutput
	if rf, ok := ret.Get(0).(func(*ec2.CancelExportTaskInput) *ec2.CancelExportTaskOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CancelExportTaskOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CancelExportTaskInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CancelExportTaskRequest provides a mock function with given fields: _a0
func (_m *EC2API) CancelExportTaskRequest(_a0 *ec2.CancelExportTaskInput) (*request.Request, *ec2.CancelExportTaskOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CancelExportTaskInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CancelExportTaskOutput
	if rf, ok := ret.Get(1).(func(*ec2.CancelExportTaskInput) *ec2.CancelExportTaskOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CancelExportTaskOutput)
		}
	}

	return r0, r1
}

// CancelExportTaskWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CancelExportTaskWithContext(_a0 aws.Context, _a1 *ec2.CancelExportTaskInput, _a2 ...request.Option) (*ec2.CancelExportTaskOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.CancelExportTaskOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CancelExportTaskInput, ...request.Option) *ec2.CancelExportTaskOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CancelExportTaskOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CancelExportTaskInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CancelImportTask provides a mock function with given fields: _a0
func (_m *EC2API) CancelImportTask(_a0 *ec2.CancelImportTaskInput) (*ec2.CancelImportTaskOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CancelImportTaskOutput
	if rf, ok := ret.Get(0).(func(*ec2.CancelImportTaskInput) *ec2.CancelImportTaskOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CancelImportTaskOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CancelImportTaskInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CancelImportTaskRequest provides a mock function with given fields: _a0
func (_m *EC2API) CancelImportTaskRequest(_a0 *ec2.CancelImportTaskInput) (*request.Request, *ec2.CancelImportTaskOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CancelImportTaskInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CancelImportTaskOutput
	if rf, ok := ret.Get(1).(func(*ec2.CancelImportTaskInput) *ec2.CancelImportTaskOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CancelImportTaskOutput)
		}
	}

	return r0, r1
}

// CancelImportTaskWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CancelImportTaskWithContext(_a0 aws.Context, _a1 *ec2.CancelImportTaskInput, _a2 ...request.Option) (*ec2.CancelImportTaskOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.CancelImportTaskOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CancelImportTaskInput, ...request.Option) *ec2.CancelImportTaskOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CancelImportTaskOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CancelImportTaskInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CancelReservedInstancesListing provides a mock function with given fields: _a0
func (_m *EC2API) CancelReservedInstancesListing(_a0 *ec2.CancelReservedInstancesListingInput) (*ec2.CancelReservedInstancesListingOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CancelReservedInstancesListingOutput
	if rf, ok := ret.Get(0).(func(*ec2.CancelReservedInstancesListingInput) *ec2.CancelReservedInstancesListingOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CancelReservedInstancesListingOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CancelReservedInstancesListingInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CancelReservedInstancesListingRequest provides a mock function with given fields: _a0
func (_m *EC2API) CancelReservedInstancesListingRequest(_a0 *ec2.CancelReservedInstancesListingInput) (*request.Request, *ec2.CancelReservedInstancesListingOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CancelReservedInstancesListingInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CancelReservedInstancesListingOutput
	if rf, ok := ret.Get(1).(func(*ec2.CancelReservedInstancesListingInput) *ec2.CancelReservedInstancesListingOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CancelReservedInstancesListingOutput)
		}
	}

	return r0, r1
}

// CancelReservedInstancesListingWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CancelReservedInstancesListingWithContext(_a0 aws.Context, _a1 *ec2.CancelReservedInstancesListingInput, _a2 ...request.Option) (*ec2.CancelReservedInstancesListingOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.CancelReservedInstancesListingOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CancelReservedInstancesListingInput, ...request.Option) *ec2.CancelReservedInstancesListingOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CancelReservedInstancesListingOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CancelReservedInstancesListingInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CancelSpotFleetRequests provides a mock function with given fields: _a0
func (_m *EC2API) CancelSpotFleetRequests(_a0 *ec2.CancelSpotFleetRequestsInput) (*ec2.CancelSpotFleetRequestsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CancelSpotFleetRequestsOutput
	if rf, ok := ret.Get(0).(func(*ec2.CancelSpotFleetRequestsInput) *ec2.CancelSpotFleetRequestsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CancelSpotFleetRequestsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CancelSpotFleetRequestsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CancelSpotFleetRequestsRequest provides a mock function with given fields: _a0
func (_m *EC2API) CancelSpotFleetRequestsRequest(_a0 *ec2.CancelSpotFleetRequestsInput) (*request.Request, *ec2.CancelSpotFleetRequestsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CancelSpotFleetRequestsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CancelSpotFleetRequestsOutput
	if rf, ok := ret.Get(1).(func(*ec2.CancelSpotFleetRequestsInput) *ec2.CancelSpotFleetRequestsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CancelSpotFleetRequestsOutput)
		}
	}

	return r0, r1
}

// CancelSpotFleetRequestsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CancelSpotFleetRequestsWithContext(_a0 aws.Context, _a1 *ec2.CancelSpotFleetRequestsInput, _a2 ...request.Option) (*ec2.CancelSpotFleetRequestsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.CancelSpotFleetRequestsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CancelSpotFleetRequestsInput, ...request.Option) *ec2.CancelSpotFleetRequestsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CancelSpotFleetRequestsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CancelSpotFleetRequestsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CancelSpotInstanceRequests provides a mock function with given fields: _a0
func (_m *EC2API) CancelSpotInstanceRequests(_a0 *ec2.CancelSpotInstanceRequestsInput) (*ec2.CancelSpotInstanceRequestsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CancelSpotInstanceRequestsOutput
	if rf, ok := ret.Get(0).(func(*ec2.CancelSpotInstanceRequestsInput) *ec2.CancelSpotInstanceRequestsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CancelSpotInstanceRequestsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CancelSpotInstanceRequestsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CancelSpotInstanceRequestsRequest provides a mock function with given fields: _a0
func (_m *EC2API) CancelSpotInstanceRequestsRequest(_a0 *ec2.CancelSpotInstanceRequestsInput) (*request.Request, *ec2.CancelSpotInstanceRequestsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CancelSpotInstanceRequestsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CancelSpotInstanceRequestsOutput
	if rf, ok := ret.Get(1).(func(*ec2.CancelSpotInstanceRequestsInput) *ec2.CancelSpotInstanceRequestsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CancelSpotInstanceRequestsOutput)
		}
	}

	return r0, r1
}

// CancelSpotInstanceRequestsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CancelSpotInstanceRequestsWithContext(_a0 aws.Context, _a1 *ec2.CancelSpotInstanceRequestsInput, _a2 ...request.Option) (*ec2.CancelSpotInstanceRequestsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.CancelSpotInstanceRequestsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CancelSpotInstanceRequestsInput, ...request.Option) *ec2.CancelSpotInstanceRequestsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CancelSpotInstanceRequestsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CancelSpotInstanceRequestsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConfirmProductInstance provides a mock function with given fields: _a0
func (_m *EC2API) ConfirmProductInstance(_a0 *ec2.ConfirmProductInstanceInput) (*ec2.ConfirmProductInstanceOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ConfirmProductInstanceOutput
	if rf, ok := ret.Get(0).(func(*ec2.ConfirmProductInstanceInput) *ec2.ConfirmProductInstanceOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ConfirmProductInstanceOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ConfirmProductInstanceInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConfirmProductInstanceRequest provides a mock function with given fields: _a0
func (_m *EC2API) ConfirmProductInstanceRequest(_a0 *ec2.ConfirmProductInstanceInput) (*request.Request, *ec2.ConfirmProductInstanceOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ConfirmProductInstanceInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ConfirmProductInstanceOutput
	if rf, ok := ret.Get(1).(func(*ec2.ConfirmProductInstanceInput) *ec2.ConfirmProductInstanceOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ConfirmProductInstanceOutput)
		}
	}

	return r0, r1
}

// ConfirmProductInstanceWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) ConfirmProductInstanceWithContext(_a0 aws.Context, _a1 *ec2.ConfirmProductInstanceInput, _a2 ...request.Option) (*ec2.ConfirmProductInstanceOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.ConfirmProductInstanceOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.ConfirmProductInstanceInput, ...request.Option) *ec2.ConfirmProductInstanceOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ConfirmProductInstanceOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.ConfirmProductInstanceInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CopyFpgaImage provides a mock function with given fields: _a0
func (_m *EC2API) CopyFpgaImage(_a0 *ec2.CopyFpgaImageInput) (*ec2.CopyFpgaImageOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CopyFpgaImageOutput
	if rf, ok := ret.Get(0).(func(*ec2.CopyFpgaImageInput) *ec2.CopyFpgaImageOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CopyFpgaImageOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CopyFpgaImageInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CopyFpgaImageRequest provides a mock function with given fields: _a0
func (_m *EC2API) CopyFpgaImageRequest(_a0 *ec2.CopyFpgaImageInput) (*request.Request, *ec2.CopyFpgaImageOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CopyFpgaImageInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CopyFpgaImageOutput
	if rf, ok := ret.Get(1).(func(*ec2.CopyFpgaImageInput) *ec2.CopyFpgaImageOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CopyFpgaImageOutput)
		}
	}

	return r0, r1
}

// CopyFpgaImageWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CopyFpgaImageWithContext(_a0 aws.Context, _a1 *ec2.CopyFpgaImageInput, _a2 ...request.Option) (*ec2.CopyFpgaImageOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.CopyFpgaImageOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CopyFpgaImageInput, ...request.Option) *ec2.CopyFpgaImageOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CopyFpgaImageOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CopyFpgaImageInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CopyImage provides a mock function with given fields: _a0
func (_m *EC2API) CopyImage(_a0 *ec2.CopyImageInput) (*ec2.CopyImageOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CopyImageOutput
	if rf, ok := ret.Get(0).(func(*ec2.CopyImageInput) *ec2.CopyImageOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CopyImageOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CopyImageInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CopyImageRequest provides a mock function with given fields: _a0
func (_m *EC2API) CopyImageRequest(_a0 *ec2.CopyImageInput) (*request.Request, *ec2.CopyImageOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CopyImageInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CopyImageOutput
	if rf, ok := ret.Get(1).(func(*ec2.CopyImageInput) *ec2.CopyImageOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CopyImageOutput)
		}
	}

	return r0, r1
}

// CopyImageWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CopyImageWithContext(_a0 aws.Context, _a1 *ec2.CopyImageInput, _a2 ...request.Option) (*ec2.CopyImageOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.CopyImageOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CopyImageInput, ...request.Option) *ec2.CopyImageOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CopyImageOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CopyImageInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CopySnapshot provides a mock function with given fields: _a0
func (_m *EC2API) CopySnapshot(_a0 *ec2.CopySnapshotInput) (*ec2.CopySnapshotOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CopySnapshotOutput
	if rf, ok := ret.Get(0).(func(*ec2.CopySnapshotInput) *ec2.CopySnapshotOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CopySnapshotOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CopySnapshotInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CopySnapshotRequest provides a mock function with given fields: _a0
func (_m *EC2API) CopySnapshotRequest(_a0 *ec2.CopySnapshotInput) (*request.Request, *ec2.CopySnapshotOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CopySnapshotInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CopySnapshotOutput
	if rf, ok := ret.Get(1).(func(*ec2.CopySnapshotInput) *ec2.CopySnapshotOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CopySnapshotOutput)
		}
	}

	return r0, r1
}

// CopySnapshotWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CopySnapshotWithContext(_a0 aws.Context, _a1 *ec2.CopySnapshotInput, _a2 ...request.Option) (*ec2.CopySnapshotOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.CopySnapshotOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CopySnapshotInput, ...request.Option) *ec2.CopySnapshotOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CopySnapshotOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CopySnapshotInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateCustomerGateway provides a mock function with given fields: _a0
func (_m *EC2API) CreateCustomerGateway(_a0 *ec2.CreateCustomerGatewayInput) (*ec2.CreateCustomerGatewayOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateCustomerGatewayOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateCustomerGatewayInput) *ec2.CreateCustomerGatewayOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateCustomerGatewayOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateCustomerGatewayInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateCustomerGatewayRequest provides a mock function with given fields: _a0
func (_m *EC2API) CreateCustomerGatewayRequest(_a0 *ec2.CreateCustomerGatewayInput) (*request.Request, *ec2.CreateCustomerGatewayOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateCustomerGatewayInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateCustomerGatewayOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateCustomerGatewayInput) *ec2.CreateCustomerGatewayOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateCustomerGatewayOutput)
		}
	}

	return r0, r1
}

// CreateCustomerGatewayWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CreateCustomerGatewayWithContext(_a0 aws.Context, _a1 *ec2.CreateCustomerGatewayInput, _a2 ...request.Option) (*ec2.CreateCustomerGatewayOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.CreateCustomerGatewayOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CreateCustomerGatewayInput, ...request.Option) *ec2.CreateCustomerGatewayOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateCustomerGatewayOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CreateCustomerGatewayInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateDefaultSubnet provides a mock function with given fields: _a0
func (_m *EC2API) CreateDefaultSubnet(_a0 *ec2.CreateDefaultSubnetInput) (*ec2.CreateDefaultSubnetOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateDefaultSubnetOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateDefaultSubnetInput) *ec2.CreateDefaultSubnetOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateDefaultSubnetOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateDefaultSubnetInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateDefaultSubnetRequest provides a mock function with given fields: _a0
func (_m *EC2API) CreateDefaultSubnetRequest(_a0 *ec2.CreateDefaultSubnetInput) (*request.Request, *ec2.CreateDefaultSubnetOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateDefaultSubnetInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateDefaultSubnetOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateDefaultSubnetInput) *ec2.CreateDefaultSubnetOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateDefaultSubnetOutput)
		}
	}

	return r0, r1
}

// CreateDefaultSubnetWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CreateDefaultSubnetWithContext(_a0 aws.Context, _a1 *ec2.CreateDefaultSubnetInput, _a2 ...request.Option) (*ec2.CreateDefaultSubnetOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.CreateDefaultSubnetOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CreateDefaultSubnetInput, ...request.Option) *ec2.CreateDefaultSubnetOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateDefaultSubnetOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CreateDefaultSubnetInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateDefaultVpc provides a mock function with given fields: _a0
func (_m *EC2API) CreateDefaultVpc(_a0 *ec2.CreateDefaultVpcInput) (*ec2.CreateDefaultVpcOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateDefaultVpcOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateDefaultVpcInput) *ec2.CreateDefaultVpcOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateDefaultVpcOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateDefaultVpcInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateDefaultVpcRequest provides a mock function with given fields: _a0
func (_m *EC2API) CreateDefaultVpcRequest(_a0 *ec2.CreateDefaultVpcInput) (*request.Request, *ec2.CreateDefaultVpcOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateDefaultVpcInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateDefaultVpcOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateDefaultVpcInput) *ec2.CreateDefaultVpcOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateDefaultVpcOutput)
		}
	}

	return r0, r1
}

// CreateDefaultVpcWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CreateDefaultVpcWithContext(_a0 aws.Context, _a1 *ec2.CreateDefaultVpcInput, _a2 ...request.Option) (*ec2.CreateDefaultVpcOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.CreateDefaultVpcOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CreateDefaultVpcInput, ...request.Option) *ec2.CreateDefaultVpcOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateDefaultVpcOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CreateDefaultVpcInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateDhcpOptions provides a mock function with given fields: _a0
func (_m *EC2API) CreateDhcpOptions(_a0 *ec2.CreateDhcpOptionsInput) (*ec2.CreateDhcpOptionsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateDhcpOptionsOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateDhcpOptionsInput) *ec2.CreateDhcpOptionsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateDhcpOptionsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateDhcpOptionsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateDhcpOptionsRequest provides a mock function with given fields: _a0
func (_m *EC2API) CreateDhcpOptionsRequest(_a0 *ec2.CreateDhcpOptionsInput) (*request.Request, *ec2.CreateDhcpOptionsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateDhcpOptionsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateDhcpOptionsOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateDhcpOptionsInput) *ec2.CreateDhcpOptionsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateDhcpOptionsOutput)
		}
	}

	return r0, r1
}

// CreateDhcpOptionsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CreateDhcpOptionsWithContext(_a0 aws.Context, _a1 *ec2.CreateDhcpOptionsInput, _a2 ...request.Option) (*ec2.CreateDhcpOptionsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.CreateDhcpOptionsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CreateDhcpOptionsInput, ...request.Option) *ec2.CreateDhcpOptionsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateDhcpOptionsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CreateDhcpOptionsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateEgressOnlyInternetGateway provides a mock function with given fields: _a0
func (_m *EC2API) CreateEgressOnlyInternetGateway(_a0 *ec2.CreateEgressOnlyInternetGatewayInput) (*ec2.CreateEgressOnlyInternetGatewayOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateEgressOnlyInternetGatewayOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateEgressOnlyInternetGatewayInput) *ec2.CreateEgressOnlyInternetGatewayOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateEgressOnlyInternetGatewayOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateEgressOnlyInternetGatewayInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateEgressOnlyInternetGatewayRequest provides a mock function with given fields: _a0
func (_m *EC2API) CreateEgressOnlyInternetGatewayRequest(_a0 *ec2.CreateEgressOnlyInternetGatewayInput) (*request.Request, *ec2.CreateEgressOnlyInternetGatewayOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateEgressOnlyInternetGatewayInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateEgressOnlyInternetGatewayOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateEgressOnlyInternetGatewayInput) *ec2.CreateEgressOnlyInternetGatewayOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateEgressOnlyInternetGatewayOutput)
		}
	}

	return r0, r1
}

// CreateEgressOnlyInternetGatewayWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CreateEgressOnlyInternetGatewayWithContext(_a0 aws.Context, _a1 *ec2.CreateEgressOnlyInternetGatewayInput, _a2 ...request.Option) (*ec2.CreateEgressOnlyInternetGatewayOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.CreateEgressOnlyInternetGatewayOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CreateEgressOnlyInternetGatewayInput, ...request.Option) *ec2.CreateEgressOnlyInternetGatewayOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateEgressOnlyInternetGatewayOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CreateEgressOnlyInternetGatewayInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateFleet provides a mock function with given fields: _a0
func (_m *EC2API) CreateFleet(_a0 *ec2.CreateFleetInput) (*ec2.CreateFleetOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateFleetOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateFleetInput) *ec2.CreateFleetOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateFleetOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateFleetInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateFleetRequest provides a mock function with given fields: _a0
func (_m *EC2API) CreateFleetRequest(_a0 *ec2.CreateFleetInput) (*request.Request, *ec2.CreateFleetOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateFleetInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateFleetOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateFleetInput) *ec2.CreateFleetOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateFleetOutput)
		}
	}

	return r0, r1
}

// CreateFleetWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CreateFleetWithContext(_a0 aws.Context, _a1 *ec2.CreateFleetInput, _a2 ...request.Option) (*ec2.CreateFleetOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.CreateFleetOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CreateFleetInput, ...request.Option) *ec2.CreateFleetOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateFleetOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CreateFleetInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateFlowLogs provides a mock function with given fields: _a0
func (_m *EC2API) CreateFlowLogs(_a0 *ec2.CreateFlowLogsInput) (*ec2.CreateFlowLogsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateFlowLogsOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateFlowLogsInput) *ec2.CreateFlowLogsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateFlowLogsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateFlowLogsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateFlowLogsRequest provides a mock function with given fields: _a0
func (_m *EC2API) CreateFlowLogsRequest(_a0 *ec2.CreateFlowLogsInput) (*request.Request, *ec2.CreateFlowLogsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateFlowLogsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateFlowLogsOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateFlowLogsInput) *ec2.CreateFlowLogsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateFlowLogsOutput)
		}
	}

	return r0, r1
}

// CreateFlowLogsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CreateFlowLogsWithContext(_a0 aws.Context, _a1 *ec2.CreateFlowLogsInput, _a2 ...request.Option) (*ec2.CreateFlowLogsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.CreateFlowLogsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CreateFlowLogsInput, ...request.Option) *ec2.CreateFlowLogsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateFlowLogsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CreateFlowLogsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateFpgaImage provides a mock function with given fields: _a0
func (_m *EC2API) CreateFpgaImage(_a0 *ec2.CreateFpgaImageInput) (*ec2.CreateFpgaImageOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateFpgaImageOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateFpgaImageInput) *ec2.CreateFpgaImageOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateFpgaImageOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateFpgaImageInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateFpgaImageRequest provides a mock function with given fields: _a0
func (_m *EC2API) CreateFpgaImageRequest(_a0 *ec2.CreateFpgaImageInput) (*request.Request, *ec2.CreateFpgaImageOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateFpgaImageInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateFpgaImageOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateFpgaImageInput) *ec2.CreateFpgaImageOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateFpgaImageOutput)
		}
	}

	return r0, r1
}

// CreateFpgaImageWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CreateFpgaImageWithContext(_a0 aws.Context, _a1 *ec2.CreateFpgaImageInput, _a2 ...request.Option) (*ec2.CreateFpgaImageOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.CreateFpgaImageOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CreateFpgaImageInput, ...request.Option) *ec2.CreateFpgaImageOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateFpgaImageOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CreateFpgaImageInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateImage provides a mock function with given fields: _a0
func (_m *EC2API) CreateImage(_a0 *ec2.CreateImageInput) (*ec2.CreateImageOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateImageOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateImageInput) *ec2.CreateImageOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateImageOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateImageInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateImageRequest provides a mock function with given fields: _a0
func (_m *EC2API) CreateImageRequest(_a0 *ec2.CreateImageInput) (*request.Request, *ec2.CreateImageOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateImageInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateImageOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateImageInput) *ec2.CreateImageOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateImageOutput)
		}
	}

	return r0, r1
}

// CreateImageWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CreateImageWithContext(_a0 aws.Context, _a1 *ec2.CreateImageInput, _a2 ...request.Option) (*ec2.CreateImageOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.CreateImageOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CreateImageInput, ...request.Option) *ec2.CreateImageOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateImageOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CreateImageInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateInstanceExportTask provides a mock function with given fields: _a0
func (_m *EC2API) CreateInstanceExportTask(_a0 *ec2.CreateInstanceExportTaskInput) (*ec2.CreateInstanceExportTaskOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateInstanceExportTaskOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateInstanceExportTaskInput) *ec2.CreateInstanceExportTaskOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateInstanceExportTaskOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateInstanceExportTaskInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateInstanceExportTaskRequest provides a mock function with given fields: _a0
func (_m *EC2API) CreateInstanceExportTaskRequest(_a0 *ec2.CreateInstanceExportTaskInput) (*request.Request, *ec2.CreateInstanceExportTaskOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateInstanceExportTaskInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateInstanceExportTaskOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateInstanceExportTaskInput) *ec2.CreateInstanceExportTaskOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateInstanceExportTaskOutput)
		}
	}

	return r0, r1
}

// CreateInstanceExportTaskWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CreateInstanceExportTaskWithContext(_a0 aws.Context, _a1 *ec2.CreateInstanceExportTaskInput, _a2 ...request.Option) (*ec2.CreateInstanceExportTaskOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.CreateInstanceExportTaskOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CreateInstanceExportTaskInput, ...request.Option) *ec2.CreateInstanceExportTaskOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateInstanceExportTaskOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CreateInstanceExportTaskInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateInternetGateway provides a mock function with given fields: _a0
func (_m *EC2API) CreateInternetGateway(_a0 *ec2.CreateInternetGatewayInput) (*ec2.CreateInternetGatewayOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateInternetGatewayOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateInternetGatewayInput) *ec2.CreateInternetGatewayOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateInternetGatewayOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateInternetGatewayInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateInternetGatewayRequest provides a mock function with given fields: _a0
func (_m *EC2API) CreateInternetGatewayRequest(_a0 *ec2.CreateInternetGatewayInput) (*request.Request, *ec2.CreateInternetGatewayOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateInternetGatewayInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateInternetGatewayOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateInternetGatewayInput) *ec2.CreateInternetGatewayOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateInternetGatewayOutput)
		}
	}

	return r0, r1
}

// CreateInternetGatewayWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CreateInternetGatewayWithContext(_a0 aws.Context, _a1 *ec2.CreateInternetGatewayInput, _a2 ...request.Option) (*ec2.CreateInternetGatewayOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.CreateInternetGatewayOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CreateInternetGatewayInput, ...request.Option) *ec2.CreateInternetGatewayOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateInternetGatewayOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CreateInternetGatewayInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateKeyPair provides a mock function with given fields: _a0
func (_m *EC2API) CreateKeyPair(_a0 *ec2.CreateKeyPairInput) (*ec2.CreateKeyPairOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateKeyPairOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateKeyPairInput) *ec2.CreateKeyPairOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateKeyPairOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateKeyPairInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateKeyPairRequest provides a mock function with given fields: _a0
func (_m *EC2API) CreateKeyPairRequest(_a0 *ec2.CreateKeyPairInput) (*request.Request, *ec2.CreateKeyPairOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateKeyPairInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateKeyPairOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateKeyPairInput) *ec2.CreateKeyPairOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateKeyPairOutput)
		}
	}

	return r0, r1
}

// CreateKeyPairWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CreateKeyPairWithContext(_a0 aws.Context, _a1 *ec2.CreateKeyPairInput, _a2 ...request.Option) (*ec2.CreateKeyPairOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.CreateKeyPairOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CreateKeyPairInput, ...request.Option) *ec2.CreateKeyPairOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateKeyPairOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CreateKeyPairInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateLaunchTemplate provides a mock function with given fields: _a0
func (_m *EC2API) CreateLaunchTemplate(_a0 *ec2.CreateLaunchTemplateInput) (*ec2.CreateLaunchTemplateOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateLaunchTemplateOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateLaunchTemplateInput) *ec2.CreateLaunchTemplateOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateLaunchTemplateOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateLaunchTemplateInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateLaunchTemplateRequest provides a mock function with given fields: _a0
func (_m *EC2API) CreateLaunchTemplateRequest(_a0 *ec2.CreateLaunchTemplateInput) (*request.Request, *ec2.CreateLaunchTemplateOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateLaunchTemplateInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateLaunchTemplateOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateLaunchTemplateInput) *ec2.CreateLaunchTemplateOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateLaunchTemplateOutput)
		}
	}

	return r0, r1
}

// CreateLaunchTemplateVersion provides a mock function with given fields: _a0
func (_m *EC2API) CreateLaunchTemplateVersion(_a0 *ec2.CreateLaunchTemplateVersionInput) (*ec2.CreateLaunchTemplateVersionOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateLaunchTemplateVersionOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateLaunchTemplateVersionInput) *ec2.CreateLaunchTemplateVersionOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateLaunchTemplateVersionOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateLaunchTemplateVersionInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateLaunchTemplateVersionRequest provides a mock function with given fields: _a0
func (_m *EC2API) CreateLaunchTemplateVersionRequest(_a0 *ec2.CreateLaunchTemplateVersionInput) (*request.Request, *ec2.CreateLaunchTemplateVersionOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateLaunchTemplateVersionInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateLaunchTemplateVersionOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateLaunchTemplateVersionInput) *ec2.CreateLaunchTemplateVersionOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateLaunchTemplateVersionOutput)
		}
	}

	return r0, r1
}

// CreateLaunchTemplateVersionWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CreateLaunchTemplateVersionWithContext(_a0 aws.Context, _a1 *ec2.CreateLaunchTemplateVersionInput, _a2 ...request.Option) (*ec2.CreateLaunchTemplateVersionOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.CreateLaunchTemplateVersionOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CreateLaunchTemplateVersionInput, ...request.Option) *ec2.CreateLaunchTemplateVersionOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateLaunchTemplateVersionOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CreateLaunchTemplateVersionInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateLaunchTemplateWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CreateLaunchTemplateWithContext(_a0 aws.Context, _a1 *ec2.CreateLaunchTemplateInput, _a2 ...request.Option) (*ec2.CreateLaunchTemplateOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.CreateLaunchTemplateOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CreateLaunchTemplateInput, ...request.Option) *ec2.CreateLaunchTemplateOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateLaunchTemplateOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CreateLaunchTemplateInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateNatGateway provides a mock function with given fields: _a0
func (_m *EC2API) CreateNatGateway(_a0 *ec2.CreateNatGatewayInput) (*ec2.CreateNatGatewayOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateNatGatewayOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateNatGatewayInput) *ec2.CreateNatGatewayOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateNatGatewayOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateNatGatewayInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateNatGatewayRequest provides a mock function with given fields: _a0
func (_m *EC2API) CreateNatGatewayRequest(_a0 *ec2.CreateNatGatewayInput) (*request.Request, *ec2.CreateNatGatewayOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateNatGatewayInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateNatGatewayOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateNatGatewayInput) *ec2.CreateNatGatewayOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateNatGatewayOutput)
		}
	}

	return r0, r1
}

// CreateNatGatewayWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CreateNatGatewayWithContext(_a0 aws.Context, _a1 *ec2.CreateNatGatewayInput, _a2 ...request.Option) (*ec2.CreateNatGatewayOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.CreateNatGatewayOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CreateNatGatewayInput, ...request.Option) *ec2.CreateNatGatewayOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateNatGatewayOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CreateNatGatewayInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateNetworkAcl provides a mock function with given fields: _a0
func (_m *EC2API) CreateNetworkAcl(_a0 *ec2.CreateNetworkAclInput) (*ec2.CreateNetworkAclOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateNetworkAclOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateNetworkAclInput) *ec2.CreateNetworkAclOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateNetworkAclOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateNetworkAclInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateNetworkAclEntry provides a mock function with given fields: _a0
func (_m *EC2API) CreateNetworkAclEntry(_a0 *ec2.CreateNetworkAclEntryInput) (*ec2.CreateNetworkAclEntryOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateNetworkAclEntryOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateNetworkAclEntryInput) *ec2.CreateNetworkAclEntryOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateNetworkAclEntryOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateNetworkAclEntryInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateNetworkAclEntryRequest provides a mock function with given fields: _a0
func (_m *EC2API) CreateNetworkAclEntryRequest(_a0 *ec2.CreateNetworkAclEntryInput) (*request.Request, *ec2.CreateNetworkAclEntryOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateNetworkAclEntryInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateNetworkAclEntryOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateNetworkAclEntryInput) *ec2.CreateNetworkAclEntryOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateNetworkAclEntryOutput)
		}
	}

	return r0, r1
}

// CreateNetworkAclEntryWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CreateNetworkAclEntryWithContext(_a0 aws.Context, _a1 *ec2.CreateNetworkAclEntryInput, _a2 ...request.Option) (*ec2.CreateNetworkAclEntryOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.CreateNetworkAclEntryOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CreateNetworkAclEntryInput, ...request.Option) *ec2.CreateNetworkAclEntryOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateNetworkAclEntryOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CreateNetworkAclEntryInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateNetworkAclRequest provides a mock function with given fields: _a0
func (_m *EC2API) CreateNetworkAclRequest(_a0 *ec2.CreateNetworkAclInput) (*request.Request, *ec2.CreateNetworkAclOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateNetworkAclInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateNetworkAclOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateNetworkAclInput) *ec2.CreateNetworkAclOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateNetworkAclOutput)
		}
	}

	return r0, r1
}

// CreateNetworkAclWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CreateNetworkAclWithContext(_a0 aws.Context, _a1 *ec2.CreateNetworkAclInput, _a2 ...request.Option) (*ec2.CreateNetworkAclOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.CreateNetworkAclOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CreateNetworkAclInput, ...request.Option) *ec2.CreateNetworkAclOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateNetworkAclOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CreateNetworkAclInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}


// CreateNetworkInterfacePermission provides a mock function with given fields: _a0
func (_m *EC2API) CreateNetworkInterfacePermission(_a0 *ec2.CreateNetworkInterfacePermissionInput) (*ec2.CreateNetworkInterfacePermissionOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateNetworkInterfacePermissionOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateNetworkInterfacePermissionInput) *ec2.CreateNetworkInterfacePermissionOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateNetworkInterfacePermissionOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateNetworkInterfacePermissionInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateNetworkInterfacePermissionRequest provides a mock function with given fields: _a0
func (_m *EC2API) CreateNetworkInterfacePermissionRequest(_a0 *ec2.CreateNetworkInterfacePermissionInput) (*request.Request, *ec2.CreateNetworkInterfacePermissionOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateNetworkInterfacePermissionInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateNetworkInterfacePermissionOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateNetworkInterfacePermissionInput) *ec2.CreateNetworkInterfacePermissionOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateNetworkInterfacePermissionOutput)
		}
	}

	return r0, r1
}

// CreateNetworkInterfacePermissionWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CreateNetworkInterfacePermissionWithContext(_a0 aws.Context, _a1 *ec2.CreateNetworkInterfacePermissionInput, _a2 ...request.Option) (*ec2.CreateNetworkInterfacePermissionOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.CreateNetworkInterfacePermissionOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CreateNetworkInterfacePermissionInput, ...request.Option) *ec2.CreateNetworkInterfacePermissionOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateNetworkInterfacePermissionOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CreateNetworkInterfacePermissionInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateNetworkInterfaceRequest provides a mock function with given fields: _a0
func (_m *EC2API) CreateNetworkInterfaceRequest(_a0 *ec2.CreateNetworkInterfaceInput) (*request.Request, *ec2.CreateNetworkInterfaceOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateNetworkInterfaceInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateNetworkInterfaceOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateNetworkInterfaceInput) *ec2.CreateNetworkInterfaceOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateNetworkInterfaceOutput)
		}
	}

	return r0, r1
}

// CreateNetworkInterfaceWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CreateNetworkInterfaceWithContext(_a0 aws.Context, _a1 *ec2.CreateNetworkInterfaceInput, _a2 ...request.Option) (*ec2.CreateNetworkInterfaceOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.CreateNetworkInterfaceOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CreateNetworkInterfaceInput, ...request.Option) *ec2.CreateNetworkInterfaceOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateNetworkInterfaceOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CreateNetworkInterfaceInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePlacementGroup provides a mock function with given fields: _a0
func (_m *EC2API) CreatePlacementGroup(_a0 *ec2.CreatePlacementGroupInput) (*ec2.CreatePlacementGroupOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreatePlacementGroupOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreatePlacementGroupInput) *ec2.CreatePlacementGroupOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreatePlacementGroupOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreatePlacementGroupInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePlacementGroupRequest provides a mock function with given fields: _a0
func (_m *EC2API) CreatePlacementGroupRequest(_a0 *ec2.CreatePlacementGroupInput) (*request.Request, *ec2.CreatePlacementGroupOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreatePlacementGroupInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreatePlacementGroupOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreatePlacementGroupInput) *ec2.CreatePlacementGroupOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreatePlacementGroupOutput)
		}
	}

	return r0, r1
}

// CreatePlacementGroupWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CreatePlacementGroupWithContext(_a0 aws.Context, _a1 *ec2.CreatePlacementGroupInput, _a2 ...request.Option) (*ec2.CreatePlacementGroupOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.CreatePlacementGroupOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CreatePlacementGroupInput, ...request.Option) *ec2.CreatePlacementGroupOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreatePlacementGroupOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CreatePlacementGroupInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateReservedInstancesListing provides a mock function with given fields: _a0
func (_m *EC2API) CreateReservedInstancesListing(_a0 *ec2.CreateReservedInstancesListingInput) (*ec2.CreateReservedInstancesListingOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateReservedInstancesListingOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateReservedInstancesListingInput) *ec2.CreateReservedInstancesListingOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateReservedInstancesListingOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateReservedInstancesListingInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateReservedInstancesListingRequest provides a mock function with given fields: _a0
func (_m *EC2API) CreateReservedInstancesListingRequest(_a0 *ec2.CreateReservedInstancesListingInput) (*request.Request, *ec2.CreateReservedInstancesListingOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateReservedInstancesListingInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateReservedInstancesListingOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateReservedInstancesListingInput) *ec2.CreateReservedInstancesListingOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateReservedInstancesListingOutput)
		}
	}

	return r0, r1
}

// CreateReservedInstancesListingWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CreateReservedInstancesListingWithContext(_a0 aws.Context, _a1 *ec2.CreateReservedInstancesListingInput, _a2 ...request.Option) (*ec2.CreateReservedInstancesListingOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.CreateReservedInstancesListingOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CreateReservedInstancesListingInput, ...request.Option) *ec2.CreateReservedInstancesListingOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateReservedInstancesListingOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CreateReservedInstancesListingInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRoute provides a mock function with given fields: _a0
func (_m *EC2API) CreateRoute(_a0 *ec2.CreateRouteInput) (*ec2.CreateRouteOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateRouteOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateRouteInput) *ec2.CreateRouteOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateRouteOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateRouteInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRouteRequest provides a mock function with given fields: _a0
func (_m *EC2API) CreateRouteRequest(_a0 *ec2.CreateRouteInput) (*request.Request, *ec2.CreateRouteOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateRouteInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateRouteOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateRouteInput) *ec2.CreateRouteOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateRouteOutput)
		}
	}

	return r0, r1
}

// CreateRouteTable provides a mock function with given fields: _a0
func (_m *EC2API) CreateRouteTable(_a0 *ec2.CreateRouteTableInput) (*ec2.CreateRouteTableOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateRouteTableOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateRouteTableInput) *ec2.CreateRouteTableOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateRouteTableOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateRouteTableInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRouteTableRequest provides a mock function with given fields: _a0
func (_m *EC2API) CreateRouteTableRequest(_a0 *ec2.CreateRouteTableInput) (*request.Request, *ec2.CreateRouteTableOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateRouteTableInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateRouteTableOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateRouteTableInput) *ec2.CreateRouteTableOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateRouteTableOutput)
		}
	}

	return r0, r1
}

// CreateRouteTableWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CreateRouteTableWithContext(_a0 aws.Context, _a1 *ec2.CreateRouteTableInput, _a2 ...request.Option) (*ec2.CreateRouteTableOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.CreateRouteTableOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CreateRouteTableInput, ...request.Option) *ec2.CreateRouteTableOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateRouteTableOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CreateRouteTableInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRouteWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CreateRouteWithContext(_a0 aws.Context, _a1 *ec2.CreateRouteInput, _a2 ...request.Option) (*ec2.CreateRouteOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.CreateRouteOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CreateRouteInput, ...request.Option) *ec2.CreateRouteOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateRouteOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CreateRouteInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}


// CreateSecurityGroupRequest provides a mock function with given fields: _a0
func (_m *EC2API) CreateSecurityGroupRequest(_a0 *ec2.CreateSecurityGroupInput) (*request.Request, *ec2.CreateSecurityGroupOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateSecurityGroupInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateSecurityGroupOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateSecurityGroupInput) *ec2.CreateSecurityGroupOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateSecurityGroupOutput)
		}
	}

	return r0, r1
}

// CreateSecurityGroupWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CreateSecurityGroupWithContext(_a0 aws.Context, _a1 *ec2.CreateSecurityGroupInput, _a2 ...request.Option) (*ec2.CreateSecurityGroupOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.CreateSecurityGroupOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CreateSecurityGroupInput, ...request.Option) *ec2.CreateSecurityGroupOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateSecurityGroupOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CreateSecurityGroupInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSnapshot provides a mock function with given fields: _a0
func (_m *EC2API) CreateSnapshot(_a0 *ec2.CreateSnapshotInput) (*ec2.Snapshot, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.Snapshot
	if rf, ok := ret.Get(0).(func(*ec2.CreateSnapshotInput) *ec2.Snapshot); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.Snapshot)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateSnapshotInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSnapshotRequest provides a mock function with given fields: _a0
func (_m *EC2API) CreateSnapshotRequest(_a0 *ec2.CreateSnapshotInput) (*request.Request, *ec2.Snapshot) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateSnapshotInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.Snapshot
	if rf, ok := ret.Get(1).(func(*ec2.CreateSnapshotInput) *ec2.Snapshot); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.Snapshot)
		}
	}

	return r0, r1
}

// CreateSnapshotWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CreateSnapshotWithContext(_a0 aws.Context, _a1 *ec2.CreateSnapshotInput, _a2 ...request.Option) (*ec2.Snapshot, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.Snapshot
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CreateSnapshotInput, ...request.Option) *ec2.Snapshot); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.Snapshot)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CreateSnapshotInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSpotDatafeedSubscription provides a mock function with given fields: _a0
func (_m *EC2API) CreateSpotDatafeedSubscription(_a0 *ec2.CreateSpotDatafeedSubscriptionInput) (*ec2.CreateSpotDatafeedSubscriptionOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateSpotDatafeedSubscriptionOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateSpotDatafeedSubscriptionInput) *ec2.CreateSpotDatafeedSubscriptionOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateSpotDatafeedSubscriptionOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateSpotDatafeedSubscriptionInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSpotDatafeedSubscriptionRequest provides a mock function with given fields: _a0
func (_m *EC2API) CreateSpotDatafeedSubscriptionRequest(_a0 *ec2.CreateSpotDatafeedSubscriptionInput) (*request.Request, *ec2.CreateSpotDatafeedSubscriptionOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateSpotDatafeedSubscriptionInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateSpotDatafeedSubscriptionOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateSpotDatafeedSubscriptionInput) *ec2.CreateSpotDatafeedSubscriptionOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateSpotDatafeedSubscriptionOutput)
		}
	}

	return r0, r1
}

// CreateSpotDatafeedSubscriptionWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CreateSpotDatafeedSubscriptionWithContext(_a0 aws.Context, _a1 *ec2.CreateSpotDatafeedSubscriptionInput, _a2 ...request.Option) (*ec2.CreateSpotDatafeedSubscriptionOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.CreateSpotDatafeedSubscriptionOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CreateSpotDatafeedSubscriptionInput, ...request.Option) *ec2.CreateSpotDatafeedSubscriptionOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateSpotDatafeedSubscriptionOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CreateSpotDatafeedSubscriptionInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSubnetRequest provides a mock function with given fields: _a0
func (_m *EC2API) CreateSubnetRequest(_a0 *ec2.CreateSubnetInput) (*request.Request, *ec2.CreateSubnetOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateSubnetInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateSubnetOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateSubnetInput) *ec2.CreateSubnetOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateSubnetOutput)
		}
	}

	return r0, r1
}

// CreateSubnetWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CreateSubnetWithContext(_a0 aws.Context, _a1 *ec2.CreateSubnetInput, _a2 ...request.Option) (*ec2.CreateSubnetOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.CreateSubnetOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CreateSubnetInput, ...request.Option) *ec2.CreateSubnetOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateSubnetOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CreateSubnetInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}



// CreateTagsRequest provides a mock function with given fields: _a0
func (_m *EC2API) CreateTagsRequest(_a0 *ec2.CreateTagsInput) (*request.Request, *ec2.CreateTagsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateTagsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateTagsOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateTagsInput) *ec2.CreateTagsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateTagsOutput)
		}
	}

	return r0, r1
}

// CreateTagsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CreateTagsWithContext(_a0 aws.Context, _a1 *ec2.CreateTagsInput, _a2 ...request.Option) (*ec2.CreateTagsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.CreateTagsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CreateTagsInput, ...request.Option) *ec2.CreateTagsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateTagsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CreateTagsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateVolume provides a mock function with given fields: _a0
func (_m *EC2API) CreateVolume(_a0 *ec2.CreateVolumeInput) (*ec2.Volume, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.Volume
	if rf, ok := ret.Get(0).(func(*ec2.CreateVolumeInput) *ec2.Volume); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.Volume)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateVolumeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateVolumeRequest provides a mock function with given fields: _a0
func (_m *EC2API) CreateVolumeRequest(_a0 *ec2.CreateVolumeInput) (*request.Request, *ec2.Volume) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateVolumeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.Volume
	if rf, ok := ret.Get(1).(func(*ec2.CreateVolumeInput) *ec2.Volume); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.Volume)
		}
	}

	return r0, r1
}

// CreateVolumeWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CreateVolumeWithContext(_a0 aws.Context, _a1 *ec2.CreateVolumeInput, _a2 ...request.Option) (*ec2.Volume, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.Volume
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CreateVolumeInput, ...request.Option) *ec2.Volume); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.Volume)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CreateVolumeInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateVpc provides a mock function with given fields: _a0
func (_m *EC2API) CreateVpc(_a0 *ec2.CreateVpcInput) (*ec2.CreateVpcOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateVpcOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateVpcInput) *ec2.CreateVpcOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateVpcOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateVpcInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateVpcEndpoint provides a mock function with given fields: _a0
func (_m *EC2API) CreateVpcEndpoint(_a0 *ec2.CreateVpcEndpointInput) (*ec2.CreateVpcEndpointOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateVpcEndpointOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateVpcEndpointInput) *ec2.CreateVpcEndpointOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateVpcEndpointOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateVpcEndpointInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateVpcEndpointConnectionNotification provides a mock function with given fields: _a0
func (_m *EC2API) CreateVpcEndpointConnectionNotification(_a0 *ec2.CreateVpcEndpointConnectionNotificationInput) (*ec2.CreateVpcEndpointConnectionNotificationOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateVpcEndpointConnectionNotificationOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateVpcEndpointConnectionNotificationInput) *ec2.CreateVpcEndpointConnectionNotificationOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateVpcEndpointConnectionNotificationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateVpcEndpointConnectionNotificationInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateVpcEndpointConnectionNotificationRequest provides a mock function with given fields: _a0
func (_m *EC2API) CreateVpcEndpointConnectionNotificationRequest(_a0 *ec2.CreateVpcEndpointConnectionNotificationInput) (*request.Request, *ec2.CreateVpcEndpointConnectionNotificationOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateVpcEndpointConnectionNotificationInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateVpcEndpointConnectionNotificationOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateVpcEndpointConnectionNotificationInput) *ec2.CreateVpcEndpointConnectionNotificationOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateVpcEndpointConnectionNotificationOutput)
		}
	}

	return r0, r1
}

// CreateVpcEndpointConnectionNotificationWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CreateVpcEndpointConnectionNotificationWithContext(_a0 aws.Context, _a1 *ec2.CreateVpcEndpointConnectionNotificationInput, _a2 ...request.Option) (*ec2.CreateVpcEndpointConnectionNotificationOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.CreateVpcEndpointConnectionNotificationOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CreateVpcEndpointConnectionNotificationInput, ...request.Option) *ec2.CreateVpcEndpointConnectionNotificationOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateVpcEndpointConnectionNotificationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CreateVpcEndpointConnectionNotificationInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateVpcEndpointRequest provides a mock function with given fields: _a0
func (_m *EC2API) CreateVpcEndpointRequest(_a0 *ec2.CreateVpcEndpointInput) (*request.Request, *ec2.CreateVpcEndpointOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateVpcEndpointInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateVpcEndpointOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateVpcEndpointInput) *ec2.CreateVpcEndpointOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateVpcEndpointOutput)
		}
	}

	return r0, r1
}

// CreateVpcEndpointServiceConfiguration provides a mock function with given fields: _a0
func (_m *EC2API) CreateVpcEndpointServiceConfiguration(_a0 *ec2.CreateVpcEndpointServiceConfigurationInput) (*ec2.CreateVpcEndpointServiceConfigurationOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateVpcEndpointServiceConfigurationOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateVpcEndpointServiceConfigurationInput) *ec2.CreateVpcEndpointServiceConfigurationOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateVpcEndpointServiceConfigurationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateVpcEndpointServiceConfigurationInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateVpcEndpointServiceConfigurationRequest provides a mock function with given fields: _a0
func (_m *EC2API) CreateVpcEndpointServiceConfigurationRequest(_a0 *ec2.CreateVpcEndpointServiceConfigurationInput) (*request.Request, *ec2.CreateVpcEndpointServiceConfigurationOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateVpcEndpointServiceConfigurationInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateVpcEndpointServiceConfigurationOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateVpcEndpointServiceConfigurationInput) *ec2.CreateVpcEndpointServiceConfigurationOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateVpcEndpointServiceConfigurationOutput)
		}
	}

	return r0, r1
}

// CreateVpcEndpointServiceConfigurationWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CreateVpcEndpointServiceConfigurationWithContext(_a0 aws.Context, _a1 *ec2.CreateVpcEndpointServiceConfigurationInput, _a2 ...request.Option) (*ec2.CreateVpcEndpointServiceConfigurationOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.CreateVpcEndpointServiceConfigurationOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CreateVpcEndpointServiceConfigurationInput, ...request.Option) *ec2.CreateVpcEndpointServiceConfigurationOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateVpcEndpointServiceConfigurationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CreateVpcEndpointServiceConfigurationInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateVpcEndpointWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CreateVpcEndpointWithContext(_a0 aws.Context, _a1 *ec2.CreateVpcEndpointInput, _a2 ...request.Option) (*ec2.CreateVpcEndpointOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.CreateVpcEndpointOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CreateVpcEndpointInput, ...request.Option) *ec2.CreateVpcEndpointOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateVpcEndpointOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CreateVpcEndpointInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateVpcPeeringConnection provides a mock function with given fields: _a0
func (_m *EC2API) CreateVpcPeeringConnection(_a0 *ec2.CreateVpcPeeringConnectionInput) (*ec2.CreateVpcPeeringConnectionOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateVpcPeeringConnectionOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateVpcPeeringConnectionInput) *ec2.CreateVpcPeeringConnectionOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateVpcPeeringConnectionOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateVpcPeeringConnectionInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateVpcPeeringConnectionRequest provides a mock function with given fields: _a0
func (_m *EC2API) CreateVpcPeeringConnectionRequest(_a0 *ec2.CreateVpcPeeringConnectionInput) (*request.Request, *ec2.CreateVpcPeeringConnectionOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateVpcPeeringConnectionInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateVpcPeeringConnectionOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateVpcPeeringConnectionInput) *ec2.CreateVpcPeeringConnectionOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateVpcPeeringConnectionOutput)
		}
	}

	return r0, r1
}

// CreateVpcPeeringConnectionWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CreateVpcPeeringConnectionWithContext(_a0 aws.Context, _a1 *ec2.CreateVpcPeeringConnectionInput, _a2 ...request.Option) (*ec2.CreateVpcPeeringConnectionOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.CreateVpcPeeringConnectionOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CreateVpcPeeringConnectionInput, ...request.Option) *ec2.CreateVpcPeeringConnectionOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateVpcPeeringConnectionOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CreateVpcPeeringConnectionInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateVpcRequest provides a mock function with given fields: _a0
func (_m *EC2API) CreateVpcRequest(_a0 *ec2.CreateVpcInput) (*request.Request, *ec2.CreateVpcOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateVpcInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateVpcOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateVpcInput) *ec2.CreateVpcOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateVpcOutput)
		}
	}

	return r0, r1
}

// CreateVpcWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CreateVpcWithContext(_a0 aws.Context, _a1 *ec2.CreateVpcInput, _a2 ...request.Option) (*ec2.CreateVpcOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.CreateVpcOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CreateVpcInput, ...request.Option) *ec2.CreateVpcOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateVpcOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CreateVpcInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateVpnConnection provides a mock function with given fields: _a0
func (_m *EC2API) CreateVpnConnection(_a0 *ec2.CreateVpnConnectionInput) (*ec2.CreateVpnConnectionOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateVpnConnectionOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateVpnConnectionInput) *ec2.CreateVpnConnectionOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateVpnConnectionOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateVpnConnectionInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateVpnConnectionRequest provides a mock function with given fields: _a0
func (_m *EC2API) CreateVpnConnectionRequest(_a0 *ec2.CreateVpnConnectionInput) (*request.Request, *ec2.CreateVpnConnectionOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateVpnConnectionInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateVpnConnectionOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateVpnConnectionInput) *ec2.CreateVpnConnectionOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateVpnConnectionOutput)
		}
	}

	return r0, r1
}

// CreateVpnConnectionRoute provides a mock function with given fields: _a0
func (_m *EC2API) CreateVpnConnectionRoute(_a0 *ec2.CreateVpnConnectionRouteInput) (*ec2.CreateVpnConnectionRouteOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateVpnConnectionRouteOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateVpnConnectionRouteInput) *ec2.CreateVpnConnectionRouteOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateVpnConnectionRouteOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateVpnConnectionRouteInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateVpnConnectionRouteRequest provides a mock function with given fields: _a0
func (_m *EC2API) CreateVpnConnectionRouteRequest(_a0 *ec2.CreateVpnConnectionRouteInput) (*request.Request, *ec2.CreateVpnConnectionRouteOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateVpnConnectionRouteInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateVpnConnectionRouteOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateVpnConnectionRouteInput) *ec2.CreateVpnConnectionRouteOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateVpnConnectionRouteOutput)
		}
	}

	return r0, r1
}

// CreateVpnConnectionRouteWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CreateVpnConnectionRouteWithContext(_a0 aws.Context, _a1 *ec2.CreateVpnConnectionRouteInput, _a2 ...request.Option) (*ec2.CreateVpnConnectionRouteOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.CreateVpnConnectionRouteOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CreateVpnConnectionRouteInput, ...request.Option) *ec2.CreateVpnConnectionRouteOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateVpnConnectionRouteOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CreateVpnConnectionRouteInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateVpnConnectionWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CreateVpnConnectionWithContext(_a0 aws.Context, _a1 *ec2.CreateVpnConnectionInput, _a2 ...request.Option) (*ec2.CreateVpnConnectionOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.CreateVpnConnectionOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CreateVpnConnectionInput, ...request.Option) *ec2.CreateVpnConnectionOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateVpnConnectionOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CreateVpnConnectionInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateVpnGateway provides a mock function with given fields: _a0
func (_m *EC2API) CreateVpnGateway(_a0 *ec2.CreateVpnGatewayInput) (*ec2.CreateVpnGatewayOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.CreateVpnGatewayOutput
	if rf, ok := ret.Get(0).(func(*ec2.CreateVpnGatewayInput) *ec2.CreateVpnGatewayOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateVpnGatewayOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.CreateVpnGatewayInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateVpnGatewayRequest provides a mock function with given fields: _a0
func (_m *EC2API) CreateVpnGatewayRequest(_a0 *ec2.CreateVpnGatewayInput) (*request.Request, *ec2.CreateVpnGatewayOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.CreateVpnGatewayInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.CreateVpnGatewayOutput
	if rf, ok := ret.Get(1).(func(*ec2.CreateVpnGatewayInput) *ec2.CreateVpnGatewayOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.CreateVpnGatewayOutput)
		}
	}

	return r0, r1
}

// CreateVpnGatewayWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) CreateVpnGatewayWithContext(_a0 aws.Context, _a1 *ec2.CreateVpnGatewayInput, _a2 ...request.Option) (*ec2.CreateVpnGatewayOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.CreateVpnGatewayOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.CreateVpnGatewayInput, ...request.Option) *ec2.CreateVpnGatewayOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.CreateVpnGatewayOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.CreateVpnGatewayInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCustomerGateway provides a mock function with given fields: _a0
func (_m *EC2API) DeleteCustomerGateway(_a0 *ec2.DeleteCustomerGatewayInput) (*ec2.DeleteCustomerGatewayOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteCustomerGatewayOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteCustomerGatewayInput) *ec2.DeleteCustomerGatewayOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteCustomerGatewayOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteCustomerGatewayInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCustomerGatewayRequest provides a mock function with given fields: _a0
func (_m *EC2API) DeleteCustomerGatewayRequest(_a0 *ec2.DeleteCustomerGatewayInput) (*request.Request, *ec2.DeleteCustomerGatewayOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteCustomerGatewayInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteCustomerGatewayOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteCustomerGatewayInput) *ec2.DeleteCustomerGatewayOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteCustomerGatewayOutput)
		}
	}

	return r0, r1
}

// DeleteCustomerGatewayWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DeleteCustomerGatewayWithContext(_a0 aws.Context, _a1 *ec2.DeleteCustomerGatewayInput, _a2 ...request.Option) (*ec2.DeleteCustomerGatewayOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DeleteCustomerGatewayOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DeleteCustomerGatewayInput, ...request.Option) *ec2.DeleteCustomerGatewayOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteCustomerGatewayOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DeleteCustomerGatewayInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteDhcpOptions provides a mock function with given fields: _a0
func (_m *EC2API) DeleteDhcpOptions(_a0 *ec2.DeleteDhcpOptionsInput) (*ec2.DeleteDhcpOptionsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteDhcpOptionsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteDhcpOptionsInput) *ec2.DeleteDhcpOptionsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteDhcpOptionsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteDhcpOptionsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteDhcpOptionsRequest provides a mock function with given fields: _a0
func (_m *EC2API) DeleteDhcpOptionsRequest(_a0 *ec2.DeleteDhcpOptionsInput) (*request.Request, *ec2.DeleteDhcpOptionsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteDhcpOptionsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteDhcpOptionsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteDhcpOptionsInput) *ec2.DeleteDhcpOptionsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteDhcpOptionsOutput)
		}
	}

	return r0, r1
}

// DeleteDhcpOptionsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DeleteDhcpOptionsWithContext(_a0 aws.Context, _a1 *ec2.DeleteDhcpOptionsInput, _a2 ...request.Option) (*ec2.DeleteDhcpOptionsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DeleteDhcpOptionsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DeleteDhcpOptionsInput, ...request.Option) *ec2.DeleteDhcpOptionsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteDhcpOptionsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DeleteDhcpOptionsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteEgressOnlyInternetGateway provides a mock function with given fields: _a0
func (_m *EC2API) DeleteEgressOnlyInternetGateway(_a0 *ec2.DeleteEgressOnlyInternetGatewayInput) (*ec2.DeleteEgressOnlyInternetGatewayOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteEgressOnlyInternetGatewayOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteEgressOnlyInternetGatewayInput) *ec2.DeleteEgressOnlyInternetGatewayOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteEgressOnlyInternetGatewayOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteEgressOnlyInternetGatewayInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteEgressOnlyInternetGatewayRequest provides a mock function with given fields: _a0
func (_m *EC2API) DeleteEgressOnlyInternetGatewayRequest(_a0 *ec2.DeleteEgressOnlyInternetGatewayInput) (*request.Request, *ec2.DeleteEgressOnlyInternetGatewayOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteEgressOnlyInternetGatewayInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteEgressOnlyInternetGatewayOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteEgressOnlyInternetGatewayInput) *ec2.DeleteEgressOnlyInternetGatewayOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteEgressOnlyInternetGatewayOutput)
		}
	}

	return r0, r1
}

// DeleteEgressOnlyInternetGatewayWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DeleteEgressOnlyInternetGatewayWithContext(_a0 aws.Context, _a1 *ec2.DeleteEgressOnlyInternetGatewayInput, _a2 ...request.Option) (*ec2.DeleteEgressOnlyInternetGatewayOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DeleteEgressOnlyInternetGatewayOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DeleteEgressOnlyInternetGatewayInput, ...request.Option) *ec2.DeleteEgressOnlyInternetGatewayOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteEgressOnlyInternetGatewayOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DeleteEgressOnlyInternetGatewayInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteFleets provides a mock function with given fields: _a0
func (_m *EC2API) DeleteFleets(_a0 *ec2.DeleteFleetsInput) (*ec2.DeleteFleetsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteFleetsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteFleetsInput) *ec2.DeleteFleetsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteFleetsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteFleetsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteFleetsRequest provides a mock function with given fields: _a0
func (_m *EC2API) DeleteFleetsRequest(_a0 *ec2.DeleteFleetsInput) (*request.Request, *ec2.DeleteFleetsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteFleetsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteFleetsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteFleetsInput) *ec2.DeleteFleetsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteFleetsOutput)
		}
	}

	return r0, r1
}

// DeleteFleetsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DeleteFleetsWithContext(_a0 aws.Context, _a1 *ec2.DeleteFleetsInput, _a2 ...request.Option) (*ec2.DeleteFleetsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DeleteFleetsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DeleteFleetsInput, ...request.Option) *ec2.DeleteFleetsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteFleetsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DeleteFleetsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteFlowLogs provides a mock function with given fields: _a0
func (_m *EC2API) DeleteFlowLogs(_a0 *ec2.DeleteFlowLogsInput) (*ec2.DeleteFlowLogsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteFlowLogsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteFlowLogsInput) *ec2.DeleteFlowLogsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteFlowLogsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteFlowLogsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteFlowLogsRequest provides a mock function with given fields: _a0
func (_m *EC2API) DeleteFlowLogsRequest(_a0 *ec2.DeleteFlowLogsInput) (*request.Request, *ec2.DeleteFlowLogsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteFlowLogsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteFlowLogsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteFlowLogsInput) *ec2.DeleteFlowLogsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteFlowLogsOutput)
		}
	}

	return r0, r1
}

// DeleteFlowLogsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DeleteFlowLogsWithContext(_a0 aws.Context, _a1 *ec2.DeleteFlowLogsInput, _a2 ...request.Option) (*ec2.DeleteFlowLogsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DeleteFlowLogsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DeleteFlowLogsInput, ...request.Option) *ec2.DeleteFlowLogsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteFlowLogsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DeleteFlowLogsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteFpgaImage provides a mock function with given fields: _a0
func (_m *EC2API) DeleteFpgaImage(_a0 *ec2.DeleteFpgaImageInput) (*ec2.DeleteFpgaImageOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteFpgaImageOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteFpgaImageInput) *ec2.DeleteFpgaImageOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteFpgaImageOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteFpgaImageInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteFpgaImageRequest provides a mock function with given fields: _a0
func (_m *EC2API) DeleteFpgaImageRequest(_a0 *ec2.DeleteFpgaImageInput) (*request.Request, *ec2.DeleteFpgaImageOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteFpgaImageInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteFpgaImageOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteFpgaImageInput) *ec2.DeleteFpgaImageOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteFpgaImageOutput)
		}
	}

	return r0, r1
}

// DeleteFpgaImageWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DeleteFpgaImageWithContext(_a0 aws.Context, _a1 *ec2.DeleteFpgaImageInput, _a2 ...request.Option) (*ec2.DeleteFpgaImageOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DeleteFpgaImageOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DeleteFpgaImageInput, ...request.Option) *ec2.DeleteFpgaImageOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteFpgaImageOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DeleteFpgaImageInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteInternetGateway provides a mock function with given fields: _a0
func (_m *EC2API) DeleteInternetGateway(_a0 *ec2.DeleteInternetGatewayInput) (*ec2.DeleteInternetGatewayOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteInternetGatewayOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteInternetGatewayInput) *ec2.DeleteInternetGatewayOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteInternetGatewayOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteInternetGatewayInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteInternetGatewayRequest provides a mock function with given fields: _a0
func (_m *EC2API) DeleteInternetGatewayRequest(_a0 *ec2.DeleteInternetGatewayInput) (*request.Request, *ec2.DeleteInternetGatewayOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteInternetGatewayInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteInternetGatewayOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteInternetGatewayInput) *ec2.DeleteInternetGatewayOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteInternetGatewayOutput)
		}
	}

	return r0, r1
}

// DeleteInternetGatewayWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DeleteInternetGatewayWithContext(_a0 aws.Context, _a1 *ec2.DeleteInternetGatewayInput, _a2 ...request.Option) (*ec2.DeleteInternetGatewayOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DeleteInternetGatewayOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DeleteInternetGatewayInput, ...request.Option) *ec2.DeleteInternetGatewayOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteInternetGatewayOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DeleteInternetGatewayInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteKeyPair provides a mock function with given fields: _a0
func (_m *EC2API) DeleteKeyPair(_a0 *ec2.DeleteKeyPairInput) (*ec2.DeleteKeyPairOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteKeyPairOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteKeyPairInput) *ec2.DeleteKeyPairOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteKeyPairOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteKeyPairInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteKeyPairRequest provides a mock function with given fields: _a0
func (_m *EC2API) DeleteKeyPairRequest(_a0 *ec2.DeleteKeyPairInput) (*request.Request, *ec2.DeleteKeyPairOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteKeyPairInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteKeyPairOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteKeyPairInput) *ec2.DeleteKeyPairOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteKeyPairOutput)
		}
	}

	return r0, r1
}

// DeleteKeyPairWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DeleteKeyPairWithContext(_a0 aws.Context, _a1 *ec2.DeleteKeyPairInput, _a2 ...request.Option) (*ec2.DeleteKeyPairOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DeleteKeyPairOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DeleteKeyPairInput, ...request.Option) *ec2.DeleteKeyPairOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteKeyPairOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DeleteKeyPairInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteLaunchTemplate provides a mock function with given fields: _a0
func (_m *EC2API) DeleteLaunchTemplate(_a0 *ec2.DeleteLaunchTemplateInput) (*ec2.DeleteLaunchTemplateOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteLaunchTemplateOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteLaunchTemplateInput) *ec2.DeleteLaunchTemplateOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteLaunchTemplateOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteLaunchTemplateInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteLaunchTemplateRequest provides a mock function with given fields: _a0
func (_m *EC2API) DeleteLaunchTemplateRequest(_a0 *ec2.DeleteLaunchTemplateInput) (*request.Request, *ec2.DeleteLaunchTemplateOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteLaunchTemplateInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteLaunchTemplateOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteLaunchTemplateInput) *ec2.DeleteLaunchTemplateOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteLaunchTemplateOutput)
		}
	}

	return r0, r1
}

// DeleteLaunchTemplateVersions provides a mock function with given fields: _a0
func (_m *EC2API) DeleteLaunchTemplateVersions(_a0 *ec2.DeleteLaunchTemplateVersionsInput) (*ec2.DeleteLaunchTemplateVersionsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteLaunchTemplateVersionsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteLaunchTemplateVersionsInput) *ec2.DeleteLaunchTemplateVersionsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteLaunchTemplateVersionsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteLaunchTemplateVersionsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteLaunchTemplateVersionsRequest provides a mock function with given fields: _a0
func (_m *EC2API) DeleteLaunchTemplateVersionsRequest(_a0 *ec2.DeleteLaunchTemplateVersionsInput) (*request.Request, *ec2.DeleteLaunchTemplateVersionsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteLaunchTemplateVersionsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteLaunchTemplateVersionsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteLaunchTemplateVersionsInput) *ec2.DeleteLaunchTemplateVersionsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteLaunchTemplateVersionsOutput)
		}
	}

	return r0, r1
}

// DeleteLaunchTemplateVersionsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DeleteLaunchTemplateVersionsWithContext(_a0 aws.Context, _a1 *ec2.DeleteLaunchTemplateVersionsInput, _a2 ...request.Option) (*ec2.DeleteLaunchTemplateVersionsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DeleteLaunchTemplateVersionsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DeleteLaunchTemplateVersionsInput, ...request.Option) *ec2.DeleteLaunchTemplateVersionsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteLaunchTemplateVersionsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DeleteLaunchTemplateVersionsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteLaunchTemplateWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DeleteLaunchTemplateWithContext(_a0 aws.Context, _a1 *ec2.DeleteLaunchTemplateInput, _a2 ...request.Option) (*ec2.DeleteLaunchTemplateOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DeleteLaunchTemplateOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DeleteLaunchTemplateInput, ...request.Option) *ec2.DeleteLaunchTemplateOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteLaunchTemplateOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DeleteLaunchTemplateInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteNatGateway provides a mock function with given fields: _a0
func (_m *EC2API) DeleteNatGateway(_a0 *ec2.DeleteNatGatewayInput) (*ec2.DeleteNatGatewayOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteNatGatewayOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteNatGatewayInput) *ec2.DeleteNatGatewayOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteNatGatewayOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteNatGatewayInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteNatGatewayRequest provides a mock function with given fields: _a0
func (_m *EC2API) DeleteNatGatewayRequest(_a0 *ec2.DeleteNatGatewayInput) (*request.Request, *ec2.DeleteNatGatewayOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteNatGatewayInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteNatGatewayOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteNatGatewayInput) *ec2.DeleteNatGatewayOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteNatGatewayOutput)
		}
	}

	return r0, r1
}

// DeleteNatGatewayWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DeleteNatGatewayWithContext(_a0 aws.Context, _a1 *ec2.DeleteNatGatewayInput, _a2 ...request.Option) (*ec2.DeleteNatGatewayOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DeleteNatGatewayOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DeleteNatGatewayInput, ...request.Option) *ec2.DeleteNatGatewayOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteNatGatewayOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DeleteNatGatewayInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteNetworkAcl provides a mock function with given fields: _a0
func (_m *EC2API) DeleteNetworkAcl(_a0 *ec2.DeleteNetworkAclInput) (*ec2.DeleteNetworkAclOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteNetworkAclOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteNetworkAclInput) *ec2.DeleteNetworkAclOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteNetworkAclOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteNetworkAclInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteNetworkAclEntry provides a mock function with given fields: _a0
func (_m *EC2API) DeleteNetworkAclEntry(_a0 *ec2.DeleteNetworkAclEntryInput) (*ec2.DeleteNetworkAclEntryOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteNetworkAclEntryOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteNetworkAclEntryInput) *ec2.DeleteNetworkAclEntryOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteNetworkAclEntryOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteNetworkAclEntryInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteNetworkAclEntryRequest provides a mock function with given fields: _a0
func (_m *EC2API) DeleteNetworkAclEntryRequest(_a0 *ec2.DeleteNetworkAclEntryInput) (*request.Request, *ec2.DeleteNetworkAclEntryOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteNetworkAclEntryInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteNetworkAclEntryOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteNetworkAclEntryInput) *ec2.DeleteNetworkAclEntryOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteNetworkAclEntryOutput)
		}
	}

	return r0, r1
}

// DeleteNetworkAclEntryWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DeleteNetworkAclEntryWithContext(_a0 aws.Context, _a1 *ec2.DeleteNetworkAclEntryInput, _a2 ...request.Option) (*ec2.DeleteNetworkAclEntryOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DeleteNetworkAclEntryOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DeleteNetworkAclEntryInput, ...request.Option) *ec2.DeleteNetworkAclEntryOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteNetworkAclEntryOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DeleteNetworkAclEntryInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteNetworkAclRequest provides a mock function with given fields: _a0
func (_m *EC2API) DeleteNetworkAclRequest(_a0 *ec2.DeleteNetworkAclInput) (*request.Request, *ec2.DeleteNetworkAclOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteNetworkAclInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteNetworkAclOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteNetworkAclInput) *ec2.DeleteNetworkAclOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteNetworkAclOutput)
		}
	}

	return r0, r1
}

// DeleteNetworkAclWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DeleteNetworkAclWithContext(_a0 aws.Context, _a1 *ec2.DeleteNetworkAclInput, _a2 ...request.Option) (*ec2.DeleteNetworkAclOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DeleteNetworkAclOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DeleteNetworkAclInput, ...request.Option) *ec2.DeleteNetworkAclOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteNetworkAclOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DeleteNetworkAclInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteNetworkInterfacePermission provides a mock function with given fields: _a0
func (_m *EC2API) DeleteNetworkInterfacePermission(_a0 *ec2.DeleteNetworkInterfacePermissionInput) (*ec2.DeleteNetworkInterfacePermissionOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteNetworkInterfacePermissionOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteNetworkInterfacePermissionInput) *ec2.DeleteNetworkInterfacePermissionOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteNetworkInterfacePermissionOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteNetworkInterfacePermissionInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteNetworkInterfacePermissionRequest provides a mock function with given fields: _a0
func (_m *EC2API) DeleteNetworkInterfacePermissionRequest(_a0 *ec2.DeleteNetworkInterfacePermissionInput) (*request.Request, *ec2.DeleteNetworkInterfacePermissionOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteNetworkInterfacePermissionInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteNetworkInterfacePermissionOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteNetworkInterfacePermissionInput) *ec2.DeleteNetworkInterfacePermissionOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteNetworkInterfacePermissionOutput)
		}
	}

	return r0, r1
}

// DeleteNetworkInterfacePermissionWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DeleteNetworkInterfacePermissionWithContext(_a0 aws.Context, _a1 *ec2.DeleteNetworkInterfacePermissionInput, _a2 ...request.Option) (*ec2.DeleteNetworkInterfacePermissionOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DeleteNetworkInterfacePermissionOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DeleteNetworkInterfacePermissionInput, ...request.Option) *ec2.DeleteNetworkInterfacePermissionOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteNetworkInterfacePermissionOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DeleteNetworkInterfacePermissionInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteNetworkInterfaceRequest provides a mock function with given fields: _a0
func (_m *EC2API) DeleteNetworkInterfaceRequest(_a0 *ec2.DeleteNetworkInterfaceInput) (*request.Request, *ec2.DeleteNetworkInterfaceOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteNetworkInterfaceInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteNetworkInterfaceOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteNetworkInterfaceInput) *ec2.DeleteNetworkInterfaceOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteNetworkInterfaceOutput)
		}
	}

	return r0, r1
}

// DeleteNetworkInterfaceWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DeleteNetworkInterfaceWithContext(_a0 aws.Context, _a1 *ec2.DeleteNetworkInterfaceInput, _a2 ...request.Option) (*ec2.DeleteNetworkInterfaceOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DeleteNetworkInterfaceOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DeleteNetworkInterfaceInput, ...request.Option) *ec2.DeleteNetworkInterfaceOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteNetworkInterfaceOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DeleteNetworkInterfaceInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeletePlacementGroup provides a mock function with given fields: _a0
func (_m *EC2API) DeletePlacementGroup(_a0 *ec2.DeletePlacementGroupInput) (*ec2.DeletePlacementGroupOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeletePlacementGroupOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeletePlacementGroupInput) *ec2.DeletePlacementGroupOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeletePlacementGroupOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeletePlacementGroupInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeletePlacementGroupRequest provides a mock function with given fields: _a0
func (_m *EC2API) DeletePlacementGroupRequest(_a0 *ec2.DeletePlacementGroupInput) (*request.Request, *ec2.DeletePlacementGroupOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeletePlacementGroupInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeletePlacementGroupOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeletePlacementGroupInput) *ec2.DeletePlacementGroupOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeletePlacementGroupOutput)
		}
	}

	return r0, r1
}

// DeletePlacementGroupWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DeletePlacementGroupWithContext(_a0 aws.Context, _a1 *ec2.DeletePlacementGroupInput, _a2 ...request.Option) (*ec2.DeletePlacementGroupOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DeletePlacementGroupOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DeletePlacementGroupInput, ...request.Option) *ec2.DeletePlacementGroupOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeletePlacementGroupOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DeletePlacementGroupInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteRoute provides a mock function with given fields: _a0
func (_m *EC2API) DeleteRoute(_a0 *ec2.DeleteRouteInput) (*ec2.DeleteRouteOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteRouteOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteRouteInput) *ec2.DeleteRouteOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteRouteOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteRouteInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteRouteRequest provides a mock function with given fields: _a0
func (_m *EC2API) DeleteRouteRequest(_a0 *ec2.DeleteRouteInput) (*request.Request, *ec2.DeleteRouteOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteRouteInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteRouteOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteRouteInput) *ec2.DeleteRouteOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteRouteOutput)
		}
	}

	return r0, r1
}

// DeleteRouteTable provides a mock function with given fields: _a0
func (_m *EC2API) DeleteRouteTable(_a0 *ec2.DeleteRouteTableInput) (*ec2.DeleteRouteTableOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteRouteTableOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteRouteTableInput) *ec2.DeleteRouteTableOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteRouteTableOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteRouteTableInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteRouteTableRequest provides a mock function with given fields: _a0
func (_m *EC2API) DeleteRouteTableRequest(_a0 *ec2.DeleteRouteTableInput) (*request.Request, *ec2.DeleteRouteTableOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteRouteTableInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteRouteTableOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteRouteTableInput) *ec2.DeleteRouteTableOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteRouteTableOutput)
		}
	}

	return r0, r1
}

// DeleteRouteTableWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DeleteRouteTableWithContext(_a0 aws.Context, _a1 *ec2.DeleteRouteTableInput, _a2 ...request.Option) (*ec2.DeleteRouteTableOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DeleteRouteTableOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DeleteRouteTableInput, ...request.Option) *ec2.DeleteRouteTableOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteRouteTableOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DeleteRouteTableInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteRouteWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DeleteRouteWithContext(_a0 aws.Context, _a1 *ec2.DeleteRouteInput, _a2 ...request.Option) (*ec2.DeleteRouteOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DeleteRouteOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DeleteRouteInput, ...request.Option) *ec2.DeleteRouteOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteRouteOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DeleteRouteInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}



// DeleteSecurityGroupRequest provides a mock function with given fields: _a0
func (_m *EC2API) DeleteSecurityGroupRequest(_a0 *ec2.DeleteSecurityGroupInput) (*request.Request, *ec2.DeleteSecurityGroupOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteSecurityGroupInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteSecurityGroupOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteSecurityGroupInput) *ec2.DeleteSecurityGroupOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteSecurityGroupOutput)
		}
	}

	return r0, r1
}

// DeleteSecurityGroupWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DeleteSecurityGroupWithContext(_a0 aws.Context, _a1 *ec2.DeleteSecurityGroupInput, _a2 ...request.Option) (*ec2.DeleteSecurityGroupOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DeleteSecurityGroupOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DeleteSecurityGroupInput, ...request.Option) *ec2.DeleteSecurityGroupOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteSecurityGroupOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DeleteSecurityGroupInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteSnapshot provides a mock function with given fields: _a0
func (_m *EC2API) DeleteSnapshot(_a0 *ec2.DeleteSnapshotInput) (*ec2.DeleteSnapshotOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteSnapshotOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteSnapshotInput) *ec2.DeleteSnapshotOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteSnapshotOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteSnapshotInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteSnapshotRequest provides a mock function with given fields: _a0
func (_m *EC2API) DeleteSnapshotRequest(_a0 *ec2.DeleteSnapshotInput) (*request.Request, *ec2.DeleteSnapshotOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteSnapshotInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteSnapshotOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteSnapshotInput) *ec2.DeleteSnapshotOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteSnapshotOutput)
		}
	}

	return r0, r1
}

// DeleteSnapshotWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DeleteSnapshotWithContext(_a0 aws.Context, _a1 *ec2.DeleteSnapshotInput, _a2 ...request.Option) (*ec2.DeleteSnapshotOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DeleteSnapshotOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DeleteSnapshotInput, ...request.Option) *ec2.DeleteSnapshotOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteSnapshotOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DeleteSnapshotInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteSpotDatafeedSubscription provides a mock function with given fields: _a0
func (_m *EC2API) DeleteSpotDatafeedSubscription(_a0 *ec2.DeleteSpotDatafeedSubscriptionInput) (*ec2.DeleteSpotDatafeedSubscriptionOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteSpotDatafeedSubscriptionOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteSpotDatafeedSubscriptionInput) *ec2.DeleteSpotDatafeedSubscriptionOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteSpotDatafeedSubscriptionOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteSpotDatafeedSubscriptionInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteSpotDatafeedSubscriptionRequest provides a mock function with given fields: _a0
func (_m *EC2API) DeleteSpotDatafeedSubscriptionRequest(_a0 *ec2.DeleteSpotDatafeedSubscriptionInput) (*request.Request, *ec2.DeleteSpotDatafeedSubscriptionOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteSpotDatafeedSubscriptionInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteSpotDatafeedSubscriptionOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteSpotDatafeedSubscriptionInput) *ec2.DeleteSpotDatafeedSubscriptionOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteSpotDatafeedSubscriptionOutput)
		}
	}

	return r0, r1
}

// DeleteSpotDatafeedSubscriptionWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DeleteSpotDatafeedSubscriptionWithContext(_a0 aws.Context, _a1 *ec2.DeleteSpotDatafeedSubscriptionInput, _a2 ...request.Option) (*ec2.DeleteSpotDatafeedSubscriptionOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DeleteSpotDatafeedSubscriptionOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DeleteSpotDatafeedSubscriptionInput, ...request.Option) *ec2.DeleteSpotDatafeedSubscriptionOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteSpotDatafeedSubscriptionOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DeleteSpotDatafeedSubscriptionInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteSubnet provides a mock function with given fields: _a0
func (_m *EC2API) DeleteSubnet(_a0 *ec2.DeleteSubnetInput) (*ec2.DeleteSubnetOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteSubnetOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteSubnetInput) *ec2.DeleteSubnetOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteSubnetOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteSubnetInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteSubnetRequest provides a mock function with given fields: _a0
func (_m *EC2API) DeleteSubnetRequest(_a0 *ec2.DeleteSubnetInput) (*request.Request, *ec2.DeleteSubnetOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteSubnetInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteSubnetOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteSubnetInput) *ec2.DeleteSubnetOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteSubnetOutput)
		}
	}

	return r0, r1
}

// DeleteSubnetWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DeleteSubnetWithContext(_a0 aws.Context, _a1 *ec2.DeleteSubnetInput, _a2 ...request.Option) (*ec2.DeleteSubnetOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DeleteSubnetOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DeleteSubnetInput, ...request.Option) *ec2.DeleteSubnetOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteSubnetOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DeleteSubnetInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTags provides a mock function with given fields: _a0
func (_m *EC2API) DeleteTags(_a0 *ec2.DeleteTagsInput) (*ec2.DeleteTagsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteTagsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteTagsInput) *ec2.DeleteTagsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteTagsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteTagsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTagsRequest provides a mock function with given fields: _a0
func (_m *EC2API) DeleteTagsRequest(_a0 *ec2.DeleteTagsInput) (*request.Request, *ec2.DeleteTagsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteTagsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteTagsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteTagsInput) *ec2.DeleteTagsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteTagsOutput)
		}
	}

	return r0, r1
}

// DeleteTagsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DeleteTagsWithContext(_a0 aws.Context, _a1 *ec2.DeleteTagsInput, _a2 ...request.Option) (*ec2.DeleteTagsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DeleteTagsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DeleteTagsInput, ...request.Option) *ec2.DeleteTagsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteTagsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DeleteTagsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteVolume provides a mock function with given fields: _a0
func (_m *EC2API) DeleteVolume(_a0 *ec2.DeleteVolumeInput) (*ec2.DeleteVolumeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteVolumeOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteVolumeInput) *ec2.DeleteVolumeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteVolumeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteVolumeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteVolumeRequest provides a mock function with given fields: _a0
func (_m *EC2API) DeleteVolumeRequest(_a0 *ec2.DeleteVolumeInput) (*request.Request, *ec2.DeleteVolumeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteVolumeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteVolumeOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteVolumeInput) *ec2.DeleteVolumeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteVolumeOutput)
		}
	}

	return r0, r1
}

// DeleteVolumeWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DeleteVolumeWithContext(_a0 aws.Context, _a1 *ec2.DeleteVolumeInput, _a2 ...request.Option) (*ec2.DeleteVolumeOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DeleteVolumeOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DeleteVolumeInput, ...request.Option) *ec2.DeleteVolumeOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteVolumeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DeleteVolumeInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteVpc provides a mock function with given fields: _a0
func (_m *EC2API) DeleteVpc(_a0 *ec2.DeleteVpcInput) (*ec2.DeleteVpcOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteVpcOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteVpcInput) *ec2.DeleteVpcOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteVpcOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteVpcInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteVpcEndpointConnectionNotifications provides a mock function with given fields: _a0
func (_m *EC2API) DeleteVpcEndpointConnectionNotifications(_a0 *ec2.DeleteVpcEndpointConnectionNotificationsInput) (*ec2.DeleteVpcEndpointConnectionNotificationsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteVpcEndpointConnectionNotificationsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteVpcEndpointConnectionNotificationsInput) *ec2.DeleteVpcEndpointConnectionNotificationsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteVpcEndpointConnectionNotificationsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteVpcEndpointConnectionNotificationsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteVpcEndpointConnectionNotificationsRequest provides a mock function with given fields: _a0
func (_m *EC2API) DeleteVpcEndpointConnectionNotificationsRequest(_a0 *ec2.DeleteVpcEndpointConnectionNotificationsInput) (*request.Request, *ec2.DeleteVpcEndpointConnectionNotificationsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteVpcEndpointConnectionNotificationsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteVpcEndpointConnectionNotificationsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteVpcEndpointConnectionNotificationsInput) *ec2.DeleteVpcEndpointConnectionNotificationsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteVpcEndpointConnectionNotificationsOutput)
		}
	}

	return r0, r1
}

// DeleteVpcEndpointConnectionNotificationsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DeleteVpcEndpointConnectionNotificationsWithContext(_a0 aws.Context, _a1 *ec2.DeleteVpcEndpointConnectionNotificationsInput, _a2 ...request.Option) (*ec2.DeleteVpcEndpointConnectionNotificationsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DeleteVpcEndpointConnectionNotificationsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DeleteVpcEndpointConnectionNotificationsInput, ...request.Option) *ec2.DeleteVpcEndpointConnectionNotificationsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteVpcEndpointConnectionNotificationsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DeleteVpcEndpointConnectionNotificationsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteVpcEndpointServiceConfigurations provides a mock function with given fields: _a0
func (_m *EC2API) DeleteVpcEndpointServiceConfigurations(_a0 *ec2.DeleteVpcEndpointServiceConfigurationsInput) (*ec2.DeleteVpcEndpointServiceConfigurationsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteVpcEndpointServiceConfigurationsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteVpcEndpointServiceConfigurationsInput) *ec2.DeleteVpcEndpointServiceConfigurationsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteVpcEndpointServiceConfigurationsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteVpcEndpointServiceConfigurationsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteVpcEndpointServiceConfigurationsRequest provides a mock function with given fields: _a0
func (_m *EC2API) DeleteVpcEndpointServiceConfigurationsRequest(_a0 *ec2.DeleteVpcEndpointServiceConfigurationsInput) (*request.Request, *ec2.DeleteVpcEndpointServiceConfigurationsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteVpcEndpointServiceConfigurationsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteVpcEndpointServiceConfigurationsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteVpcEndpointServiceConfigurationsInput) *ec2.DeleteVpcEndpointServiceConfigurationsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteVpcEndpointServiceConfigurationsOutput)
		}
	}

	return r0, r1
}

// DeleteVpcEndpointServiceConfigurationsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DeleteVpcEndpointServiceConfigurationsWithContext(_a0 aws.Context, _a1 *ec2.DeleteVpcEndpointServiceConfigurationsInput, _a2 ...request.Option) (*ec2.DeleteVpcEndpointServiceConfigurationsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DeleteVpcEndpointServiceConfigurationsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DeleteVpcEndpointServiceConfigurationsInput, ...request.Option) *ec2.DeleteVpcEndpointServiceConfigurationsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteVpcEndpointServiceConfigurationsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DeleteVpcEndpointServiceConfigurationsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteVpcEndpoints provides a mock function with given fields: _a0
func (_m *EC2API) DeleteVpcEndpoints(_a0 *ec2.DeleteVpcEndpointsInput) (*ec2.DeleteVpcEndpointsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteVpcEndpointsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteVpcEndpointsInput) *ec2.DeleteVpcEndpointsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteVpcEndpointsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteVpcEndpointsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteVpcEndpointsRequest provides a mock function with given fields: _a0
func (_m *EC2API) DeleteVpcEndpointsRequest(_a0 *ec2.DeleteVpcEndpointsInput) (*request.Request, *ec2.DeleteVpcEndpointsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteVpcEndpointsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteVpcEndpointsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteVpcEndpointsInput) *ec2.DeleteVpcEndpointsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteVpcEndpointsOutput)
		}
	}

	return r0, r1
}

// DeleteVpcEndpointsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DeleteVpcEndpointsWithContext(_a0 aws.Context, _a1 *ec2.DeleteVpcEndpointsInput, _a2 ...request.Option) (*ec2.DeleteVpcEndpointsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DeleteVpcEndpointsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DeleteVpcEndpointsInput, ...request.Option) *ec2.DeleteVpcEndpointsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteVpcEndpointsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DeleteVpcEndpointsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteVpcPeeringConnection provides a mock function with given fields: _a0
func (_m *EC2API) DeleteVpcPeeringConnection(_a0 *ec2.DeleteVpcPeeringConnectionInput) (*ec2.DeleteVpcPeeringConnectionOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteVpcPeeringConnectionOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteVpcPeeringConnectionInput) *ec2.DeleteVpcPeeringConnectionOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteVpcPeeringConnectionOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteVpcPeeringConnectionInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteVpcPeeringConnectionRequest provides a mock function with given fields: _a0
func (_m *EC2API) DeleteVpcPeeringConnectionRequest(_a0 *ec2.DeleteVpcPeeringConnectionInput) (*request.Request, *ec2.DeleteVpcPeeringConnectionOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteVpcPeeringConnectionInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteVpcPeeringConnectionOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteVpcPeeringConnectionInput) *ec2.DeleteVpcPeeringConnectionOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteVpcPeeringConnectionOutput)
		}
	}

	return r0, r1
}

// DeleteVpcPeeringConnectionWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DeleteVpcPeeringConnectionWithContext(_a0 aws.Context, _a1 *ec2.DeleteVpcPeeringConnectionInput, _a2 ...request.Option) (*ec2.DeleteVpcPeeringConnectionOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DeleteVpcPeeringConnectionOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DeleteVpcPeeringConnectionInput, ...request.Option) *ec2.DeleteVpcPeeringConnectionOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteVpcPeeringConnectionOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DeleteVpcPeeringConnectionInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteVpcRequest provides a mock function with given fields: _a0
func (_m *EC2API) DeleteVpcRequest(_a0 *ec2.DeleteVpcInput) (*request.Request, *ec2.DeleteVpcOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteVpcInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteVpcOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteVpcInput) *ec2.DeleteVpcOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteVpcOutput)
		}
	}

	return r0, r1
}

// DeleteVpcWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DeleteVpcWithContext(_a0 aws.Context, _a1 *ec2.DeleteVpcInput, _a2 ...request.Option) (*ec2.DeleteVpcOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DeleteVpcOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DeleteVpcInput, ...request.Option) *ec2.DeleteVpcOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteVpcOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DeleteVpcInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteVpnConnection provides a mock function with given fields: _a0
func (_m *EC2API) DeleteVpnConnection(_a0 *ec2.DeleteVpnConnectionInput) (*ec2.DeleteVpnConnectionOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteVpnConnectionOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteVpnConnectionInput) *ec2.DeleteVpnConnectionOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteVpnConnectionOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteVpnConnectionInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteVpnConnectionRequest provides a mock function with given fields: _a0
func (_m *EC2API) DeleteVpnConnectionRequest(_a0 *ec2.DeleteVpnConnectionInput) (*request.Request, *ec2.DeleteVpnConnectionOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteVpnConnectionInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteVpnConnectionOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteVpnConnectionInput) *ec2.DeleteVpnConnectionOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteVpnConnectionOutput)
		}
	}

	return r0, r1
}

// DeleteVpnConnectionRoute provides a mock function with given fields: _a0
func (_m *EC2API) DeleteVpnConnectionRoute(_a0 *ec2.DeleteVpnConnectionRouteInput) (*ec2.DeleteVpnConnectionRouteOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteVpnConnectionRouteOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteVpnConnectionRouteInput) *ec2.DeleteVpnConnectionRouteOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteVpnConnectionRouteOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteVpnConnectionRouteInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteVpnConnectionRouteRequest provides a mock function with given fields: _a0
func (_m *EC2API) DeleteVpnConnectionRouteRequest(_a0 *ec2.DeleteVpnConnectionRouteInput) (*request.Request, *ec2.DeleteVpnConnectionRouteOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteVpnConnectionRouteInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteVpnConnectionRouteOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteVpnConnectionRouteInput) *ec2.DeleteVpnConnectionRouteOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteVpnConnectionRouteOutput)
		}
	}

	return r0, r1
}

// DeleteVpnConnectionRouteWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DeleteVpnConnectionRouteWithContext(_a0 aws.Context, _a1 *ec2.DeleteVpnConnectionRouteInput, _a2 ...request.Option) (*ec2.DeleteVpnConnectionRouteOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DeleteVpnConnectionRouteOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DeleteVpnConnectionRouteInput, ...request.Option) *ec2.DeleteVpnConnectionRouteOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteVpnConnectionRouteOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DeleteVpnConnectionRouteInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteVpnConnectionWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DeleteVpnConnectionWithContext(_a0 aws.Context, _a1 *ec2.DeleteVpnConnectionInput, _a2 ...request.Option) (*ec2.DeleteVpnConnectionOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DeleteVpnConnectionOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DeleteVpnConnectionInput, ...request.Option) *ec2.DeleteVpnConnectionOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteVpnConnectionOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DeleteVpnConnectionInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteVpnGateway provides a mock function with given fields: _a0
func (_m *EC2API) DeleteVpnGateway(_a0 *ec2.DeleteVpnGatewayInput) (*ec2.DeleteVpnGatewayOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeleteVpnGatewayOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeleteVpnGatewayInput) *ec2.DeleteVpnGatewayOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteVpnGatewayOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeleteVpnGatewayInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteVpnGatewayRequest provides a mock function with given fields: _a0
func (_m *EC2API) DeleteVpnGatewayRequest(_a0 *ec2.DeleteVpnGatewayInput) (*request.Request, *ec2.DeleteVpnGatewayOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeleteVpnGatewayInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeleteVpnGatewayOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeleteVpnGatewayInput) *ec2.DeleteVpnGatewayOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeleteVpnGatewayOutput)
		}
	}

	return r0, r1
}

// DeleteVpnGatewayWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DeleteVpnGatewayWithContext(_a0 aws.Context, _a1 *ec2.DeleteVpnGatewayInput, _a2 ...request.Option) (*ec2.DeleteVpnGatewayOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DeleteVpnGatewayOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DeleteVpnGatewayInput, ...request.Option) *ec2.DeleteVpnGatewayOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeleteVpnGatewayOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DeleteVpnGatewayInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeregisterImage provides a mock function with given fields: _a0
func (_m *EC2API) DeregisterImage(_a0 *ec2.DeregisterImageInput) (*ec2.DeregisterImageOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DeregisterImageOutput
	if rf, ok := ret.Get(0).(func(*ec2.DeregisterImageInput) *ec2.DeregisterImageOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeregisterImageOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DeregisterImageInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeregisterImageRequest provides a mock function with given fields: _a0
func (_m *EC2API) DeregisterImageRequest(_a0 *ec2.DeregisterImageInput) (*request.Request, *ec2.DeregisterImageOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DeregisterImageInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DeregisterImageOutput
	if rf, ok := ret.Get(1).(func(*ec2.DeregisterImageInput) *ec2.DeregisterImageOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DeregisterImageOutput)
		}
	}

	return r0, r1
}

// DeregisterImageWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DeregisterImageWithContext(_a0 aws.Context, _a1 *ec2.DeregisterImageInput, _a2 ...request.Option) (*ec2.DeregisterImageOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DeregisterImageOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DeregisterImageInput, ...request.Option) *ec2.DeregisterImageOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DeregisterImageOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DeregisterImageInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeAccountAttributes provides a mock function with given fields: _a0
func (_m *EC2API) DescribeAccountAttributes(_a0 *ec2.DescribeAccountAttributesInput) (*ec2.DescribeAccountAttributesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeAccountAttributesOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeAccountAttributesInput) *ec2.DescribeAccountAttributesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeAccountAttributesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeAccountAttributesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeAccountAttributesRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeAccountAttributesRequest(_a0 *ec2.DescribeAccountAttributesInput) (*request.Request, *ec2.DescribeAccountAttributesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeAccountAttributesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeAccountAttributesOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeAccountAttributesInput) *ec2.DescribeAccountAttributesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeAccountAttributesOutput)
		}
	}

	return r0, r1
}

// DescribeAccountAttributesWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeAccountAttributesWithContext(_a0 aws.Context, _a1 *ec2.DescribeAccountAttributesInput, _a2 ...request.Option) (*ec2.DescribeAccountAttributesOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeAccountAttributesOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeAccountAttributesInput, ...request.Option) *ec2.DescribeAccountAttributesOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeAccountAttributesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeAccountAttributesInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}



// DescribeAddressesRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeAddressesRequest(_a0 *ec2.DescribeAddressesInput) (*request.Request, *ec2.DescribeAddressesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeAddressesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeAddressesOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeAddressesInput) *ec2.DescribeAddressesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeAddressesOutput)
		}
	}

	return r0, r1
}

// DescribeAddressesWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeAddressesWithContext(_a0 aws.Context, _a1 *ec2.DescribeAddressesInput, _a2 ...request.Option) (*ec2.DescribeAddressesOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeAddressesOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeAddressesInput, ...request.Option) *ec2.DescribeAddressesOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeAddressesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeAddressesInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeAggregateIdFormat provides a mock function with given fields: _a0
func (_m *EC2API) DescribeAggregateIdFormat(_a0 *ec2.DescribeAggregateIdFormatInput) (*ec2.DescribeAggregateIdFormatOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeAggregateIdFormatOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeAggregateIdFormatInput) *ec2.DescribeAggregateIdFormatOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeAggregateIdFormatOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeAggregateIdFormatInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeAggregateIdFormatRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeAggregateIdFormatRequest(_a0 *ec2.DescribeAggregateIdFormatInput) (*request.Request, *ec2.DescribeAggregateIdFormatOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeAggregateIdFormatInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeAggregateIdFormatOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeAggregateIdFormatInput) *ec2.DescribeAggregateIdFormatOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeAggregateIdFormatOutput)
		}
	}

	return r0, r1
}

// DescribeAggregateIdFormatWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeAggregateIdFormatWithContext(_a0 aws.Context, _a1 *ec2.DescribeAggregateIdFormatInput, _a2 ...request.Option) (*ec2.DescribeAggregateIdFormatOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeAggregateIdFormatOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeAggregateIdFormatInput, ...request.Option) *ec2.DescribeAggregateIdFormatOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeAggregateIdFormatOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeAggregateIdFormatInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeAvailabilityZones provides a mock function with given fields: _a0
func (_m *EC2API) DescribeAvailabilityZones(_a0 *ec2.DescribeAvailabilityZonesInput) (*ec2.DescribeAvailabilityZonesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeAvailabilityZonesOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeAvailabilityZonesInput) *ec2.DescribeAvailabilityZonesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeAvailabilityZonesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeAvailabilityZonesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeAvailabilityZonesRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeAvailabilityZonesRequest(_a0 *ec2.DescribeAvailabilityZonesInput) (*request.Request, *ec2.DescribeAvailabilityZonesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeAvailabilityZonesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeAvailabilityZonesOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeAvailabilityZonesInput) *ec2.DescribeAvailabilityZonesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeAvailabilityZonesOutput)
		}
	}

	return r0, r1
}

// DescribeAvailabilityZonesWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeAvailabilityZonesWithContext(_a0 aws.Context, _a1 *ec2.DescribeAvailabilityZonesInput, _a2 ...request.Option) (*ec2.DescribeAvailabilityZonesOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeAvailabilityZonesOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeAvailabilityZonesInput, ...request.Option) *ec2.DescribeAvailabilityZonesOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeAvailabilityZonesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeAvailabilityZonesInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeBundleTasks provides a mock function with given fields: _a0
func (_m *EC2API) DescribeBundleTasks(_a0 *ec2.DescribeBundleTasksInput) (*ec2.DescribeBundleTasksOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeBundleTasksOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeBundleTasksInput) *ec2.DescribeBundleTasksOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeBundleTasksOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeBundleTasksInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeBundleTasksRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeBundleTasksRequest(_a0 *ec2.DescribeBundleTasksInput) (*request.Request, *ec2.DescribeBundleTasksOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeBundleTasksInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeBundleTasksOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeBundleTasksInput) *ec2.DescribeBundleTasksOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeBundleTasksOutput)
		}
	}

	return r0, r1
}

// DescribeBundleTasksWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeBundleTasksWithContext(_a0 aws.Context, _a1 *ec2.DescribeBundleTasksInput, _a2 ...request.Option) (*ec2.DescribeBundleTasksOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeBundleTasksOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeBundleTasksInput, ...request.Option) *ec2.DescribeBundleTasksOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeBundleTasksOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeBundleTasksInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeClassicLinkInstances provides a mock function with given fields: _a0
func (_m *EC2API) DescribeClassicLinkInstances(_a0 *ec2.DescribeClassicLinkInstancesInput) (*ec2.DescribeClassicLinkInstancesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeClassicLinkInstancesOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeClassicLinkInstancesInput) *ec2.DescribeClassicLinkInstancesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeClassicLinkInstancesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeClassicLinkInstancesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeClassicLinkInstancesRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeClassicLinkInstancesRequest(_a0 *ec2.DescribeClassicLinkInstancesInput) (*request.Request, *ec2.DescribeClassicLinkInstancesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeClassicLinkInstancesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeClassicLinkInstancesOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeClassicLinkInstancesInput) *ec2.DescribeClassicLinkInstancesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeClassicLinkInstancesOutput)
		}
	}

	return r0, r1
}

// DescribeClassicLinkInstancesWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeClassicLinkInstancesWithContext(_a0 aws.Context, _a1 *ec2.DescribeClassicLinkInstancesInput, _a2 ...request.Option) (*ec2.DescribeClassicLinkInstancesOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeClassicLinkInstancesOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeClassicLinkInstancesInput, ...request.Option) *ec2.DescribeClassicLinkInstancesOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeClassicLinkInstancesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeClassicLinkInstancesInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeConversionTasks provides a mock function with given fields: _a0
func (_m *EC2API) DescribeConversionTasks(_a0 *ec2.DescribeConversionTasksInput) (*ec2.DescribeConversionTasksOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeConversionTasksOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeConversionTasksInput) *ec2.DescribeConversionTasksOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeConversionTasksOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeConversionTasksInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeConversionTasksRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeConversionTasksRequest(_a0 *ec2.DescribeConversionTasksInput) (*request.Request, *ec2.DescribeConversionTasksOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeConversionTasksInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeConversionTasksOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeConversionTasksInput) *ec2.DescribeConversionTasksOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeConversionTasksOutput)
		}
	}

	return r0, r1
}

// DescribeConversionTasksWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeConversionTasksWithContext(_a0 aws.Context, _a1 *ec2.DescribeConversionTasksInput, _a2 ...request.Option) (*ec2.DescribeConversionTasksOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeConversionTasksOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeConversionTasksInput, ...request.Option) *ec2.DescribeConversionTasksOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeConversionTasksOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeConversionTasksInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeCustomerGateways provides a mock function with given fields: _a0
func (_m *EC2API) DescribeCustomerGateways(_a0 *ec2.DescribeCustomerGatewaysInput) (*ec2.DescribeCustomerGatewaysOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeCustomerGatewaysOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeCustomerGatewaysInput) *ec2.DescribeCustomerGatewaysOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeCustomerGatewaysOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeCustomerGatewaysInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeCustomerGatewaysRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeCustomerGatewaysRequest(_a0 *ec2.DescribeCustomerGatewaysInput) (*request.Request, *ec2.DescribeCustomerGatewaysOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeCustomerGatewaysInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeCustomerGatewaysOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeCustomerGatewaysInput) *ec2.DescribeCustomerGatewaysOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeCustomerGatewaysOutput)
		}
	}

	return r0, r1
}

// DescribeCustomerGatewaysWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeCustomerGatewaysWithContext(_a0 aws.Context, _a1 *ec2.DescribeCustomerGatewaysInput, _a2 ...request.Option) (*ec2.DescribeCustomerGatewaysOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeCustomerGatewaysOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeCustomerGatewaysInput, ...request.Option) *ec2.DescribeCustomerGatewaysOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeCustomerGatewaysOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeCustomerGatewaysInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeDhcpOptions provides a mock function with given fields: _a0
func (_m *EC2API) DescribeDhcpOptions(_a0 *ec2.DescribeDhcpOptionsInput) (*ec2.DescribeDhcpOptionsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeDhcpOptionsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeDhcpOptionsInput) *ec2.DescribeDhcpOptionsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeDhcpOptionsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeDhcpOptionsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeDhcpOptionsRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeDhcpOptionsRequest(_a0 *ec2.DescribeDhcpOptionsInput) (*request.Request, *ec2.DescribeDhcpOptionsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeDhcpOptionsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeDhcpOptionsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeDhcpOptionsInput) *ec2.DescribeDhcpOptionsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeDhcpOptionsOutput)
		}
	}

	return r0, r1
}

// DescribeDhcpOptionsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeDhcpOptionsWithContext(_a0 aws.Context, _a1 *ec2.DescribeDhcpOptionsInput, _a2 ...request.Option) (*ec2.DescribeDhcpOptionsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeDhcpOptionsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeDhcpOptionsInput, ...request.Option) *ec2.DescribeDhcpOptionsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeDhcpOptionsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeDhcpOptionsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeEgressOnlyInternetGateways provides a mock function with given fields: _a0
func (_m *EC2API) DescribeEgressOnlyInternetGateways(_a0 *ec2.DescribeEgressOnlyInternetGatewaysInput) (*ec2.DescribeEgressOnlyInternetGatewaysOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeEgressOnlyInternetGatewaysOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeEgressOnlyInternetGatewaysInput) *ec2.DescribeEgressOnlyInternetGatewaysOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeEgressOnlyInternetGatewaysOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeEgressOnlyInternetGatewaysInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeEgressOnlyInternetGatewaysRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeEgressOnlyInternetGatewaysRequest(_a0 *ec2.DescribeEgressOnlyInternetGatewaysInput) (*request.Request, *ec2.DescribeEgressOnlyInternetGatewaysOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeEgressOnlyInternetGatewaysInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeEgressOnlyInternetGatewaysOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeEgressOnlyInternetGatewaysInput) *ec2.DescribeEgressOnlyInternetGatewaysOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeEgressOnlyInternetGatewaysOutput)
		}
	}

	return r0, r1
}

// DescribeEgressOnlyInternetGatewaysWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeEgressOnlyInternetGatewaysWithContext(_a0 aws.Context, _a1 *ec2.DescribeEgressOnlyInternetGatewaysInput, _a2 ...request.Option) (*ec2.DescribeEgressOnlyInternetGatewaysOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeEgressOnlyInternetGatewaysOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeEgressOnlyInternetGatewaysInput, ...request.Option) *ec2.DescribeEgressOnlyInternetGatewaysOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeEgressOnlyInternetGatewaysOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeEgressOnlyInternetGatewaysInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeElasticGpus provides a mock function with given fields: _a0
func (_m *EC2API) DescribeElasticGpus(_a0 *ec2.DescribeElasticGpusInput) (*ec2.DescribeElasticGpusOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeElasticGpusOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeElasticGpusInput) *ec2.DescribeElasticGpusOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeElasticGpusOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeElasticGpusInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeElasticGpusRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeElasticGpusRequest(_a0 *ec2.DescribeElasticGpusInput) (*request.Request, *ec2.DescribeElasticGpusOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeElasticGpusInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeElasticGpusOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeElasticGpusInput) *ec2.DescribeElasticGpusOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeElasticGpusOutput)
		}
	}

	return r0, r1
}

// DescribeElasticGpusWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeElasticGpusWithContext(_a0 aws.Context, _a1 *ec2.DescribeElasticGpusInput, _a2 ...request.Option) (*ec2.DescribeElasticGpusOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeElasticGpusOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeElasticGpusInput, ...request.Option) *ec2.DescribeElasticGpusOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeElasticGpusOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeElasticGpusInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeExportTasks provides a mock function with given fields: _a0
func (_m *EC2API) DescribeExportTasks(_a0 *ec2.DescribeExportTasksInput) (*ec2.DescribeExportTasksOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeExportTasksOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeExportTasksInput) *ec2.DescribeExportTasksOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeExportTasksOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeExportTasksInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeExportTasksRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeExportTasksRequest(_a0 *ec2.DescribeExportTasksInput) (*request.Request, *ec2.DescribeExportTasksOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeExportTasksInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeExportTasksOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeExportTasksInput) *ec2.DescribeExportTasksOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeExportTasksOutput)
		}
	}

	return r0, r1
}

// DescribeExportTasksWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeExportTasksWithContext(_a0 aws.Context, _a1 *ec2.DescribeExportTasksInput, _a2 ...request.Option) (*ec2.DescribeExportTasksOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeExportTasksOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeExportTasksInput, ...request.Option) *ec2.DescribeExportTasksOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeExportTasksOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeExportTasksInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeFleetHistory provides a mock function with given fields: _a0
func (_m *EC2API) DescribeFleetHistory(_a0 *ec2.DescribeFleetHistoryInput) (*ec2.DescribeFleetHistoryOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeFleetHistoryOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeFleetHistoryInput) *ec2.DescribeFleetHistoryOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeFleetHistoryOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeFleetHistoryInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeFleetHistoryRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeFleetHistoryRequest(_a0 *ec2.DescribeFleetHistoryInput) (*request.Request, *ec2.DescribeFleetHistoryOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeFleetHistoryInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeFleetHistoryOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeFleetHistoryInput) *ec2.DescribeFleetHistoryOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeFleetHistoryOutput)
		}
	}

	return r0, r1
}

// DescribeFleetHistoryWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeFleetHistoryWithContext(_a0 aws.Context, _a1 *ec2.DescribeFleetHistoryInput, _a2 ...request.Option) (*ec2.DescribeFleetHistoryOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeFleetHistoryOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeFleetHistoryInput, ...request.Option) *ec2.DescribeFleetHistoryOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeFleetHistoryOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeFleetHistoryInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeFleetInstances provides a mock function with given fields: _a0
func (_m *EC2API) DescribeFleetInstances(_a0 *ec2.DescribeFleetInstancesInput) (*ec2.DescribeFleetInstancesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeFleetInstancesOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeFleetInstancesInput) *ec2.DescribeFleetInstancesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeFleetInstancesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeFleetInstancesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeFleetInstancesRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeFleetInstancesRequest(_a0 *ec2.DescribeFleetInstancesInput) (*request.Request, *ec2.DescribeFleetInstancesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeFleetInstancesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeFleetInstancesOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeFleetInstancesInput) *ec2.DescribeFleetInstancesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeFleetInstancesOutput)
		}
	}

	return r0, r1
}

// DescribeFleetInstancesWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeFleetInstancesWithContext(_a0 aws.Context, _a1 *ec2.DescribeFleetInstancesInput, _a2 ...request.Option) (*ec2.DescribeFleetInstancesOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeFleetInstancesOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeFleetInstancesInput, ...request.Option) *ec2.DescribeFleetInstancesOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeFleetInstancesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeFleetInstancesInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeFleets provides a mock function with given fields: _a0
func (_m *EC2API) DescribeFleets(_a0 *ec2.DescribeFleetsInput) (*ec2.DescribeFleetsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeFleetsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeFleetsInput) *ec2.DescribeFleetsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeFleetsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeFleetsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeFleetsRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeFleetsRequest(_a0 *ec2.DescribeFleetsInput) (*request.Request, *ec2.DescribeFleetsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeFleetsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeFleetsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeFleetsInput) *ec2.DescribeFleetsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeFleetsOutput)
		}
	}

	return r0, r1
}

// DescribeFleetsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeFleetsWithContext(_a0 aws.Context, _a1 *ec2.DescribeFleetsInput, _a2 ...request.Option) (*ec2.DescribeFleetsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeFleetsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeFleetsInput, ...request.Option) *ec2.DescribeFleetsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeFleetsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeFleetsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeFlowLogs provides a mock function with given fields: _a0
func (_m *EC2API) DescribeFlowLogs(_a0 *ec2.DescribeFlowLogsInput) (*ec2.DescribeFlowLogsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeFlowLogsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeFlowLogsInput) *ec2.DescribeFlowLogsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeFlowLogsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeFlowLogsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeFlowLogsRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeFlowLogsRequest(_a0 *ec2.DescribeFlowLogsInput) (*request.Request, *ec2.DescribeFlowLogsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeFlowLogsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeFlowLogsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeFlowLogsInput) *ec2.DescribeFlowLogsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeFlowLogsOutput)
		}
	}

	return r0, r1
}

// DescribeFlowLogsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeFlowLogsWithContext(_a0 aws.Context, _a1 *ec2.DescribeFlowLogsInput, _a2 ...request.Option) (*ec2.DescribeFlowLogsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeFlowLogsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeFlowLogsInput, ...request.Option) *ec2.DescribeFlowLogsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeFlowLogsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeFlowLogsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeFpgaImageAttribute provides a mock function with given fields: _a0
func (_m *EC2API) DescribeFpgaImageAttribute(_a0 *ec2.DescribeFpgaImageAttributeInput) (*ec2.DescribeFpgaImageAttributeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeFpgaImageAttributeOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeFpgaImageAttributeInput) *ec2.DescribeFpgaImageAttributeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeFpgaImageAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeFpgaImageAttributeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeFpgaImageAttributeRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeFpgaImageAttributeRequest(_a0 *ec2.DescribeFpgaImageAttributeInput) (*request.Request, *ec2.DescribeFpgaImageAttributeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeFpgaImageAttributeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeFpgaImageAttributeOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeFpgaImageAttributeInput) *ec2.DescribeFpgaImageAttributeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeFpgaImageAttributeOutput)
		}
	}

	return r0, r1
}

// DescribeFpgaImageAttributeWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeFpgaImageAttributeWithContext(_a0 aws.Context, _a1 *ec2.DescribeFpgaImageAttributeInput, _a2 ...request.Option) (*ec2.DescribeFpgaImageAttributeOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeFpgaImageAttributeOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeFpgaImageAttributeInput, ...request.Option) *ec2.DescribeFpgaImageAttributeOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeFpgaImageAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeFpgaImageAttributeInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeFpgaImages provides a mock function with given fields: _a0
func (_m *EC2API) DescribeFpgaImages(_a0 *ec2.DescribeFpgaImagesInput) (*ec2.DescribeFpgaImagesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeFpgaImagesOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeFpgaImagesInput) *ec2.DescribeFpgaImagesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeFpgaImagesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeFpgaImagesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeFpgaImagesRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeFpgaImagesRequest(_a0 *ec2.DescribeFpgaImagesInput) (*request.Request, *ec2.DescribeFpgaImagesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeFpgaImagesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeFpgaImagesOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeFpgaImagesInput) *ec2.DescribeFpgaImagesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeFpgaImagesOutput)
		}
	}

	return r0, r1
}

// DescribeFpgaImagesWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeFpgaImagesWithContext(_a0 aws.Context, _a1 *ec2.DescribeFpgaImagesInput, _a2 ...request.Option) (*ec2.DescribeFpgaImagesOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeFpgaImagesOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeFpgaImagesInput, ...request.Option) *ec2.DescribeFpgaImagesOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeFpgaImagesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeFpgaImagesInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeHostReservationOfferings provides a mock function with given fields: _a0
func (_m *EC2API) DescribeHostReservationOfferings(_a0 *ec2.DescribeHostReservationOfferingsInput) (*ec2.DescribeHostReservationOfferingsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeHostReservationOfferingsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeHostReservationOfferingsInput) *ec2.DescribeHostReservationOfferingsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeHostReservationOfferingsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeHostReservationOfferingsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeHostReservationOfferingsRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeHostReservationOfferingsRequest(_a0 *ec2.DescribeHostReservationOfferingsInput) (*request.Request, *ec2.DescribeHostReservationOfferingsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeHostReservationOfferingsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeHostReservationOfferingsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeHostReservationOfferingsInput) *ec2.DescribeHostReservationOfferingsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeHostReservationOfferingsOutput)
		}
	}

	return r0, r1
}

// DescribeHostReservationOfferingsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeHostReservationOfferingsWithContext(_a0 aws.Context, _a1 *ec2.DescribeHostReservationOfferingsInput, _a2 ...request.Option) (*ec2.DescribeHostReservationOfferingsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeHostReservationOfferingsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeHostReservationOfferingsInput, ...request.Option) *ec2.DescribeHostReservationOfferingsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeHostReservationOfferingsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeHostReservationOfferingsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeHostReservations provides a mock function with given fields: _a0
func (_m *EC2API) DescribeHostReservations(_a0 *ec2.DescribeHostReservationsInput) (*ec2.DescribeHostReservationsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeHostReservationsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeHostReservationsInput) *ec2.DescribeHostReservationsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeHostReservationsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeHostReservationsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeHostReservationsRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeHostReservationsRequest(_a0 *ec2.DescribeHostReservationsInput) (*request.Request, *ec2.DescribeHostReservationsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeHostReservationsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeHostReservationsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeHostReservationsInput) *ec2.DescribeHostReservationsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeHostReservationsOutput)
		}
	}

	return r0, r1
}

// DescribeHostReservationsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeHostReservationsWithContext(_a0 aws.Context, _a1 *ec2.DescribeHostReservationsInput, _a2 ...request.Option) (*ec2.DescribeHostReservationsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeHostReservationsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeHostReservationsInput, ...request.Option) *ec2.DescribeHostReservationsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeHostReservationsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeHostReservationsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeHosts provides a mock function with given fields: _a0
func (_m *EC2API) DescribeHosts(_a0 *ec2.DescribeHostsInput) (*ec2.DescribeHostsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeHostsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeHostsInput) *ec2.DescribeHostsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeHostsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeHostsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeHostsRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeHostsRequest(_a0 *ec2.DescribeHostsInput) (*request.Request, *ec2.DescribeHostsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeHostsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeHostsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeHostsInput) *ec2.DescribeHostsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeHostsOutput)
		}
	}

	return r0, r1
}

// DescribeHostsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeHostsWithContext(_a0 aws.Context, _a1 *ec2.DescribeHostsInput, _a2 ...request.Option) (*ec2.DescribeHostsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeHostsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeHostsInput, ...request.Option) *ec2.DescribeHostsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeHostsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeHostsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeIamInstanceProfileAssociations provides a mock function with given fields: _a0
func (_m *EC2API) DescribeIamInstanceProfileAssociations(_a0 *ec2.DescribeIamInstanceProfileAssociationsInput) (*ec2.DescribeIamInstanceProfileAssociationsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeIamInstanceProfileAssociationsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeIamInstanceProfileAssociationsInput) *ec2.DescribeIamInstanceProfileAssociationsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeIamInstanceProfileAssociationsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeIamInstanceProfileAssociationsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeIamInstanceProfileAssociationsRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeIamInstanceProfileAssociationsRequest(_a0 *ec2.DescribeIamInstanceProfileAssociationsInput) (*request.Request, *ec2.DescribeIamInstanceProfileAssociationsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeIamInstanceProfileAssociationsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeIamInstanceProfileAssociationsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeIamInstanceProfileAssociationsInput) *ec2.DescribeIamInstanceProfileAssociationsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeIamInstanceProfileAssociationsOutput)
		}
	}

	return r0, r1
}

// DescribeIamInstanceProfileAssociationsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeIamInstanceProfileAssociationsWithContext(_a0 aws.Context, _a1 *ec2.DescribeIamInstanceProfileAssociationsInput, _a2 ...request.Option) (*ec2.DescribeIamInstanceProfileAssociationsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeIamInstanceProfileAssociationsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeIamInstanceProfileAssociationsInput, ...request.Option) *ec2.DescribeIamInstanceProfileAssociationsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeIamInstanceProfileAssociationsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeIamInstanceProfileAssociationsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeIdFormat provides a mock function with given fields: _a0
func (_m *EC2API) DescribeIdFormat(_a0 *ec2.DescribeIdFormatInput) (*ec2.DescribeIdFormatOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeIdFormatOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeIdFormatInput) *ec2.DescribeIdFormatOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeIdFormatOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeIdFormatInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeIdFormatRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeIdFormatRequest(_a0 *ec2.DescribeIdFormatInput) (*request.Request, *ec2.DescribeIdFormatOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeIdFormatInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeIdFormatOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeIdFormatInput) *ec2.DescribeIdFormatOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeIdFormatOutput)
		}
	}

	return r0, r1
}

// DescribeIdFormatWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeIdFormatWithContext(_a0 aws.Context, _a1 *ec2.DescribeIdFormatInput, _a2 ...request.Option) (*ec2.DescribeIdFormatOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeIdFormatOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeIdFormatInput, ...request.Option) *ec2.DescribeIdFormatOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeIdFormatOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeIdFormatInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeIdentityIdFormat provides a mock function with given fields: _a0
func (_m *EC2API) DescribeIdentityIdFormat(_a0 *ec2.DescribeIdentityIdFormatInput) (*ec2.DescribeIdentityIdFormatOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeIdentityIdFormatOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeIdentityIdFormatInput) *ec2.DescribeIdentityIdFormatOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeIdentityIdFormatOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeIdentityIdFormatInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeIdentityIdFormatRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeIdentityIdFormatRequest(_a0 *ec2.DescribeIdentityIdFormatInput) (*request.Request, *ec2.DescribeIdentityIdFormatOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeIdentityIdFormatInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeIdentityIdFormatOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeIdentityIdFormatInput) *ec2.DescribeIdentityIdFormatOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeIdentityIdFormatOutput)
		}
	}

	return r0, r1
}

// DescribeIdentityIdFormatWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeIdentityIdFormatWithContext(_a0 aws.Context, _a1 *ec2.DescribeIdentityIdFormatInput, _a2 ...request.Option) (*ec2.DescribeIdentityIdFormatOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeIdentityIdFormatOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeIdentityIdFormatInput, ...request.Option) *ec2.DescribeIdentityIdFormatOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeIdentityIdFormatOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeIdentityIdFormatInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeImageAttribute provides a mock function with given fields: _a0
func (_m *EC2API) DescribeImageAttribute(_a0 *ec2.DescribeImageAttributeInput) (*ec2.DescribeImageAttributeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeImageAttributeOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeImageAttributeInput) *ec2.DescribeImageAttributeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeImageAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeImageAttributeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeImageAttributeRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeImageAttributeRequest(_a0 *ec2.DescribeImageAttributeInput) (*request.Request, *ec2.DescribeImageAttributeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeImageAttributeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeImageAttributeOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeImageAttributeInput) *ec2.DescribeImageAttributeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeImageAttributeOutput)
		}
	}

	return r0, r1
}

// DescribeImageAttributeWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeImageAttributeWithContext(_a0 aws.Context, _a1 *ec2.DescribeImageAttributeInput, _a2 ...request.Option) (*ec2.DescribeImageAttributeOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeImageAttributeOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeImageAttributeInput, ...request.Option) *ec2.DescribeImageAttributeOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeImageAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeImageAttributeInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeImages provides a mock function with given fields: _a0
func (_m *EC2API) DescribeImages(_a0 *ec2.DescribeImagesInput) (*ec2.DescribeImagesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeImagesOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeImagesInput) *ec2.DescribeImagesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeImagesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeImagesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeImagesRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeImagesRequest(_a0 *ec2.DescribeImagesInput) (*request.Request, *ec2.DescribeImagesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeImagesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeImagesOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeImagesInput) *ec2.DescribeImagesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeImagesOutput)
		}
	}

	return r0, r1
}

// DescribeImagesWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeImagesWithContext(_a0 aws.Context, _a1 *ec2.DescribeImagesInput, _a2 ...request.Option) (*ec2.DescribeImagesOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeImagesOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeImagesInput, ...request.Option) *ec2.DescribeImagesOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeImagesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeImagesInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeImportImageTasks provides a mock function with given fields: _a0
func (_m *EC2API) DescribeImportImageTasks(_a0 *ec2.DescribeImportImageTasksInput) (*ec2.DescribeImportImageTasksOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeImportImageTasksOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeImportImageTasksInput) *ec2.DescribeImportImageTasksOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeImportImageTasksOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeImportImageTasksInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeImportImageTasksRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeImportImageTasksRequest(_a0 *ec2.DescribeImportImageTasksInput) (*request.Request, *ec2.DescribeImportImageTasksOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeImportImageTasksInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeImportImageTasksOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeImportImageTasksInput) *ec2.DescribeImportImageTasksOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeImportImageTasksOutput)
		}
	}

	return r0, r1
}

// DescribeImportImageTasksWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeImportImageTasksWithContext(_a0 aws.Context, _a1 *ec2.DescribeImportImageTasksInput, _a2 ...request.Option) (*ec2.DescribeImportImageTasksOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeImportImageTasksOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeImportImageTasksInput, ...request.Option) *ec2.DescribeImportImageTasksOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeImportImageTasksOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeImportImageTasksInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeImportSnapshotTasks provides a mock function with given fields: _a0
func (_m *EC2API) DescribeImportSnapshotTasks(_a0 *ec2.DescribeImportSnapshotTasksInput) (*ec2.DescribeImportSnapshotTasksOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeImportSnapshotTasksOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeImportSnapshotTasksInput) *ec2.DescribeImportSnapshotTasksOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeImportSnapshotTasksOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeImportSnapshotTasksInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeImportSnapshotTasksRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeImportSnapshotTasksRequest(_a0 *ec2.DescribeImportSnapshotTasksInput) (*request.Request, *ec2.DescribeImportSnapshotTasksOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeImportSnapshotTasksInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeImportSnapshotTasksOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeImportSnapshotTasksInput) *ec2.DescribeImportSnapshotTasksOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeImportSnapshotTasksOutput)
		}
	}

	return r0, r1
}

// DescribeImportSnapshotTasksWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeImportSnapshotTasksWithContext(_a0 aws.Context, _a1 *ec2.DescribeImportSnapshotTasksInput, _a2 ...request.Option) (*ec2.DescribeImportSnapshotTasksOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeImportSnapshotTasksOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeImportSnapshotTasksInput, ...request.Option) *ec2.DescribeImportSnapshotTasksOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeImportSnapshotTasksOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeImportSnapshotTasksInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}



// DescribeInstanceAttributeRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeInstanceAttributeRequest(_a0 *ec2.DescribeInstanceAttributeInput) (*request.Request, *ec2.DescribeInstanceAttributeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeInstanceAttributeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeInstanceAttributeOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeInstanceAttributeInput) *ec2.DescribeInstanceAttributeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeInstanceAttributeOutput)
		}
	}

	return r0, r1
}

// DescribeInstanceAttributeWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeInstanceAttributeWithContext(_a0 aws.Context, _a1 *ec2.DescribeInstanceAttributeInput, _a2 ...request.Option) (*ec2.DescribeInstanceAttributeOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeInstanceAttributeOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeInstanceAttributeInput, ...request.Option) *ec2.DescribeInstanceAttributeOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeInstanceAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeInstanceAttributeInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeInstanceCreditSpecifications provides a mock function with given fields: _a0
func (_m *EC2API) DescribeInstanceCreditSpecifications(_a0 *ec2.DescribeInstanceCreditSpecificationsInput) (*ec2.DescribeInstanceCreditSpecificationsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeInstanceCreditSpecificationsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeInstanceCreditSpecificationsInput) *ec2.DescribeInstanceCreditSpecificationsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeInstanceCreditSpecificationsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeInstanceCreditSpecificationsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeInstanceCreditSpecificationsRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeInstanceCreditSpecificationsRequest(_a0 *ec2.DescribeInstanceCreditSpecificationsInput) (*request.Request, *ec2.DescribeInstanceCreditSpecificationsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeInstanceCreditSpecificationsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeInstanceCreditSpecificationsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeInstanceCreditSpecificationsInput) *ec2.DescribeInstanceCreditSpecificationsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeInstanceCreditSpecificationsOutput)
		}
	}

	return r0, r1
}

// DescribeInstanceCreditSpecificationsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeInstanceCreditSpecificationsWithContext(_a0 aws.Context, _a1 *ec2.DescribeInstanceCreditSpecificationsInput, _a2 ...request.Option) (*ec2.DescribeInstanceCreditSpecificationsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeInstanceCreditSpecificationsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeInstanceCreditSpecificationsInput, ...request.Option) *ec2.DescribeInstanceCreditSpecificationsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeInstanceCreditSpecificationsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeInstanceCreditSpecificationsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeInstanceStatus provides a mock function with given fields: _a0
func (_m *EC2API) DescribeInstanceStatus(_a0 *ec2.DescribeInstanceStatusInput) (*ec2.DescribeInstanceStatusOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeInstanceStatusOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeInstanceStatusInput) *ec2.DescribeInstanceStatusOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeInstanceStatusOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeInstanceStatusInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeInstanceStatusPages provides a mock function with given fields: _a0, _a1
func (_m *EC2API) DescribeInstanceStatusPages(_a0 *ec2.DescribeInstanceStatusInput, _a1 func(*ec2.DescribeInstanceStatusOutput, bool) bool) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeInstanceStatusInput, func(*ec2.DescribeInstanceStatusOutput, bool) bool) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DescribeInstanceStatusPagesWithContext provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *EC2API) DescribeInstanceStatusPagesWithContext(_a0 aws.Context, _a1 *ec2.DescribeInstanceStatusInput, _a2 func(*ec2.DescribeInstanceStatusOutput, bool) bool, _a3 ...request.Option) error {
	_va := make([]interface{}, len(_a3))
	for _i := range _a3 {
		_va[_i] = _a3[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1, _a2)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeInstanceStatusInput, func(*ec2.DescribeInstanceStatusOutput, bool) bool, ...request.Option) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DescribeInstanceStatusRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeInstanceStatusRequest(_a0 *ec2.DescribeInstanceStatusInput) (*request.Request, *ec2.DescribeInstanceStatusOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeInstanceStatusInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeInstanceStatusOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeInstanceStatusInput) *ec2.DescribeInstanceStatusOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeInstanceStatusOutput)
		}
	}

	return r0, r1
}

// DescribeInstanceStatusWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeInstanceStatusWithContext(_a0 aws.Context, _a1 *ec2.DescribeInstanceStatusInput, _a2 ...request.Option) (*ec2.DescribeInstanceStatusOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeInstanceStatusOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeInstanceStatusInput, ...request.Option) *ec2.DescribeInstanceStatusOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeInstanceStatusOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeInstanceStatusInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}



// DescribeInstancesPages provides a mock function with given fields: _a0, _a1
func (_m *EC2API) DescribeInstancesPages(_a0 *ec2.DescribeInstancesInput, _a1 func(*ec2.DescribeInstancesOutput, bool) bool) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeInstancesInput, func(*ec2.DescribeInstancesOutput, bool) bool) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DescribeInstancesPagesWithContext provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *EC2API) DescribeInstancesPagesWithContext(_a0 aws.Context, _a1 *ec2.DescribeInstancesInput, _a2 func(*ec2.DescribeInstancesOutput, bool) bool, _a3 ...request.Option) error {
	_va := make([]interface{}, len(_a3))
	for _i := range _a3 {
		_va[_i] = _a3[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1, _a2)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeInstancesInput, func(*ec2.DescribeInstancesOutput, bool) bool, ...request.Option) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DescribeInstancesRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeInstancesRequest(_a0 *ec2.DescribeInstancesInput) (*request.Request, *ec2.DescribeInstancesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeInstancesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeInstancesOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeInstancesInput) *ec2.DescribeInstancesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeInstancesOutput)
		}
	}

	return r0, r1
}

// DescribeInstancesWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeInstancesWithContext(_a0 aws.Context, _a1 *ec2.DescribeInstancesInput, _a2 ...request.Option) (*ec2.DescribeInstancesOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeInstancesOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeInstancesInput, ...request.Option) *ec2.DescribeInstancesOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeInstancesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeInstancesInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeInternetGateways provides a mock function with given fields: _a0
func (_m *EC2API) DescribeInternetGateways(_a0 *ec2.DescribeInternetGatewaysInput) (*ec2.DescribeInternetGatewaysOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeInternetGatewaysOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeInternetGatewaysInput) *ec2.DescribeInternetGatewaysOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeInternetGatewaysOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeInternetGatewaysInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeInternetGatewaysRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeInternetGatewaysRequest(_a0 *ec2.DescribeInternetGatewaysInput) (*request.Request, *ec2.DescribeInternetGatewaysOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeInternetGatewaysInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeInternetGatewaysOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeInternetGatewaysInput) *ec2.DescribeInternetGatewaysOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeInternetGatewaysOutput)
		}
	}

	return r0, r1
}

// DescribeInternetGatewaysWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeInternetGatewaysWithContext(_a0 aws.Context, _a1 *ec2.DescribeInternetGatewaysInput, _a2 ...request.Option) (*ec2.DescribeInternetGatewaysOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeInternetGatewaysOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeInternetGatewaysInput, ...request.Option) *ec2.DescribeInternetGatewaysOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeInternetGatewaysOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeInternetGatewaysInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeKeyPairs provides a mock function with given fields: _a0
func (_m *EC2API) DescribeKeyPairs(_a0 *ec2.DescribeKeyPairsInput) (*ec2.DescribeKeyPairsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeKeyPairsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeKeyPairsInput) *ec2.DescribeKeyPairsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeKeyPairsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeKeyPairsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeKeyPairsRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeKeyPairsRequest(_a0 *ec2.DescribeKeyPairsInput) (*request.Request, *ec2.DescribeKeyPairsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeKeyPairsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeKeyPairsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeKeyPairsInput) *ec2.DescribeKeyPairsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeKeyPairsOutput)
		}
	}

	return r0, r1
}

// DescribeKeyPairsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeKeyPairsWithContext(_a0 aws.Context, _a1 *ec2.DescribeKeyPairsInput, _a2 ...request.Option) (*ec2.DescribeKeyPairsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeKeyPairsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeKeyPairsInput, ...request.Option) *ec2.DescribeKeyPairsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeKeyPairsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeKeyPairsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeLaunchTemplateVersions provides a mock function with given fields: _a0
func (_m *EC2API) DescribeLaunchTemplateVersions(_a0 *ec2.DescribeLaunchTemplateVersionsInput) (*ec2.DescribeLaunchTemplateVersionsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeLaunchTemplateVersionsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeLaunchTemplateVersionsInput) *ec2.DescribeLaunchTemplateVersionsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeLaunchTemplateVersionsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeLaunchTemplateVersionsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeLaunchTemplateVersionsRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeLaunchTemplateVersionsRequest(_a0 *ec2.DescribeLaunchTemplateVersionsInput) (*request.Request, *ec2.DescribeLaunchTemplateVersionsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeLaunchTemplateVersionsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeLaunchTemplateVersionsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeLaunchTemplateVersionsInput) *ec2.DescribeLaunchTemplateVersionsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeLaunchTemplateVersionsOutput)
		}
	}

	return r0, r1
}

// DescribeLaunchTemplateVersionsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeLaunchTemplateVersionsWithContext(_a0 aws.Context, _a1 *ec2.DescribeLaunchTemplateVersionsInput, _a2 ...request.Option) (*ec2.DescribeLaunchTemplateVersionsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeLaunchTemplateVersionsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeLaunchTemplateVersionsInput, ...request.Option) *ec2.DescribeLaunchTemplateVersionsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeLaunchTemplateVersionsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeLaunchTemplateVersionsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeLaunchTemplates provides a mock function with given fields: _a0
func (_m *EC2API) DescribeLaunchTemplates(_a0 *ec2.DescribeLaunchTemplatesInput) (*ec2.DescribeLaunchTemplatesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeLaunchTemplatesOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeLaunchTemplatesInput) *ec2.DescribeLaunchTemplatesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeLaunchTemplatesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeLaunchTemplatesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeLaunchTemplatesRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeLaunchTemplatesRequest(_a0 *ec2.DescribeLaunchTemplatesInput) (*request.Request, *ec2.DescribeLaunchTemplatesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeLaunchTemplatesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeLaunchTemplatesOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeLaunchTemplatesInput) *ec2.DescribeLaunchTemplatesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeLaunchTemplatesOutput)
		}
	}

	return r0, r1
}

// DescribeLaunchTemplatesWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeLaunchTemplatesWithContext(_a0 aws.Context, _a1 *ec2.DescribeLaunchTemplatesInput, _a2 ...request.Option) (*ec2.DescribeLaunchTemplatesOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeLaunchTemplatesOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeLaunchTemplatesInput, ...request.Option) *ec2.DescribeLaunchTemplatesOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeLaunchTemplatesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeLaunchTemplatesInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeMovingAddresses provides a mock function with given fields: _a0
func (_m *EC2API) DescribeMovingAddresses(_a0 *ec2.DescribeMovingAddressesInput) (*ec2.DescribeMovingAddressesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeMovingAddressesOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeMovingAddressesInput) *ec2.DescribeMovingAddressesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeMovingAddressesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeMovingAddressesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeMovingAddressesRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeMovingAddressesRequest(_a0 *ec2.DescribeMovingAddressesInput) (*request.Request, *ec2.DescribeMovingAddressesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeMovingAddressesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeMovingAddressesOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeMovingAddressesInput) *ec2.DescribeMovingAddressesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeMovingAddressesOutput)
		}
	}

	return r0, r1
}

// DescribeMovingAddressesWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeMovingAddressesWithContext(_a0 aws.Context, _a1 *ec2.DescribeMovingAddressesInput, _a2 ...request.Option) (*ec2.DescribeMovingAddressesOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeMovingAddressesOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeMovingAddressesInput, ...request.Option) *ec2.DescribeMovingAddressesOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeMovingAddressesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeMovingAddressesInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeNatGateways provides a mock function with given fields: _a0
func (_m *EC2API) DescribeNatGateways(_a0 *ec2.DescribeNatGatewaysInput) (*ec2.DescribeNatGatewaysOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeNatGatewaysOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeNatGatewaysInput) *ec2.DescribeNatGatewaysOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeNatGatewaysOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeNatGatewaysInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeNatGatewaysPages provides a mock function with given fields: _a0, _a1
func (_m *EC2API) DescribeNatGatewaysPages(_a0 *ec2.DescribeNatGatewaysInput, _a1 func(*ec2.DescribeNatGatewaysOutput, bool) bool) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeNatGatewaysInput, func(*ec2.DescribeNatGatewaysOutput, bool) bool) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DescribeNatGatewaysPagesWithContext provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *EC2API) DescribeNatGatewaysPagesWithContext(_a0 aws.Context, _a1 *ec2.DescribeNatGatewaysInput, _a2 func(*ec2.DescribeNatGatewaysOutput, bool) bool, _a3 ...request.Option) error {
	_va := make([]interface{}, len(_a3))
	for _i := range _a3 {
		_va[_i] = _a3[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1, _a2)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeNatGatewaysInput, func(*ec2.DescribeNatGatewaysOutput, bool) bool, ...request.Option) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DescribeNatGatewaysRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeNatGatewaysRequest(_a0 *ec2.DescribeNatGatewaysInput) (*request.Request, *ec2.DescribeNatGatewaysOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeNatGatewaysInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeNatGatewaysOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeNatGatewaysInput) *ec2.DescribeNatGatewaysOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeNatGatewaysOutput)
		}
	}

	return r0, r1
}

// DescribeNatGatewaysWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeNatGatewaysWithContext(_a0 aws.Context, _a1 *ec2.DescribeNatGatewaysInput, _a2 ...request.Option) (*ec2.DescribeNatGatewaysOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeNatGatewaysOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeNatGatewaysInput, ...request.Option) *ec2.DescribeNatGatewaysOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeNatGatewaysOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeNatGatewaysInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeNetworkAcls provides a mock function with given fields: _a0
func (_m *EC2API) DescribeNetworkAcls(_a0 *ec2.DescribeNetworkAclsInput) (*ec2.DescribeNetworkAclsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeNetworkAclsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeNetworkAclsInput) *ec2.DescribeNetworkAclsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeNetworkAclsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeNetworkAclsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeNetworkAclsRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeNetworkAclsRequest(_a0 *ec2.DescribeNetworkAclsInput) (*request.Request, *ec2.DescribeNetworkAclsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeNetworkAclsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeNetworkAclsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeNetworkAclsInput) *ec2.DescribeNetworkAclsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeNetworkAclsOutput)
		}
	}

	return r0, r1
}

// DescribeNetworkAclsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeNetworkAclsWithContext(_a0 aws.Context, _a1 *ec2.DescribeNetworkAclsInput, _a2 ...request.Option) (*ec2.DescribeNetworkAclsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeNetworkAclsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeNetworkAclsInput, ...request.Option) *ec2.DescribeNetworkAclsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeNetworkAclsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeNetworkAclsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeNetworkInterfaceAttribute provides a mock function with given fields: _a0
func (_m *EC2API) DescribeNetworkInterfaceAttribute(_a0 *ec2.DescribeNetworkInterfaceAttributeInput) (*ec2.DescribeNetworkInterfaceAttributeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeNetworkInterfaceAttributeOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeNetworkInterfaceAttributeInput) *ec2.DescribeNetworkInterfaceAttributeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeNetworkInterfaceAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeNetworkInterfaceAttributeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeNetworkInterfaceAttributeRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeNetworkInterfaceAttributeRequest(_a0 *ec2.DescribeNetworkInterfaceAttributeInput) (*request.Request, *ec2.DescribeNetworkInterfaceAttributeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeNetworkInterfaceAttributeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeNetworkInterfaceAttributeOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeNetworkInterfaceAttributeInput) *ec2.DescribeNetworkInterfaceAttributeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeNetworkInterfaceAttributeOutput)
		}
	}

	return r0, r1
}

// DescribeNetworkInterfaceAttributeWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeNetworkInterfaceAttributeWithContext(_a0 aws.Context, _a1 *ec2.DescribeNetworkInterfaceAttributeInput, _a2 ...request.Option) (*ec2.DescribeNetworkInterfaceAttributeOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeNetworkInterfaceAttributeOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeNetworkInterfaceAttributeInput, ...request.Option) *ec2.DescribeNetworkInterfaceAttributeOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeNetworkInterfaceAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeNetworkInterfaceAttributeInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeNetworkInterfacePermissions provides a mock function with given fields: _a0
func (_m *EC2API) DescribeNetworkInterfacePermissions(_a0 *ec2.DescribeNetworkInterfacePermissionsInput) (*ec2.DescribeNetworkInterfacePermissionsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeNetworkInterfacePermissionsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeNetworkInterfacePermissionsInput) *ec2.DescribeNetworkInterfacePermissionsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeNetworkInterfacePermissionsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeNetworkInterfacePermissionsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeNetworkInterfacePermissionsRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeNetworkInterfacePermissionsRequest(_a0 *ec2.DescribeNetworkInterfacePermissionsInput) (*request.Request, *ec2.DescribeNetworkInterfacePermissionsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeNetworkInterfacePermissionsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeNetworkInterfacePermissionsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeNetworkInterfacePermissionsInput) *ec2.DescribeNetworkInterfacePermissionsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeNetworkInterfacePermissionsOutput)
		}
	}

	return r0, r1
}

// DescribeNetworkInterfacePermissionsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeNetworkInterfacePermissionsWithContext(_a0 aws.Context, _a1 *ec2.DescribeNetworkInterfacePermissionsInput, _a2 ...request.Option) (*ec2.DescribeNetworkInterfacePermissionsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeNetworkInterfacePermissionsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeNetworkInterfacePermissionsInput, ...request.Option) *ec2.DescribeNetworkInterfacePermissionsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeNetworkInterfacePermissionsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeNetworkInterfacePermissionsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}


// DescribeNetworkInterfacesRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeNetworkInterfacesRequest(_a0 *ec2.DescribeNetworkInterfacesInput) (*request.Request, *ec2.DescribeNetworkInterfacesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeNetworkInterfacesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeNetworkInterfacesOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeNetworkInterfacesInput) *ec2.DescribeNetworkInterfacesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeNetworkInterfacesOutput)
		}
	}

	return r0, r1
}

// DescribeNetworkInterfacesWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeNetworkInterfacesWithContext(_a0 aws.Context, _a1 *ec2.DescribeNetworkInterfacesInput, _a2 ...request.Option) (*ec2.DescribeNetworkInterfacesOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeNetworkInterfacesOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeNetworkInterfacesInput, ...request.Option) *ec2.DescribeNetworkInterfacesOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeNetworkInterfacesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeNetworkInterfacesInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribePlacementGroups provides a mock function with given fields: _a0
func (_m *EC2API) DescribePlacementGroups(_a0 *ec2.DescribePlacementGroupsInput) (*ec2.DescribePlacementGroupsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribePlacementGroupsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribePlacementGroupsInput) *ec2.DescribePlacementGroupsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribePlacementGroupsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribePlacementGroupsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribePlacementGroupsRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribePlacementGroupsRequest(_a0 *ec2.DescribePlacementGroupsInput) (*request.Request, *ec2.DescribePlacementGroupsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribePlacementGroupsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribePlacementGroupsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribePlacementGroupsInput) *ec2.DescribePlacementGroupsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribePlacementGroupsOutput)
		}
	}

	return r0, r1
}

// DescribePlacementGroupsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribePlacementGroupsWithContext(_a0 aws.Context, _a1 *ec2.DescribePlacementGroupsInput, _a2 ...request.Option) (*ec2.DescribePlacementGroupsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribePlacementGroupsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribePlacementGroupsInput, ...request.Option) *ec2.DescribePlacementGroupsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribePlacementGroupsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribePlacementGroupsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribePrefixLists provides a mock function with given fields: _a0
func (_m *EC2API) DescribePrefixLists(_a0 *ec2.DescribePrefixListsInput) (*ec2.DescribePrefixListsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribePrefixListsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribePrefixListsInput) *ec2.DescribePrefixListsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribePrefixListsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribePrefixListsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribePrefixListsRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribePrefixListsRequest(_a0 *ec2.DescribePrefixListsInput) (*request.Request, *ec2.DescribePrefixListsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribePrefixListsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribePrefixListsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribePrefixListsInput) *ec2.DescribePrefixListsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribePrefixListsOutput)
		}
	}

	return r0, r1
}

// DescribePrefixListsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribePrefixListsWithContext(_a0 aws.Context, _a1 *ec2.DescribePrefixListsInput, _a2 ...request.Option) (*ec2.DescribePrefixListsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribePrefixListsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribePrefixListsInput, ...request.Option) *ec2.DescribePrefixListsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribePrefixListsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribePrefixListsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribePrincipalIdFormat provides a mock function with given fields: _a0
func (_m *EC2API) DescribePrincipalIdFormat(_a0 *ec2.DescribePrincipalIdFormatInput) (*ec2.DescribePrincipalIdFormatOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribePrincipalIdFormatOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribePrincipalIdFormatInput) *ec2.DescribePrincipalIdFormatOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribePrincipalIdFormatOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribePrincipalIdFormatInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribePrincipalIdFormatRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribePrincipalIdFormatRequest(_a0 *ec2.DescribePrincipalIdFormatInput) (*request.Request, *ec2.DescribePrincipalIdFormatOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribePrincipalIdFormatInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribePrincipalIdFormatOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribePrincipalIdFormatInput) *ec2.DescribePrincipalIdFormatOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribePrincipalIdFormatOutput)
		}
	}

	return r0, r1
}

// DescribePrincipalIdFormatWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribePrincipalIdFormatWithContext(_a0 aws.Context, _a1 *ec2.DescribePrincipalIdFormatInput, _a2 ...request.Option) (*ec2.DescribePrincipalIdFormatOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribePrincipalIdFormatOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribePrincipalIdFormatInput, ...request.Option) *ec2.DescribePrincipalIdFormatOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribePrincipalIdFormatOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribePrincipalIdFormatInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeRegions provides a mock function with given fields: _a0
func (_m *EC2API) DescribeRegions(_a0 *ec2.DescribeRegionsInput) (*ec2.DescribeRegionsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeRegionsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeRegionsInput) *ec2.DescribeRegionsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeRegionsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeRegionsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeRegionsRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeRegionsRequest(_a0 *ec2.DescribeRegionsInput) (*request.Request, *ec2.DescribeRegionsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeRegionsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeRegionsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeRegionsInput) *ec2.DescribeRegionsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeRegionsOutput)
		}
	}

	return r0, r1
}

// DescribeRegionsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeRegionsWithContext(_a0 aws.Context, _a1 *ec2.DescribeRegionsInput, _a2 ...request.Option) (*ec2.DescribeRegionsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeRegionsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeRegionsInput, ...request.Option) *ec2.DescribeRegionsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeRegionsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeRegionsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeReservedInstances provides a mock function with given fields: _a0
func (_m *EC2API) DescribeReservedInstances(_a0 *ec2.DescribeReservedInstancesInput) (*ec2.DescribeReservedInstancesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeReservedInstancesOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeReservedInstancesInput) *ec2.DescribeReservedInstancesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeReservedInstancesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeReservedInstancesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeReservedInstancesListings provides a mock function with given fields: _a0
func (_m *EC2API) DescribeReservedInstancesListings(_a0 *ec2.DescribeReservedInstancesListingsInput) (*ec2.DescribeReservedInstancesListingsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeReservedInstancesListingsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeReservedInstancesListingsInput) *ec2.DescribeReservedInstancesListingsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeReservedInstancesListingsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeReservedInstancesListingsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeReservedInstancesListingsRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeReservedInstancesListingsRequest(_a0 *ec2.DescribeReservedInstancesListingsInput) (*request.Request, *ec2.DescribeReservedInstancesListingsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeReservedInstancesListingsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeReservedInstancesListingsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeReservedInstancesListingsInput) *ec2.DescribeReservedInstancesListingsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeReservedInstancesListingsOutput)
		}
	}

	return r0, r1
}

// DescribeReservedInstancesListingsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeReservedInstancesListingsWithContext(_a0 aws.Context, _a1 *ec2.DescribeReservedInstancesListingsInput, _a2 ...request.Option) (*ec2.DescribeReservedInstancesListingsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeReservedInstancesListingsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeReservedInstancesListingsInput, ...request.Option) *ec2.DescribeReservedInstancesListingsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeReservedInstancesListingsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeReservedInstancesListingsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeReservedInstancesModifications provides a mock function with given fields: _a0
func (_m *EC2API) DescribeReservedInstancesModifications(_a0 *ec2.DescribeReservedInstancesModificationsInput) (*ec2.DescribeReservedInstancesModificationsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeReservedInstancesModificationsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeReservedInstancesModificationsInput) *ec2.DescribeReservedInstancesModificationsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeReservedInstancesModificationsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeReservedInstancesModificationsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeReservedInstancesModificationsPages provides a mock function with given fields: _a0, _a1
func (_m *EC2API) DescribeReservedInstancesModificationsPages(_a0 *ec2.DescribeReservedInstancesModificationsInput, _a1 func(*ec2.DescribeReservedInstancesModificationsOutput, bool) bool) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeReservedInstancesModificationsInput, func(*ec2.DescribeReservedInstancesModificationsOutput, bool) bool) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DescribeReservedInstancesModificationsPagesWithContext provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *EC2API) DescribeReservedInstancesModificationsPagesWithContext(_a0 aws.Context, _a1 *ec2.DescribeReservedInstancesModificationsInput, _a2 func(*ec2.DescribeReservedInstancesModificationsOutput, bool) bool, _a3 ...request.Option) error {
	_va := make([]interface{}, len(_a3))
	for _i := range _a3 {
		_va[_i] = _a3[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1, _a2)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeReservedInstancesModificationsInput, func(*ec2.DescribeReservedInstancesModificationsOutput, bool) bool, ...request.Option) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DescribeReservedInstancesModificationsRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeReservedInstancesModificationsRequest(_a0 *ec2.DescribeReservedInstancesModificationsInput) (*request.Request, *ec2.DescribeReservedInstancesModificationsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeReservedInstancesModificationsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeReservedInstancesModificationsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeReservedInstancesModificationsInput) *ec2.DescribeReservedInstancesModificationsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeReservedInstancesModificationsOutput)
		}
	}

	return r0, r1
}

// DescribeReservedInstancesModificationsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeReservedInstancesModificationsWithContext(_a0 aws.Context, _a1 *ec2.DescribeReservedInstancesModificationsInput, _a2 ...request.Option) (*ec2.DescribeReservedInstancesModificationsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeReservedInstancesModificationsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeReservedInstancesModificationsInput, ...request.Option) *ec2.DescribeReservedInstancesModificationsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeReservedInstancesModificationsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeReservedInstancesModificationsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeReservedInstancesOfferings provides a mock function with given fields: _a0
func (_m *EC2API) DescribeReservedInstancesOfferings(_a0 *ec2.DescribeReservedInstancesOfferingsInput) (*ec2.DescribeReservedInstancesOfferingsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeReservedInstancesOfferingsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeReservedInstancesOfferingsInput) *ec2.DescribeReservedInstancesOfferingsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeReservedInstancesOfferingsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeReservedInstancesOfferingsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeReservedInstancesOfferingsPages provides a mock function with given fields: _a0, _a1
func (_m *EC2API) DescribeReservedInstancesOfferingsPages(_a0 *ec2.DescribeReservedInstancesOfferingsInput, _a1 func(*ec2.DescribeReservedInstancesOfferingsOutput, bool) bool) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeReservedInstancesOfferingsInput, func(*ec2.DescribeReservedInstancesOfferingsOutput, bool) bool) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DescribeReservedInstancesOfferingsPagesWithContext provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *EC2API) DescribeReservedInstancesOfferingsPagesWithContext(_a0 aws.Context, _a1 *ec2.DescribeReservedInstancesOfferingsInput, _a2 func(*ec2.DescribeReservedInstancesOfferingsOutput, bool) bool, _a3 ...request.Option) error {
	_va := make([]interface{}, len(_a3))
	for _i := range _a3 {
		_va[_i] = _a3[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1, _a2)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeReservedInstancesOfferingsInput, func(*ec2.DescribeReservedInstancesOfferingsOutput, bool) bool, ...request.Option) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DescribeReservedInstancesOfferingsRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeReservedInstancesOfferingsRequest(_a0 *ec2.DescribeReservedInstancesOfferingsInput) (*request.Request, *ec2.DescribeReservedInstancesOfferingsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeReservedInstancesOfferingsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeReservedInstancesOfferingsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeReservedInstancesOfferingsInput) *ec2.DescribeReservedInstancesOfferingsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeReservedInstancesOfferingsOutput)
		}
	}

	return r0, r1
}

// DescribeReservedInstancesOfferingsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeReservedInstancesOfferingsWithContext(_a0 aws.Context, _a1 *ec2.DescribeReservedInstancesOfferingsInput, _a2 ...request.Option) (*ec2.DescribeReservedInstancesOfferingsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeReservedInstancesOfferingsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeReservedInstancesOfferingsInput, ...request.Option) *ec2.DescribeReservedInstancesOfferingsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeReservedInstancesOfferingsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeReservedInstancesOfferingsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeReservedInstancesRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeReservedInstancesRequest(_a0 *ec2.DescribeReservedInstancesInput) (*request.Request, *ec2.DescribeReservedInstancesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeReservedInstancesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeReservedInstancesOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeReservedInstancesInput) *ec2.DescribeReservedInstancesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeReservedInstancesOutput)
		}
	}

	return r0, r1
}

// DescribeReservedInstancesWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeReservedInstancesWithContext(_a0 aws.Context, _a1 *ec2.DescribeReservedInstancesInput, _a2 ...request.Option) (*ec2.DescribeReservedInstancesOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeReservedInstancesOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeReservedInstancesInput, ...request.Option) *ec2.DescribeReservedInstancesOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeReservedInstancesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeReservedInstancesInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeRouteTablesRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeRouteTablesRequest(_a0 *ec2.DescribeRouteTablesInput) (*request.Request, *ec2.DescribeRouteTablesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeRouteTablesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeRouteTablesOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeRouteTablesInput) *ec2.DescribeRouteTablesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeRouteTablesOutput)
		}
	}

	return r0, r1
}

// DescribeRouteTablesWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeRouteTablesWithContext(_a0 aws.Context, _a1 *ec2.DescribeRouteTablesInput, _a2 ...request.Option) (*ec2.DescribeRouteTablesOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeRouteTablesOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeRouteTablesInput, ...request.Option) *ec2.DescribeRouteTablesOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeRouteTablesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeRouteTablesInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeScheduledInstanceAvailability provides a mock function with given fields: _a0
func (_m *EC2API) DescribeScheduledInstanceAvailability(_a0 *ec2.DescribeScheduledInstanceAvailabilityInput) (*ec2.DescribeScheduledInstanceAvailabilityOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeScheduledInstanceAvailabilityOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeScheduledInstanceAvailabilityInput) *ec2.DescribeScheduledInstanceAvailabilityOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeScheduledInstanceAvailabilityOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeScheduledInstanceAvailabilityInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeScheduledInstanceAvailabilityRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeScheduledInstanceAvailabilityRequest(_a0 *ec2.DescribeScheduledInstanceAvailabilityInput) (*request.Request, *ec2.DescribeScheduledInstanceAvailabilityOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeScheduledInstanceAvailabilityInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeScheduledInstanceAvailabilityOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeScheduledInstanceAvailabilityInput) *ec2.DescribeScheduledInstanceAvailabilityOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeScheduledInstanceAvailabilityOutput)
		}
	}

	return r0, r1
}

// DescribeScheduledInstanceAvailabilityWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeScheduledInstanceAvailabilityWithContext(_a0 aws.Context, _a1 *ec2.DescribeScheduledInstanceAvailabilityInput, _a2 ...request.Option) (*ec2.DescribeScheduledInstanceAvailabilityOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeScheduledInstanceAvailabilityOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeScheduledInstanceAvailabilityInput, ...request.Option) *ec2.DescribeScheduledInstanceAvailabilityOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeScheduledInstanceAvailabilityOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeScheduledInstanceAvailabilityInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeScheduledInstances provides a mock function with given fields: _a0
func (_m *EC2API) DescribeScheduledInstances(_a0 *ec2.DescribeScheduledInstancesInput) (*ec2.DescribeScheduledInstancesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeScheduledInstancesOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeScheduledInstancesInput) *ec2.DescribeScheduledInstancesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeScheduledInstancesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeScheduledInstancesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeScheduledInstancesRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeScheduledInstancesRequest(_a0 *ec2.DescribeScheduledInstancesInput) (*request.Request, *ec2.DescribeScheduledInstancesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeScheduledInstancesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeScheduledInstancesOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeScheduledInstancesInput) *ec2.DescribeScheduledInstancesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeScheduledInstancesOutput)
		}
	}

	return r0, r1
}

// DescribeScheduledInstancesWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeScheduledInstancesWithContext(_a0 aws.Context, _a1 *ec2.DescribeScheduledInstancesInput, _a2 ...request.Option) (*ec2.DescribeScheduledInstancesOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeScheduledInstancesOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeScheduledInstancesInput, ...request.Option) *ec2.DescribeScheduledInstancesOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeScheduledInstancesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeScheduledInstancesInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeSecurityGroupReferences provides a mock function with given fields: _a0
func (_m *EC2API) DescribeSecurityGroupReferences(_a0 *ec2.DescribeSecurityGroupReferencesInput) (*ec2.DescribeSecurityGroupReferencesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeSecurityGroupReferencesOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSecurityGroupReferencesInput) *ec2.DescribeSecurityGroupReferencesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeSecurityGroupReferencesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeSecurityGroupReferencesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeSecurityGroupReferencesRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeSecurityGroupReferencesRequest(_a0 *ec2.DescribeSecurityGroupReferencesInput) (*request.Request, *ec2.DescribeSecurityGroupReferencesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSecurityGroupReferencesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeSecurityGroupReferencesOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeSecurityGroupReferencesInput) *ec2.DescribeSecurityGroupReferencesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeSecurityGroupReferencesOutput)
		}
	}

	return r0, r1
}

// DescribeSecurityGroupReferencesWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeSecurityGroupReferencesWithContext(_a0 aws.Context, _a1 *ec2.DescribeSecurityGroupReferencesInput, _a2 ...request.Option) (*ec2.DescribeSecurityGroupReferencesOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeSecurityGroupReferencesOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeSecurityGroupReferencesInput, ...request.Option) *ec2.DescribeSecurityGroupReferencesOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeSecurityGroupReferencesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeSecurityGroupReferencesInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}



// DescribeSecurityGroupsRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeSecurityGroupsRequest(_a0 *ec2.DescribeSecurityGroupsInput) (*request.Request, *ec2.DescribeSecurityGroupsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSecurityGroupsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeSecurityGroupsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeSecurityGroupsInput) *ec2.DescribeSecurityGroupsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeSecurityGroupsOutput)
		}
	}

	return r0, r1
}

// DescribeSecurityGroupsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeSecurityGroupsWithContext(_a0 aws.Context, _a1 *ec2.DescribeSecurityGroupsInput, _a2 ...request.Option) (*ec2.DescribeSecurityGroupsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeSecurityGroupsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeSecurityGroupsInput, ...request.Option) *ec2.DescribeSecurityGroupsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeSecurityGroupsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeSecurityGroupsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeSnapshotAttribute provides a mock function with given fields: _a0
func (_m *EC2API) DescribeSnapshotAttribute(_a0 *ec2.DescribeSnapshotAttributeInput) (*ec2.DescribeSnapshotAttributeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeSnapshotAttributeOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSnapshotAttributeInput) *ec2.DescribeSnapshotAttributeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeSnapshotAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeSnapshotAttributeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeSnapshotAttributeRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeSnapshotAttributeRequest(_a0 *ec2.DescribeSnapshotAttributeInput) (*request.Request, *ec2.DescribeSnapshotAttributeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSnapshotAttributeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeSnapshotAttributeOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeSnapshotAttributeInput) *ec2.DescribeSnapshotAttributeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeSnapshotAttributeOutput)
		}
	}

	return r0, r1
}

// DescribeSnapshotAttributeWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeSnapshotAttributeWithContext(_a0 aws.Context, _a1 *ec2.DescribeSnapshotAttributeInput, _a2 ...request.Option) (*ec2.DescribeSnapshotAttributeOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeSnapshotAttributeOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeSnapshotAttributeInput, ...request.Option) *ec2.DescribeSnapshotAttributeOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeSnapshotAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeSnapshotAttributeInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeSnapshots provides a mock function with given fields: _a0
func (_m *EC2API) DescribeSnapshots(_a0 *ec2.DescribeSnapshotsInput) (*ec2.DescribeSnapshotsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeSnapshotsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSnapshotsInput) *ec2.DescribeSnapshotsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeSnapshotsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeSnapshotsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeSnapshotsPages provides a mock function with given fields: _a0, _a1
func (_m *EC2API) DescribeSnapshotsPages(_a0 *ec2.DescribeSnapshotsInput, _a1 func(*ec2.DescribeSnapshotsOutput, bool) bool) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSnapshotsInput, func(*ec2.DescribeSnapshotsOutput, bool) bool) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DescribeSnapshotsPagesWithContext provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *EC2API) DescribeSnapshotsPagesWithContext(_a0 aws.Context, _a1 *ec2.DescribeSnapshotsInput, _a2 func(*ec2.DescribeSnapshotsOutput, bool) bool, _a3 ...request.Option) error {
	_va := make([]interface{}, len(_a3))
	for _i := range _a3 {
		_va[_i] = _a3[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1, _a2)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeSnapshotsInput, func(*ec2.DescribeSnapshotsOutput, bool) bool, ...request.Option) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DescribeSnapshotsRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeSnapshotsRequest(_a0 *ec2.DescribeSnapshotsInput) (*request.Request, *ec2.DescribeSnapshotsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSnapshotsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeSnapshotsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeSnapshotsInput) *ec2.DescribeSnapshotsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeSnapshotsOutput)
		}
	}

	return r0, r1
}

// DescribeSnapshotsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeSnapshotsWithContext(_a0 aws.Context, _a1 *ec2.DescribeSnapshotsInput, _a2 ...request.Option) (*ec2.DescribeSnapshotsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeSnapshotsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeSnapshotsInput, ...request.Option) *ec2.DescribeSnapshotsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeSnapshotsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeSnapshotsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeSpotDatafeedSubscription provides a mock function with given fields: _a0
func (_m *EC2API) DescribeSpotDatafeedSubscription(_a0 *ec2.DescribeSpotDatafeedSubscriptionInput) (*ec2.DescribeSpotDatafeedSubscriptionOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeSpotDatafeedSubscriptionOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSpotDatafeedSubscriptionInput) *ec2.DescribeSpotDatafeedSubscriptionOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeSpotDatafeedSubscriptionOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeSpotDatafeedSubscriptionInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeSpotDatafeedSubscriptionRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeSpotDatafeedSubscriptionRequest(_a0 *ec2.DescribeSpotDatafeedSubscriptionInput) (*request.Request, *ec2.DescribeSpotDatafeedSubscriptionOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSpotDatafeedSubscriptionInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeSpotDatafeedSubscriptionOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeSpotDatafeedSubscriptionInput) *ec2.DescribeSpotDatafeedSubscriptionOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeSpotDatafeedSubscriptionOutput)
		}
	}

	return r0, r1
}

// DescribeSpotDatafeedSubscriptionWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeSpotDatafeedSubscriptionWithContext(_a0 aws.Context, _a1 *ec2.DescribeSpotDatafeedSubscriptionInput, _a2 ...request.Option) (*ec2.DescribeSpotDatafeedSubscriptionOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeSpotDatafeedSubscriptionOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeSpotDatafeedSubscriptionInput, ...request.Option) *ec2.DescribeSpotDatafeedSubscriptionOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeSpotDatafeedSubscriptionOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeSpotDatafeedSubscriptionInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeSpotFleetInstances provides a mock function with given fields: _a0
func (_m *EC2API) DescribeSpotFleetInstances(_a0 *ec2.DescribeSpotFleetInstancesInput) (*ec2.DescribeSpotFleetInstancesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeSpotFleetInstancesOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSpotFleetInstancesInput) *ec2.DescribeSpotFleetInstancesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeSpotFleetInstancesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeSpotFleetInstancesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeSpotFleetInstancesRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeSpotFleetInstancesRequest(_a0 *ec2.DescribeSpotFleetInstancesInput) (*request.Request, *ec2.DescribeSpotFleetInstancesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSpotFleetInstancesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeSpotFleetInstancesOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeSpotFleetInstancesInput) *ec2.DescribeSpotFleetInstancesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeSpotFleetInstancesOutput)
		}
	}

	return r0, r1
}

// DescribeSpotFleetInstancesWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeSpotFleetInstancesWithContext(_a0 aws.Context, _a1 *ec2.DescribeSpotFleetInstancesInput, _a2 ...request.Option) (*ec2.DescribeSpotFleetInstancesOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeSpotFleetInstancesOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeSpotFleetInstancesInput, ...request.Option) *ec2.DescribeSpotFleetInstancesOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeSpotFleetInstancesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeSpotFleetInstancesInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeSpotFleetRequestHistory provides a mock function with given fields: _a0
func (_m *EC2API) DescribeSpotFleetRequestHistory(_a0 *ec2.DescribeSpotFleetRequestHistoryInput) (*ec2.DescribeSpotFleetRequestHistoryOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeSpotFleetRequestHistoryOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSpotFleetRequestHistoryInput) *ec2.DescribeSpotFleetRequestHistoryOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeSpotFleetRequestHistoryOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeSpotFleetRequestHistoryInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeSpotFleetRequestHistoryRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeSpotFleetRequestHistoryRequest(_a0 *ec2.DescribeSpotFleetRequestHistoryInput) (*request.Request, *ec2.DescribeSpotFleetRequestHistoryOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSpotFleetRequestHistoryInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeSpotFleetRequestHistoryOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeSpotFleetRequestHistoryInput) *ec2.DescribeSpotFleetRequestHistoryOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeSpotFleetRequestHistoryOutput)
		}
	}

	return r0, r1
}

// DescribeSpotFleetRequestHistoryWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeSpotFleetRequestHistoryWithContext(_a0 aws.Context, _a1 *ec2.DescribeSpotFleetRequestHistoryInput, _a2 ...request.Option) (*ec2.DescribeSpotFleetRequestHistoryOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeSpotFleetRequestHistoryOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeSpotFleetRequestHistoryInput, ...request.Option) *ec2.DescribeSpotFleetRequestHistoryOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeSpotFleetRequestHistoryOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeSpotFleetRequestHistoryInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeSpotFleetRequests provides a mock function with given fields: _a0
func (_m *EC2API) DescribeSpotFleetRequests(_a0 *ec2.DescribeSpotFleetRequestsInput) (*ec2.DescribeSpotFleetRequestsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeSpotFleetRequestsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSpotFleetRequestsInput) *ec2.DescribeSpotFleetRequestsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeSpotFleetRequestsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeSpotFleetRequestsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeSpotFleetRequestsPages provides a mock function with given fields: _a0, _a1
func (_m *EC2API) DescribeSpotFleetRequestsPages(_a0 *ec2.DescribeSpotFleetRequestsInput, _a1 func(*ec2.DescribeSpotFleetRequestsOutput, bool) bool) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSpotFleetRequestsInput, func(*ec2.DescribeSpotFleetRequestsOutput, bool) bool) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DescribeSpotFleetRequestsPagesWithContext provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *EC2API) DescribeSpotFleetRequestsPagesWithContext(_a0 aws.Context, _a1 *ec2.DescribeSpotFleetRequestsInput, _a2 func(*ec2.DescribeSpotFleetRequestsOutput, bool) bool, _a3 ...request.Option) error {
	_va := make([]interface{}, len(_a3))
	for _i := range _a3 {
		_va[_i] = _a3[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1, _a2)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeSpotFleetRequestsInput, func(*ec2.DescribeSpotFleetRequestsOutput, bool) bool, ...request.Option) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DescribeSpotFleetRequestsRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeSpotFleetRequestsRequest(_a0 *ec2.DescribeSpotFleetRequestsInput) (*request.Request, *ec2.DescribeSpotFleetRequestsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSpotFleetRequestsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeSpotFleetRequestsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeSpotFleetRequestsInput) *ec2.DescribeSpotFleetRequestsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeSpotFleetRequestsOutput)
		}
	}

	return r0, r1
}

// DescribeSpotFleetRequestsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeSpotFleetRequestsWithContext(_a0 aws.Context, _a1 *ec2.DescribeSpotFleetRequestsInput, _a2 ...request.Option) (*ec2.DescribeSpotFleetRequestsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeSpotFleetRequestsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeSpotFleetRequestsInput, ...request.Option) *ec2.DescribeSpotFleetRequestsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeSpotFleetRequestsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeSpotFleetRequestsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeSpotInstanceRequests provides a mock function with given fields: _a0
func (_m *EC2API) DescribeSpotInstanceRequests(_a0 *ec2.DescribeSpotInstanceRequestsInput) (*ec2.DescribeSpotInstanceRequestsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeSpotInstanceRequestsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSpotInstanceRequestsInput) *ec2.DescribeSpotInstanceRequestsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeSpotInstanceRequestsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeSpotInstanceRequestsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeSpotInstanceRequestsRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeSpotInstanceRequestsRequest(_a0 *ec2.DescribeSpotInstanceRequestsInput) (*request.Request, *ec2.DescribeSpotInstanceRequestsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSpotInstanceRequestsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeSpotInstanceRequestsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeSpotInstanceRequestsInput) *ec2.DescribeSpotInstanceRequestsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeSpotInstanceRequestsOutput)
		}
	}

	return r0, r1
}

// DescribeSpotInstanceRequestsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeSpotInstanceRequestsWithContext(_a0 aws.Context, _a1 *ec2.DescribeSpotInstanceRequestsInput, _a2 ...request.Option) (*ec2.DescribeSpotInstanceRequestsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeSpotInstanceRequestsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeSpotInstanceRequestsInput, ...request.Option) *ec2.DescribeSpotInstanceRequestsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeSpotInstanceRequestsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeSpotInstanceRequestsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeSpotPriceHistory provides a mock function with given fields: _a0
func (_m *EC2API) DescribeSpotPriceHistory(_a0 *ec2.DescribeSpotPriceHistoryInput) (*ec2.DescribeSpotPriceHistoryOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeSpotPriceHistoryOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSpotPriceHistoryInput) *ec2.DescribeSpotPriceHistoryOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeSpotPriceHistoryOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeSpotPriceHistoryInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeSpotPriceHistoryPages provides a mock function with given fields: _a0, _a1
func (_m *EC2API) DescribeSpotPriceHistoryPages(_a0 *ec2.DescribeSpotPriceHistoryInput, _a1 func(*ec2.DescribeSpotPriceHistoryOutput, bool) bool) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSpotPriceHistoryInput, func(*ec2.DescribeSpotPriceHistoryOutput, bool) bool) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DescribeSpotPriceHistoryPagesWithContext provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *EC2API) DescribeSpotPriceHistoryPagesWithContext(_a0 aws.Context, _a1 *ec2.DescribeSpotPriceHistoryInput, _a2 func(*ec2.DescribeSpotPriceHistoryOutput, bool) bool, _a3 ...request.Option) error {
	_va := make([]interface{}, len(_a3))
	for _i := range _a3 {
		_va[_i] = _a3[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1, _a2)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeSpotPriceHistoryInput, func(*ec2.DescribeSpotPriceHistoryOutput, bool) bool, ...request.Option) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DescribeSpotPriceHistoryRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeSpotPriceHistoryRequest(_a0 *ec2.DescribeSpotPriceHistoryInput) (*request.Request, *ec2.DescribeSpotPriceHistoryOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSpotPriceHistoryInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeSpotPriceHistoryOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeSpotPriceHistoryInput) *ec2.DescribeSpotPriceHistoryOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeSpotPriceHistoryOutput)
		}
	}

	return r0, r1
}

// DescribeSpotPriceHistoryWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeSpotPriceHistoryWithContext(_a0 aws.Context, _a1 *ec2.DescribeSpotPriceHistoryInput, _a2 ...request.Option) (*ec2.DescribeSpotPriceHistoryOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeSpotPriceHistoryOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeSpotPriceHistoryInput, ...request.Option) *ec2.DescribeSpotPriceHistoryOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeSpotPriceHistoryOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeSpotPriceHistoryInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeStaleSecurityGroups provides a mock function with given fields: _a0
func (_m *EC2API) DescribeStaleSecurityGroups(_a0 *ec2.DescribeStaleSecurityGroupsInput) (*ec2.DescribeStaleSecurityGroupsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeStaleSecurityGroupsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeStaleSecurityGroupsInput) *ec2.DescribeStaleSecurityGroupsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeStaleSecurityGroupsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeStaleSecurityGroupsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeStaleSecurityGroupsRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeStaleSecurityGroupsRequest(_a0 *ec2.DescribeStaleSecurityGroupsInput) (*request.Request, *ec2.DescribeStaleSecurityGroupsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeStaleSecurityGroupsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeStaleSecurityGroupsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeStaleSecurityGroupsInput) *ec2.DescribeStaleSecurityGroupsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeStaleSecurityGroupsOutput)
		}
	}

	return r0, r1
}

// DescribeStaleSecurityGroupsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeStaleSecurityGroupsWithContext(_a0 aws.Context, _a1 *ec2.DescribeStaleSecurityGroupsInput, _a2 ...request.Option) (*ec2.DescribeStaleSecurityGroupsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeStaleSecurityGroupsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeStaleSecurityGroupsInput, ...request.Option) *ec2.DescribeStaleSecurityGroupsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeStaleSecurityGroupsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeStaleSecurityGroupsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}



// DescribeSubnetsRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeSubnetsRequest(_a0 *ec2.DescribeSubnetsInput) (*request.Request, *ec2.DescribeSubnetsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSubnetsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeSubnetsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeSubnetsInput) *ec2.DescribeSubnetsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeSubnetsOutput)
		}
	}

	return r0, r1
}

// DescribeSubnetsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeSubnetsWithContext(_a0 aws.Context, _a1 *ec2.DescribeSubnetsInput, _a2 ...request.Option) (*ec2.DescribeSubnetsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeSubnetsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeSubnetsInput, ...request.Option) *ec2.DescribeSubnetsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeSubnetsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeSubnetsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeTags provides a mock function with given fields: _a0
func (_m *EC2API) DescribeTags(_a0 *ec2.DescribeTagsInput) (*ec2.DescribeTagsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeTagsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeTagsInput) *ec2.DescribeTagsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeTagsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeTagsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeTagsPages provides a mock function with given fields: _a0, _a1
func (_m *EC2API) DescribeTagsPages(_a0 *ec2.DescribeTagsInput, _a1 func(*ec2.DescribeTagsOutput, bool) bool) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeTagsInput, func(*ec2.DescribeTagsOutput, bool) bool) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DescribeTagsPagesWithContext provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *EC2API) DescribeTagsPagesWithContext(_a0 aws.Context, _a1 *ec2.DescribeTagsInput, _a2 func(*ec2.DescribeTagsOutput, bool) bool, _a3 ...request.Option) error {
	_va := make([]interface{}, len(_a3))
	for _i := range _a3 {
		_va[_i] = _a3[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1, _a2)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeTagsInput, func(*ec2.DescribeTagsOutput, bool) bool, ...request.Option) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DescribeTagsRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeTagsRequest(_a0 *ec2.DescribeTagsInput) (*request.Request, *ec2.DescribeTagsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeTagsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeTagsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeTagsInput) *ec2.DescribeTagsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeTagsOutput)
		}
	}

	return r0, r1
}

// DescribeTagsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeTagsWithContext(_a0 aws.Context, _a1 *ec2.DescribeTagsInput, _a2 ...request.Option) (*ec2.DescribeTagsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeTagsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeTagsInput, ...request.Option) *ec2.DescribeTagsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeTagsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeTagsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeVolumeAttribute provides a mock function with given fields: _a0
func (_m *EC2API) DescribeVolumeAttribute(_a0 *ec2.DescribeVolumeAttributeInput) (*ec2.DescribeVolumeAttributeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeVolumeAttributeOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVolumeAttributeInput) *ec2.DescribeVolumeAttributeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeVolumeAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVolumeAttributeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeVolumeAttributeRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeVolumeAttributeRequest(_a0 *ec2.DescribeVolumeAttributeInput) (*request.Request, *ec2.DescribeVolumeAttributeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVolumeAttributeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeVolumeAttributeOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVolumeAttributeInput) *ec2.DescribeVolumeAttributeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeVolumeAttributeOutput)
		}
	}

	return r0, r1
}

// DescribeVolumeAttributeWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeVolumeAttributeWithContext(_a0 aws.Context, _a1 *ec2.DescribeVolumeAttributeInput, _a2 ...request.Option) (*ec2.DescribeVolumeAttributeOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeVolumeAttributeOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeVolumeAttributeInput, ...request.Option) *ec2.DescribeVolumeAttributeOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeVolumeAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeVolumeAttributeInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeVolumeStatus provides a mock function with given fields: _a0
func (_m *EC2API) DescribeVolumeStatus(_a0 *ec2.DescribeVolumeStatusInput) (*ec2.DescribeVolumeStatusOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeVolumeStatusOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVolumeStatusInput) *ec2.DescribeVolumeStatusOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeVolumeStatusOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVolumeStatusInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeVolumeStatusPages provides a mock function with given fields: _a0, _a1
func (_m *EC2API) DescribeVolumeStatusPages(_a0 *ec2.DescribeVolumeStatusInput, _a1 func(*ec2.DescribeVolumeStatusOutput, bool) bool) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVolumeStatusInput, func(*ec2.DescribeVolumeStatusOutput, bool) bool) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DescribeVolumeStatusPagesWithContext provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *EC2API) DescribeVolumeStatusPagesWithContext(_a0 aws.Context, _a1 *ec2.DescribeVolumeStatusInput, _a2 func(*ec2.DescribeVolumeStatusOutput, bool) bool, _a3 ...request.Option) error {
	_va := make([]interface{}, len(_a3))
	for _i := range _a3 {
		_va[_i] = _a3[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1, _a2)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeVolumeStatusInput, func(*ec2.DescribeVolumeStatusOutput, bool) bool, ...request.Option) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DescribeVolumeStatusRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeVolumeStatusRequest(_a0 *ec2.DescribeVolumeStatusInput) (*request.Request, *ec2.DescribeVolumeStatusOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVolumeStatusInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeVolumeStatusOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVolumeStatusInput) *ec2.DescribeVolumeStatusOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeVolumeStatusOutput)
		}
	}

	return r0, r1
}

// DescribeVolumeStatusWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeVolumeStatusWithContext(_a0 aws.Context, _a1 *ec2.DescribeVolumeStatusInput, _a2 ...request.Option) (*ec2.DescribeVolumeStatusOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeVolumeStatusOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeVolumeStatusInput, ...request.Option) *ec2.DescribeVolumeStatusOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeVolumeStatusOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeVolumeStatusInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeVolumes provides a mock function with given fields: _a0
func (_m *EC2API) DescribeVolumes(_a0 *ec2.DescribeVolumesInput) (*ec2.DescribeVolumesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeVolumesOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVolumesInput) *ec2.DescribeVolumesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeVolumesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVolumesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeVolumesModifications provides a mock function with given fields: _a0
func (_m *EC2API) DescribeVolumesModifications(_a0 *ec2.DescribeVolumesModificationsInput) (*ec2.DescribeVolumesModificationsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeVolumesModificationsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVolumesModificationsInput) *ec2.DescribeVolumesModificationsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeVolumesModificationsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVolumesModificationsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeVolumesModificationsRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeVolumesModificationsRequest(_a0 *ec2.DescribeVolumesModificationsInput) (*request.Request, *ec2.DescribeVolumesModificationsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVolumesModificationsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeVolumesModificationsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVolumesModificationsInput) *ec2.DescribeVolumesModificationsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeVolumesModificationsOutput)
		}
	}

	return r0, r1
}

// DescribeVolumesModificationsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeVolumesModificationsWithContext(_a0 aws.Context, _a1 *ec2.DescribeVolumesModificationsInput, _a2 ...request.Option) (*ec2.DescribeVolumesModificationsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeVolumesModificationsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeVolumesModificationsInput, ...request.Option) *ec2.DescribeVolumesModificationsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeVolumesModificationsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeVolumesModificationsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeVolumesPages provides a mock function with given fields: _a0, _a1
func (_m *EC2API) DescribeVolumesPages(_a0 *ec2.DescribeVolumesInput, _a1 func(*ec2.DescribeVolumesOutput, bool) bool) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVolumesInput, func(*ec2.DescribeVolumesOutput, bool) bool) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DescribeVolumesPagesWithContext provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *EC2API) DescribeVolumesPagesWithContext(_a0 aws.Context, _a1 *ec2.DescribeVolumesInput, _a2 func(*ec2.DescribeVolumesOutput, bool) bool, _a3 ...request.Option) error {
	_va := make([]interface{}, len(_a3))
	for _i := range _a3 {
		_va[_i] = _a3[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1, _a2)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeVolumesInput, func(*ec2.DescribeVolumesOutput, bool) bool, ...request.Option) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DescribeVolumesRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeVolumesRequest(_a0 *ec2.DescribeVolumesInput) (*request.Request, *ec2.DescribeVolumesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVolumesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeVolumesOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVolumesInput) *ec2.DescribeVolumesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeVolumesOutput)
		}
	}

	return r0, r1
}

// DescribeVolumesWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeVolumesWithContext(_a0 aws.Context, _a1 *ec2.DescribeVolumesInput, _a2 ...request.Option) (*ec2.DescribeVolumesOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeVolumesOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeVolumesInput, ...request.Option) *ec2.DescribeVolumesOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeVolumesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeVolumesInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeVpcAttribute provides a mock function with given fields: _a0
func (_m *EC2API) DescribeVpcAttribute(_a0 *ec2.DescribeVpcAttributeInput) (*ec2.DescribeVpcAttributeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeVpcAttributeOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpcAttributeInput) *ec2.DescribeVpcAttributeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeVpcAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVpcAttributeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeVpcAttributeRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeVpcAttributeRequest(_a0 *ec2.DescribeVpcAttributeInput) (*request.Request, *ec2.DescribeVpcAttributeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpcAttributeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeVpcAttributeOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVpcAttributeInput) *ec2.DescribeVpcAttributeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeVpcAttributeOutput)
		}
	}

	return r0, r1
}

// DescribeVpcAttributeWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeVpcAttributeWithContext(_a0 aws.Context, _a1 *ec2.DescribeVpcAttributeInput, _a2 ...request.Option) (*ec2.DescribeVpcAttributeOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeVpcAttributeOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeVpcAttributeInput, ...request.Option) *ec2.DescribeVpcAttributeOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeVpcAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeVpcAttributeInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeVpcClassicLink provides a mock function with given fields: _a0
func (_m *EC2API) DescribeVpcClassicLink(_a0 *ec2.DescribeVpcClassicLinkInput) (*ec2.DescribeVpcClassicLinkOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeVpcClassicLinkOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpcClassicLinkInput) *ec2.DescribeVpcClassicLinkOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeVpcClassicLinkOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVpcClassicLinkInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeVpcClassicLinkDnsSupport provides a mock function with given fields: _a0
func (_m *EC2API) DescribeVpcClassicLinkDnsSupport(_a0 *ec2.DescribeVpcClassicLinkDnsSupportInput) (*ec2.DescribeVpcClassicLinkDnsSupportOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeVpcClassicLinkDnsSupportOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpcClassicLinkDnsSupportInput) *ec2.DescribeVpcClassicLinkDnsSupportOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeVpcClassicLinkDnsSupportOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVpcClassicLinkDnsSupportInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeVpcClassicLinkDnsSupportRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeVpcClassicLinkDnsSupportRequest(_a0 *ec2.DescribeVpcClassicLinkDnsSupportInput) (*request.Request, *ec2.DescribeVpcClassicLinkDnsSupportOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpcClassicLinkDnsSupportInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeVpcClassicLinkDnsSupportOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVpcClassicLinkDnsSupportInput) *ec2.DescribeVpcClassicLinkDnsSupportOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeVpcClassicLinkDnsSupportOutput)
		}
	}

	return r0, r1
}

// DescribeVpcClassicLinkDnsSupportWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeVpcClassicLinkDnsSupportWithContext(_a0 aws.Context, _a1 *ec2.DescribeVpcClassicLinkDnsSupportInput, _a2 ...request.Option) (*ec2.DescribeVpcClassicLinkDnsSupportOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeVpcClassicLinkDnsSupportOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeVpcClassicLinkDnsSupportInput, ...request.Option) *ec2.DescribeVpcClassicLinkDnsSupportOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeVpcClassicLinkDnsSupportOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeVpcClassicLinkDnsSupportInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeVpcClassicLinkRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeVpcClassicLinkRequest(_a0 *ec2.DescribeVpcClassicLinkInput) (*request.Request, *ec2.DescribeVpcClassicLinkOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpcClassicLinkInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeVpcClassicLinkOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVpcClassicLinkInput) *ec2.DescribeVpcClassicLinkOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeVpcClassicLinkOutput)
		}
	}

	return r0, r1
}

// DescribeVpcClassicLinkWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeVpcClassicLinkWithContext(_a0 aws.Context, _a1 *ec2.DescribeVpcClassicLinkInput, _a2 ...request.Option) (*ec2.DescribeVpcClassicLinkOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeVpcClassicLinkOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeVpcClassicLinkInput, ...request.Option) *ec2.DescribeVpcClassicLinkOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeVpcClassicLinkOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeVpcClassicLinkInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeVpcEndpointConnectionNotifications provides a mock function with given fields: _a0
func (_m *EC2API) DescribeVpcEndpointConnectionNotifications(_a0 *ec2.DescribeVpcEndpointConnectionNotificationsInput) (*ec2.DescribeVpcEndpointConnectionNotificationsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeVpcEndpointConnectionNotificationsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpcEndpointConnectionNotificationsInput) *ec2.DescribeVpcEndpointConnectionNotificationsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeVpcEndpointConnectionNotificationsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVpcEndpointConnectionNotificationsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeVpcEndpointConnectionNotificationsRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeVpcEndpointConnectionNotificationsRequest(_a0 *ec2.DescribeVpcEndpointConnectionNotificationsInput) (*request.Request, *ec2.DescribeVpcEndpointConnectionNotificationsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpcEndpointConnectionNotificationsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeVpcEndpointConnectionNotificationsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVpcEndpointConnectionNotificationsInput) *ec2.DescribeVpcEndpointConnectionNotificationsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeVpcEndpointConnectionNotificationsOutput)
		}
	}

	return r0, r1
}

// DescribeVpcEndpointConnectionNotificationsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeVpcEndpointConnectionNotificationsWithContext(_a0 aws.Context, _a1 *ec2.DescribeVpcEndpointConnectionNotificationsInput, _a2 ...request.Option) (*ec2.DescribeVpcEndpointConnectionNotificationsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeVpcEndpointConnectionNotificationsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeVpcEndpointConnectionNotificationsInput, ...request.Option) *ec2.DescribeVpcEndpointConnectionNotificationsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeVpcEndpointConnectionNotificationsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeVpcEndpointConnectionNotificationsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeVpcEndpointConnections provides a mock function with given fields: _a0
func (_m *EC2API) DescribeVpcEndpointConnections(_a0 *ec2.DescribeVpcEndpointConnectionsInput) (*ec2.DescribeVpcEndpointConnectionsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeVpcEndpointConnectionsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpcEndpointConnectionsInput) *ec2.DescribeVpcEndpointConnectionsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeVpcEndpointConnectionsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVpcEndpointConnectionsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeVpcEndpointConnectionsRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeVpcEndpointConnectionsRequest(_a0 *ec2.DescribeVpcEndpointConnectionsInput) (*request.Request, *ec2.DescribeVpcEndpointConnectionsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpcEndpointConnectionsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeVpcEndpointConnectionsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVpcEndpointConnectionsInput) *ec2.DescribeVpcEndpointConnectionsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeVpcEndpointConnectionsOutput)
		}
	}

	return r0, r1
}

// DescribeVpcEndpointConnectionsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeVpcEndpointConnectionsWithContext(_a0 aws.Context, _a1 *ec2.DescribeVpcEndpointConnectionsInput, _a2 ...request.Option) (*ec2.DescribeVpcEndpointConnectionsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeVpcEndpointConnectionsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeVpcEndpointConnectionsInput, ...request.Option) *ec2.DescribeVpcEndpointConnectionsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeVpcEndpointConnectionsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeVpcEndpointConnectionsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeVpcEndpointServiceConfigurations provides a mock function with given fields: _a0
func (_m *EC2API) DescribeVpcEndpointServiceConfigurations(_a0 *ec2.DescribeVpcEndpointServiceConfigurationsInput) (*ec2.DescribeVpcEndpointServiceConfigurationsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeVpcEndpointServiceConfigurationsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpcEndpointServiceConfigurationsInput) *ec2.DescribeVpcEndpointServiceConfigurationsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeVpcEndpointServiceConfigurationsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVpcEndpointServiceConfigurationsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeVpcEndpointServiceConfigurationsRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeVpcEndpointServiceConfigurationsRequest(_a0 *ec2.DescribeVpcEndpointServiceConfigurationsInput) (*request.Request, *ec2.DescribeVpcEndpointServiceConfigurationsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpcEndpointServiceConfigurationsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeVpcEndpointServiceConfigurationsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVpcEndpointServiceConfigurationsInput) *ec2.DescribeVpcEndpointServiceConfigurationsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeVpcEndpointServiceConfigurationsOutput)
		}
	}

	return r0, r1
}

// DescribeVpcEndpointServiceConfigurationsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeVpcEndpointServiceConfigurationsWithContext(_a0 aws.Context, _a1 *ec2.DescribeVpcEndpointServiceConfigurationsInput, _a2 ...request.Option) (*ec2.DescribeVpcEndpointServiceConfigurationsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeVpcEndpointServiceConfigurationsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeVpcEndpointServiceConfigurationsInput, ...request.Option) *ec2.DescribeVpcEndpointServiceConfigurationsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeVpcEndpointServiceConfigurationsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeVpcEndpointServiceConfigurationsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeVpcEndpointServicePermissions provides a mock function with given fields: _a0
func (_m *EC2API) DescribeVpcEndpointServicePermissions(_a0 *ec2.DescribeVpcEndpointServicePermissionsInput) (*ec2.DescribeVpcEndpointServicePermissionsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeVpcEndpointServicePermissionsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpcEndpointServicePermissionsInput) *ec2.DescribeVpcEndpointServicePermissionsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeVpcEndpointServicePermissionsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVpcEndpointServicePermissionsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeVpcEndpointServicePermissionsRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeVpcEndpointServicePermissionsRequest(_a0 *ec2.DescribeVpcEndpointServicePermissionsInput) (*request.Request, *ec2.DescribeVpcEndpointServicePermissionsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpcEndpointServicePermissionsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeVpcEndpointServicePermissionsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVpcEndpointServicePermissionsInput) *ec2.DescribeVpcEndpointServicePermissionsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeVpcEndpointServicePermissionsOutput)
		}
	}

	return r0, r1
}

// DescribeVpcEndpointServicePermissionsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeVpcEndpointServicePermissionsWithContext(_a0 aws.Context, _a1 *ec2.DescribeVpcEndpointServicePermissionsInput, _a2 ...request.Option) (*ec2.DescribeVpcEndpointServicePermissionsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeVpcEndpointServicePermissionsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeVpcEndpointServicePermissionsInput, ...request.Option) *ec2.DescribeVpcEndpointServicePermissionsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeVpcEndpointServicePermissionsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeVpcEndpointServicePermissionsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeVpcEndpointServices provides a mock function with given fields: _a0
func (_m *EC2API) DescribeVpcEndpointServices(_a0 *ec2.DescribeVpcEndpointServicesInput) (*ec2.DescribeVpcEndpointServicesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeVpcEndpointServicesOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpcEndpointServicesInput) *ec2.DescribeVpcEndpointServicesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeVpcEndpointServicesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVpcEndpointServicesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeVpcEndpointServicesRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeVpcEndpointServicesRequest(_a0 *ec2.DescribeVpcEndpointServicesInput) (*request.Request, *ec2.DescribeVpcEndpointServicesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpcEndpointServicesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeVpcEndpointServicesOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVpcEndpointServicesInput) *ec2.DescribeVpcEndpointServicesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeVpcEndpointServicesOutput)
		}
	}

	return r0, r1
}

// DescribeVpcEndpointServicesWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeVpcEndpointServicesWithContext(_a0 aws.Context, _a1 *ec2.DescribeVpcEndpointServicesInput, _a2 ...request.Option) (*ec2.DescribeVpcEndpointServicesOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeVpcEndpointServicesOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeVpcEndpointServicesInput, ...request.Option) *ec2.DescribeVpcEndpointServicesOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeVpcEndpointServicesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeVpcEndpointServicesInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeVpcEndpoints provides a mock function with given fields: _a0
func (_m *EC2API) DescribeVpcEndpoints(_a0 *ec2.DescribeVpcEndpointsInput) (*ec2.DescribeVpcEndpointsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeVpcEndpointsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpcEndpointsInput) *ec2.DescribeVpcEndpointsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeVpcEndpointsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVpcEndpointsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeVpcEndpointsRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeVpcEndpointsRequest(_a0 *ec2.DescribeVpcEndpointsInput) (*request.Request, *ec2.DescribeVpcEndpointsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpcEndpointsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeVpcEndpointsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVpcEndpointsInput) *ec2.DescribeVpcEndpointsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeVpcEndpointsOutput)
		}
	}

	return r0, r1
}

// DescribeVpcEndpointsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeVpcEndpointsWithContext(_a0 aws.Context, _a1 *ec2.DescribeVpcEndpointsInput, _a2 ...request.Option) (*ec2.DescribeVpcEndpointsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeVpcEndpointsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeVpcEndpointsInput, ...request.Option) *ec2.DescribeVpcEndpointsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeVpcEndpointsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeVpcEndpointsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeVpcPeeringConnections provides a mock function with given fields: _a0
func (_m *EC2API) DescribeVpcPeeringConnections(_a0 *ec2.DescribeVpcPeeringConnectionsInput) (*ec2.DescribeVpcPeeringConnectionsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeVpcPeeringConnectionsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpcPeeringConnectionsInput) *ec2.DescribeVpcPeeringConnectionsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeVpcPeeringConnectionsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVpcPeeringConnectionsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeVpcPeeringConnectionsRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeVpcPeeringConnectionsRequest(_a0 *ec2.DescribeVpcPeeringConnectionsInput) (*request.Request, *ec2.DescribeVpcPeeringConnectionsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpcPeeringConnectionsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeVpcPeeringConnectionsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVpcPeeringConnectionsInput) *ec2.DescribeVpcPeeringConnectionsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeVpcPeeringConnectionsOutput)
		}
	}

	return r0, r1
}

// DescribeVpcPeeringConnectionsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeVpcPeeringConnectionsWithContext(_a0 aws.Context, _a1 *ec2.DescribeVpcPeeringConnectionsInput, _a2 ...request.Option) (*ec2.DescribeVpcPeeringConnectionsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeVpcPeeringConnectionsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeVpcPeeringConnectionsInput, ...request.Option) *ec2.DescribeVpcPeeringConnectionsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeVpcPeeringConnectionsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeVpcPeeringConnectionsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}



// DescribeVpcsRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeVpcsRequest(_a0 *ec2.DescribeVpcsInput) (*request.Request, *ec2.DescribeVpcsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpcsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeVpcsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVpcsInput) *ec2.DescribeVpcsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeVpcsOutput)
		}
	}

	return r0, r1
}

// DescribeVpcsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeVpcsWithContext(_a0 aws.Context, _a1 *ec2.DescribeVpcsInput, _a2 ...request.Option) (*ec2.DescribeVpcsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeVpcsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeVpcsInput, ...request.Option) *ec2.DescribeVpcsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeVpcsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeVpcsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeVpnConnections provides a mock function with given fields: _a0
func (_m *EC2API) DescribeVpnConnections(_a0 *ec2.DescribeVpnConnectionsInput) (*ec2.DescribeVpnConnectionsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeVpnConnectionsOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpnConnectionsInput) *ec2.DescribeVpnConnectionsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeVpnConnectionsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVpnConnectionsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeVpnConnectionsRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeVpnConnectionsRequest(_a0 *ec2.DescribeVpnConnectionsInput) (*request.Request, *ec2.DescribeVpnConnectionsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpnConnectionsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeVpnConnectionsOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVpnConnectionsInput) *ec2.DescribeVpnConnectionsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeVpnConnectionsOutput)
		}
	}

	return r0, r1
}

// DescribeVpnConnectionsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeVpnConnectionsWithContext(_a0 aws.Context, _a1 *ec2.DescribeVpnConnectionsInput, _a2 ...request.Option) (*ec2.DescribeVpnConnectionsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeVpnConnectionsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeVpnConnectionsInput, ...request.Option) *ec2.DescribeVpnConnectionsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeVpnConnectionsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeVpnConnectionsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeVpnGateways provides a mock function with given fields: _a0
func (_m *EC2API) DescribeVpnGateways(_a0 *ec2.DescribeVpnGatewaysInput) (*ec2.DescribeVpnGatewaysOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DescribeVpnGatewaysOutput
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpnGatewaysInput) *ec2.DescribeVpnGatewaysOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeVpnGatewaysOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVpnGatewaysInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeVpnGatewaysRequest provides a mock function with given fields: _a0
func (_m *EC2API) DescribeVpnGatewaysRequest(_a0 *ec2.DescribeVpnGatewaysInput) (*request.Request, *ec2.DescribeVpnGatewaysOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpnGatewaysInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DescribeVpnGatewaysOutput
	if rf, ok := ret.Get(1).(func(*ec2.DescribeVpnGatewaysInput) *ec2.DescribeVpnGatewaysOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DescribeVpnGatewaysOutput)
		}
	}

	return r0, r1
}

// DescribeVpnGatewaysWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DescribeVpnGatewaysWithContext(_a0 aws.Context, _a1 *ec2.DescribeVpnGatewaysInput, _a2 ...request.Option) (*ec2.DescribeVpnGatewaysOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeVpnGatewaysOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeVpnGatewaysInput, ...request.Option) *ec2.DescribeVpnGatewaysOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeVpnGatewaysOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DescribeVpnGatewaysInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DetachClassicLinkVpc provides a mock function with given fields: _a0
func (_m *EC2API) DetachClassicLinkVpc(_a0 *ec2.DetachClassicLinkVpcInput) (*ec2.DetachClassicLinkVpcOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DetachClassicLinkVpcOutput
	if rf, ok := ret.Get(0).(func(*ec2.DetachClassicLinkVpcInput) *ec2.DetachClassicLinkVpcOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DetachClassicLinkVpcOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DetachClassicLinkVpcInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DetachClassicLinkVpcRequest provides a mock function with given fields: _a0
func (_m *EC2API) DetachClassicLinkVpcRequest(_a0 *ec2.DetachClassicLinkVpcInput) (*request.Request, *ec2.DetachClassicLinkVpcOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DetachClassicLinkVpcInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DetachClassicLinkVpcOutput
	if rf, ok := ret.Get(1).(func(*ec2.DetachClassicLinkVpcInput) *ec2.DetachClassicLinkVpcOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DetachClassicLinkVpcOutput)
		}
	}

	return r0, r1
}

// DetachClassicLinkVpcWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DetachClassicLinkVpcWithContext(_a0 aws.Context, _a1 *ec2.DetachClassicLinkVpcInput, _a2 ...request.Option) (*ec2.DetachClassicLinkVpcOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DetachClassicLinkVpcOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DetachClassicLinkVpcInput, ...request.Option) *ec2.DetachClassicLinkVpcOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DetachClassicLinkVpcOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DetachClassicLinkVpcInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DetachInternetGateway provides a mock function with given fields: _a0
func (_m *EC2API) DetachInternetGateway(_a0 *ec2.DetachInternetGatewayInput) (*ec2.DetachInternetGatewayOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DetachInternetGatewayOutput
	if rf, ok := ret.Get(0).(func(*ec2.DetachInternetGatewayInput) *ec2.DetachInternetGatewayOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DetachInternetGatewayOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DetachInternetGatewayInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DetachInternetGatewayRequest provides a mock function with given fields: _a0
func (_m *EC2API) DetachInternetGatewayRequest(_a0 *ec2.DetachInternetGatewayInput) (*request.Request, *ec2.DetachInternetGatewayOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DetachInternetGatewayInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DetachInternetGatewayOutput
	if rf, ok := ret.Get(1).(func(*ec2.DetachInternetGatewayInput) *ec2.DetachInternetGatewayOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DetachInternetGatewayOutput)
		}
	}

	return r0, r1
}

// DetachInternetGatewayWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DetachInternetGatewayWithContext(_a0 aws.Context, _a1 *ec2.DetachInternetGatewayInput, _a2 ...request.Option) (*ec2.DetachInternetGatewayOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DetachInternetGatewayOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DetachInternetGatewayInput, ...request.Option) *ec2.DetachInternetGatewayOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DetachInternetGatewayOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DetachInternetGatewayInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DetachNetworkInterface provides a mock function with given fields: _a0
func (_m *EC2API) DetachNetworkInterface(_a0 *ec2.DetachNetworkInterfaceInput) (*ec2.DetachNetworkInterfaceOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DetachNetworkInterfaceOutput
	if rf, ok := ret.Get(0).(func(*ec2.DetachNetworkInterfaceInput) *ec2.DetachNetworkInterfaceOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DetachNetworkInterfaceOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DetachNetworkInterfaceInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DetachNetworkInterfaceRequest provides a mock function with given fields: _a0
func (_m *EC2API) DetachNetworkInterfaceRequest(_a0 *ec2.DetachNetworkInterfaceInput) (*request.Request, *ec2.DetachNetworkInterfaceOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DetachNetworkInterfaceInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DetachNetworkInterfaceOutput
	if rf, ok := ret.Get(1).(func(*ec2.DetachNetworkInterfaceInput) *ec2.DetachNetworkInterfaceOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DetachNetworkInterfaceOutput)
		}
	}

	return r0, r1
}

// DetachNetworkInterfaceWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DetachNetworkInterfaceWithContext(_a0 aws.Context, _a1 *ec2.DetachNetworkInterfaceInput, _a2 ...request.Option) (*ec2.DetachNetworkInterfaceOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DetachNetworkInterfaceOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DetachNetworkInterfaceInput, ...request.Option) *ec2.DetachNetworkInterfaceOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DetachNetworkInterfaceOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DetachNetworkInterfaceInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DetachVolume provides a mock function with given fields: _a0
func (_m *EC2API) DetachVolume(_a0 *ec2.DetachVolumeInput) (*ec2.VolumeAttachment, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.VolumeAttachment
	if rf, ok := ret.Get(0).(func(*ec2.DetachVolumeInput) *ec2.VolumeAttachment); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.VolumeAttachment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DetachVolumeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DetachVolumeRequest provides a mock function with given fields: _a0
func (_m *EC2API) DetachVolumeRequest(_a0 *ec2.DetachVolumeInput) (*request.Request, *ec2.VolumeAttachment) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DetachVolumeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.VolumeAttachment
	if rf, ok := ret.Get(1).(func(*ec2.DetachVolumeInput) *ec2.VolumeAttachment); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.VolumeAttachment)
		}
	}

	return r0, r1
}

// DetachVolumeWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DetachVolumeWithContext(_a0 aws.Context, _a1 *ec2.DetachVolumeInput, _a2 ...request.Option) (*ec2.VolumeAttachment, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.VolumeAttachment
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DetachVolumeInput, ...request.Option) *ec2.VolumeAttachment); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.VolumeAttachment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DetachVolumeInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DetachVpnGateway provides a mock function with given fields: _a0
func (_m *EC2API) DetachVpnGateway(_a0 *ec2.DetachVpnGatewayInput) (*ec2.DetachVpnGatewayOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DetachVpnGatewayOutput
	if rf, ok := ret.Get(0).(func(*ec2.DetachVpnGatewayInput) *ec2.DetachVpnGatewayOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DetachVpnGatewayOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DetachVpnGatewayInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DetachVpnGatewayRequest provides a mock function with given fields: _a0
func (_m *EC2API) DetachVpnGatewayRequest(_a0 *ec2.DetachVpnGatewayInput) (*request.Request, *ec2.DetachVpnGatewayOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DetachVpnGatewayInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DetachVpnGatewayOutput
	if rf, ok := ret.Get(1).(func(*ec2.DetachVpnGatewayInput) *ec2.DetachVpnGatewayOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DetachVpnGatewayOutput)
		}
	}

	return r0, r1
}

// DetachVpnGatewayWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DetachVpnGatewayWithContext(_a0 aws.Context, _a1 *ec2.DetachVpnGatewayInput, _a2 ...request.Option) (*ec2.DetachVpnGatewayOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DetachVpnGatewayOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DetachVpnGatewayInput, ...request.Option) *ec2.DetachVpnGatewayOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DetachVpnGatewayOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DetachVpnGatewayInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisableVgwRoutePropagation provides a mock function with given fields: _a0
func (_m *EC2API) DisableVgwRoutePropagation(_a0 *ec2.DisableVgwRoutePropagationInput) (*ec2.DisableVgwRoutePropagationOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DisableVgwRoutePropagationOutput
	if rf, ok := ret.Get(0).(func(*ec2.DisableVgwRoutePropagationInput) *ec2.DisableVgwRoutePropagationOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DisableVgwRoutePropagationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DisableVgwRoutePropagationInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisableVgwRoutePropagationRequest provides a mock function with given fields: _a0
func (_m *EC2API) DisableVgwRoutePropagationRequest(_a0 *ec2.DisableVgwRoutePropagationInput) (*request.Request, *ec2.DisableVgwRoutePropagationOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DisableVgwRoutePropagationInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DisableVgwRoutePropagationOutput
	if rf, ok := ret.Get(1).(func(*ec2.DisableVgwRoutePropagationInput) *ec2.DisableVgwRoutePropagationOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DisableVgwRoutePropagationOutput)
		}
	}

	return r0, r1
}

// DisableVgwRoutePropagationWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DisableVgwRoutePropagationWithContext(_a0 aws.Context, _a1 *ec2.DisableVgwRoutePropagationInput, _a2 ...request.Option) (*ec2.DisableVgwRoutePropagationOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DisableVgwRoutePropagationOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DisableVgwRoutePropagationInput, ...request.Option) *ec2.DisableVgwRoutePropagationOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DisableVgwRoutePropagationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DisableVgwRoutePropagationInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisableVpcClassicLink provides a mock function with given fields: _a0
func (_m *EC2API) DisableVpcClassicLink(_a0 *ec2.DisableVpcClassicLinkInput) (*ec2.DisableVpcClassicLinkOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DisableVpcClassicLinkOutput
	if rf, ok := ret.Get(0).(func(*ec2.DisableVpcClassicLinkInput) *ec2.DisableVpcClassicLinkOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DisableVpcClassicLinkOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DisableVpcClassicLinkInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisableVpcClassicLinkDnsSupport provides a mock function with given fields: _a0
func (_m *EC2API) DisableVpcClassicLinkDnsSupport(_a0 *ec2.DisableVpcClassicLinkDnsSupportInput) (*ec2.DisableVpcClassicLinkDnsSupportOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DisableVpcClassicLinkDnsSupportOutput
	if rf, ok := ret.Get(0).(func(*ec2.DisableVpcClassicLinkDnsSupportInput) *ec2.DisableVpcClassicLinkDnsSupportOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DisableVpcClassicLinkDnsSupportOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DisableVpcClassicLinkDnsSupportInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisableVpcClassicLinkDnsSupportRequest provides a mock function with given fields: _a0
func (_m *EC2API) DisableVpcClassicLinkDnsSupportRequest(_a0 *ec2.DisableVpcClassicLinkDnsSupportInput) (*request.Request, *ec2.DisableVpcClassicLinkDnsSupportOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DisableVpcClassicLinkDnsSupportInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DisableVpcClassicLinkDnsSupportOutput
	if rf, ok := ret.Get(1).(func(*ec2.DisableVpcClassicLinkDnsSupportInput) *ec2.DisableVpcClassicLinkDnsSupportOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DisableVpcClassicLinkDnsSupportOutput)
		}
	}

	return r0, r1
}

// DisableVpcClassicLinkDnsSupportWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DisableVpcClassicLinkDnsSupportWithContext(_a0 aws.Context, _a1 *ec2.DisableVpcClassicLinkDnsSupportInput, _a2 ...request.Option) (*ec2.DisableVpcClassicLinkDnsSupportOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DisableVpcClassicLinkDnsSupportOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DisableVpcClassicLinkDnsSupportInput, ...request.Option) *ec2.DisableVpcClassicLinkDnsSupportOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DisableVpcClassicLinkDnsSupportOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DisableVpcClassicLinkDnsSupportInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisableVpcClassicLinkRequest provides a mock function with given fields: _a0
func (_m *EC2API) DisableVpcClassicLinkRequest(_a0 *ec2.DisableVpcClassicLinkInput) (*request.Request, *ec2.DisableVpcClassicLinkOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DisableVpcClassicLinkInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DisableVpcClassicLinkOutput
	if rf, ok := ret.Get(1).(func(*ec2.DisableVpcClassicLinkInput) *ec2.DisableVpcClassicLinkOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DisableVpcClassicLinkOutput)
		}
	}

	return r0, r1
}

// DisableVpcClassicLinkWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DisableVpcClassicLinkWithContext(_a0 aws.Context, _a1 *ec2.DisableVpcClassicLinkInput, _a2 ...request.Option) (*ec2.DisableVpcClassicLinkOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DisableVpcClassicLinkOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DisableVpcClassicLinkInput, ...request.Option) *ec2.DisableVpcClassicLinkOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DisableVpcClassicLinkOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DisableVpcClassicLinkInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}



// DisassociateAddressRequest provides a mock function with given fields: _a0
func (_m *EC2API) DisassociateAddressRequest(_a0 *ec2.DisassociateAddressInput) (*request.Request, *ec2.DisassociateAddressOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DisassociateAddressInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DisassociateAddressOutput
	if rf, ok := ret.Get(1).(func(*ec2.DisassociateAddressInput) *ec2.DisassociateAddressOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DisassociateAddressOutput)
		}
	}

	return r0, r1
}

// DisassociateAddressWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DisassociateAddressWithContext(_a0 aws.Context, _a1 *ec2.DisassociateAddressInput, _a2 ...request.Option) (*ec2.DisassociateAddressOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DisassociateAddressOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DisassociateAddressInput, ...request.Option) *ec2.DisassociateAddressOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DisassociateAddressOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DisassociateAddressInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisassociateIamInstanceProfile provides a mock function with given fields: _a0
func (_m *EC2API) DisassociateIamInstanceProfile(_a0 *ec2.DisassociateIamInstanceProfileInput) (*ec2.DisassociateIamInstanceProfileOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DisassociateIamInstanceProfileOutput
	if rf, ok := ret.Get(0).(func(*ec2.DisassociateIamInstanceProfileInput) *ec2.DisassociateIamInstanceProfileOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DisassociateIamInstanceProfileOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DisassociateIamInstanceProfileInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisassociateIamInstanceProfileRequest provides a mock function with given fields: _a0
func (_m *EC2API) DisassociateIamInstanceProfileRequest(_a0 *ec2.DisassociateIamInstanceProfileInput) (*request.Request, *ec2.DisassociateIamInstanceProfileOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DisassociateIamInstanceProfileInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DisassociateIamInstanceProfileOutput
	if rf, ok := ret.Get(1).(func(*ec2.DisassociateIamInstanceProfileInput) *ec2.DisassociateIamInstanceProfileOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DisassociateIamInstanceProfileOutput)
		}
	}

	return r0, r1
}

// DisassociateIamInstanceProfileWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DisassociateIamInstanceProfileWithContext(_a0 aws.Context, _a1 *ec2.DisassociateIamInstanceProfileInput, _a2 ...request.Option) (*ec2.DisassociateIamInstanceProfileOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DisassociateIamInstanceProfileOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DisassociateIamInstanceProfileInput, ...request.Option) *ec2.DisassociateIamInstanceProfileOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DisassociateIamInstanceProfileOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DisassociateIamInstanceProfileInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisassociateRouteTable provides a mock function with given fields: _a0
func (_m *EC2API) DisassociateRouteTable(_a0 *ec2.DisassociateRouteTableInput) (*ec2.DisassociateRouteTableOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DisassociateRouteTableOutput
	if rf, ok := ret.Get(0).(func(*ec2.DisassociateRouteTableInput) *ec2.DisassociateRouteTableOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DisassociateRouteTableOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DisassociateRouteTableInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisassociateRouteTableRequest provides a mock function with given fields: _a0
func (_m *EC2API) DisassociateRouteTableRequest(_a0 *ec2.DisassociateRouteTableInput) (*request.Request, *ec2.DisassociateRouteTableOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DisassociateRouteTableInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DisassociateRouteTableOutput
	if rf, ok := ret.Get(1).(func(*ec2.DisassociateRouteTableInput) *ec2.DisassociateRouteTableOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DisassociateRouteTableOutput)
		}
	}

	return r0, r1
}

// DisassociateRouteTableWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DisassociateRouteTableWithContext(_a0 aws.Context, _a1 *ec2.DisassociateRouteTableInput, _a2 ...request.Option) (*ec2.DisassociateRouteTableOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DisassociateRouteTableOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DisassociateRouteTableInput, ...request.Option) *ec2.DisassociateRouteTableOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DisassociateRouteTableOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DisassociateRouteTableInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisassociateSubnetCidrBlock provides a mock function with given fields: _a0
func (_m *EC2API) DisassociateSubnetCidrBlock(_a0 *ec2.DisassociateSubnetCidrBlockInput) (*ec2.DisassociateSubnetCidrBlockOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DisassociateSubnetCidrBlockOutput
	if rf, ok := ret.Get(0).(func(*ec2.DisassociateSubnetCidrBlockInput) *ec2.DisassociateSubnetCidrBlockOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DisassociateSubnetCidrBlockOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DisassociateSubnetCidrBlockInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisassociateSubnetCidrBlockRequest provides a mock function with given fields: _a0
func (_m *EC2API) DisassociateSubnetCidrBlockRequest(_a0 *ec2.DisassociateSubnetCidrBlockInput) (*request.Request, *ec2.DisassociateSubnetCidrBlockOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DisassociateSubnetCidrBlockInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DisassociateSubnetCidrBlockOutput
	if rf, ok := ret.Get(1).(func(*ec2.DisassociateSubnetCidrBlockInput) *ec2.DisassociateSubnetCidrBlockOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DisassociateSubnetCidrBlockOutput)
		}
	}

	return r0, r1
}

// DisassociateSubnetCidrBlockWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DisassociateSubnetCidrBlockWithContext(_a0 aws.Context, _a1 *ec2.DisassociateSubnetCidrBlockInput, _a2 ...request.Option) (*ec2.DisassociateSubnetCidrBlockOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DisassociateSubnetCidrBlockOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DisassociateSubnetCidrBlockInput, ...request.Option) *ec2.DisassociateSubnetCidrBlockOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DisassociateSubnetCidrBlockOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DisassociateSubnetCidrBlockInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisassociateVpcCidrBlock provides a mock function with given fields: _a0
func (_m *EC2API) DisassociateVpcCidrBlock(_a0 *ec2.DisassociateVpcCidrBlockInput) (*ec2.DisassociateVpcCidrBlockOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.DisassociateVpcCidrBlockOutput
	if rf, ok := ret.Get(0).(func(*ec2.DisassociateVpcCidrBlockInput) *ec2.DisassociateVpcCidrBlockOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DisassociateVpcCidrBlockOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.DisassociateVpcCidrBlockInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisassociateVpcCidrBlockRequest provides a mock function with given fields: _a0
func (_m *EC2API) DisassociateVpcCidrBlockRequest(_a0 *ec2.DisassociateVpcCidrBlockInput) (*request.Request, *ec2.DisassociateVpcCidrBlockOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.DisassociateVpcCidrBlockInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.DisassociateVpcCidrBlockOutput
	if rf, ok := ret.Get(1).(func(*ec2.DisassociateVpcCidrBlockInput) *ec2.DisassociateVpcCidrBlockOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.DisassociateVpcCidrBlockOutput)
		}
	}

	return r0, r1
}

// DisassociateVpcCidrBlockWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) DisassociateVpcCidrBlockWithContext(_a0 aws.Context, _a1 *ec2.DisassociateVpcCidrBlockInput, _a2 ...request.Option) (*ec2.DisassociateVpcCidrBlockOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DisassociateVpcCidrBlockOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DisassociateVpcCidrBlockInput, ...request.Option) *ec2.DisassociateVpcCidrBlockOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DisassociateVpcCidrBlockOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.DisassociateVpcCidrBlockInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EnableVgwRoutePropagation provides a mock function with given fields: _a0
func (_m *EC2API) EnableVgwRoutePropagation(_a0 *ec2.EnableVgwRoutePropagationInput) (*ec2.EnableVgwRoutePropagationOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.EnableVgwRoutePropagationOutput
	if rf, ok := ret.Get(0).(func(*ec2.EnableVgwRoutePropagationInput) *ec2.EnableVgwRoutePropagationOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.EnableVgwRoutePropagationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.EnableVgwRoutePropagationInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EnableVgwRoutePropagationRequest provides a mock function with given fields: _a0
func (_m *EC2API) EnableVgwRoutePropagationRequest(_a0 *ec2.EnableVgwRoutePropagationInput) (*request.Request, *ec2.EnableVgwRoutePropagationOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.EnableVgwRoutePropagationInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.EnableVgwRoutePropagationOutput
	if rf, ok := ret.Get(1).(func(*ec2.EnableVgwRoutePropagationInput) *ec2.EnableVgwRoutePropagationOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.EnableVgwRoutePropagationOutput)
		}
	}

	return r0, r1
}

// EnableVgwRoutePropagationWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) EnableVgwRoutePropagationWithContext(_a0 aws.Context, _a1 *ec2.EnableVgwRoutePropagationInput, _a2 ...request.Option) (*ec2.EnableVgwRoutePropagationOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.EnableVgwRoutePropagationOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.EnableVgwRoutePropagationInput, ...request.Option) *ec2.EnableVgwRoutePropagationOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.EnableVgwRoutePropagationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.EnableVgwRoutePropagationInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EnableVolumeIO provides a mock function with given fields: _a0
func (_m *EC2API) EnableVolumeIO(_a0 *ec2.EnableVolumeIOInput) (*ec2.EnableVolumeIOOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.EnableVolumeIOOutput
	if rf, ok := ret.Get(0).(func(*ec2.EnableVolumeIOInput) *ec2.EnableVolumeIOOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.EnableVolumeIOOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.EnableVolumeIOInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EnableVolumeIORequest provides a mock function with given fields: _a0
func (_m *EC2API) EnableVolumeIORequest(_a0 *ec2.EnableVolumeIOInput) (*request.Request, *ec2.EnableVolumeIOOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.EnableVolumeIOInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.EnableVolumeIOOutput
	if rf, ok := ret.Get(1).(func(*ec2.EnableVolumeIOInput) *ec2.EnableVolumeIOOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.EnableVolumeIOOutput)
		}
	}

	return r0, r1
}

// EnableVolumeIOWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) EnableVolumeIOWithContext(_a0 aws.Context, _a1 *ec2.EnableVolumeIOInput, _a2 ...request.Option) (*ec2.EnableVolumeIOOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.EnableVolumeIOOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.EnableVolumeIOInput, ...request.Option) *ec2.EnableVolumeIOOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.EnableVolumeIOOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.EnableVolumeIOInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EnableVpcClassicLink provides a mock function with given fields: _a0
func (_m *EC2API) EnableVpcClassicLink(_a0 *ec2.EnableVpcClassicLinkInput) (*ec2.EnableVpcClassicLinkOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.EnableVpcClassicLinkOutput
	if rf, ok := ret.Get(0).(func(*ec2.EnableVpcClassicLinkInput) *ec2.EnableVpcClassicLinkOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.EnableVpcClassicLinkOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.EnableVpcClassicLinkInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EnableVpcClassicLinkDnsSupport provides a mock function with given fields: _a0
func (_m *EC2API) EnableVpcClassicLinkDnsSupport(_a0 *ec2.EnableVpcClassicLinkDnsSupportInput) (*ec2.EnableVpcClassicLinkDnsSupportOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.EnableVpcClassicLinkDnsSupportOutput
	if rf, ok := ret.Get(0).(func(*ec2.EnableVpcClassicLinkDnsSupportInput) *ec2.EnableVpcClassicLinkDnsSupportOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.EnableVpcClassicLinkDnsSupportOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.EnableVpcClassicLinkDnsSupportInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EnableVpcClassicLinkDnsSupportRequest provides a mock function with given fields: _a0
func (_m *EC2API) EnableVpcClassicLinkDnsSupportRequest(_a0 *ec2.EnableVpcClassicLinkDnsSupportInput) (*request.Request, *ec2.EnableVpcClassicLinkDnsSupportOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.EnableVpcClassicLinkDnsSupportInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.EnableVpcClassicLinkDnsSupportOutput
	if rf, ok := ret.Get(1).(func(*ec2.EnableVpcClassicLinkDnsSupportInput) *ec2.EnableVpcClassicLinkDnsSupportOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.EnableVpcClassicLinkDnsSupportOutput)
		}
	}

	return r0, r1
}

// EnableVpcClassicLinkDnsSupportWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) EnableVpcClassicLinkDnsSupportWithContext(_a0 aws.Context, _a1 *ec2.EnableVpcClassicLinkDnsSupportInput, _a2 ...request.Option) (*ec2.EnableVpcClassicLinkDnsSupportOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.EnableVpcClassicLinkDnsSupportOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.EnableVpcClassicLinkDnsSupportInput, ...request.Option) *ec2.EnableVpcClassicLinkDnsSupportOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.EnableVpcClassicLinkDnsSupportOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.EnableVpcClassicLinkDnsSupportInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EnableVpcClassicLinkRequest provides a mock function with given fields: _a0
func (_m *EC2API) EnableVpcClassicLinkRequest(_a0 *ec2.EnableVpcClassicLinkInput) (*request.Request, *ec2.EnableVpcClassicLinkOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.EnableVpcClassicLinkInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.EnableVpcClassicLinkOutput
	if rf, ok := ret.Get(1).(func(*ec2.EnableVpcClassicLinkInput) *ec2.EnableVpcClassicLinkOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.EnableVpcClassicLinkOutput)
		}
	}

	return r0, r1
}

// EnableVpcClassicLinkWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) EnableVpcClassicLinkWithContext(_a0 aws.Context, _a1 *ec2.EnableVpcClassicLinkInput, _a2 ...request.Option) (*ec2.EnableVpcClassicLinkOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.EnableVpcClassicLinkOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.EnableVpcClassicLinkInput, ...request.Option) *ec2.EnableVpcClassicLinkOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.EnableVpcClassicLinkOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.EnableVpcClassicLinkInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConsoleOutput provides a mock function with given fields: _a0
func (_m *EC2API) GetConsoleOutput(_a0 *ec2.GetConsoleOutputInput) (*ec2.GetConsoleOutputOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.GetConsoleOutputOutput
	if rf, ok := ret.Get(0).(func(*ec2.GetConsoleOutputInput) *ec2.GetConsoleOutputOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.GetConsoleOutputOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.GetConsoleOutputInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConsoleOutputRequest provides a mock function with given fields: _a0
func (_m *EC2API) GetConsoleOutputRequest(_a0 *ec2.GetConsoleOutputInput) (*request.Request, *ec2.GetConsoleOutputOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.GetConsoleOutputInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.GetConsoleOutputOutput
	if rf, ok := ret.Get(1).(func(*ec2.GetConsoleOutputInput) *ec2.GetConsoleOutputOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.GetConsoleOutputOutput)
		}
	}

	return r0, r1
}

// GetConsoleOutputWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) GetConsoleOutputWithContext(_a0 aws.Context, _a1 *ec2.GetConsoleOutputInput, _a2 ...request.Option) (*ec2.GetConsoleOutputOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.GetConsoleOutputOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.GetConsoleOutputInput, ...request.Option) *ec2.GetConsoleOutputOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.GetConsoleOutputOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.GetConsoleOutputInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConsoleScreenshot provides a mock function with given fields: _a0
func (_m *EC2API) GetConsoleScreenshot(_a0 *ec2.GetConsoleScreenshotInput) (*ec2.GetConsoleScreenshotOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.GetConsoleScreenshotOutput
	if rf, ok := ret.Get(0).(func(*ec2.GetConsoleScreenshotInput) *ec2.GetConsoleScreenshotOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.GetConsoleScreenshotOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.GetConsoleScreenshotInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConsoleScreenshotRequest provides a mock function with given fields: _a0
func (_m *EC2API) GetConsoleScreenshotRequest(_a0 *ec2.GetConsoleScreenshotInput) (*request.Request, *ec2.GetConsoleScreenshotOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.GetConsoleScreenshotInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.GetConsoleScreenshotOutput
	if rf, ok := ret.Get(1).(func(*ec2.GetConsoleScreenshotInput) *ec2.GetConsoleScreenshotOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.GetConsoleScreenshotOutput)
		}
	}

	return r0, r1
}

// GetConsoleScreenshotWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) GetConsoleScreenshotWithContext(_a0 aws.Context, _a1 *ec2.GetConsoleScreenshotInput, _a2 ...request.Option) (*ec2.GetConsoleScreenshotOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.GetConsoleScreenshotOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.GetConsoleScreenshotInput, ...request.Option) *ec2.GetConsoleScreenshotOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.GetConsoleScreenshotOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.GetConsoleScreenshotInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetHostReservationPurchasePreview provides a mock function with given fields: _a0
func (_m *EC2API) GetHostReservationPurchasePreview(_a0 *ec2.GetHostReservationPurchasePreviewInput) (*ec2.GetHostReservationPurchasePreviewOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.GetHostReservationPurchasePreviewOutput
	if rf, ok := ret.Get(0).(func(*ec2.GetHostReservationPurchasePreviewInput) *ec2.GetHostReservationPurchasePreviewOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.GetHostReservationPurchasePreviewOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.GetHostReservationPurchasePreviewInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetHostReservationPurchasePreviewRequest provides a mock function with given fields: _a0
func (_m *EC2API) GetHostReservationPurchasePreviewRequest(_a0 *ec2.GetHostReservationPurchasePreviewInput) (*request.Request, *ec2.GetHostReservationPurchasePreviewOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.GetHostReservationPurchasePreviewInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.GetHostReservationPurchasePreviewOutput
	if rf, ok := ret.Get(1).(func(*ec2.GetHostReservationPurchasePreviewInput) *ec2.GetHostReservationPurchasePreviewOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.GetHostReservationPurchasePreviewOutput)
		}
	}

	return r0, r1
}

// GetHostReservationPurchasePreviewWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) GetHostReservationPurchasePreviewWithContext(_a0 aws.Context, _a1 *ec2.GetHostReservationPurchasePreviewInput, _a2 ...request.Option) (*ec2.GetHostReservationPurchasePreviewOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.GetHostReservationPurchasePreviewOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.GetHostReservationPurchasePreviewInput, ...request.Option) *ec2.GetHostReservationPurchasePreviewOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.GetHostReservationPurchasePreviewOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.GetHostReservationPurchasePreviewInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLaunchTemplateData provides a mock function with given fields: _a0
func (_m *EC2API) GetLaunchTemplateData(_a0 *ec2.GetLaunchTemplateDataInput) (*ec2.GetLaunchTemplateDataOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.GetLaunchTemplateDataOutput
	if rf, ok := ret.Get(0).(func(*ec2.GetLaunchTemplateDataInput) *ec2.GetLaunchTemplateDataOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.GetLaunchTemplateDataOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.GetLaunchTemplateDataInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLaunchTemplateDataRequest provides a mock function with given fields: _a0
func (_m *EC2API) GetLaunchTemplateDataRequest(_a0 *ec2.GetLaunchTemplateDataInput) (*request.Request, *ec2.GetLaunchTemplateDataOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.GetLaunchTemplateDataInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.GetLaunchTemplateDataOutput
	if rf, ok := ret.Get(1).(func(*ec2.GetLaunchTemplateDataInput) *ec2.GetLaunchTemplateDataOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.GetLaunchTemplateDataOutput)
		}
	}

	return r0, r1
}

// GetLaunchTemplateDataWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) GetLaunchTemplateDataWithContext(_a0 aws.Context, _a1 *ec2.GetLaunchTemplateDataInput, _a2 ...request.Option) (*ec2.GetLaunchTemplateDataOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.GetLaunchTemplateDataOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.GetLaunchTemplateDataInput, ...request.Option) *ec2.GetLaunchTemplateDataOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.GetLaunchTemplateDataOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.GetLaunchTemplateDataInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPasswordData provides a mock function with given fields: _a0
func (_m *EC2API) GetPasswordData(_a0 *ec2.GetPasswordDataInput) (*ec2.GetPasswordDataOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.GetPasswordDataOutput
	if rf, ok := ret.Get(0).(func(*ec2.GetPasswordDataInput) *ec2.GetPasswordDataOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.GetPasswordDataOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.GetPasswordDataInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPasswordDataRequest provides a mock function with given fields: _a0
func (_m *EC2API) GetPasswordDataRequest(_a0 *ec2.GetPasswordDataInput) (*request.Request, *ec2.GetPasswordDataOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.GetPasswordDataInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.GetPasswordDataOutput
	if rf, ok := ret.Get(1).(func(*ec2.GetPasswordDataInput) *ec2.GetPasswordDataOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.GetPasswordDataOutput)
		}
	}

	return r0, r1
}

// GetPasswordDataWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) GetPasswordDataWithContext(_a0 aws.Context, _a1 *ec2.GetPasswordDataInput, _a2 ...request.Option) (*ec2.GetPasswordDataOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.GetPasswordDataOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.GetPasswordDataInput, ...request.Option) *ec2.GetPasswordDataOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.GetPasswordDataOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.GetPasswordDataInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReservedInstancesExchangeQuote provides a mock function with given fields: _a0
func (_m *EC2API) GetReservedInstancesExchangeQuote(_a0 *ec2.GetReservedInstancesExchangeQuoteInput) (*ec2.GetReservedInstancesExchangeQuoteOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.GetReservedInstancesExchangeQuoteOutput
	if rf, ok := ret.Get(0).(func(*ec2.GetReservedInstancesExchangeQuoteInput) *ec2.GetReservedInstancesExchangeQuoteOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.GetReservedInstancesExchangeQuoteOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.GetReservedInstancesExchangeQuoteInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReservedInstancesExchangeQuoteRequest provides a mock function with given fields: _a0
func (_m *EC2API) GetReservedInstancesExchangeQuoteRequest(_a0 *ec2.GetReservedInstancesExchangeQuoteInput) (*request.Request, *ec2.GetReservedInstancesExchangeQuoteOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.GetReservedInstancesExchangeQuoteInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.GetReservedInstancesExchangeQuoteOutput
	if rf, ok := ret.Get(1).(func(*ec2.GetReservedInstancesExchangeQuoteInput) *ec2.GetReservedInstancesExchangeQuoteOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.GetReservedInstancesExchangeQuoteOutput)
		}
	}

	return r0, r1
}

// GetReservedInstancesExchangeQuoteWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) GetReservedInstancesExchangeQuoteWithContext(_a0 aws.Context, _a1 *ec2.GetReservedInstancesExchangeQuoteInput, _a2 ...request.Option) (*ec2.GetReservedInstancesExchangeQuoteOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.GetReservedInstancesExchangeQuoteOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.GetReservedInstancesExchangeQuoteInput, ...request.Option) *ec2.GetReservedInstancesExchangeQuoteOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.GetReservedInstancesExchangeQuoteOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.GetReservedInstancesExchangeQuoteInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImportImage provides a mock function with given fields: _a0
func (_m *EC2API) ImportImage(_a0 *ec2.ImportImageInput) (*ec2.ImportImageOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ImportImageOutput
	if rf, ok := ret.Get(0).(func(*ec2.ImportImageInput) *ec2.ImportImageOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ImportImageOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ImportImageInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImportImageRequest provides a mock function with given fields: _a0
func (_m *EC2API) ImportImageRequest(_a0 *ec2.ImportImageInput) (*request.Request, *ec2.ImportImageOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ImportImageInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ImportImageOutput
	if rf, ok := ret.Get(1).(func(*ec2.ImportImageInput) *ec2.ImportImageOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ImportImageOutput)
		}
	}

	return r0, r1
}

// ImportImageWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) ImportImageWithContext(_a0 aws.Context, _a1 *ec2.ImportImageInput, _a2 ...request.Option) (*ec2.ImportImageOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.ImportImageOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.ImportImageInput, ...request.Option) *ec2.ImportImageOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ImportImageOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.ImportImageInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImportInstance provides a mock function with given fields: _a0
func (_m *EC2API) ImportInstance(_a0 *ec2.ImportInstanceInput) (*ec2.ImportInstanceOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ImportInstanceOutput
	if rf, ok := ret.Get(0).(func(*ec2.ImportInstanceInput) *ec2.ImportInstanceOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ImportInstanceOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ImportInstanceInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImportInstanceRequest provides a mock function with given fields: _a0
func (_m *EC2API) ImportInstanceRequest(_a0 *ec2.ImportInstanceInput) (*request.Request, *ec2.ImportInstanceOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ImportInstanceInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ImportInstanceOutput
	if rf, ok := ret.Get(1).(func(*ec2.ImportInstanceInput) *ec2.ImportInstanceOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ImportInstanceOutput)
		}
	}

	return r0, r1
}

// ImportInstanceWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) ImportInstanceWithContext(_a0 aws.Context, _a1 *ec2.ImportInstanceInput, _a2 ...request.Option) (*ec2.ImportInstanceOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.ImportInstanceOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.ImportInstanceInput, ...request.Option) *ec2.ImportInstanceOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ImportInstanceOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.ImportInstanceInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImportKeyPair provides a mock function with given fields: _a0
func (_m *EC2API) ImportKeyPair(_a0 *ec2.ImportKeyPairInput) (*ec2.ImportKeyPairOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ImportKeyPairOutput
	if rf, ok := ret.Get(0).(func(*ec2.ImportKeyPairInput) *ec2.ImportKeyPairOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ImportKeyPairOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ImportKeyPairInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImportKeyPairRequest provides a mock function with given fields: _a0
func (_m *EC2API) ImportKeyPairRequest(_a0 *ec2.ImportKeyPairInput) (*request.Request, *ec2.ImportKeyPairOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ImportKeyPairInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ImportKeyPairOutput
	if rf, ok := ret.Get(1).(func(*ec2.ImportKeyPairInput) *ec2.ImportKeyPairOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ImportKeyPairOutput)
		}
	}

	return r0, r1
}

// ImportKeyPairWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) ImportKeyPairWithContext(_a0 aws.Context, _a1 *ec2.ImportKeyPairInput, _a2 ...request.Option) (*ec2.ImportKeyPairOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.ImportKeyPairOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.ImportKeyPairInput, ...request.Option) *ec2.ImportKeyPairOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ImportKeyPairOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.ImportKeyPairInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImportSnapshot provides a mock function with given fields: _a0
func (_m *EC2API) ImportSnapshot(_a0 *ec2.ImportSnapshotInput) (*ec2.ImportSnapshotOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ImportSnapshotOutput
	if rf, ok := ret.Get(0).(func(*ec2.ImportSnapshotInput) *ec2.ImportSnapshotOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ImportSnapshotOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ImportSnapshotInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImportSnapshotRequest provides a mock function with given fields: _a0
func (_m *EC2API) ImportSnapshotRequest(_a0 *ec2.ImportSnapshotInput) (*request.Request, *ec2.ImportSnapshotOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ImportSnapshotInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ImportSnapshotOutput
	if rf, ok := ret.Get(1).(func(*ec2.ImportSnapshotInput) *ec2.ImportSnapshotOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ImportSnapshotOutput)
		}
	}

	return r0, r1
}

// ImportSnapshotWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) ImportSnapshotWithContext(_a0 aws.Context, _a1 *ec2.ImportSnapshotInput, _a2 ...request.Option) (*ec2.ImportSnapshotOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.ImportSnapshotOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.ImportSnapshotInput, ...request.Option) *ec2.ImportSnapshotOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ImportSnapshotOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.ImportSnapshotInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImportVolume provides a mock function with given fields: _a0
func (_m *EC2API) ImportVolume(_a0 *ec2.ImportVolumeInput) (*ec2.ImportVolumeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ImportVolumeOutput
	if rf, ok := ret.Get(0).(func(*ec2.ImportVolumeInput) *ec2.ImportVolumeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ImportVolumeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ImportVolumeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImportVolumeRequest provides a mock function with given fields: _a0
func (_m *EC2API) ImportVolumeRequest(_a0 *ec2.ImportVolumeInput) (*request.Request, *ec2.ImportVolumeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ImportVolumeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ImportVolumeOutput
	if rf, ok := ret.Get(1).(func(*ec2.ImportVolumeInput) *ec2.ImportVolumeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ImportVolumeOutput)
		}
	}

	return r0, r1
}

// ImportVolumeWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) ImportVolumeWithContext(_a0 aws.Context, _a1 *ec2.ImportVolumeInput, _a2 ...request.Option) (*ec2.ImportVolumeOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.ImportVolumeOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.ImportVolumeInput, ...request.Option) *ec2.ImportVolumeOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ImportVolumeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.ImportVolumeInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyFleet provides a mock function with given fields: _a0
func (_m *EC2API) ModifyFleet(_a0 *ec2.ModifyFleetInput) (*ec2.ModifyFleetOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ModifyFleetOutput
	if rf, ok := ret.Get(0).(func(*ec2.ModifyFleetInput) *ec2.ModifyFleetOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyFleetOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ModifyFleetInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyFleetRequest provides a mock function with given fields: _a0
func (_m *EC2API) ModifyFleetRequest(_a0 *ec2.ModifyFleetInput) (*request.Request, *ec2.ModifyFleetOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ModifyFleetInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ModifyFleetOutput
	if rf, ok := ret.Get(1).(func(*ec2.ModifyFleetInput) *ec2.ModifyFleetOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ModifyFleetOutput)
		}
	}

	return r0, r1
}

// ModifyFleetWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) ModifyFleetWithContext(_a0 aws.Context, _a1 *ec2.ModifyFleetInput, _a2 ...request.Option) (*ec2.ModifyFleetOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.ModifyFleetOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.ModifyFleetInput, ...request.Option) *ec2.ModifyFleetOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyFleetOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.ModifyFleetInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyFpgaImageAttribute provides a mock function with given fields: _a0
func (_m *EC2API) ModifyFpgaImageAttribute(_a0 *ec2.ModifyFpgaImageAttributeInput) (*ec2.ModifyFpgaImageAttributeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ModifyFpgaImageAttributeOutput
	if rf, ok := ret.Get(0).(func(*ec2.ModifyFpgaImageAttributeInput) *ec2.ModifyFpgaImageAttributeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyFpgaImageAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ModifyFpgaImageAttributeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyFpgaImageAttributeRequest provides a mock function with given fields: _a0
func (_m *EC2API) ModifyFpgaImageAttributeRequest(_a0 *ec2.ModifyFpgaImageAttributeInput) (*request.Request, *ec2.ModifyFpgaImageAttributeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ModifyFpgaImageAttributeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ModifyFpgaImageAttributeOutput
	if rf, ok := ret.Get(1).(func(*ec2.ModifyFpgaImageAttributeInput) *ec2.ModifyFpgaImageAttributeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ModifyFpgaImageAttributeOutput)
		}
	}

	return r0, r1
}

// ModifyFpgaImageAttributeWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) ModifyFpgaImageAttributeWithContext(_a0 aws.Context, _a1 *ec2.ModifyFpgaImageAttributeInput, _a2 ...request.Option) (*ec2.ModifyFpgaImageAttributeOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.ModifyFpgaImageAttributeOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.ModifyFpgaImageAttributeInput, ...request.Option) *ec2.ModifyFpgaImageAttributeOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyFpgaImageAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.ModifyFpgaImageAttributeInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyHosts provides a mock function with given fields: _a0
func (_m *EC2API) ModifyHosts(_a0 *ec2.ModifyHostsInput) (*ec2.ModifyHostsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ModifyHostsOutput
	if rf, ok := ret.Get(0).(func(*ec2.ModifyHostsInput) *ec2.ModifyHostsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyHostsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ModifyHostsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyHostsRequest provides a mock function with given fields: _a0
func (_m *EC2API) ModifyHostsRequest(_a0 *ec2.ModifyHostsInput) (*request.Request, *ec2.ModifyHostsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ModifyHostsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ModifyHostsOutput
	if rf, ok := ret.Get(1).(func(*ec2.ModifyHostsInput) *ec2.ModifyHostsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ModifyHostsOutput)
		}
	}

	return r0, r1
}

// ModifyHostsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) ModifyHostsWithContext(_a0 aws.Context, _a1 *ec2.ModifyHostsInput, _a2 ...request.Option) (*ec2.ModifyHostsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.ModifyHostsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.ModifyHostsInput, ...request.Option) *ec2.ModifyHostsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyHostsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.ModifyHostsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyIdFormat provides a mock function with given fields: _a0
func (_m *EC2API) ModifyIdFormat(_a0 *ec2.ModifyIdFormatInput) (*ec2.ModifyIdFormatOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ModifyIdFormatOutput
	if rf, ok := ret.Get(0).(func(*ec2.ModifyIdFormatInput) *ec2.ModifyIdFormatOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyIdFormatOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ModifyIdFormatInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyIdFormatRequest provides a mock function with given fields: _a0
func (_m *EC2API) ModifyIdFormatRequest(_a0 *ec2.ModifyIdFormatInput) (*request.Request, *ec2.ModifyIdFormatOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ModifyIdFormatInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ModifyIdFormatOutput
	if rf, ok := ret.Get(1).(func(*ec2.ModifyIdFormatInput) *ec2.ModifyIdFormatOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ModifyIdFormatOutput)
		}
	}

	return r0, r1
}

// ModifyIdFormatWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) ModifyIdFormatWithContext(_a0 aws.Context, _a1 *ec2.ModifyIdFormatInput, _a2 ...request.Option) (*ec2.ModifyIdFormatOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.ModifyIdFormatOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.ModifyIdFormatInput, ...request.Option) *ec2.ModifyIdFormatOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyIdFormatOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.ModifyIdFormatInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyIdentityIdFormat provides a mock function with given fields: _a0
func (_m *EC2API) ModifyIdentityIdFormat(_a0 *ec2.ModifyIdentityIdFormatInput) (*ec2.ModifyIdentityIdFormatOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ModifyIdentityIdFormatOutput
	if rf, ok := ret.Get(0).(func(*ec2.ModifyIdentityIdFormatInput) *ec2.ModifyIdentityIdFormatOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyIdentityIdFormatOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ModifyIdentityIdFormatInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyIdentityIdFormatRequest provides a mock function with given fields: _a0
func (_m *EC2API) ModifyIdentityIdFormatRequest(_a0 *ec2.ModifyIdentityIdFormatInput) (*request.Request, *ec2.ModifyIdentityIdFormatOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ModifyIdentityIdFormatInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ModifyIdentityIdFormatOutput
	if rf, ok := ret.Get(1).(func(*ec2.ModifyIdentityIdFormatInput) *ec2.ModifyIdentityIdFormatOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ModifyIdentityIdFormatOutput)
		}
	}

	return r0, r1
}

// ModifyIdentityIdFormatWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) ModifyIdentityIdFormatWithContext(_a0 aws.Context, _a1 *ec2.ModifyIdentityIdFormatInput, _a2 ...request.Option) (*ec2.ModifyIdentityIdFormatOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.ModifyIdentityIdFormatOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.ModifyIdentityIdFormatInput, ...request.Option) *ec2.ModifyIdentityIdFormatOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyIdentityIdFormatOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.ModifyIdentityIdFormatInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyImageAttribute provides a mock function with given fields: _a0
func (_m *EC2API) ModifyImageAttribute(_a0 *ec2.ModifyImageAttributeInput) (*ec2.ModifyImageAttributeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ModifyImageAttributeOutput
	if rf, ok := ret.Get(0).(func(*ec2.ModifyImageAttributeInput) *ec2.ModifyImageAttributeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyImageAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ModifyImageAttributeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyImageAttributeRequest provides a mock function with given fields: _a0
func (_m *EC2API) ModifyImageAttributeRequest(_a0 *ec2.ModifyImageAttributeInput) (*request.Request, *ec2.ModifyImageAttributeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ModifyImageAttributeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ModifyImageAttributeOutput
	if rf, ok := ret.Get(1).(func(*ec2.ModifyImageAttributeInput) *ec2.ModifyImageAttributeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ModifyImageAttributeOutput)
		}
	}

	return r0, r1
}

// ModifyImageAttributeWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) ModifyImageAttributeWithContext(_a0 aws.Context, _a1 *ec2.ModifyImageAttributeInput, _a2 ...request.Option) (*ec2.ModifyImageAttributeOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.ModifyImageAttributeOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.ModifyImageAttributeInput, ...request.Option) *ec2.ModifyImageAttributeOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyImageAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.ModifyImageAttributeInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyInstanceAttribute provides a mock function with given fields: _a0
func (_m *EC2API) ModifyInstanceAttribute(_a0 *ec2.ModifyInstanceAttributeInput) (*ec2.ModifyInstanceAttributeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ModifyInstanceAttributeOutput
	if rf, ok := ret.Get(0).(func(*ec2.ModifyInstanceAttributeInput) *ec2.ModifyInstanceAttributeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyInstanceAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ModifyInstanceAttributeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyInstanceAttributeRequest provides a mock function with given fields: _a0
func (_m *EC2API) ModifyInstanceAttributeRequest(_a0 *ec2.ModifyInstanceAttributeInput) (*request.Request, *ec2.ModifyInstanceAttributeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ModifyInstanceAttributeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ModifyInstanceAttributeOutput
	if rf, ok := ret.Get(1).(func(*ec2.ModifyInstanceAttributeInput) *ec2.ModifyInstanceAttributeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ModifyInstanceAttributeOutput)
		}
	}

	return r0, r1
}

// ModifyInstanceAttributeWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) ModifyInstanceAttributeWithContext(_a0 aws.Context, _a1 *ec2.ModifyInstanceAttributeInput, _a2 ...request.Option) (*ec2.ModifyInstanceAttributeOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.ModifyInstanceAttributeOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.ModifyInstanceAttributeInput, ...request.Option) *ec2.ModifyInstanceAttributeOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyInstanceAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.ModifyInstanceAttributeInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyInstanceCreditSpecification provides a mock function with given fields: _a0
func (_m *EC2API) ModifyInstanceCreditSpecification(_a0 *ec2.ModifyInstanceCreditSpecificationInput) (*ec2.ModifyInstanceCreditSpecificationOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ModifyInstanceCreditSpecificationOutput
	if rf, ok := ret.Get(0).(func(*ec2.ModifyInstanceCreditSpecificationInput) *ec2.ModifyInstanceCreditSpecificationOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyInstanceCreditSpecificationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ModifyInstanceCreditSpecificationInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyInstanceCreditSpecificationRequest provides a mock function with given fields: _a0
func (_m *EC2API) ModifyInstanceCreditSpecificationRequest(_a0 *ec2.ModifyInstanceCreditSpecificationInput) (*request.Request, *ec2.ModifyInstanceCreditSpecificationOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ModifyInstanceCreditSpecificationInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ModifyInstanceCreditSpecificationOutput
	if rf, ok := ret.Get(1).(func(*ec2.ModifyInstanceCreditSpecificationInput) *ec2.ModifyInstanceCreditSpecificationOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ModifyInstanceCreditSpecificationOutput)
		}
	}

	return r0, r1
}

// ModifyInstanceCreditSpecificationWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) ModifyInstanceCreditSpecificationWithContext(_a0 aws.Context, _a1 *ec2.ModifyInstanceCreditSpecificationInput, _a2 ...request.Option) (*ec2.ModifyInstanceCreditSpecificationOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.ModifyInstanceCreditSpecificationOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.ModifyInstanceCreditSpecificationInput, ...request.Option) *ec2.ModifyInstanceCreditSpecificationOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyInstanceCreditSpecificationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.ModifyInstanceCreditSpecificationInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyInstancePlacement provides a mock function with given fields: _a0
func (_m *EC2API) ModifyInstancePlacement(_a0 *ec2.ModifyInstancePlacementInput) (*ec2.ModifyInstancePlacementOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ModifyInstancePlacementOutput
	if rf, ok := ret.Get(0).(func(*ec2.ModifyInstancePlacementInput) *ec2.ModifyInstancePlacementOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyInstancePlacementOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ModifyInstancePlacementInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyInstancePlacementRequest provides a mock function with given fields: _a0
func (_m *EC2API) ModifyInstancePlacementRequest(_a0 *ec2.ModifyInstancePlacementInput) (*request.Request, *ec2.ModifyInstancePlacementOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ModifyInstancePlacementInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ModifyInstancePlacementOutput
	if rf, ok := ret.Get(1).(func(*ec2.ModifyInstancePlacementInput) *ec2.ModifyInstancePlacementOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ModifyInstancePlacementOutput)
		}
	}

	return r0, r1
}

// ModifyInstancePlacementWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) ModifyInstancePlacementWithContext(_a0 aws.Context, _a1 *ec2.ModifyInstancePlacementInput, _a2 ...request.Option) (*ec2.ModifyInstancePlacementOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.ModifyInstancePlacementOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.ModifyInstancePlacementInput, ...request.Option) *ec2.ModifyInstancePlacementOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyInstancePlacementOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.ModifyInstancePlacementInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyLaunchTemplate provides a mock function with given fields: _a0
func (_m *EC2API) ModifyLaunchTemplate(_a0 *ec2.ModifyLaunchTemplateInput) (*ec2.ModifyLaunchTemplateOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ModifyLaunchTemplateOutput
	if rf, ok := ret.Get(0).(func(*ec2.ModifyLaunchTemplateInput) *ec2.ModifyLaunchTemplateOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyLaunchTemplateOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ModifyLaunchTemplateInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyLaunchTemplateRequest provides a mock function with given fields: _a0
func (_m *EC2API) ModifyLaunchTemplateRequest(_a0 *ec2.ModifyLaunchTemplateInput) (*request.Request, *ec2.ModifyLaunchTemplateOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ModifyLaunchTemplateInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ModifyLaunchTemplateOutput
	if rf, ok := ret.Get(1).(func(*ec2.ModifyLaunchTemplateInput) *ec2.ModifyLaunchTemplateOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ModifyLaunchTemplateOutput)
		}
	}

	return r0, r1
}

// ModifyLaunchTemplateWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) ModifyLaunchTemplateWithContext(_a0 aws.Context, _a1 *ec2.ModifyLaunchTemplateInput, _a2 ...request.Option) (*ec2.ModifyLaunchTemplateOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.ModifyLaunchTemplateOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.ModifyLaunchTemplateInput, ...request.Option) *ec2.ModifyLaunchTemplateOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyLaunchTemplateOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.ModifyLaunchTemplateInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyNetworkInterfaceAttribute provides a mock function with given fields: _a0
func (_m *EC2API) ModifyNetworkInterfaceAttribute(_a0 *ec2.ModifyNetworkInterfaceAttributeInput) (*ec2.ModifyNetworkInterfaceAttributeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ModifyNetworkInterfaceAttributeOutput
	if rf, ok := ret.Get(0).(func(*ec2.ModifyNetworkInterfaceAttributeInput) *ec2.ModifyNetworkInterfaceAttributeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyNetworkInterfaceAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ModifyNetworkInterfaceAttributeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyNetworkInterfaceAttributeRequest provides a mock function with given fields: _a0
func (_m *EC2API) ModifyNetworkInterfaceAttributeRequest(_a0 *ec2.ModifyNetworkInterfaceAttributeInput) (*request.Request, *ec2.ModifyNetworkInterfaceAttributeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ModifyNetworkInterfaceAttributeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ModifyNetworkInterfaceAttributeOutput
	if rf, ok := ret.Get(1).(func(*ec2.ModifyNetworkInterfaceAttributeInput) *ec2.ModifyNetworkInterfaceAttributeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ModifyNetworkInterfaceAttributeOutput)
		}
	}

	return r0, r1
}

// ModifyNetworkInterfaceAttributeWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) ModifyNetworkInterfaceAttributeWithContext(_a0 aws.Context, _a1 *ec2.ModifyNetworkInterfaceAttributeInput, _a2 ...request.Option) (*ec2.ModifyNetworkInterfaceAttributeOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.ModifyNetworkInterfaceAttributeOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.ModifyNetworkInterfaceAttributeInput, ...request.Option) *ec2.ModifyNetworkInterfaceAttributeOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyNetworkInterfaceAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.ModifyNetworkInterfaceAttributeInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyReservedInstances provides a mock function with given fields: _a0
func (_m *EC2API) ModifyReservedInstances(_a0 *ec2.ModifyReservedInstancesInput) (*ec2.ModifyReservedInstancesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ModifyReservedInstancesOutput
	if rf, ok := ret.Get(0).(func(*ec2.ModifyReservedInstancesInput) *ec2.ModifyReservedInstancesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyReservedInstancesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ModifyReservedInstancesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyReservedInstancesRequest provides a mock function with given fields: _a0
func (_m *EC2API) ModifyReservedInstancesRequest(_a0 *ec2.ModifyReservedInstancesInput) (*request.Request, *ec2.ModifyReservedInstancesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ModifyReservedInstancesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ModifyReservedInstancesOutput
	if rf, ok := ret.Get(1).(func(*ec2.ModifyReservedInstancesInput) *ec2.ModifyReservedInstancesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ModifyReservedInstancesOutput)
		}
	}

	return r0, r1
}

// ModifyReservedInstancesWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) ModifyReservedInstancesWithContext(_a0 aws.Context, _a1 *ec2.ModifyReservedInstancesInput, _a2 ...request.Option) (*ec2.ModifyReservedInstancesOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.ModifyReservedInstancesOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.ModifyReservedInstancesInput, ...request.Option) *ec2.ModifyReservedInstancesOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyReservedInstancesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.ModifyReservedInstancesInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifySnapshotAttribute provides a mock function with given fields: _a0
func (_m *EC2API) ModifySnapshotAttribute(_a0 *ec2.ModifySnapshotAttributeInput) (*ec2.ModifySnapshotAttributeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ModifySnapshotAttributeOutput
	if rf, ok := ret.Get(0).(func(*ec2.ModifySnapshotAttributeInput) *ec2.ModifySnapshotAttributeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifySnapshotAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ModifySnapshotAttributeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifySnapshotAttributeRequest provides a mock function with given fields: _a0
func (_m *EC2API) ModifySnapshotAttributeRequest(_a0 *ec2.ModifySnapshotAttributeInput) (*request.Request, *ec2.ModifySnapshotAttributeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ModifySnapshotAttributeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ModifySnapshotAttributeOutput
	if rf, ok := ret.Get(1).(func(*ec2.ModifySnapshotAttributeInput) *ec2.ModifySnapshotAttributeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ModifySnapshotAttributeOutput)
		}
	}

	return r0, r1
}

// ModifySnapshotAttributeWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) ModifySnapshotAttributeWithContext(_a0 aws.Context, _a1 *ec2.ModifySnapshotAttributeInput, _a2 ...request.Option) (*ec2.ModifySnapshotAttributeOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.ModifySnapshotAttributeOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.ModifySnapshotAttributeInput, ...request.Option) *ec2.ModifySnapshotAttributeOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifySnapshotAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.ModifySnapshotAttributeInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifySpotFleetRequest provides a mock function with given fields: _a0
func (_m *EC2API) ModifySpotFleetRequest(_a0 *ec2.ModifySpotFleetRequestInput) (*ec2.ModifySpotFleetRequestOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ModifySpotFleetRequestOutput
	if rf, ok := ret.Get(0).(func(*ec2.ModifySpotFleetRequestInput) *ec2.ModifySpotFleetRequestOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifySpotFleetRequestOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ModifySpotFleetRequestInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifySpotFleetRequestRequest provides a mock function with given fields: _a0
func (_m *EC2API) ModifySpotFleetRequestRequest(_a0 *ec2.ModifySpotFleetRequestInput) (*request.Request, *ec2.ModifySpotFleetRequestOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ModifySpotFleetRequestInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ModifySpotFleetRequestOutput
	if rf, ok := ret.Get(1).(func(*ec2.ModifySpotFleetRequestInput) *ec2.ModifySpotFleetRequestOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ModifySpotFleetRequestOutput)
		}
	}

	return r0, r1
}

// ModifySpotFleetRequestWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) ModifySpotFleetRequestWithContext(_a0 aws.Context, _a1 *ec2.ModifySpotFleetRequestInput, _a2 ...request.Option) (*ec2.ModifySpotFleetRequestOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.ModifySpotFleetRequestOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.ModifySpotFleetRequestInput, ...request.Option) *ec2.ModifySpotFleetRequestOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifySpotFleetRequestOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.ModifySpotFleetRequestInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifySubnetAttribute provides a mock function with given fields: _a0
func (_m *EC2API) ModifySubnetAttribute(_a0 *ec2.ModifySubnetAttributeInput) (*ec2.ModifySubnetAttributeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ModifySubnetAttributeOutput
	if rf, ok := ret.Get(0).(func(*ec2.ModifySubnetAttributeInput) *ec2.ModifySubnetAttributeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifySubnetAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ModifySubnetAttributeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifySubnetAttributeRequest provides a mock function with given fields: _a0
func (_m *EC2API) ModifySubnetAttributeRequest(_a0 *ec2.ModifySubnetAttributeInput) (*request.Request, *ec2.ModifySubnetAttributeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ModifySubnetAttributeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ModifySubnetAttributeOutput
	if rf, ok := ret.Get(1).(func(*ec2.ModifySubnetAttributeInput) *ec2.ModifySubnetAttributeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ModifySubnetAttributeOutput)
		}
	}

	return r0, r1
}

// ModifySubnetAttributeWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) ModifySubnetAttributeWithContext(_a0 aws.Context, _a1 *ec2.ModifySubnetAttributeInput, _a2 ...request.Option) (*ec2.ModifySubnetAttributeOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.ModifySubnetAttributeOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.ModifySubnetAttributeInput, ...request.Option) *ec2.ModifySubnetAttributeOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifySubnetAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.ModifySubnetAttributeInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyVolume provides a mock function with given fields: _a0
func (_m *EC2API) ModifyVolume(_a0 *ec2.ModifyVolumeInput) (*ec2.ModifyVolumeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ModifyVolumeOutput
	if rf, ok := ret.Get(0).(func(*ec2.ModifyVolumeInput) *ec2.ModifyVolumeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyVolumeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ModifyVolumeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyVolumeAttribute provides a mock function with given fields: _a0
func (_m *EC2API) ModifyVolumeAttribute(_a0 *ec2.ModifyVolumeAttributeInput) (*ec2.ModifyVolumeAttributeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ModifyVolumeAttributeOutput
	if rf, ok := ret.Get(0).(func(*ec2.ModifyVolumeAttributeInput) *ec2.ModifyVolumeAttributeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyVolumeAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ModifyVolumeAttributeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyVolumeAttributeRequest provides a mock function with given fields: _a0
func (_m *EC2API) ModifyVolumeAttributeRequest(_a0 *ec2.ModifyVolumeAttributeInput) (*request.Request, *ec2.ModifyVolumeAttributeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ModifyVolumeAttributeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ModifyVolumeAttributeOutput
	if rf, ok := ret.Get(1).(func(*ec2.ModifyVolumeAttributeInput) *ec2.ModifyVolumeAttributeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ModifyVolumeAttributeOutput)
		}
	}

	return r0, r1
}

// ModifyVolumeAttributeWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) ModifyVolumeAttributeWithContext(_a0 aws.Context, _a1 *ec2.ModifyVolumeAttributeInput, _a2 ...request.Option) (*ec2.ModifyVolumeAttributeOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.ModifyVolumeAttributeOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.ModifyVolumeAttributeInput, ...request.Option) *ec2.ModifyVolumeAttributeOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyVolumeAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.ModifyVolumeAttributeInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyVolumeRequest provides a mock function with given fields: _a0
func (_m *EC2API) ModifyVolumeRequest(_a0 *ec2.ModifyVolumeInput) (*request.Request, *ec2.ModifyVolumeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ModifyVolumeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ModifyVolumeOutput
	if rf, ok := ret.Get(1).(func(*ec2.ModifyVolumeInput) *ec2.ModifyVolumeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ModifyVolumeOutput)
		}
	}

	return r0, r1
}

// ModifyVolumeWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) ModifyVolumeWithContext(_a0 aws.Context, _a1 *ec2.ModifyVolumeInput, _a2 ...request.Option) (*ec2.ModifyVolumeOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.ModifyVolumeOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.ModifyVolumeInput, ...request.Option) *ec2.ModifyVolumeOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyVolumeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.ModifyVolumeInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyVpcAttribute provides a mock function with given fields: _a0
func (_m *EC2API) ModifyVpcAttribute(_a0 *ec2.ModifyVpcAttributeInput) (*ec2.ModifyVpcAttributeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ModifyVpcAttributeOutput
	if rf, ok := ret.Get(0).(func(*ec2.ModifyVpcAttributeInput) *ec2.ModifyVpcAttributeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyVpcAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ModifyVpcAttributeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyVpcAttributeRequest provides a mock function with given fields: _a0
func (_m *EC2API) ModifyVpcAttributeRequest(_a0 *ec2.ModifyVpcAttributeInput) (*request.Request, *ec2.ModifyVpcAttributeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ModifyVpcAttributeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ModifyVpcAttributeOutput
	if rf, ok := ret.Get(1).(func(*ec2.ModifyVpcAttributeInput) *ec2.ModifyVpcAttributeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ModifyVpcAttributeOutput)
		}
	}

	return r0, r1
}

// ModifyVpcAttributeWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) ModifyVpcAttributeWithContext(_a0 aws.Context, _a1 *ec2.ModifyVpcAttributeInput, _a2 ...request.Option) (*ec2.ModifyVpcAttributeOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.ModifyVpcAttributeOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.ModifyVpcAttributeInput, ...request.Option) *ec2.ModifyVpcAttributeOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyVpcAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.ModifyVpcAttributeInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyVpcEndpoint provides a mock function with given fields: _a0
func (_m *EC2API) ModifyVpcEndpoint(_a0 *ec2.ModifyVpcEndpointInput) (*ec2.ModifyVpcEndpointOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ModifyVpcEndpointOutput
	if rf, ok := ret.Get(0).(func(*ec2.ModifyVpcEndpointInput) *ec2.ModifyVpcEndpointOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyVpcEndpointOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ModifyVpcEndpointInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyVpcEndpointConnectionNotification provides a mock function with given fields: _a0
func (_m *EC2API) ModifyVpcEndpointConnectionNotification(_a0 *ec2.ModifyVpcEndpointConnectionNotificationInput) (*ec2.ModifyVpcEndpointConnectionNotificationOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ModifyVpcEndpointConnectionNotificationOutput
	if rf, ok := ret.Get(0).(func(*ec2.ModifyVpcEndpointConnectionNotificationInput) *ec2.ModifyVpcEndpointConnectionNotificationOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyVpcEndpointConnectionNotificationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ModifyVpcEndpointConnectionNotificationInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyVpcEndpointConnectionNotificationRequest provides a mock function with given fields: _a0
func (_m *EC2API) ModifyVpcEndpointConnectionNotificationRequest(_a0 *ec2.ModifyVpcEndpointConnectionNotificationInput) (*request.Request, *ec2.ModifyVpcEndpointConnectionNotificationOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ModifyVpcEndpointConnectionNotificationInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ModifyVpcEndpointConnectionNotificationOutput
	if rf, ok := ret.Get(1).(func(*ec2.ModifyVpcEndpointConnectionNotificationInput) *ec2.ModifyVpcEndpointConnectionNotificationOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ModifyVpcEndpointConnectionNotificationOutput)
		}
	}

	return r0, r1
}

// ModifyVpcEndpointConnectionNotificationWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) ModifyVpcEndpointConnectionNotificationWithContext(_a0 aws.Context, _a1 *ec2.ModifyVpcEndpointConnectionNotificationInput, _a2 ...request.Option) (*ec2.ModifyVpcEndpointConnectionNotificationOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.ModifyVpcEndpointConnectionNotificationOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.ModifyVpcEndpointConnectionNotificationInput, ...request.Option) *ec2.ModifyVpcEndpointConnectionNotificationOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyVpcEndpointConnectionNotificationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.ModifyVpcEndpointConnectionNotificationInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyVpcEndpointRequest provides a mock function with given fields: _a0
func (_m *EC2API) ModifyVpcEndpointRequest(_a0 *ec2.ModifyVpcEndpointInput) (*request.Request, *ec2.ModifyVpcEndpointOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ModifyVpcEndpointInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ModifyVpcEndpointOutput
	if rf, ok := ret.Get(1).(func(*ec2.ModifyVpcEndpointInput) *ec2.ModifyVpcEndpointOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ModifyVpcEndpointOutput)
		}
	}

	return r0, r1
}

// ModifyVpcEndpointServiceConfiguration provides a mock function with given fields: _a0
func (_m *EC2API) ModifyVpcEndpointServiceConfiguration(_a0 *ec2.ModifyVpcEndpointServiceConfigurationInput) (*ec2.ModifyVpcEndpointServiceConfigurationOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ModifyVpcEndpointServiceConfigurationOutput
	if rf, ok := ret.Get(0).(func(*ec2.ModifyVpcEndpointServiceConfigurationInput) *ec2.ModifyVpcEndpointServiceConfigurationOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyVpcEndpointServiceConfigurationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ModifyVpcEndpointServiceConfigurationInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyVpcEndpointServiceConfigurationRequest provides a mock function with given fields: _a0
func (_m *EC2API) ModifyVpcEndpointServiceConfigurationRequest(_a0 *ec2.ModifyVpcEndpointServiceConfigurationInput) (*request.Request, *ec2.ModifyVpcEndpointServiceConfigurationOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ModifyVpcEndpointServiceConfigurationInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ModifyVpcEndpointServiceConfigurationOutput
	if rf, ok := ret.Get(1).(func(*ec2.ModifyVpcEndpointServiceConfigurationInput) *ec2.ModifyVpcEndpointServiceConfigurationOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ModifyVpcEndpointServiceConfigurationOutput)
		}
	}

	return r0, r1
}

// ModifyVpcEndpointServiceConfigurationWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) ModifyVpcEndpointServiceConfigurationWithContext(_a0 aws.Context, _a1 *ec2.ModifyVpcEndpointServiceConfigurationInput, _a2 ...request.Option) (*ec2.ModifyVpcEndpointServiceConfigurationOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.ModifyVpcEndpointServiceConfigurationOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.ModifyVpcEndpointServiceConfigurationInput, ...request.Option) *ec2.ModifyVpcEndpointServiceConfigurationOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyVpcEndpointServiceConfigurationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.ModifyVpcEndpointServiceConfigurationInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyVpcEndpointServicePermissions provides a mock function with given fields: _a0
func (_m *EC2API) ModifyVpcEndpointServicePermissions(_a0 *ec2.ModifyVpcEndpointServicePermissionsInput) (*ec2.ModifyVpcEndpointServicePermissionsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ModifyVpcEndpointServicePermissionsOutput
	if rf, ok := ret.Get(0).(func(*ec2.ModifyVpcEndpointServicePermissionsInput) *ec2.ModifyVpcEndpointServicePermissionsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyVpcEndpointServicePermissionsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ModifyVpcEndpointServicePermissionsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyVpcEndpointServicePermissionsRequest provides a mock function with given fields: _a0
func (_m *EC2API) ModifyVpcEndpointServicePermissionsRequest(_a0 *ec2.ModifyVpcEndpointServicePermissionsInput) (*request.Request, *ec2.ModifyVpcEndpointServicePermissionsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ModifyVpcEndpointServicePermissionsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ModifyVpcEndpointServicePermissionsOutput
	if rf, ok := ret.Get(1).(func(*ec2.ModifyVpcEndpointServicePermissionsInput) *ec2.ModifyVpcEndpointServicePermissionsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ModifyVpcEndpointServicePermissionsOutput)
		}
	}

	return r0, r1
}

// ModifyVpcEndpointServicePermissionsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) ModifyVpcEndpointServicePermissionsWithContext(_a0 aws.Context, _a1 *ec2.ModifyVpcEndpointServicePermissionsInput, _a2 ...request.Option) (*ec2.ModifyVpcEndpointServicePermissionsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.ModifyVpcEndpointServicePermissionsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.ModifyVpcEndpointServicePermissionsInput, ...request.Option) *ec2.ModifyVpcEndpointServicePermissionsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyVpcEndpointServicePermissionsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.ModifyVpcEndpointServicePermissionsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyVpcEndpointWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) ModifyVpcEndpointWithContext(_a0 aws.Context, _a1 *ec2.ModifyVpcEndpointInput, _a2 ...request.Option) (*ec2.ModifyVpcEndpointOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.ModifyVpcEndpointOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.ModifyVpcEndpointInput, ...request.Option) *ec2.ModifyVpcEndpointOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyVpcEndpointOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.ModifyVpcEndpointInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyVpcPeeringConnectionOptions provides a mock function with given fields: _a0
func (_m *EC2API) ModifyVpcPeeringConnectionOptions(_a0 *ec2.ModifyVpcPeeringConnectionOptionsInput) (*ec2.ModifyVpcPeeringConnectionOptionsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ModifyVpcPeeringConnectionOptionsOutput
	if rf, ok := ret.Get(0).(func(*ec2.ModifyVpcPeeringConnectionOptionsInput) *ec2.ModifyVpcPeeringConnectionOptionsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyVpcPeeringConnectionOptionsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ModifyVpcPeeringConnectionOptionsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyVpcPeeringConnectionOptionsRequest provides a mock function with given fields: _a0
func (_m *EC2API) ModifyVpcPeeringConnectionOptionsRequest(_a0 *ec2.ModifyVpcPeeringConnectionOptionsInput) (*request.Request, *ec2.ModifyVpcPeeringConnectionOptionsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ModifyVpcPeeringConnectionOptionsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ModifyVpcPeeringConnectionOptionsOutput
	if rf, ok := ret.Get(1).(func(*ec2.ModifyVpcPeeringConnectionOptionsInput) *ec2.ModifyVpcPeeringConnectionOptionsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ModifyVpcPeeringConnectionOptionsOutput)
		}
	}

	return r0, r1
}

// ModifyVpcPeeringConnectionOptionsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) ModifyVpcPeeringConnectionOptionsWithContext(_a0 aws.Context, _a1 *ec2.ModifyVpcPeeringConnectionOptionsInput, _a2 ...request.Option) (*ec2.ModifyVpcPeeringConnectionOptionsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.ModifyVpcPeeringConnectionOptionsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.ModifyVpcPeeringConnectionOptionsInput, ...request.Option) *ec2.ModifyVpcPeeringConnectionOptionsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyVpcPeeringConnectionOptionsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.ModifyVpcPeeringConnectionOptionsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyVpcTenancy provides a mock function with given fields: _a0
func (_m *EC2API) ModifyVpcTenancy(_a0 *ec2.ModifyVpcTenancyInput) (*ec2.ModifyVpcTenancyOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ModifyVpcTenancyOutput
	if rf, ok := ret.Get(0).(func(*ec2.ModifyVpcTenancyInput) *ec2.ModifyVpcTenancyOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyVpcTenancyOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ModifyVpcTenancyInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyVpcTenancyRequest provides a mock function with given fields: _a0
func (_m *EC2API) ModifyVpcTenancyRequest(_a0 *ec2.ModifyVpcTenancyInput) (*request.Request, *ec2.ModifyVpcTenancyOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ModifyVpcTenancyInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ModifyVpcTenancyOutput
	if rf, ok := ret.Get(1).(func(*ec2.ModifyVpcTenancyInput) *ec2.ModifyVpcTenancyOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ModifyVpcTenancyOutput)
		}
	}

	return r0, r1
}

// ModifyVpcTenancyWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) ModifyVpcTenancyWithContext(_a0 aws.Context, _a1 *ec2.ModifyVpcTenancyInput, _a2 ...request.Option) (*ec2.ModifyVpcTenancyOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.ModifyVpcTenancyOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.ModifyVpcTenancyInput, ...request.Option) *ec2.ModifyVpcTenancyOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ModifyVpcTenancyOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.ModifyVpcTenancyInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MonitorInstances provides a mock function with given fields: _a0
func (_m *EC2API) MonitorInstances(_a0 *ec2.MonitorInstancesInput) (*ec2.MonitorInstancesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.MonitorInstancesOutput
	if rf, ok := ret.Get(0).(func(*ec2.MonitorInstancesInput) *ec2.MonitorInstancesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.MonitorInstancesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.MonitorInstancesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MonitorInstancesRequest provides a mock function with given fields: _a0
func (_m *EC2API) MonitorInstancesRequest(_a0 *ec2.MonitorInstancesInput) (*request.Request, *ec2.MonitorInstancesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.MonitorInstancesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.MonitorInstancesOutput
	if rf, ok := ret.Get(1).(func(*ec2.MonitorInstancesInput) *ec2.MonitorInstancesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.MonitorInstancesOutput)
		}
	}

	return r0, r1
}

// MonitorInstancesWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) MonitorInstancesWithContext(_a0 aws.Context, _a1 *ec2.MonitorInstancesInput, _a2 ...request.Option) (*ec2.MonitorInstancesOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.MonitorInstancesOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.MonitorInstancesInput, ...request.Option) *ec2.MonitorInstancesOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.MonitorInstancesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.MonitorInstancesInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MoveAddressToVpc provides a mock function with given fields: _a0
func (_m *EC2API) MoveAddressToVpc(_a0 *ec2.MoveAddressToVpcInput) (*ec2.MoveAddressToVpcOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.MoveAddressToVpcOutput
	if rf, ok := ret.Get(0).(func(*ec2.MoveAddressToVpcInput) *ec2.MoveAddressToVpcOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.MoveAddressToVpcOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.MoveAddressToVpcInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MoveAddressToVpcRequest provides a mock function with given fields: _a0
func (_m *EC2API) MoveAddressToVpcRequest(_a0 *ec2.MoveAddressToVpcInput) (*request.Request, *ec2.MoveAddressToVpcOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.MoveAddressToVpcInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.MoveAddressToVpcOutput
	if rf, ok := ret.Get(1).(func(*ec2.MoveAddressToVpcInput) *ec2.MoveAddressToVpcOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.MoveAddressToVpcOutput)
		}
	}

	return r0, r1
}

// MoveAddressToVpcWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) MoveAddressToVpcWithContext(_a0 aws.Context, _a1 *ec2.MoveAddressToVpcInput, _a2 ...request.Option) (*ec2.MoveAddressToVpcOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.MoveAddressToVpcOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.MoveAddressToVpcInput, ...request.Option) *ec2.MoveAddressToVpcOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.MoveAddressToVpcOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.MoveAddressToVpcInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PurchaseHostReservation provides a mock function with given fields: _a0
func (_m *EC2API) PurchaseHostReservation(_a0 *ec2.PurchaseHostReservationInput) (*ec2.PurchaseHostReservationOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.PurchaseHostReservationOutput
	if rf, ok := ret.Get(0).(func(*ec2.PurchaseHostReservationInput) *ec2.PurchaseHostReservationOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.PurchaseHostReservationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.PurchaseHostReservationInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PurchaseHostReservationRequest provides a mock function with given fields: _a0
func (_m *EC2API) PurchaseHostReservationRequest(_a0 *ec2.PurchaseHostReservationInput) (*request.Request, *ec2.PurchaseHostReservationOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.PurchaseHostReservationInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.PurchaseHostReservationOutput
	if rf, ok := ret.Get(1).(func(*ec2.PurchaseHostReservationInput) *ec2.PurchaseHostReservationOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.PurchaseHostReservationOutput)
		}
	}

	return r0, r1
}

// PurchaseHostReservationWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) PurchaseHostReservationWithContext(_a0 aws.Context, _a1 *ec2.PurchaseHostReservationInput, _a2 ...request.Option) (*ec2.PurchaseHostReservationOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.PurchaseHostReservationOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.PurchaseHostReservationInput, ...request.Option) *ec2.PurchaseHostReservationOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.PurchaseHostReservationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.PurchaseHostReservationInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PurchaseReservedInstancesOffering provides a mock function with given fields: _a0
func (_m *EC2API) PurchaseReservedInstancesOffering(_a0 *ec2.PurchaseReservedInstancesOfferingInput) (*ec2.PurchaseReservedInstancesOfferingOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.PurchaseReservedInstancesOfferingOutput
	if rf, ok := ret.Get(0).(func(*ec2.PurchaseReservedInstancesOfferingInput) *ec2.PurchaseReservedInstancesOfferingOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.PurchaseReservedInstancesOfferingOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.PurchaseReservedInstancesOfferingInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PurchaseReservedInstancesOfferingRequest provides a mock function with given fields: _a0
func (_m *EC2API) PurchaseReservedInstancesOfferingRequest(_a0 *ec2.PurchaseReservedInstancesOfferingInput) (*request.Request, *ec2.PurchaseReservedInstancesOfferingOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.PurchaseReservedInstancesOfferingInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.PurchaseReservedInstancesOfferingOutput
	if rf, ok := ret.Get(1).(func(*ec2.PurchaseReservedInstancesOfferingInput) *ec2.PurchaseReservedInstancesOfferingOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.PurchaseReservedInstancesOfferingOutput)
		}
	}

	return r0, r1
}

// PurchaseReservedInstancesOfferingWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) PurchaseReservedInstancesOfferingWithContext(_a0 aws.Context, _a1 *ec2.PurchaseReservedInstancesOfferingInput, _a2 ...request.Option) (*ec2.PurchaseReservedInstancesOfferingOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.PurchaseReservedInstancesOfferingOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.PurchaseReservedInstancesOfferingInput, ...request.Option) *ec2.PurchaseReservedInstancesOfferingOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.PurchaseReservedInstancesOfferingOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.PurchaseReservedInstancesOfferingInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PurchaseScheduledInstances provides a mock function with given fields: _a0
func (_m *EC2API) PurchaseScheduledInstances(_a0 *ec2.PurchaseScheduledInstancesInput) (*ec2.PurchaseScheduledInstancesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.PurchaseScheduledInstancesOutput
	if rf, ok := ret.Get(0).(func(*ec2.PurchaseScheduledInstancesInput) *ec2.PurchaseScheduledInstancesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.PurchaseScheduledInstancesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.PurchaseScheduledInstancesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PurchaseScheduledInstancesRequest provides a mock function with given fields: _a0
func (_m *EC2API) PurchaseScheduledInstancesRequest(_a0 *ec2.PurchaseScheduledInstancesInput) (*request.Request, *ec2.PurchaseScheduledInstancesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.PurchaseScheduledInstancesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.PurchaseScheduledInstancesOutput
	if rf, ok := ret.Get(1).(func(*ec2.PurchaseScheduledInstancesInput) *ec2.PurchaseScheduledInstancesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.PurchaseScheduledInstancesOutput)
		}
	}

	return r0, r1
}

// PurchaseScheduledInstancesWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) PurchaseScheduledInstancesWithContext(_a0 aws.Context, _a1 *ec2.PurchaseScheduledInstancesInput, _a2 ...request.Option) (*ec2.PurchaseScheduledInstancesOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.PurchaseScheduledInstancesOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.PurchaseScheduledInstancesInput, ...request.Option) *ec2.PurchaseScheduledInstancesOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.PurchaseScheduledInstancesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.PurchaseScheduledInstancesInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RebootInstances provides a mock function with given fields: _a0
func (_m *EC2API) RebootInstances(_a0 *ec2.RebootInstancesInput) (*ec2.RebootInstancesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.RebootInstancesOutput
	if rf, ok := ret.Get(0).(func(*ec2.RebootInstancesInput) *ec2.RebootInstancesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.RebootInstancesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.RebootInstancesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RebootInstancesRequest provides a mock function with given fields: _a0
func (_m *EC2API) RebootInstancesRequest(_a0 *ec2.RebootInstancesInput) (*request.Request, *ec2.RebootInstancesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.RebootInstancesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.RebootInstancesOutput
	if rf, ok := ret.Get(1).(func(*ec2.RebootInstancesInput) *ec2.RebootInstancesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.RebootInstancesOutput)
		}
	}

	return r0, r1
}

// RebootInstancesWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) RebootInstancesWithContext(_a0 aws.Context, _a1 *ec2.RebootInstancesInput, _a2 ...request.Option) (*ec2.RebootInstancesOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.RebootInstancesOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.RebootInstancesInput, ...request.Option) *ec2.RebootInstancesOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.RebootInstancesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.RebootInstancesInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterImage provides a mock function with given fields: _a0
func (_m *EC2API) RegisterImage(_a0 *ec2.RegisterImageInput) (*ec2.RegisterImageOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.RegisterImageOutput
	if rf, ok := ret.Get(0).(func(*ec2.RegisterImageInput) *ec2.RegisterImageOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.RegisterImageOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.RegisterImageInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterImageRequest provides a mock function with given fields: _a0
func (_m *EC2API) RegisterImageRequest(_a0 *ec2.RegisterImageInput) (*request.Request, *ec2.RegisterImageOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.RegisterImageInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.RegisterImageOutput
	if rf, ok := ret.Get(1).(func(*ec2.RegisterImageInput) *ec2.RegisterImageOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.RegisterImageOutput)
		}
	}

	return r0, r1
}

// RegisterImageWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) RegisterImageWithContext(_a0 aws.Context, _a1 *ec2.RegisterImageInput, _a2 ...request.Option) (*ec2.RegisterImageOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.RegisterImageOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.RegisterImageInput, ...request.Option) *ec2.RegisterImageOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.RegisterImageOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.RegisterImageInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RejectVpcEndpointConnections provides a mock function with given fields: _a0
func (_m *EC2API) RejectVpcEndpointConnections(_a0 *ec2.RejectVpcEndpointConnectionsInput) (*ec2.RejectVpcEndpointConnectionsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.RejectVpcEndpointConnectionsOutput
	if rf, ok := ret.Get(0).(func(*ec2.RejectVpcEndpointConnectionsInput) *ec2.RejectVpcEndpointConnectionsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.RejectVpcEndpointConnectionsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.RejectVpcEndpointConnectionsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RejectVpcEndpointConnectionsRequest provides a mock function with given fields: _a0
func (_m *EC2API) RejectVpcEndpointConnectionsRequest(_a0 *ec2.RejectVpcEndpointConnectionsInput) (*request.Request, *ec2.RejectVpcEndpointConnectionsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.RejectVpcEndpointConnectionsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.RejectVpcEndpointConnectionsOutput
	if rf, ok := ret.Get(1).(func(*ec2.RejectVpcEndpointConnectionsInput) *ec2.RejectVpcEndpointConnectionsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.RejectVpcEndpointConnectionsOutput)
		}
	}

	return r0, r1
}

// RejectVpcEndpointConnectionsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) RejectVpcEndpointConnectionsWithContext(_a0 aws.Context, _a1 *ec2.RejectVpcEndpointConnectionsInput, _a2 ...request.Option) (*ec2.RejectVpcEndpointConnectionsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.RejectVpcEndpointConnectionsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.RejectVpcEndpointConnectionsInput, ...request.Option) *ec2.RejectVpcEndpointConnectionsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.RejectVpcEndpointConnectionsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.RejectVpcEndpointConnectionsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RejectVpcPeeringConnection provides a mock function with given fields: _a0
func (_m *EC2API) RejectVpcPeeringConnection(_a0 *ec2.RejectVpcPeeringConnectionInput) (*ec2.RejectVpcPeeringConnectionOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.RejectVpcPeeringConnectionOutput
	if rf, ok := ret.Get(0).(func(*ec2.RejectVpcPeeringConnectionInput) *ec2.RejectVpcPeeringConnectionOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.RejectVpcPeeringConnectionOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.RejectVpcPeeringConnectionInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RejectVpcPeeringConnectionRequest provides a mock function with given fields: _a0
func (_m *EC2API) RejectVpcPeeringConnectionRequest(_a0 *ec2.RejectVpcPeeringConnectionInput) (*request.Request, *ec2.RejectVpcPeeringConnectionOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.RejectVpcPeeringConnectionInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.RejectVpcPeeringConnectionOutput
	if rf, ok := ret.Get(1).(func(*ec2.RejectVpcPeeringConnectionInput) *ec2.RejectVpcPeeringConnectionOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.RejectVpcPeeringConnectionOutput)
		}
	}

	return r0, r1
}

// RejectVpcPeeringConnectionWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) RejectVpcPeeringConnectionWithContext(_a0 aws.Context, _a1 *ec2.RejectVpcPeeringConnectionInput, _a2 ...request.Option) (*ec2.RejectVpcPeeringConnectionOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.RejectVpcPeeringConnectionOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.RejectVpcPeeringConnectionInput, ...request.Option) *ec2.RejectVpcPeeringConnectionOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.RejectVpcPeeringConnectionOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.RejectVpcPeeringConnectionInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}




// ReleaseAddressWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) ReleaseAddressWithContext(_a0 aws.Context, _a1 *ec2.ReleaseAddressInput, _a2 ...request.Option) (*ec2.ReleaseAddressOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.ReleaseAddressOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.ReleaseAddressInput, ...request.Option) *ec2.ReleaseAddressOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ReleaseAddressOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.ReleaseAddressInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReleaseHosts provides a mock function with given fields: _a0
func (_m *EC2API) ReleaseHosts(_a0 *ec2.ReleaseHostsInput) (*ec2.ReleaseHostsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ReleaseHostsOutput
	if rf, ok := ret.Get(0).(func(*ec2.ReleaseHostsInput) *ec2.ReleaseHostsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ReleaseHostsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ReleaseHostsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReleaseHostsRequest provides a mock function with given fields: _a0
func (_m *EC2API) ReleaseHostsRequest(_a0 *ec2.ReleaseHostsInput) (*request.Request, *ec2.ReleaseHostsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ReleaseHostsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ReleaseHostsOutput
	if rf, ok := ret.Get(1).(func(*ec2.ReleaseHostsInput) *ec2.ReleaseHostsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ReleaseHostsOutput)
		}
	}

	return r0, r1
}

// ReleaseHostsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) ReleaseHostsWithContext(_a0 aws.Context, _a1 *ec2.ReleaseHostsInput, _a2 ...request.Option) (*ec2.ReleaseHostsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.ReleaseHostsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.ReleaseHostsInput, ...request.Option) *ec2.ReleaseHostsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ReleaseHostsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.ReleaseHostsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReplaceIamInstanceProfileAssociation provides a mock function with given fields: _a0
func (_m *EC2API) ReplaceIamInstanceProfileAssociation(_a0 *ec2.ReplaceIamInstanceProfileAssociationInput) (*ec2.ReplaceIamInstanceProfileAssociationOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ReplaceIamInstanceProfileAssociationOutput
	if rf, ok := ret.Get(0).(func(*ec2.ReplaceIamInstanceProfileAssociationInput) *ec2.ReplaceIamInstanceProfileAssociationOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ReplaceIamInstanceProfileAssociationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ReplaceIamInstanceProfileAssociationInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReplaceIamInstanceProfileAssociationRequest provides a mock function with given fields: _a0
func (_m *EC2API) ReplaceIamInstanceProfileAssociationRequest(_a0 *ec2.ReplaceIamInstanceProfileAssociationInput) (*request.Request, *ec2.ReplaceIamInstanceProfileAssociationOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ReplaceIamInstanceProfileAssociationInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ReplaceIamInstanceProfileAssociationOutput
	if rf, ok := ret.Get(1).(func(*ec2.ReplaceIamInstanceProfileAssociationInput) *ec2.ReplaceIamInstanceProfileAssociationOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ReplaceIamInstanceProfileAssociationOutput)
		}
	}

	return r0, r1
}

// ReplaceIamInstanceProfileAssociationWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) ReplaceIamInstanceProfileAssociationWithContext(_a0 aws.Context, _a1 *ec2.ReplaceIamInstanceProfileAssociationInput, _a2 ...request.Option) (*ec2.ReplaceIamInstanceProfileAssociationOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.ReplaceIamInstanceProfileAssociationOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.ReplaceIamInstanceProfileAssociationInput, ...request.Option) *ec2.ReplaceIamInstanceProfileAssociationOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ReplaceIamInstanceProfileAssociationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.ReplaceIamInstanceProfileAssociationInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReplaceNetworkAclAssociation provides a mock function with given fields: _a0
func (_m *EC2API) ReplaceNetworkAclAssociation(_a0 *ec2.ReplaceNetworkAclAssociationInput) (*ec2.ReplaceNetworkAclAssociationOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ReplaceNetworkAclAssociationOutput
	if rf, ok := ret.Get(0).(func(*ec2.ReplaceNetworkAclAssociationInput) *ec2.ReplaceNetworkAclAssociationOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ReplaceNetworkAclAssociationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ReplaceNetworkAclAssociationInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReplaceNetworkAclAssociationRequest provides a mock function with given fields: _a0
func (_m *EC2API) ReplaceNetworkAclAssociationRequest(_a0 *ec2.ReplaceNetworkAclAssociationInput) (*request.Request, *ec2.ReplaceNetworkAclAssociationOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ReplaceNetworkAclAssociationInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ReplaceNetworkAclAssociationOutput
	if rf, ok := ret.Get(1).(func(*ec2.ReplaceNetworkAclAssociationInput) *ec2.ReplaceNetworkAclAssociationOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ReplaceNetworkAclAssociationOutput)
		}
	}

	return r0, r1
}

// ReplaceNetworkAclAssociationWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) ReplaceNetworkAclAssociationWithContext(_a0 aws.Context, _a1 *ec2.ReplaceNetworkAclAssociationInput, _a2 ...request.Option) (*ec2.ReplaceNetworkAclAssociationOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.ReplaceNetworkAclAssociationOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.ReplaceNetworkAclAssociationInput, ...request.Option) *ec2.ReplaceNetworkAclAssociationOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ReplaceNetworkAclAssociationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.ReplaceNetworkAclAssociationInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReplaceNetworkAclEntry provides a mock function with given fields: _a0
func (_m *EC2API) ReplaceNetworkAclEntry(_a0 *ec2.ReplaceNetworkAclEntryInput) (*ec2.ReplaceNetworkAclEntryOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ReplaceNetworkAclEntryOutput
	if rf, ok := ret.Get(0).(func(*ec2.ReplaceNetworkAclEntryInput) *ec2.ReplaceNetworkAclEntryOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ReplaceNetworkAclEntryOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ReplaceNetworkAclEntryInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReplaceNetworkAclEntryRequest provides a mock function with given fields: _a0
func (_m *EC2API) ReplaceNetworkAclEntryRequest(_a0 *ec2.ReplaceNetworkAclEntryInput) (*request.Request, *ec2.ReplaceNetworkAclEntryOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ReplaceNetworkAclEntryInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ReplaceNetworkAclEntryOutput
	if rf, ok := ret.Get(1).(func(*ec2.ReplaceNetworkAclEntryInput) *ec2.ReplaceNetworkAclEntryOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ReplaceNetworkAclEntryOutput)
		}
	}

	return r0, r1
}

// ReplaceNetworkAclEntryWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) ReplaceNetworkAclEntryWithContext(_a0 aws.Context, _a1 *ec2.ReplaceNetworkAclEntryInput, _a2 ...request.Option) (*ec2.ReplaceNetworkAclEntryOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.ReplaceNetworkAclEntryOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.ReplaceNetworkAclEntryInput, ...request.Option) *ec2.ReplaceNetworkAclEntryOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ReplaceNetworkAclEntryOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.ReplaceNetworkAclEntryInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReplaceRoute provides a mock function with given fields: _a0
func (_m *EC2API) ReplaceRoute(_a0 *ec2.ReplaceRouteInput) (*ec2.ReplaceRouteOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ReplaceRouteOutput
	if rf, ok := ret.Get(0).(func(*ec2.ReplaceRouteInput) *ec2.ReplaceRouteOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ReplaceRouteOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ReplaceRouteInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReplaceRouteRequest provides a mock function with given fields: _a0
func (_m *EC2API) ReplaceRouteRequest(_a0 *ec2.ReplaceRouteInput) (*request.Request, *ec2.ReplaceRouteOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ReplaceRouteInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ReplaceRouteOutput
	if rf, ok := ret.Get(1).(func(*ec2.ReplaceRouteInput) *ec2.ReplaceRouteOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ReplaceRouteOutput)
		}
	}

	return r0, r1
}

// ReplaceRouteTableAssociation provides a mock function with given fields: _a0
func (_m *EC2API) ReplaceRouteTableAssociation(_a0 *ec2.ReplaceRouteTableAssociationInput) (*ec2.ReplaceRouteTableAssociationOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ReplaceRouteTableAssociationOutput
	if rf, ok := ret.Get(0).(func(*ec2.ReplaceRouteTableAssociationInput) *ec2.ReplaceRouteTableAssociationOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ReplaceRouteTableAssociationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ReplaceRouteTableAssociationInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReplaceRouteTableAssociationRequest provides a mock function with given fields: _a0
func (_m *EC2API) ReplaceRouteTableAssociationRequest(_a0 *ec2.ReplaceRouteTableAssociationInput) (*request.Request, *ec2.ReplaceRouteTableAssociationOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ReplaceRouteTableAssociationInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ReplaceRouteTableAssociationOutput
	if rf, ok := ret.Get(1).(func(*ec2.ReplaceRouteTableAssociationInput) *ec2.ReplaceRouteTableAssociationOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ReplaceRouteTableAssociationOutput)
		}
	}

	return r0, r1
}

// ReplaceRouteTableAssociationWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) ReplaceRouteTableAssociationWithContext(_a0 aws.Context, _a1 *ec2.ReplaceRouteTableAssociationInput, _a2 ...request.Option) (*ec2.ReplaceRouteTableAssociationOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.ReplaceRouteTableAssociationOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.ReplaceRouteTableAssociationInput, ...request.Option) *ec2.ReplaceRouteTableAssociationOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ReplaceRouteTableAssociationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.ReplaceRouteTableAssociationInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReplaceRouteWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) ReplaceRouteWithContext(_a0 aws.Context, _a1 *ec2.ReplaceRouteInput, _a2 ...request.Option) (*ec2.ReplaceRouteOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.ReplaceRouteOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.ReplaceRouteInput, ...request.Option) *ec2.ReplaceRouteOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ReplaceRouteOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.ReplaceRouteInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReportInstanceStatus provides a mock function with given fields: _a0
func (_m *EC2API) ReportInstanceStatus(_a0 *ec2.ReportInstanceStatusInput) (*ec2.ReportInstanceStatusOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ReportInstanceStatusOutput
	if rf, ok := ret.Get(0).(func(*ec2.ReportInstanceStatusInput) *ec2.ReportInstanceStatusOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ReportInstanceStatusOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ReportInstanceStatusInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReportInstanceStatusRequest provides a mock function with given fields: _a0
func (_m *EC2API) ReportInstanceStatusRequest(_a0 *ec2.ReportInstanceStatusInput) (*request.Request, *ec2.ReportInstanceStatusOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ReportInstanceStatusInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ReportInstanceStatusOutput
	if rf, ok := ret.Get(1).(func(*ec2.ReportInstanceStatusInput) *ec2.ReportInstanceStatusOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ReportInstanceStatusOutput)
		}
	}

	return r0, r1
}

// ReportInstanceStatusWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) ReportInstanceStatusWithContext(_a0 aws.Context, _a1 *ec2.ReportInstanceStatusInput, _a2 ...request.Option) (*ec2.ReportInstanceStatusOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.ReportInstanceStatusOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.ReportInstanceStatusInput, ...request.Option) *ec2.ReportInstanceStatusOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ReportInstanceStatusOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.ReportInstanceStatusInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RequestSpotFleet provides a mock function with given fields: _a0
func (_m *EC2API) RequestSpotFleet(_a0 *ec2.RequestSpotFleetInput) (*ec2.RequestSpotFleetOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.RequestSpotFleetOutput
	if rf, ok := ret.Get(0).(func(*ec2.RequestSpotFleetInput) *ec2.RequestSpotFleetOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.RequestSpotFleetOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.RequestSpotFleetInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RequestSpotFleetRequest provides a mock function with given fields: _a0
func (_m *EC2API) RequestSpotFleetRequest(_a0 *ec2.RequestSpotFleetInput) (*request.Request, *ec2.RequestSpotFleetOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.RequestSpotFleetInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.RequestSpotFleetOutput
	if rf, ok := ret.Get(1).(func(*ec2.RequestSpotFleetInput) *ec2.RequestSpotFleetOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.RequestSpotFleetOutput)
		}
	}

	return r0, r1
}

// RequestSpotFleetWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) RequestSpotFleetWithContext(_a0 aws.Context, _a1 *ec2.RequestSpotFleetInput, _a2 ...request.Option) (*ec2.RequestSpotFleetOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.RequestSpotFleetOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.RequestSpotFleetInput, ...request.Option) *ec2.RequestSpotFleetOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.RequestSpotFleetOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.RequestSpotFleetInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RequestSpotInstances provides a mock function with given fields: _a0
func (_m *EC2API) RequestSpotInstances(_a0 *ec2.RequestSpotInstancesInput) (*ec2.RequestSpotInstancesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.RequestSpotInstancesOutput
	if rf, ok := ret.Get(0).(func(*ec2.RequestSpotInstancesInput) *ec2.RequestSpotInstancesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.RequestSpotInstancesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.RequestSpotInstancesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RequestSpotInstancesRequest provides a mock function with given fields: _a0
func (_m *EC2API) RequestSpotInstancesRequest(_a0 *ec2.RequestSpotInstancesInput) (*request.Request, *ec2.RequestSpotInstancesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.RequestSpotInstancesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.RequestSpotInstancesOutput
	if rf, ok := ret.Get(1).(func(*ec2.RequestSpotInstancesInput) *ec2.RequestSpotInstancesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.RequestSpotInstancesOutput)
		}
	}

	return r0, r1
}

// RequestSpotInstancesWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) RequestSpotInstancesWithContext(_a0 aws.Context, _a1 *ec2.RequestSpotInstancesInput, _a2 ...request.Option) (*ec2.RequestSpotInstancesOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.RequestSpotInstancesOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.RequestSpotInstancesInput, ...request.Option) *ec2.RequestSpotInstancesOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.RequestSpotInstancesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.RequestSpotInstancesInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResetFpgaImageAttribute provides a mock function with given fields: _a0
func (_m *EC2API) ResetFpgaImageAttribute(_a0 *ec2.ResetFpgaImageAttributeInput) (*ec2.ResetFpgaImageAttributeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ResetFpgaImageAttributeOutput
	if rf, ok := ret.Get(0).(func(*ec2.ResetFpgaImageAttributeInput) *ec2.ResetFpgaImageAttributeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ResetFpgaImageAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ResetFpgaImageAttributeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResetFpgaImageAttributeRequest provides a mock function with given fields: _a0
func (_m *EC2API) ResetFpgaImageAttributeRequest(_a0 *ec2.ResetFpgaImageAttributeInput) (*request.Request, *ec2.ResetFpgaImageAttributeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ResetFpgaImageAttributeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ResetFpgaImageAttributeOutput
	if rf, ok := ret.Get(1).(func(*ec2.ResetFpgaImageAttributeInput) *ec2.ResetFpgaImageAttributeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ResetFpgaImageAttributeOutput)
		}
	}

	return r0, r1
}

// ResetFpgaImageAttributeWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) ResetFpgaImageAttributeWithContext(_a0 aws.Context, _a1 *ec2.ResetFpgaImageAttributeInput, _a2 ...request.Option) (*ec2.ResetFpgaImageAttributeOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.ResetFpgaImageAttributeOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.ResetFpgaImageAttributeInput, ...request.Option) *ec2.ResetFpgaImageAttributeOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ResetFpgaImageAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.ResetFpgaImageAttributeInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResetImageAttribute provides a mock function with given fields: _a0
func (_m *EC2API) ResetImageAttribute(_a0 *ec2.ResetImageAttributeInput) (*ec2.ResetImageAttributeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ResetImageAttributeOutput
	if rf, ok := ret.Get(0).(func(*ec2.ResetImageAttributeInput) *ec2.ResetImageAttributeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ResetImageAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ResetImageAttributeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResetImageAttributeRequest provides a mock function with given fields: _a0
func (_m *EC2API) ResetImageAttributeRequest(_a0 *ec2.ResetImageAttributeInput) (*request.Request, *ec2.ResetImageAttributeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ResetImageAttributeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ResetImageAttributeOutput
	if rf, ok := ret.Get(1).(func(*ec2.ResetImageAttributeInput) *ec2.ResetImageAttributeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ResetImageAttributeOutput)
		}
	}

	return r0, r1
}

// ResetImageAttributeWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) ResetImageAttributeWithContext(_a0 aws.Context, _a1 *ec2.ResetImageAttributeInput, _a2 ...request.Option) (*ec2.ResetImageAttributeOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.ResetImageAttributeOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.ResetImageAttributeInput, ...request.Option) *ec2.ResetImageAttributeOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ResetImageAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.ResetImageAttributeInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResetInstanceAttribute provides a mock function with given fields: _a0
func (_m *EC2API) ResetInstanceAttribute(_a0 *ec2.ResetInstanceAttributeInput) (*ec2.ResetInstanceAttributeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ResetInstanceAttributeOutput
	if rf, ok := ret.Get(0).(func(*ec2.ResetInstanceAttributeInput) *ec2.ResetInstanceAttributeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ResetInstanceAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ResetInstanceAttributeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResetInstanceAttributeRequest provides a mock function with given fields: _a0
func (_m *EC2API) ResetInstanceAttributeRequest(_a0 *ec2.ResetInstanceAttributeInput) (*request.Request, *ec2.ResetInstanceAttributeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ResetInstanceAttributeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ResetInstanceAttributeOutput
	if rf, ok := ret.Get(1).(func(*ec2.ResetInstanceAttributeInput) *ec2.ResetInstanceAttributeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ResetInstanceAttributeOutput)
		}
	}

	return r0, r1
}

// ResetInstanceAttributeWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) ResetInstanceAttributeWithContext(_a0 aws.Context, _a1 *ec2.ResetInstanceAttributeInput, _a2 ...request.Option) (*ec2.ResetInstanceAttributeOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.ResetInstanceAttributeOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.ResetInstanceAttributeInput, ...request.Option) *ec2.ResetInstanceAttributeOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ResetInstanceAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.ResetInstanceAttributeInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResetNetworkInterfaceAttribute provides a mock function with given fields: _a0
func (_m *EC2API) ResetNetworkInterfaceAttribute(_a0 *ec2.ResetNetworkInterfaceAttributeInput) (*ec2.ResetNetworkInterfaceAttributeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ResetNetworkInterfaceAttributeOutput
	if rf, ok := ret.Get(0).(func(*ec2.ResetNetworkInterfaceAttributeInput) *ec2.ResetNetworkInterfaceAttributeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ResetNetworkInterfaceAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ResetNetworkInterfaceAttributeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResetNetworkInterfaceAttributeRequest provides a mock function with given fields: _a0
func (_m *EC2API) ResetNetworkInterfaceAttributeRequest(_a0 *ec2.ResetNetworkInterfaceAttributeInput) (*request.Request, *ec2.ResetNetworkInterfaceAttributeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ResetNetworkInterfaceAttributeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ResetNetworkInterfaceAttributeOutput
	if rf, ok := ret.Get(1).(func(*ec2.ResetNetworkInterfaceAttributeInput) *ec2.ResetNetworkInterfaceAttributeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ResetNetworkInterfaceAttributeOutput)
		}
	}

	return r0, r1
}

// ResetNetworkInterfaceAttributeWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) ResetNetworkInterfaceAttributeWithContext(_a0 aws.Context, _a1 *ec2.ResetNetworkInterfaceAttributeInput, _a2 ...request.Option) (*ec2.ResetNetworkInterfaceAttributeOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.ResetNetworkInterfaceAttributeOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.ResetNetworkInterfaceAttributeInput, ...request.Option) *ec2.ResetNetworkInterfaceAttributeOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ResetNetworkInterfaceAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.ResetNetworkInterfaceAttributeInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResetSnapshotAttribute provides a mock function with given fields: _a0
func (_m *EC2API) ResetSnapshotAttribute(_a0 *ec2.ResetSnapshotAttributeInput) (*ec2.ResetSnapshotAttributeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.ResetSnapshotAttributeOutput
	if rf, ok := ret.Get(0).(func(*ec2.ResetSnapshotAttributeInput) *ec2.ResetSnapshotAttributeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ResetSnapshotAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.ResetSnapshotAttributeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResetSnapshotAttributeRequest provides a mock function with given fields: _a0
func (_m *EC2API) ResetSnapshotAttributeRequest(_a0 *ec2.ResetSnapshotAttributeInput) (*request.Request, *ec2.ResetSnapshotAttributeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.ResetSnapshotAttributeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.ResetSnapshotAttributeOutput
	if rf, ok := ret.Get(1).(func(*ec2.ResetSnapshotAttributeInput) *ec2.ResetSnapshotAttributeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.ResetSnapshotAttributeOutput)
		}
	}

	return r0, r1
}

// ResetSnapshotAttributeWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) ResetSnapshotAttributeWithContext(_a0 aws.Context, _a1 *ec2.ResetSnapshotAttributeInput, _a2 ...request.Option) (*ec2.ResetSnapshotAttributeOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.ResetSnapshotAttributeOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.ResetSnapshotAttributeInput, ...request.Option) *ec2.ResetSnapshotAttributeOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.ResetSnapshotAttributeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.ResetSnapshotAttributeInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RestoreAddressToClassic provides a mock function with given fields: _a0
func (_m *EC2API) RestoreAddressToClassic(_a0 *ec2.RestoreAddressToClassicInput) (*ec2.RestoreAddressToClassicOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.RestoreAddressToClassicOutput
	if rf, ok := ret.Get(0).(func(*ec2.RestoreAddressToClassicInput) *ec2.RestoreAddressToClassicOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.RestoreAddressToClassicOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.RestoreAddressToClassicInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RestoreAddressToClassicRequest provides a mock function with given fields: _a0
func (_m *EC2API) RestoreAddressToClassicRequest(_a0 *ec2.RestoreAddressToClassicInput) (*request.Request, *ec2.RestoreAddressToClassicOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.RestoreAddressToClassicInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.RestoreAddressToClassicOutput
	if rf, ok := ret.Get(1).(func(*ec2.RestoreAddressToClassicInput) *ec2.RestoreAddressToClassicOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.RestoreAddressToClassicOutput)
		}
	}

	return r0, r1
}

// RestoreAddressToClassicWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) RestoreAddressToClassicWithContext(_a0 aws.Context, _a1 *ec2.RestoreAddressToClassicInput, _a2 ...request.Option) (*ec2.RestoreAddressToClassicOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.RestoreAddressToClassicOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.RestoreAddressToClassicInput, ...request.Option) *ec2.RestoreAddressToClassicOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.RestoreAddressToClassicOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.RestoreAddressToClassicInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RevokeSecurityGroupEgress provides a mock function with given fields: _a0
func (_m *EC2API) RevokeSecurityGroupEgress(_a0 *ec2.RevokeSecurityGroupEgressInput) (*ec2.RevokeSecurityGroupEgressOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.RevokeSecurityGroupEgressOutput
	if rf, ok := ret.Get(0).(func(*ec2.RevokeSecurityGroupEgressInput) *ec2.RevokeSecurityGroupEgressOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.RevokeSecurityGroupEgressOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.RevokeSecurityGroupEgressInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RevokeSecurityGroupEgressRequest provides a mock function with given fields: _a0
func (_m *EC2API) RevokeSecurityGroupEgressRequest(_a0 *ec2.RevokeSecurityGroupEgressInput) (*request.Request, *ec2.RevokeSecurityGroupEgressOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.RevokeSecurityGroupEgressInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.RevokeSecurityGroupEgressOutput
	if rf, ok := ret.Get(1).(func(*ec2.RevokeSecurityGroupEgressInput) *ec2.RevokeSecurityGroupEgressOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.RevokeSecurityGroupEgressOutput)
		}
	}

	return r0, r1
}

// RevokeSecurityGroupEgressWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) RevokeSecurityGroupEgressWithContext(_a0 aws.Context, _a1 *ec2.RevokeSecurityGroupEgressInput, _a2 ...request.Option) (*ec2.RevokeSecurityGroupEgressOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.RevokeSecurityGroupEgressOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.RevokeSecurityGroupEgressInput, ...request.Option) *ec2.RevokeSecurityGroupEgressOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.RevokeSecurityGroupEgressOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.RevokeSecurityGroupEgressInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}



// RevokeSecurityGroupIngressRequest provides a mock function with given fields: _a0
func (_m *EC2API) RevokeSecurityGroupIngressRequest(_a0 *ec2.RevokeSecurityGroupIngressInput) (*request.Request, *ec2.RevokeSecurityGroupIngressOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.RevokeSecurityGroupIngressInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.RevokeSecurityGroupIngressOutput
	if rf, ok := ret.Get(1).(func(*ec2.RevokeSecurityGroupIngressInput) *ec2.RevokeSecurityGroupIngressOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.RevokeSecurityGroupIngressOutput)
		}
	}

	return r0, r1
}

// RevokeSecurityGroupIngressWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) RevokeSecurityGroupIngressWithContext(_a0 aws.Context, _a1 *ec2.RevokeSecurityGroupIngressInput, _a2 ...request.Option) (*ec2.RevokeSecurityGroupIngressOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.RevokeSecurityGroupIngressOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.RevokeSecurityGroupIngressInput, ...request.Option) *ec2.RevokeSecurityGroupIngressOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.RevokeSecurityGroupIngressOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.RevokeSecurityGroupIngressInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RunInstances provides a mock function with given fields: _a0
func (_m *EC2API) RunInstances(_a0 *ec2.RunInstancesInput) (*ec2.Reservation, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.Reservation
	if rf, ok := ret.Get(0).(func(*ec2.RunInstancesInput) *ec2.Reservation); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.Reservation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.RunInstancesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RunInstancesRequest provides a mock function with given fields: _a0
func (_m *EC2API) RunInstancesRequest(_a0 *ec2.RunInstancesInput) (*request.Request, *ec2.Reservation) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.RunInstancesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.Reservation
	if rf, ok := ret.Get(1).(func(*ec2.RunInstancesInput) *ec2.Reservation); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.Reservation)
		}
	}

	return r0, r1
}

// RunInstancesWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) RunInstancesWithContext(_a0 aws.Context, _a1 *ec2.RunInstancesInput, _a2 ...request.Option) (*ec2.Reservation, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.Reservation
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.RunInstancesInput, ...request.Option) *ec2.Reservation); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.Reservation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.RunInstancesInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RunScheduledInstances provides a mock function with given fields: _a0
func (_m *EC2API) RunScheduledInstances(_a0 *ec2.RunScheduledInstancesInput) (*ec2.RunScheduledInstancesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.RunScheduledInstancesOutput
	if rf, ok := ret.Get(0).(func(*ec2.RunScheduledInstancesInput) *ec2.RunScheduledInstancesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.RunScheduledInstancesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.RunScheduledInstancesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RunScheduledInstancesRequest provides a mock function with given fields: _a0
func (_m *EC2API) RunScheduledInstancesRequest(_a0 *ec2.RunScheduledInstancesInput) (*request.Request, *ec2.RunScheduledInstancesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.RunScheduledInstancesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.RunScheduledInstancesOutput
	if rf, ok := ret.Get(1).(func(*ec2.RunScheduledInstancesInput) *ec2.RunScheduledInstancesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.RunScheduledInstancesOutput)
		}
	}

	return r0, r1
}

// RunScheduledInstancesWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) RunScheduledInstancesWithContext(_a0 aws.Context, _a1 *ec2.RunScheduledInstancesInput, _a2 ...request.Option) (*ec2.RunScheduledInstancesOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.RunScheduledInstancesOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.RunScheduledInstancesInput, ...request.Option) *ec2.RunScheduledInstancesOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.RunScheduledInstancesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.RunScheduledInstancesInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartInstances provides a mock function with given fields: _a0
func (_m *EC2API) StartInstances(_a0 *ec2.StartInstancesInput) (*ec2.StartInstancesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.StartInstancesOutput
	if rf, ok := ret.Get(0).(func(*ec2.StartInstancesInput) *ec2.StartInstancesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.StartInstancesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.StartInstancesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartInstancesRequest provides a mock function with given fields: _a0
func (_m *EC2API) StartInstancesRequest(_a0 *ec2.StartInstancesInput) (*request.Request, *ec2.StartInstancesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.StartInstancesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.StartInstancesOutput
	if rf, ok := ret.Get(1).(func(*ec2.StartInstancesInput) *ec2.StartInstancesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.StartInstancesOutput)
		}
	}

	return r0, r1
}

// StartInstancesWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) StartInstancesWithContext(_a0 aws.Context, _a1 *ec2.StartInstancesInput, _a2 ...request.Option) (*ec2.StartInstancesOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.StartInstancesOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.StartInstancesInput, ...request.Option) *ec2.StartInstancesOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.StartInstancesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.StartInstancesInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StopInstances provides a mock function with given fields: _a0
func (_m *EC2API) StopInstances(_a0 *ec2.StopInstancesInput) (*ec2.StopInstancesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.StopInstancesOutput
	if rf, ok := ret.Get(0).(func(*ec2.StopInstancesInput) *ec2.StopInstancesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.StopInstancesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.StopInstancesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StopInstancesRequest provides a mock function with given fields: _a0
func (_m *EC2API) StopInstancesRequest(_a0 *ec2.StopInstancesInput) (*request.Request, *ec2.StopInstancesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.StopInstancesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.StopInstancesOutput
	if rf, ok := ret.Get(1).(func(*ec2.StopInstancesInput) *ec2.StopInstancesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.StopInstancesOutput)
		}
	}

	return r0, r1
}

// StopInstancesWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) StopInstancesWithContext(_a0 aws.Context, _a1 *ec2.StopInstancesInput, _a2 ...request.Option) (*ec2.StopInstancesOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.StopInstancesOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.StopInstancesInput, ...request.Option) *ec2.StopInstancesOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.StopInstancesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.StopInstancesInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TerminateInstances provides a mock function with given fields: _a0
func (_m *EC2API) TerminateInstances(_a0 *ec2.TerminateInstancesInput) (*ec2.TerminateInstancesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.TerminateInstancesOutput
	if rf, ok := ret.Get(0).(func(*ec2.TerminateInstancesInput) *ec2.TerminateInstancesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.TerminateInstancesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.TerminateInstancesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TerminateInstancesRequest provides a mock function with given fields: _a0
func (_m *EC2API) TerminateInstancesRequest(_a0 *ec2.TerminateInstancesInput) (*request.Request, *ec2.TerminateInstancesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.TerminateInstancesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.TerminateInstancesOutput
	if rf, ok := ret.Get(1).(func(*ec2.TerminateInstancesInput) *ec2.TerminateInstancesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.TerminateInstancesOutput)
		}
	}

	return r0, r1
}

// TerminateInstancesWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) TerminateInstancesWithContext(_a0 aws.Context, _a1 *ec2.TerminateInstancesInput, _a2 ...request.Option) (*ec2.TerminateInstancesOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.TerminateInstancesOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.TerminateInstancesInput, ...request.Option) *ec2.TerminateInstancesOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.TerminateInstancesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.TerminateInstancesInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UnassignIpv6Addresses provides a mock function with given fields: _a0
func (_m *EC2API) UnassignIpv6Addresses(_a0 *ec2.UnassignIpv6AddressesInput) (*ec2.UnassignIpv6AddressesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.UnassignIpv6AddressesOutput
	if rf, ok := ret.Get(0).(func(*ec2.UnassignIpv6AddressesInput) *ec2.UnassignIpv6AddressesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.UnassignIpv6AddressesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.UnassignIpv6AddressesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UnassignIpv6AddressesRequest provides a mock function with given fields: _a0
func (_m *EC2API) UnassignIpv6AddressesRequest(_a0 *ec2.UnassignIpv6AddressesInput) (*request.Request, *ec2.UnassignIpv6AddressesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.UnassignIpv6AddressesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.UnassignIpv6AddressesOutput
	if rf, ok := ret.Get(1).(func(*ec2.UnassignIpv6AddressesInput) *ec2.UnassignIpv6AddressesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.UnassignIpv6AddressesOutput)
		}
	}

	return r0, r1
}

// UnassignIpv6AddressesWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) UnassignIpv6AddressesWithContext(_a0 aws.Context, _a1 *ec2.UnassignIpv6AddressesInput, _a2 ...request.Option) (*ec2.UnassignIpv6AddressesOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.UnassignIpv6AddressesOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.UnassignIpv6AddressesInput, ...request.Option) *ec2.UnassignIpv6AddressesOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.UnassignIpv6AddressesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.UnassignIpv6AddressesInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}



// UnassignPrivateIpAddressesRequest provides a mock function with given fields: _a0
func (_m *EC2API) UnassignPrivateIpAddressesRequest(_a0 *ec2.UnassignPrivateIpAddressesInput) (*request.Request, *ec2.UnassignPrivateIpAddressesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.UnassignPrivateIpAddressesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.UnassignPrivateIpAddressesOutput
	if rf, ok := ret.Get(1).(func(*ec2.UnassignPrivateIpAddressesInput) *ec2.UnassignPrivateIpAddressesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.UnassignPrivateIpAddressesOutput)
		}
	}

	return r0, r1
}

// UnassignPrivateIpAddressesWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) UnassignPrivateIpAddressesWithContext(_a0 aws.Context, _a1 *ec2.UnassignPrivateIpAddressesInput, _a2 ...request.Option) (*ec2.UnassignPrivateIpAddressesOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.UnassignPrivateIpAddressesOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.UnassignPrivateIpAddressesInput, ...request.Option) *ec2.UnassignPrivateIpAddressesOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.UnassignPrivateIpAddressesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.UnassignPrivateIpAddressesInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UnmonitorInstances provides a mock function with given fields: _a0
func (_m *EC2API) UnmonitorInstances(_a0 *ec2.UnmonitorInstancesInput) (*ec2.UnmonitorInstancesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.UnmonitorInstancesOutput
	if rf, ok := ret.Get(0).(func(*ec2.UnmonitorInstancesInput) *ec2.UnmonitorInstancesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.UnmonitorInstancesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.UnmonitorInstancesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UnmonitorInstancesRequest provides a mock function with given fields: _a0
func (_m *EC2API) UnmonitorInstancesRequest(_a0 *ec2.UnmonitorInstancesInput) (*request.Request, *ec2.UnmonitorInstancesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.UnmonitorInstancesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.UnmonitorInstancesOutput
	if rf, ok := ret.Get(1).(func(*ec2.UnmonitorInstancesInput) *ec2.UnmonitorInstancesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.UnmonitorInstancesOutput)
		}
	}

	return r0, r1
}

// UnmonitorInstancesWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) UnmonitorInstancesWithContext(_a0 aws.Context, _a1 *ec2.UnmonitorInstancesInput, _a2 ...request.Option) (*ec2.UnmonitorInstancesOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.UnmonitorInstancesOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.UnmonitorInstancesInput, ...request.Option) *ec2.UnmonitorInstancesOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.UnmonitorInstancesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.UnmonitorInstancesInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateSecurityGroupRuleDescriptionsEgress provides a mock function with given fields: _a0
func (_m *EC2API) UpdateSecurityGroupRuleDescriptionsEgress(_a0 *ec2.UpdateSecurityGroupRuleDescriptionsEgressInput) (*ec2.UpdateSecurityGroupRuleDescriptionsEgressOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.UpdateSecurityGroupRuleDescriptionsEgressOutput
	if rf, ok := ret.Get(0).(func(*ec2.UpdateSecurityGroupRuleDescriptionsEgressInput) *ec2.UpdateSecurityGroupRuleDescriptionsEgressOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.UpdateSecurityGroupRuleDescriptionsEgressOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.UpdateSecurityGroupRuleDescriptionsEgressInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateSecurityGroupRuleDescriptionsEgressRequest provides a mock function with given fields: _a0
func (_m *EC2API) UpdateSecurityGroupRuleDescriptionsEgressRequest(_a0 *ec2.UpdateSecurityGroupRuleDescriptionsEgressInput) (*request.Request, *ec2.UpdateSecurityGroupRuleDescriptionsEgressOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.UpdateSecurityGroupRuleDescriptionsEgressInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.UpdateSecurityGroupRuleDescriptionsEgressOutput
	if rf, ok := ret.Get(1).(func(*ec2.UpdateSecurityGroupRuleDescriptionsEgressInput) *ec2.UpdateSecurityGroupRuleDescriptionsEgressOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.UpdateSecurityGroupRuleDescriptionsEgressOutput)
		}
	}

	return r0, r1
}

// UpdateSecurityGroupRuleDescriptionsEgressWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) UpdateSecurityGroupRuleDescriptionsEgressWithContext(_a0 aws.Context, _a1 *ec2.UpdateSecurityGroupRuleDescriptionsEgressInput, _a2 ...request.Option) (*ec2.UpdateSecurityGroupRuleDescriptionsEgressOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.UpdateSecurityGroupRuleDescriptionsEgressOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.UpdateSecurityGroupRuleDescriptionsEgressInput, ...request.Option) *ec2.UpdateSecurityGroupRuleDescriptionsEgressOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.UpdateSecurityGroupRuleDescriptionsEgressOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.UpdateSecurityGroupRuleDescriptionsEgressInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateSecurityGroupRuleDescriptionsIngress provides a mock function with given fields: _a0
func (_m *EC2API) UpdateSecurityGroupRuleDescriptionsIngress(_a0 *ec2.UpdateSecurityGroupRuleDescriptionsIngressInput) (*ec2.UpdateSecurityGroupRuleDescriptionsIngressOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ec2.UpdateSecurityGroupRuleDescriptionsIngressOutput
	if rf, ok := ret.Get(0).(func(*ec2.UpdateSecurityGroupRuleDescriptionsIngressInput) *ec2.UpdateSecurityGroupRuleDescriptionsIngressOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.UpdateSecurityGroupRuleDescriptionsIngressOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ec2.UpdateSecurityGroupRuleDescriptionsIngressInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateSecurityGroupRuleDescriptionsIngressRequest provides a mock function with given fields: _a0
func (_m *EC2API) UpdateSecurityGroupRuleDescriptionsIngressRequest(_a0 *ec2.UpdateSecurityGroupRuleDescriptionsIngressInput) (*request.Request, *ec2.UpdateSecurityGroupRuleDescriptionsIngressOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ec2.UpdateSecurityGroupRuleDescriptionsIngressInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ec2.UpdateSecurityGroupRuleDescriptionsIngressOutput
	if rf, ok := ret.Get(1).(func(*ec2.UpdateSecurityGroupRuleDescriptionsIngressInput) *ec2.UpdateSecurityGroupRuleDescriptionsIngressOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ec2.UpdateSecurityGroupRuleDescriptionsIngressOutput)
		}
	}

	return r0, r1
}

// UpdateSecurityGroupRuleDescriptionsIngressWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) UpdateSecurityGroupRuleDescriptionsIngressWithContext(_a0 aws.Context, _a1 *ec2.UpdateSecurityGroupRuleDescriptionsIngressInput, _a2 ...request.Option) (*ec2.UpdateSecurityGroupRuleDescriptionsIngressOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.UpdateSecurityGroupRuleDescriptionsIngressOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.UpdateSecurityGroupRuleDescriptionsIngressInput, ...request.Option) *ec2.UpdateSecurityGroupRuleDescriptionsIngressOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.UpdateSecurityGroupRuleDescriptionsIngressOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *ec2.UpdateSecurityGroupRuleDescriptionsIngressInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WaitUntilBundleTaskComplete provides a mock function with given fields: _a0
func (_m *EC2API) WaitUntilBundleTaskComplete(_a0 *ec2.DescribeBundleTasksInput) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeBundleTasksInput) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilBundleTaskCompleteWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) WaitUntilBundleTaskCompleteWithContext(_a0 aws.Context, _a1 *ec2.DescribeBundleTasksInput, _a2 ...request.WaiterOption) error {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeBundleTasksInput, ...request.WaiterOption) error); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilConversionTaskCancelled provides a mock function with given fields: _a0
func (_m *EC2API) WaitUntilConversionTaskCancelled(_a0 *ec2.DescribeConversionTasksInput) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeConversionTasksInput) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilConversionTaskCancelledWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) WaitUntilConversionTaskCancelledWithContext(_a0 aws.Context, _a1 *ec2.DescribeConversionTasksInput, _a2 ...request.WaiterOption) error {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeConversionTasksInput, ...request.WaiterOption) error); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilConversionTaskCompleted provides a mock function with given fields: _a0
func (_m *EC2API) WaitUntilConversionTaskCompleted(_a0 *ec2.DescribeConversionTasksInput) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeConversionTasksInput) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilConversionTaskCompletedWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) WaitUntilConversionTaskCompletedWithContext(_a0 aws.Context, _a1 *ec2.DescribeConversionTasksInput, _a2 ...request.WaiterOption) error {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeConversionTasksInput, ...request.WaiterOption) error); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilConversionTaskDeleted provides a mock function with given fields: _a0
func (_m *EC2API) WaitUntilConversionTaskDeleted(_a0 *ec2.DescribeConversionTasksInput) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeConversionTasksInput) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilConversionTaskDeletedWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) WaitUntilConversionTaskDeletedWithContext(_a0 aws.Context, _a1 *ec2.DescribeConversionTasksInput, _a2 ...request.WaiterOption) error {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeConversionTasksInput, ...request.WaiterOption) error); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilCustomerGatewayAvailable provides a mock function with given fields: _a0
func (_m *EC2API) WaitUntilCustomerGatewayAvailable(_a0 *ec2.DescribeCustomerGatewaysInput) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeCustomerGatewaysInput) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilCustomerGatewayAvailableWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) WaitUntilCustomerGatewayAvailableWithContext(_a0 aws.Context, _a1 *ec2.DescribeCustomerGatewaysInput, _a2 ...request.WaiterOption) error {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeCustomerGatewaysInput, ...request.WaiterOption) error); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilExportTaskCancelled provides a mock function with given fields: _a0
func (_m *EC2API) WaitUntilExportTaskCancelled(_a0 *ec2.DescribeExportTasksInput) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeExportTasksInput) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilExportTaskCancelledWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) WaitUntilExportTaskCancelledWithContext(_a0 aws.Context, _a1 *ec2.DescribeExportTasksInput, _a2 ...request.WaiterOption) error {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeExportTasksInput, ...request.WaiterOption) error); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilExportTaskCompleted provides a mock function with given fields: _a0
func (_m *EC2API) WaitUntilExportTaskCompleted(_a0 *ec2.DescribeExportTasksInput) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeExportTasksInput) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilExportTaskCompletedWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) WaitUntilExportTaskCompletedWithContext(_a0 aws.Context, _a1 *ec2.DescribeExportTasksInput, _a2 ...request.WaiterOption) error {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeExportTasksInput, ...request.WaiterOption) error); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilImageAvailable provides a mock function with given fields: _a0
func (_m *EC2API) WaitUntilImageAvailable(_a0 *ec2.DescribeImagesInput) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeImagesInput) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilImageAvailableWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) WaitUntilImageAvailableWithContext(_a0 aws.Context, _a1 *ec2.DescribeImagesInput, _a2 ...request.WaiterOption) error {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeImagesInput, ...request.WaiterOption) error); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilImageExists provides a mock function with given fields: _a0
func (_m *EC2API) WaitUntilImageExists(_a0 *ec2.DescribeImagesInput) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeImagesInput) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilImageExistsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) WaitUntilImageExistsWithContext(_a0 aws.Context, _a1 *ec2.DescribeImagesInput, _a2 ...request.WaiterOption) error {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeImagesInput, ...request.WaiterOption) error); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilInstanceExists provides a mock function with given fields: _a0
func (_m *EC2API) WaitUntilInstanceExists(_a0 *ec2.DescribeInstancesInput) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeInstancesInput) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilInstanceExistsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) WaitUntilInstanceExistsWithContext(_a0 aws.Context, _a1 *ec2.DescribeInstancesInput, _a2 ...request.WaiterOption) error {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeInstancesInput, ...request.WaiterOption) error); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilInstanceRunning provides a mock function with given fields: _a0
func (_m *EC2API) WaitUntilInstanceRunning(_a0 *ec2.DescribeInstancesInput) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeInstancesInput) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilInstanceRunningWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) WaitUntilInstanceRunningWithContext(_a0 aws.Context, _a1 *ec2.DescribeInstancesInput, _a2 ...request.WaiterOption) error {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeInstancesInput, ...request.WaiterOption) error); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilInstanceStatusOk provides a mock function with given fields: _a0
func (_m *EC2API) WaitUntilInstanceStatusOk(_a0 *ec2.DescribeInstanceStatusInput) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeInstanceStatusInput) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilInstanceStatusOkWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) WaitUntilInstanceStatusOkWithContext(_a0 aws.Context, _a1 *ec2.DescribeInstanceStatusInput, _a2 ...request.WaiterOption) error {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeInstanceStatusInput, ...request.WaiterOption) error); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilInstanceStopped provides a mock function with given fields: _a0
func (_m *EC2API) WaitUntilInstanceStopped(_a0 *ec2.DescribeInstancesInput) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeInstancesInput) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilInstanceStoppedWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) WaitUntilInstanceStoppedWithContext(_a0 aws.Context, _a1 *ec2.DescribeInstancesInput, _a2 ...request.WaiterOption) error {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeInstancesInput, ...request.WaiterOption) error); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilInstanceTerminated provides a mock function with given fields: _a0
func (_m *EC2API) WaitUntilInstanceTerminated(_a0 *ec2.DescribeInstancesInput) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeInstancesInput) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilInstanceTerminatedWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) WaitUntilInstanceTerminatedWithContext(_a0 aws.Context, _a1 *ec2.DescribeInstancesInput, _a2 ...request.WaiterOption) error {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeInstancesInput, ...request.WaiterOption) error); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilKeyPairExists provides a mock function with given fields: _a0
func (_m *EC2API) WaitUntilKeyPairExists(_a0 *ec2.DescribeKeyPairsInput) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeKeyPairsInput) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilKeyPairExistsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) WaitUntilKeyPairExistsWithContext(_a0 aws.Context, _a1 *ec2.DescribeKeyPairsInput, _a2 ...request.WaiterOption) error {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeKeyPairsInput, ...request.WaiterOption) error); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilNatGatewayAvailable provides a mock function with given fields: _a0
func (_m *EC2API) WaitUntilNatGatewayAvailable(_a0 *ec2.DescribeNatGatewaysInput) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeNatGatewaysInput) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilNatGatewayAvailableWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) WaitUntilNatGatewayAvailableWithContext(_a0 aws.Context, _a1 *ec2.DescribeNatGatewaysInput, _a2 ...request.WaiterOption) error {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeNatGatewaysInput, ...request.WaiterOption) error); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilNetworkInterfaceAvailable provides a mock function with given fields: _a0
func (_m *EC2API) WaitUntilNetworkInterfaceAvailable(_a0 *ec2.DescribeNetworkInterfacesInput) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeNetworkInterfacesInput) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilNetworkInterfaceAvailableWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) WaitUntilNetworkInterfaceAvailableWithContext(_a0 aws.Context, _a1 *ec2.DescribeNetworkInterfacesInput, _a2 ...request.WaiterOption) error {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeNetworkInterfacesInput, ...request.WaiterOption) error); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilPasswordDataAvailable provides a mock function with given fields: _a0
func (_m *EC2API) WaitUntilPasswordDataAvailable(_a0 *ec2.GetPasswordDataInput) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.GetPasswordDataInput) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilPasswordDataAvailableWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) WaitUntilPasswordDataAvailableWithContext(_a0 aws.Context, _a1 *ec2.GetPasswordDataInput, _a2 ...request.WaiterOption) error {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.GetPasswordDataInput, ...request.WaiterOption) error); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilSnapshotCompleted provides a mock function with given fields: _a0
func (_m *EC2API) WaitUntilSnapshotCompleted(_a0 *ec2.DescribeSnapshotsInput) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSnapshotsInput) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilSnapshotCompletedWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) WaitUntilSnapshotCompletedWithContext(_a0 aws.Context, _a1 *ec2.DescribeSnapshotsInput, _a2 ...request.WaiterOption) error {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeSnapshotsInput, ...request.WaiterOption) error); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilSpotInstanceRequestFulfilled provides a mock function with given fields: _a0
func (_m *EC2API) WaitUntilSpotInstanceRequestFulfilled(_a0 *ec2.DescribeSpotInstanceRequestsInput) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSpotInstanceRequestsInput) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilSpotInstanceRequestFulfilledWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) WaitUntilSpotInstanceRequestFulfilledWithContext(_a0 aws.Context, _a1 *ec2.DescribeSpotInstanceRequestsInput, _a2 ...request.WaiterOption) error {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeSpotInstanceRequestsInput, ...request.WaiterOption) error); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilSubnetAvailable provides a mock function with given fields: _a0
func (_m *EC2API) WaitUntilSubnetAvailable(_a0 *ec2.DescribeSubnetsInput) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeSubnetsInput) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilSubnetAvailableWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) WaitUntilSubnetAvailableWithContext(_a0 aws.Context, _a1 *ec2.DescribeSubnetsInput, _a2 ...request.WaiterOption) error {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeSubnetsInput, ...request.WaiterOption) error); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilSystemStatusOk provides a mock function with given fields: _a0
func (_m *EC2API) WaitUntilSystemStatusOk(_a0 *ec2.DescribeInstanceStatusInput) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeInstanceStatusInput) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilSystemStatusOkWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) WaitUntilSystemStatusOkWithContext(_a0 aws.Context, _a1 *ec2.DescribeInstanceStatusInput, _a2 ...request.WaiterOption) error {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeInstanceStatusInput, ...request.WaiterOption) error); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilVolumeAvailable provides a mock function with given fields: _a0
func (_m *EC2API) WaitUntilVolumeAvailable(_a0 *ec2.DescribeVolumesInput) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVolumesInput) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilVolumeAvailableWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) WaitUntilVolumeAvailableWithContext(_a0 aws.Context, _a1 *ec2.DescribeVolumesInput, _a2 ...request.WaiterOption) error {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeVolumesInput, ...request.WaiterOption) error); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilVolumeDeleted provides a mock function with given fields: _a0
func (_m *EC2API) WaitUntilVolumeDeleted(_a0 *ec2.DescribeVolumesInput) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVolumesInput) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilVolumeDeletedWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) WaitUntilVolumeDeletedWithContext(_a0 aws.Context, _a1 *ec2.DescribeVolumesInput, _a2 ...request.WaiterOption) error {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeVolumesInput, ...request.WaiterOption) error); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilVolumeInUse provides a mock function with given fields: _a0
func (_m *EC2API) WaitUntilVolumeInUse(_a0 *ec2.DescribeVolumesInput) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVolumesInput) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilVolumeInUseWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) WaitUntilVolumeInUseWithContext(_a0 aws.Context, _a1 *ec2.DescribeVolumesInput, _a2 ...request.WaiterOption) error {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeVolumesInput, ...request.WaiterOption) error); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilVpcAvailable provides a mock function with given fields: _a0
func (_m *EC2API) WaitUntilVpcAvailable(_a0 *ec2.DescribeVpcsInput) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpcsInput) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilVpcAvailableWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) WaitUntilVpcAvailableWithContext(_a0 aws.Context, _a1 *ec2.DescribeVpcsInput, _a2 ...request.WaiterOption) error {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeVpcsInput, ...request.WaiterOption) error); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilVpcExists provides a mock function with given fields: _a0
func (_m *EC2API) WaitUntilVpcExists(_a0 *ec2.DescribeVpcsInput) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpcsInput) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilVpcExistsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) WaitUntilVpcExistsWithContext(_a0 aws.Context, _a1 *ec2.DescribeVpcsInput, _a2 ...request.WaiterOption) error {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeVpcsInput, ...request.WaiterOption) error); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilVpcPeeringConnectionDeleted provides a mock function with given fields: _a0
func (_m *EC2API) WaitUntilVpcPeeringConnectionDeleted(_a0 *ec2.DescribeVpcPeeringConnectionsInput) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpcPeeringConnectionsInput) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilVpcPeeringConnectionDeletedWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) WaitUntilVpcPeeringConnectionDeletedWithContext(_a0 aws.Context, _a1 *ec2.DescribeVpcPeeringConnectionsInput, _a2 ...request.WaiterOption) error {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeVpcPeeringConnectionsInput, ...request.WaiterOption) error); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilVpcPeeringConnectionExists provides a mock function with given fields: _a0
func (_m *EC2API) WaitUntilVpcPeeringConnectionExists(_a0 *ec2.DescribeVpcPeeringConnectionsInput) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpcPeeringConnectionsInput) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilVpcPeeringConnectionExistsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) WaitUntilVpcPeeringConnectionExistsWithContext(_a0 aws.Context, _a1 *ec2.DescribeVpcPeeringConnectionsInput, _a2 ...request.WaiterOption) error {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeVpcPeeringConnectionsInput, ...request.WaiterOption) error); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilVpnConnectionAvailable provides a mock function with given fields: _a0
func (_m *EC2API) WaitUntilVpnConnectionAvailable(_a0 *ec2.DescribeVpnConnectionsInput) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpnConnectionsInput) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilVpnConnectionAvailableWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) WaitUntilVpnConnectionAvailableWithContext(_a0 aws.Context, _a1 *ec2.DescribeVpnConnectionsInput, _a2 ...request.WaiterOption) error {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeVpnConnectionsInput, ...request.WaiterOption) error); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilVpnConnectionDeleted provides a mock function with given fields: _a0
func (_m *EC2API) WaitUntilVpnConnectionDeleted(_a0 *ec2.DescribeVpnConnectionsInput) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ec2.DescribeVpnConnectionsInput) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilVpnConnectionDeletedWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *EC2API) WaitUntilVpnConnectionDeletedWithContext(_a0 aws.Context, _a1 *ec2.DescribeVpnConnectionsInput, _a2 ...request.WaiterOption) error {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *ec2.DescribeVpnConnectionsInput, ...request.WaiterOption) error); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
