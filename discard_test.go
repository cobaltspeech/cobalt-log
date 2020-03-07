/*
   Copyright (2020) Cobalt Speech and Language Inc.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package log

import "testing"

func TestDiscardLogger(t *testing.T) {
	// this test only asserts the implementation of Logger.  There are no
	// values to test.

	l := NewDiscardLogger()
	l.Trace("msg", "trace message")
	l.Debug("msg", "debug message")
	l.Info("msg", "info message")
	l.Error("msg", "error message")
}
