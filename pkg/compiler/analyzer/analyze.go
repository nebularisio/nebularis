//  Copyright 2019 Nebularis Authors.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package analyzer

func Analyze(o Options) Result {
	c := newContext(o.Sources)

	c.create()
	if c.m.HasErrors() {
		return Result{
			Messages: c.m,
		}
	}

	c.resolve()
	if c.m.HasErrors() {
		return Result{
			Messages: c.m,
		}
	}

	c.check()
	if c.m.HasErrors() {
		return Result{
			Messages: c.m,
		}
	}

	return Result{
		Semantics: c.s,
		Messages:  c.m,
	}
}
