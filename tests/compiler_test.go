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

package tests

import (
	"flag"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"nebularis.io/nebularis/tests/pkg/model"
)

func TestCompiler(t *testing.T) {
	flag.Parse()
	t.Logf("UpdateBaselines: %v", model.UpdateBaselines())
	outfolder := model.OutFolder()
	if outfolder == "" {
		var err error
		if outfolder, err = ioutil.TempDir(os.TempDir(), "TestCompiler"); err != nil {
			t.Fatalf("error creating temporary folder: %v", err)
		}
	}

	t.Logf("Output Temp folder: %v", outfolder)
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Unable get current working dir: %v", err)
	}
	t.Logf("${DIFF} %s/testdata/compiler %s", cwd, outfolder)
	runTests(t, "testdata/compiler", outfolder)
}

func runTests(t *testing.T, folder, tmpfolder string) {
	runSubTests(t, folder, tmpfolder)

	cases, err := model.Load(folder)
	if err != nil {
		t.Fatalf("Error loading test cases from %q: %v", folder, err)
	}

	for _, tc := range cases {
		tc.Run(t, tmpfolder)
	}
}

func runSubTests(t *testing.T, folder, outFolder string) {
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		t.Fatalf("Reading directory failed: %v", err)
	}

	for _, f := range files {
		if f.IsDir() {
			t.Run(f.Name(), func(t *testing.T) {
				subFolder := path.Join(folder, f.Name())
				subOutFolder := path.Join(outFolder, f.Name())
				err := os.MkdirAll(subOutFolder, os.ModePerm)
				if err != nil {
					t.Fatalf("error creating output folder %q: %v", subOutFolder, err)
				}
				runTests(t, subFolder, subOutFolder)
			})
		}
		continue
	}
}
