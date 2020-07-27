package main

import "testing"

func Test_miniMaxSum(t *testing.T) {
	tests := []struct {
		name string
		args []int32
	}{
		{name: "valid", args: []int32{1, 3, 5, 7, 9}},
		{name: "valid2", args: []int32{7, 69, 2, 221, 8974}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			miniMaxSum(tt.args)
		})
	}
}
