package main

import "testing"

func Test_diagonalDifference(t *testing.T) {
	tests := []struct {
		name string
		args [][]int32
		want int32
	}{
		{name: "valid", args: [][]int32{
			{1, 2, 3},
			{4, 5, 6},
			{9, 8, 9}}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diagonalDifference(tt.args); got != tt.want {
				t.Errorf("diagonalDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
