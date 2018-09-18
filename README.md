# aws-mock
Mock library for AWS go SDK 

It is in WIP. The Primary focus would be mocking the ec2 service.

# Contribution Guideline
 - If you are planing to send a PR, please raise a issue first and discuss with us, before sending the code.
 - Feel free to send PR. Happy Coding :)

 # Usage

 ```go
 package main
 import(
     mockec2 "github.com/avi-internal/cloud/awsmock/service/ec2"
 )

 func main() {
     mockedEC2 := mockec2.New()
     // Few helpers are built to interact with the client
     // In order to deal with the helpers, you have to do the initial seeding
     mockedEC2.InitialSeeding()
     defaultVPCID:= mockedEC2.GetDefaultVPCID()
     // New gives you the instance of the ec2 api interface
     // Then you can interact as normal go aws ec2 sdk
     vpc, _ := mockedEC2.DescribeVpcs(&ec2.DescribeVpcsInput{
		VpcIds: []*string{
			&defaultVpcID,
		},
    })
    fmt.Println(vpc)
 }
 ```