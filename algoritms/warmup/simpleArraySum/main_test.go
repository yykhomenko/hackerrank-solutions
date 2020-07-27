package main

import "testing"

func Test_simpleArraySum(t *testing.T) {
	tests := []struct {
		name string
		args []int32
		want int32
	}{
		{name: "valid", args: []int32{1, 2, 3, 4, 5}, want: 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simpleArraySum(tt.args); got != tt.want {
				t.Errorf("simpleArraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
