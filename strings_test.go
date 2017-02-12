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

import "testing"

func TestSlugify(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"áéíóú", "aeiou"},
		{"DOBROSLAWZYBORT", "dobroslawzybort"},
		{"Dobroslaw Zybort", "dobroslaw-zybort"},
		{"  Dobroslaw     Zybort  ?", "dobroslaw-zybort"},
		{"Dobrosław Żybort", "dobroslaw-zybort"},
		{"Ala ma 6 kotów.", "ala-ma-6-kotow"},
		{"áÁàÀãÃâÂäÄąĄą̊Ą̊", "aaaaaaaaaaaaaa"},
		{"ćĆĉĈçÇ", "cccccc"},
		{"éÉèÈẽẼêÊëËęĘ", "eeeeeeeeeeee"},
		{"íÍìÌĩĨîÎïÏįĮ", "iiiiiiiiiiii"},
		{"łŁ", "ll"},
		{"ńŃ", "nn"},
		{"óÓòÒõÕôÔöÖǫǪǭǬø", "ooooooooooooooo"},
		{"śŚ", "ss"},
		{"úÚùÙũŨûÛüÜųŲ", "uuuuuuuuuuuu"},
		{"y̨Y̨", "yy"},
		{"źŹżŹ", "zzzz"},
		{"·/,:;`˜'\"", ""},
		{"2000–2013", "2000-2013"},
		{"style—not", "style-not"},
		{"test_slug", "test-slug"},
		{"Æ", "ae"},
		{"Ich heiße", "ich-heisse"},
		{"fácil €", "facil-eu"},
		{"smile ☺", "smile"},
		{"Hellö Wörld хелло ворлд", "hello-world-khello-vorld"},
		{"\"C'est déjà l’été.\"", "c-est-deja-l-ete"},
		{"jaja---lol-méméméoo--a", "jaja-lol-mememeoo-a"},
		{"影師", "ying-shi"},
		{"  Foo bar  ", "foo-bar"},
		{"Foo.Bar/foo_Bar-Foo", "foo-bar-foo-bar-foo"},
		{"fOO,bar:foobAR", "foo-bar-foobar"},
		{"FOo/BaR.html", "foo-bar-html"},
		{"трям/трям", "triam-triam"},
		{"은행", "eunhaeng"},
		{"Банковский кассир", "bankovskii-kassir"},
		{"♪ Short musical interlude ♬ ", "short-musical-interlude"},
		{"Foo & bar", "foo-bar"},
		{"संस्कृत", "snskrt"},
		{"a%C3%B1ame", "a-c3-b1ame"},
		{"this+is+a+test", "this-is-a-test"},
		{"~foo", "foo"},
	}

	for i, test := range tests {
		output := Strings.Slugify(test.input)
		if output != test.expected {
			t.Errorf("Test: %d: Expected %#v, got %#v\n", i, test.expected, output)
		}
	}
}
