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

package diag

import "fmt"

type Level int

const (
	Info Level = iota
	Warning
	Error
)

func (l Level) String() string {
	switch l {
	case Info:
		return "I"
	case Warning:
		return "W"
	case Error:
		return "E"
	default:
		panic(fmt.Sprintf("Unrecognized level: %d", int(l)))
	}
}
