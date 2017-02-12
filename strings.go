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
	"regexp"
	"strings"

	"github.com/rainycape/unidecode"
)

// Strings type basically namespaces the helpers under the "strs" section
var Strings = strs{
	reMultiSeparators:  regexp.MustCompile("[\\s-_]+"),
	reNotWordSpaceDash: regexp.MustCompile("[^\\w\\s-]"),
}

type strs struct {
	// Regexp used by Slugify() to remove multiple spaces, underscores, and dashes
	reMultiSeparators *regexp.Regexp
	// Regexp used by Slugify() to remove runes that are not words, spaces, or dashes
	reNotWordSpaceDash *regexp.Regexp
}

func (s strs) Slugify(str string) string {
	// Given an unicode encoded string, returns
	// another string with non-ASCII characters replaced
	// with their closest ASCII counterparts.
	// e.g. Unicode("áéíóú") => "aeiou"
	result := unidecode.Unidecode(str)

	// remove runes that aren't words, spaces, dashes
	result = s.reNotWordSpaceDash.ReplaceAllString(result, "-")
	// make lowercase
	result = strings.ToLower(result)
	// remove multiple spaces and dashes
	result = s.reMultiSeparators.ReplaceAllString(result, "-")
	// return string with leading and trailing dashes removed
	return strings.Trim(result, "-")
}
