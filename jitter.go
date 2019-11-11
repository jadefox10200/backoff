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

import (
	"math/rand"
	"time"
)

// Jitter is a backoff policy that applies jitter to the delay of another
// backoff policy.
type Jitter struct {
	Source Policy
	Factor float64
}

// NewJitter returns a new Jitter backoff policy given a source backoff policy
// and a jitter factor.
func NewJitter(source Policy, factor float64) *Jitter {
	return &Jitter{source, factor}
}

// Decrease implements the Policy interface.
func (jit *Jitter) Decrease() {
	jit.Source.Decrease()
}

// Increase implements the Policy interface.
func (jit *Jitter) Increase() time.Duration {
	delay := float64(jit.Source.Increase())
	delta := float64(delay) * jit.Factor
	min := delay - delta
	max := delay + delta
	return time.Duration(min + (rand.Float64() * (max - min + 1)))
}

// Sleep implements the Policy interface.
func (jit *Jitter) Sleep() {
	time.Sleep(jit.Increase())
}
