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

// Exponential is a backoff policy with an exponentially changing delay.
type Exponential struct {
	Min    time.Duration
	Factor float64
	Max    time.Duration
	delay  time.Duration
}

// NewExponential returns a new Exponential backoff policy.
func NewExponential(min time.Duration, factor float64, max time.Duration) *Exponential {
	return &Exponential{Min: min, Factor: factor, Max: max}
}

// Decrease implements the Policy interface.
func (exp *Exponential) Decrease() {
	exp.delay = time.Duration(float64(exp.delay) / exp.Factor)
	if exp.delay < exp.Min {
		exp.delay = exp.Min
	}
}

// Increase implements the Policy interface.
func (exp *Exponential) Increase() time.Duration {
	delay := exp.delay
	exp.delay = time.Duration(float64(exp.delay) * exp.Factor)
	if exp.delay > exp.Max {
		exp.delay = exp.Max
	}
	return delay
}

// Sleep implements the Policy interface.
func (exp *Exponential) Sleep() {
	time.Sleep(exp.Increase())
}
