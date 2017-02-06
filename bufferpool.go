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
	"bytes"
	"sync"
)

// BufferPool type basically namespaces the helpers under the "bufferPool" section
var BufferPool bufferPool

type bufferPool struct {
	pool *sync.Pool
}

// GetBuffer returns a buffer from the pool.
func (bp bufferPool) Get() (buf *bytes.Buffer) {
	return bp.pool.Get().(*bytes.Buffer)
}

// PutBuffer returns a buffer to the pool.
// The buffer is reset before it is put back into circulation.
func (bp bufferPool) Put(buf *bytes.Buffer) {
	buf.Reset()
	bp.pool.Put(buf)
}

func init() {
	BufferPool = bufferPool{
		&sync.Pool{
			New: func() interface{} {
				return &bytes.Buffer{}
			},
		},
	}
}
