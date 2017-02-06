// Copyright 2016 Bryan Jeal <bryan@jeal.ca>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// 	http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package helpers

import (
	"reflect"
	"time"
)

// Reflect type basically namespaces the helpers under the "reflect" section
var Reflect = rflct{
	ErrorType: reflect.TypeOf((*error)(nil)).Elem(),
	TimeType:  reflect.TypeOf((*time.Time)(nil)).Elem(),
}

type rflct struct {
	// Zero is an empty reflect.Value
	Zero reflect.Value
	// ErrorType is TypeOf error
	ErrorType reflect.Type

	// TimeType is TypeOf time.Time
	TimeType reflect.Type
}

// toInt returns the int value if possible, -1 if not.
func (r rflct) ToInt(v reflect.Value) int64 {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int()
	case reflect.Interface:
		return r.ToInt(v.Elem())
	}
	return -1
}

// toString returns the string value if possible, "" if not.
func (r rflct) ToString(v reflect.Value) string {
	switch v.Kind() {
	case reflect.String:
		return v.String()
	case reflect.Interface:
		return r.ToString(v.Elem())
	}
	return ""
}

func (r rflct) ToTimeUnix(v reflect.Value) int64 {
	if v.Kind() == reflect.Interface {
		return r.ToTimeUnix(v.Elem())
	}
	if v.Type() != r.TimeType {
		panic("coding error: argument must be time.Time type reflect Value")
	}
	return v.MethodByName("Unix").Call([]reflect.Value{})[0].Int()
}

// indirect is taken from 'text/template/exec.go'
func (r rflct) Indirect(v reflect.Value) (rv reflect.Value, isNil bool) {
	for ; v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface; v = v.Elem() {
		if v.IsNil() {
			return v, true
		}
		if v.Kind() == reflect.Interface && v.NumMethod() > 0 {
			break
		}
	}
	return v, false
}
