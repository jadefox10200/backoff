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

// Constant is a backoff policy with a constant delay.
type Constant struct {
	delay time.Duration
}

// Nil is a constant backoff policy with zero delay.
var Nil = &Constant{}

// NewConstant returns a new Constant backoff policy.
func NewConstant(delay time.Duration) *Constant {
	return &Constant{delay}
}

// Decrease implements the Policy interface.
func (con *Constant) Decrease() {
}

// Increase implements the Policy interface.
func (con *Constant) Increase() time.Duration {
	if con == nil {
		return 0
	}
	return con.delay
}

// Sleep implements the Policy interface.
func (con *Constant) Sleep() {
	if con == nil {
		time.Sleep(0)
		return
	}
	time.Sleep(con.Increase())
}
