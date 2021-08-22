// SPDX-License-Identifier: Apache-2.0
// Copyright 2020-2021 Authors of Cilium

package utils

import "testing"

func TestCheckVersion(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"v1.9.6", true},
		{"v1.9.6.5", false},
		{"1.9.6", true},
		{"1.9.6.5", false},
		{"v1.10.0-rc1", true},
		{"1.10.0-rc1", true},
		{"1.9", false},
		{"v1.9", false},
		{"1", false},
		{"a01..0..0", false},
		{".1.9", false},
		{"..1.9", false},
		{"1...9", false},
		{"ddd", false},
		{"v.1.9", false},
		{"v..1.9", false},
		{":latest", true},
		{":92ff7ffa762f6f8bc397a28e6f3147906e20e8fa", true},
		{":92ff7ffa762f6f8bc397a28e6f3147906e20e8fa@sha256:4fde4abc19a1cbedb5084f683f5d91c0ea04b964a029e6d0ba43961e1ff5b5d8", true},
		{"-ci:92ff7ffa762f6f8bc397a28e6f3147906e20e8fa", true},
		{"-ci:92ff7ffa762f6f8bc397a28e6f3147906e20e8fa@sha256:4fde4abc19a1cbedb5084f683f5d91c0ea04b964a029e6d0ba43961e1ff5b5d8", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckVersion(tt.name); got != tt.want {
				t.Errorf("CheckVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
