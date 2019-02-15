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

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
func rangeIn(low, hi int) string {
	return strconv.Itoa(low + rand.Intn(hi-low))
}

func rangeInInt(low, hi int) int {
	//rand.Seed(time.Now().Unix())
	return low + rand.Intn(hi-low)
}

// https://codereview.stackexchange.com/questions/60074/in-array-in-go
func in_array(val string, array []string) (exists bool, index int) {
	exists = false
	index = -1

	for i, v := range array {
		if val == v {
			index = i
			exists = true
			return
		}
	}

	return
}

// https://gist.github.com/kotakanbe/d3059af990252ba89a82
func Hosts(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}
	// remove network address and broadcast address
	return ips[1 : len(ips)-1], nil
}

// https://gist.github.com/kotakanbe/d3059af990252ba89a82
//  http://play.golang.org/p/m8TNTtygK0
func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func PickRandomHostFromCIDR(cidr string) (host string, err error) {

	hosts, err := Hosts(cidr)
	if err != nil {
		return
	}
	// first 4 and last 1 is reserved by aws
	n := rangeInInt(4, len(hosts)-1)
	host = hosts[n]
	return
}

// https://stackoverflow.com/questions/21018729/generate-mac-address-in-go
func GiveRandMacAddress() (string, error) {
	buf := make([]byte, 6)
	_, err := rand.Read(buf)
	if err != nil {
		return "", err
	}
	// Set the local bit
	buf[0] |= 2

	return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", buf[0], buf[1], buf[2], buf[3], buf[4], buf[5]), nil
}

func GiveRandomId(prefix string) string {
	buf := make([]byte, 6)
	_, err := rand.Read(buf)
	if err != nil {
		return prefix + "1234"
	}
	rsp := prefix
	for _, b := range buf {
		rsp = fmt.Sprintf(rsp+"%x", b)
	}
	return rsp
}
