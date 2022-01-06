/*
 * Copyright 2012-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package util_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/go-spring/spring-base/assert"
	"github.com/go-spring/spring-base/util"
	"github.com/go-spring/spring-base/util/testdata"
)

func fnNoArgs() {}

func fnWithArgs(i int) {}

type receiver struct{}

func (r receiver) fnNoArgs() {}

func (r receiver) fnWithArgs(i int) {}

func (r *receiver) ptrFnNoArgs() {}

func (r *receiver) ptrFnWithArgs(i int) {}

func TestFileLine(t *testing.T) {

	fmt.Println(util.FileLine(fnNoArgs))
	fmt.Println(util.FileLine(fnWithArgs))
	fmt.Println(util.FileLine(receiver{}.fnNoArgs))
	fmt.Println(util.FileLine(receiver{}.fnWithArgs))
	fmt.Println(util.FileLine((&receiver{}).ptrFnNoArgs))
	fmt.Println(util.FileLine((&receiver{}).ptrFnWithArgs))
	fmt.Println(util.FileLine(receiver.fnNoArgs))
	fmt.Println(util.FileLine(receiver.fnWithArgs))
	fmt.Println(util.FileLine((*receiver).ptrFnNoArgs))
	fmt.Println(util.FileLine((*receiver).ptrFnWithArgs))

	fmt.Println(util.FileLine(testdata.FnNoArgs))
	fmt.Println(util.FileLine(testdata.FnWithArgs))
	fmt.Println(util.FileLine(testdata.Receiver{}.FnNoArgs))
	fmt.Println(util.FileLine(testdata.Receiver{}.FnWithArgs))
	fmt.Println(util.FileLine((&testdata.Receiver{}).PtrFnNoArgs))
	fmt.Println(util.FileLine((&testdata.Receiver{}).PtrFnWithArgs))
	fmt.Println(util.FileLine(testdata.Receiver.FnNoArgs))
	fmt.Println(util.FileLine(testdata.Receiver.FnWithArgs))
	fmt.Println(util.FileLine((*testdata.Receiver).PtrFnNoArgs))
	fmt.Println(util.FileLine((*testdata.Receiver).PtrFnWithArgs))
}

func TestIsNil(t *testing.T) {

	a := (*int)(nil)        // %T == *int
	b := (interface{})(nil) // %T == <nil>
	assert.False(t, a == b)

	v := reflect.ValueOf((*int)(nil)) // %T == *int
	assert.True(t, util.IsNil(v))

	assert.True(t, util.IsNil(reflect.Value{}))              // %T == <nil>
	assert.True(t, util.IsNil(reflect.ValueOf(nil)))         // %T == <nil>
	assert.True(t, util.IsNil(reflect.ValueOf((*int)(nil)))) // %T == *int

	assert.False(t, util.IsNil(reflect.ValueOf(3)))
	assert.False(t, util.IsNil(reflect.ValueOf("3")))
}
