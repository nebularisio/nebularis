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

import "flag"

var (
	updateBaselines bool   // nolint:gochecknoglobals
	outFolder       string // nolint:gochecknoglobals
)

func init() { // nolint:gochecknoinits
	flag.BoolVar(&updateBaselines, "ubl", false, "update the baseline files with the test results")
	flag.StringVar(&outFolder, "diffout", "", "output folder for diffing")
}

func UpdateBaselines() bool {
	return updateBaselines
}

func OutFolder() string {
	return outFolder
}
