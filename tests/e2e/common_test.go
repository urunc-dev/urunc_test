// Copyright (c) 2023-2026, Nubificus LTD
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

package urunce2etesting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindValOfKey(t *testing.T) {
	tests := []struct {
		name       string
		searchArea string
		key        string
		wantVal    string
		wantErr    bool
	}{
		{
			name:       "scalar string value",
			searchArea: `[{"IPAddress":"172.17.0.2","Name":"test"}]`,
			key:        "IPAddress",
			wantVal:    "172.17.0.2",
		},
		{
			name:       "scalar integer value",
			searchArea: `[{"Pid":12345,"Name":"test"}]`,
			key:        "Pid",
			wantVal:    "12345",
		},
		{
			name:       "value containing comma is truncated",
			searchArea: `[{"Name":"test","Label":"foo,bar"}]`,
			key:        "Label",
			wantVal:    "foo,bar",
		},
		{
			name:       "nested object value",
			searchArea: `[{"NetworkSettings":{"IPAddress":"172.17.0.2"}}]`,
			key:        "NetworkSettings",
			wantVal:    `{"IPAddress":"172.17.0.2"}`,
		},
		{
			name:       "key not found",
			searchArea: `[{"Name":"test"}]`,
			key:        "Missing",
			wantErr:    true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			val, err := findValOfKey(tc.searchArea, tc.key)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.wantVal, val)
		})
	}
}
