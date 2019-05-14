// Copyright 2017 Bryan Jeal <bryan@jeal.ca>

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
	"path"
	"path/filepath"
	"strings"
)

// Path type basically namespaces the helpers under the "pth" section
var Path pth

type pth struct{}

func (p pth) StopEscape(str string) string {
	str = path.Clean(str)
	str = strings.TrimPrefix(str, "../")
	if path.IsAbs(str) {
		return strings.TrimPrefix(str, "/")
	}
	return str
}

func (p pth) ReplaceExt(filename, ext string) string {
	ext = strings.Trim(ext, ".")
	return strings.TrimSuffix(filename, filepath.Ext(filename)) + "." + ext
}

func (p pth) AppendBeforeExt(filename, str string) string {
	ext := filepath.Ext(filename)
	base := strings.TrimSuffix(filename, ext)
	return base + str + ext
}
