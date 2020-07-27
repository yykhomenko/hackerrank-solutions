package main

import (
	"testing"
)

func Test_plusMinus(t *testing.T) {
	tests := []struct {
		name string
		args []int32
	}{
		{name: "valid", args: []int32{1, 1, 0, -1, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			plusMinus(tt.args)
		})
	}
}
