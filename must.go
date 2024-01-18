/*
 * Copyright (c) 2023 sixwaaaay.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package must

// the defer statement is executed even if the function panics
// so in the main() function,defer and Must() can be used to simplify the error handling

// do not use in library code, only in top level main() functions where you want to avoid error handling boilerplate

// Must panics if err is not nil.
// use this to avoid if err != nil { panic(err) } boilerplate in main()
func Must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

// RunE panics if err is not nil.
// use this to avoid if err != nil { panic(err) } boilerplate in main()
func RunE(err error) {
	if err != nil {
		panic(err)
	}
}

// Run panics if err is not nil.
// use this to avoid if err != nil { panic(err) } boilerplate in main()
func Run(f func() error) {
	if err := f(); err != nil {
		panic(err)
	}
}
