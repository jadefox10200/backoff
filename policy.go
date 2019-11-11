// Copyright 2019 Miles Barr <milesbarr2@gmail.com>
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

package backoff

import "time"

// A Policy is an interface for implementing backoff policies.
type Policy interface {
	Decrease()
	Increase() time.Duration
	Sleep()
}

// Default returns default "safe" backoff policy.
func Default() Policy {
	return NewJitter(NewExponential(100*time.Millisecond, 1.5, 60*time.Second), 0.5)
}
