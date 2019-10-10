// Copyright 2019 Nebularis Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	"strings"
	"testing"
)

func compare(actual, expected string) bool {
	return trim(actual) == trim(expected)
}

func trim(s string) string {
	return strings.TrimSpace(strings.Replace(strings.Replace(s, "\t", "", -1), " ", "", -1))
}

func log(t *testing.T, context, actual, expected string) {
	t.Logf("----")
	t.Logf("Expected %s: \n%s\n", context, expected)
	t.Logf("----")
	t.Logf("Actual %s: \n%s\n", context, actual)
	t.Logf("----")
}

//func annotateSpaces(s string) string {
//	s = strings.Replace(s, " ", ".", -1)
//	s = strings.Replace(s, "\r", "\\r", -1)
//	s = strings.Replace(s, "\n", "\\n", -1)
//	return s
//}
