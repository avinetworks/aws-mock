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
package main

import (
	"fmt"

	mockec2 "github.com/avi-internal/cloud/awsmock/service/ec2"
	ec2 "github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	mockedEC2 := mockec2.New()
	defaultVpcID := "avi-seeding-vpc"
	defaultCidrBlock := "10.0.0.0/16"
	defaultVpcState := "available"
	mockedEC2.AppendVpcs(&ec2.Vpc{
		VpcId:     &defaultVpcID,
		CidrBlock: &defaultCidrBlock,
		State:     &defaultVpcState,
	})
	vpc, err := mockedEC2.DescribeVpcs(&ec2.DescribeVpcsInput{
		VpcIds: []*string{
			&defaultVpcID,
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(vpc)

	vpc, err = mockedEC2.DescribeVpcs(&ec2.DescribeVpcsInput{
		VpcIds: []*string{
			&defaultVpcID,
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(vpc)
	vpc, err = mockedEC2.DescribeVpcs(&ec2.DescribeVpcsInput{
		VpcIds: []*string{
			&defaultVpcID,
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(vpc)
	vpc, err = mockedEC2.DescribeVpcs(&ec2.DescribeVpcsInput{
		VpcIds: []*string{
			&defaultVpcID,
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(vpc)
}
