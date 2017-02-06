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
	"context"
	"errors"
	"net/http"
)

// Errors
var (
	ErrCtxNoValue = errors.New("no value exists in the context for the supplied key")
)

// Ctx type basically namespaces the helpers under the "ctx" section
var Ctx ctx

type ctx struct {
	Http ctxHttp
}

type ctxHttp struct{}

func (h ctxHttp) CtxGet(r *http.Request, key string) (interface{}, error) {
	val := r.Context().Value(key)
	if val == nil {
		return nil, ErrCtxNoValue
	}

	return val, nil
}

func (h ctxHttp) CtxSave(r *http.Request, key string, val interface{}) *http.Request {
	ctx := r.Context()
	ctx = context.WithValue(ctx, key, val)
	return r.WithContext(ctx)
}
