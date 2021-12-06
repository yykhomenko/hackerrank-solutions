package main

import "testing"

func Test_miniMaxSum(t *testing.T) {
	tests := []struct {
		name string
		s    string
		n    int64
		want int64
	}{
		{"valid", "abcac", 10, 4},
		{"valid2", "aba", 10, 7},
		{"valid3", "a", 1000000000000, 1000000000000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if repeatedString(tt.s, tt.n) != tt.want {
				t.Fail()
			}
		})
	}
}
