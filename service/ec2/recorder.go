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
	"errors"
	"fmt"
	"reflect"
	"sync"
	"sync/atomic"
	"time"
)

type ApiErrorByCount struct {
	errorStr string
	count    int64
}

type methodArgs struct {
	Input  int
	Return int
}

type expectedArgs struct {
	Input  []interface{}
	Return []interface{}
}
type Recorder struct {
	totalApiCall                          int64
	countByApiName                        map[string]int64 //key will be api name
    activeByApiName                       map[string]bool
	giveForTimeOutError                   bool
	giveErrorByApiName                    map[string]ApiErrorByCount // key will be api name
	giveErrorNow                          bool
	globalErrorString                     string
	apiErrorString                        map[string]string
	flagApiCountSet                       bool
	flaggedApiCount                       int64
	methodArgs                            map[string]methodArgs
	expectedInput                         map[string]expectedArgs
	NthErrorCheck                         map[string]int64
	assignPrivateIpFailNetworkInterfaceId string
	delay                                 time.Duration
	delayByApiName                        map[string]time.Duration
	sync.Mutex
}

func (r *Recorder) init() {
	calcVar := &EC2API{}
	ec2Type := reflect.TypeOf(calcVar)
	totalMethods := ec2Type.NumMethod()
	r.methodArgs = make(map[string]methodArgs, 0)
	r.expectedInput = make(map[string]expectedArgs, 0)
	r.countByApiName = make(map[string]int64, 0)
	r.delayByApiName = make(map[string]time.Duration, 0)
	r.activeByApiName = make(map[string]bool, 0)
	r.apiErrorString = make(map[string]string, 0)
	r.giveErrorByApiName = make(map[string]ApiErrorByCount, 0)
	r.giveErrorNow = false
	r.flagApiCountSet = false
	r.totalApiCall = 0
	r.flaggedApiCount = 0
	r.globalErrorString = "dummy error, please fill with your's :)"
	r.delay = time.Second * 0
	r.NthErrorCheck = make(map[string]int64)
	r.assignPrivateIpFailNetworkInterfaceId = ""
	for i := 0; i < totalMethods; i++ {
		r.methodArgs[ec2Type.Method(i).Name] = methodArgs{
			Input:  ec2Type.Method(i).Type.NumIn() - 1,
			Return: ec2Type.Method(i).Type.NumOut(),
		}
        r.countByApiName[ec2Type.Method(i).Name] = 0
        r.delayByApiName[ec2Type.Method(i).Name] = time.Second * 0
        r.activeByApiName[ec2Type.Method(i).Name] = true
        r.apiErrorString[ec2Type.Method(i).Name] = "mocked error"
	}
}

type ReturnExpecter struct {
	recorder    *Recorder
	ReturnCount int
	Input       []interface{}
	methodName  string
}

func (r *Recorder) SetApiActive(apiName string) {
    r.activeByApiName[apiName] = true
    r.apiErrorString[apiName] = "mocked error"
}

func (r *Recorder) SetApiInactive(apiName string, errorStr string) {
    r.activeByApiName[apiName] = false
    r.apiErrorString[apiName] = errorStr
}

func (r *Recorder) SetError(errorStr string) {
	r.globalErrorString = errorStr
}

func (r *Recorder) SetAssignIpFail(nicID string) {
	r.assignPrivateIpFailNetworkInterfaceId = nicID
}

func (r *Recorder) GetAssignIpFailNetworkInterfaceId() string {
	return r.assignPrivateIpFailNetworkInterfaceId
}
func (re *ReturnExpecter) Return(args ...interface{}) {
	if re.ReturnCount != len(args) {
		panic(fmt.Sprintf("expected output arg %d and got %d", re.ReturnCount, len(args)))
	}
	// calcVar := &EC2API{}
	// ec2Type := reflect.TypeOf(calcVar)
	//method, ok := ec2Type.MethodByName(re.methodName)
	// if !ok {
	// 	panic(fmt.Sprintf("%s method doesn't exist", re.methodName))
	// }
	// for i := 0; i < len(args); i++ {
	// 	if method.Type.Out(i) != reflect.TypeOf(args[i]) {
	// 		panic("return argument is not of method's return type")
	// 	}
	// }
	expectedArguments := expectedArgs{
		Input:  re.Input,
		Return: args,
	}
	re.recorder.expectedInput[re.methodName] = expectedArguments
}

func (r *Recorder) Input(methodName string, args ...interface{}) *ReturnExpecter {
	val, ok := r.methodArgs[methodName]
	if !ok {
		panic(fmt.Sprintf("%s is not in the ec2 api method", methodName))
	}
	if len(args) != val.Input {
		panic(fmt.Sprintf("expected input arg %d and got %d", val.Input, len(args)))
	}
	return &ReturnExpecter{
		methodName:  methodName,
		Input:       args,
		recorder:    r,
		ReturnCount: val.Return,
	}
}

func (r *Recorder) giveRecordedOutput(methodName string, args ...interface{}) ([]interface{}, bool) {
	expectedArguments, ok := r.expectedInput[methodName]
	if !ok {
		return []interface{}{}, false
	}
	for i, val := range args {
		if !reflect.DeepEqual(val, expectedArguments.Input[i]) {
			return []interface{}{}, false
		}
	}
	return expectedArguments.Return, true
}

func (r *Recorder) GiveErrorNthTime(apiName string, nth int64) {
	r.NthErrorCheck[apiName] = nth
}

func (r *Recorder) Record(apiName string) {
	atomic.AddInt64(&r.totalApiCall, 1)
	r.Lock()
	defer r.Unlock()
	_, ok := r.countByApiName[apiName]
	if !ok {
		r.countByApiName[apiName] = 0
	}
	r.countByApiName[apiName] = r.countByApiName[apiName] + 1
}

func (r *Recorder) GiveErrorByTimeOut(t time.Duration, errorString string) {
	go func() {
		time.Sleep(t)
		r.Lock()
		defer r.Unlock()
		r.giveErrorNow = true
		r.globalErrorString = errorString
	}()
}

func (r *Recorder) GiveErrorByApiNameByTimeOut(apiName string, t time.Duration, errorString string) {
	go func() {
		time.Sleep(t)
		r.Lock()
		defer r.Unlock()
		r.activeByApiName[apiName] = false
		r.apiErrorString[apiName] = errorString
	}()
}

func (r *Recorder) GiveErrorNow(errorStr string) {
	r.giveErrorNow = true
	r.globalErrorString = errorStr
}

func (r *Recorder) GiveErrorForApiByTime(apiName, errorString string, t time.Duration) {
	go func() {
		time.Sleep(t)
		r.Lock()
		defer r.Unlock()
		r.giveErrorByApiName[apiName] = ApiErrorByCount{
			count:    0,
			errorStr: errorString,
		}
	}()
}

func (r *Recorder) GiveErrorByApiNameCount(apiName string, count int64, errorString string) {
	r.Lock()
	defer r.Unlock()
	r.giveErrorByApiName[apiName] = ApiErrorByCount{
		count:    count,
		errorStr: errorString,
	}
}

func (r *Recorder) GiveErrorByApiCount(limit int64) {
	r.flagApiCountSet = true
	atomic.StoreInt64(&r.flaggedApiCount, limit)
}

func (r *Recorder) SetDelay(duration time.Duration) {
	r.delay = duration
}

func (r *Recorder) SetDelayByApiName(apiName string, duration time.Duration) {
    r.delayByApiName[apiName] = duration
}

func (r *Recorder) CheckError(apiName string) error {
    time.Sleep(r.delay)
    time.Sleep(r.delayByApiName[apiName])
	if r.giveErrorNow == true {
		return errors.New(r.globalErrorString)
	}

    if r.activeByApiName[apiName] == false {
        return errors.New(r.apiErrorString[apiName])
    }

	if r.flagApiCountSet {
		if r.totalApiCall >= r.flaggedApiCount {
			return errors.New(r.globalErrorString)
		}
	}

	nth, ok := r.NthErrorCheck[apiName]

	if ok {
		countByApiName, ok := r.countByApiName[apiName]
		if ok {
			if countByApiName%nth == 0 {
				r.countByApiName[apiName] = r.countByApiName[apiName] + 1
				return errors.New(r.globalErrorString)
			}
		}
	}

	val, ok := r.giveErrorByApiName[apiName]

	if ok {
		countByApiName, ok := r.countByApiName[apiName]
		if ok {
			if countByApiName >= val.count {
				return errors.New(val.errorStr)
			}
		}
	}
	return nil
}
