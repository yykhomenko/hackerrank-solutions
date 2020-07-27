package main

import "testing"

func Test_aVeryBigSum(t *testing.T) {
	tests := []struct {
		name string
		args []int64
		want int64
	}{
		{name: "valid", args: []int64{1, 2, 3, 4, 5}, want: 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := aVeryBigSum(tt.args); got != tt.want {
				t.Errorf("aVeryBigSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
