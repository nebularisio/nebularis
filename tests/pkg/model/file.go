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
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"testing"
)

func isInputFile(name string) bool {
	return path.Ext(name) == ".ns"
}

func parseInputName(s string) (baseName string, index *int, err error) {
	baseName = s[:len(s)-len(".ns")]

	indexExt := path.Ext(baseName)
	if indexSuffixPattern.MatchString(indexExt) {
		baseName = baseName[:len(baseName)-len(indexExt)]
		var idx int
		if idx, err = strconv.Atoi(indexExt[1:]); err != nil {
			return
		}
		index = &idx
	}

	return
}

func readFile(name string) (string, error) {
	by, err := ioutil.ReadFile(name)
	if err != nil {
		return "", err
	}

	return string(by), nil
}

func tryReadFile(name string) (*string, error) {
	by, err := ioutil.ReadFile(name)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}

	s := string(by)
	return &s, nil
}

func writeFile(name, contents string) error {
	return ioutil.WriteFile(name, []byte(contents), os.ModePerm)
}

func writeFileOrFail(t *testing.T, name, contents string) {
	if err := writeFile(name, contents); err != nil {
		t.Fatalf("error writing output file %q: %v", name, err)
	}
}
