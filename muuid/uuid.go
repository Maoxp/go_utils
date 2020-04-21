// Copyright 2020 com authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package muuid

import (
	"github.com/google/uuid"
)

//Go package for UUIDs based on RFC 4122 and DCE 1.1: Authentication and Security Services.
//A UUID is a 16 byte (128 bit) array. UUIDs may be used as keys to maps or compared directly.
func GetUUID() string {
	u, _ := uuid.NewRandom()
	return u.String()
}