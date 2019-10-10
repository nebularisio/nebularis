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

import (
	"bytes"
	"fmt"
)

type Messages []Message

var _ error = Messages{}

func (ms *Messages) Append(m Message) {
	*ms = append(*ms, m)
}

func (ms Messages) HasErrors() bool {
	for _, m := range ms {
		if m.Level == Error {
			return true
		}
	}

	return false
}

func (ms Messages) String() string {
	var b bytes.Buffer

	for _, m := range ms {
		_, _ = fmt.Fprintf(&b, "%v\n", m)
	}

	return b.String()
}

func (ms Messages) Error() string {
	return ms.String()
}
