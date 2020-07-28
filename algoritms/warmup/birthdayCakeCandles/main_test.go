package main

import "testing"

func Test_birthdayCakeCandles(t *testing.T) {
	tests := []struct {
		name string
		args []int32
		want int32
	}{
		{name: "valid", args: []int32{3, 2, 1, 3}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := birthdayCakeCandles(tt.args); got != tt.want {
				t.Errorf("birthdayCakeCandles() = %v, want %v", got, tt.want)
			}
		})
	}
}
