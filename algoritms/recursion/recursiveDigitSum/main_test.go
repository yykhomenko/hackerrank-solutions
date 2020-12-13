package main

import "testing"

func TestSuperDigit(t *testing.T) {
	tests := []struct {
		args string
		want int32
	}{
		{"1", 1},
		{"5", 5},
		{"9875", 2},
		{"9875987598759875", 8},
	}
	for _, tt := range tests {
		if got := superDigit(tt.args, int32(len(tt.args))); got != tt.want {
			t.Errorf("superDigit() = %v, want %v", got, tt.want)
		}
	}
}
