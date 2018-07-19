// Copyright 2018 Comcast Cable Communications Management, LLC
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pulsar

import "sync/atomic"

// monotonicID handles unique id generation
type monotonicID struct {
	id uint64
}

// next returns the next ID
func (r *monotonicID) next() *uint64 {
	nid := atomic.AddUint64(&r.id, 1) - 1
	return &nid
}
