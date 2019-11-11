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

// Linear is a backoff policy with a linearly changing delay.
type Linear struct {
	Min   time.Duration
	Step  time.Duration
	Max   time.Duration
	delay time.Duration
}

// NewLinear returns a new Linear backoff policy.
func NewLinear(min, step, max time.Duration) *Linear {
	return &Linear{Min: min, Step: step, Max: max}
}

// Decrease implements the Policy interface.
func (lin *Linear) Decrease() {
	lin.delay -= lin.Step
	if lin.delay < lin.Min {
		lin.delay = lin.Min
	}
}

// Increase implements the Policy interface.
func (lin *Linear) Increase() time.Duration {
	delay := lin.delay
	lin.delay += lin.Step
	if lin.delay > lin.Max {
		lin.delay = lin.Max
	}
	return delay
}

// Sleep implements the Policy interface.
func (lin *Linear) Sleep() {
	time.Sleep(lin.Increase())
}
