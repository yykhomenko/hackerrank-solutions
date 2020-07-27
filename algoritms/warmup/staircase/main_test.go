package main

import "testing"

func Test_staircase(t *testing.T) {
	tests := []struct {
		name string
		args int32
	}{
		{name: "valid", args: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			staircase(tt.args)
		})
	}
}
