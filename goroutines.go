// Copyright 2022 Harald Albrecht.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not
// use this file except in compliance with the License. You may obtain a copy
// of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package noleak

import "github.com/thediveo/noleak/goroutine"

// Goroutines returns information about all goroutines: their goroutine IDs, the
// names of the topmost functions in the backtraces, and finally the goroutine
// backtraces.
func Goroutines() []goroutine.Goroutine {
	return goroutine.Goroutines()
}
