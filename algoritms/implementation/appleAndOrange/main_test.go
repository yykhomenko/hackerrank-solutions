package main

import "testing"

func Test_countApplesAndOranges(t *testing.T) {
	type args struct {
		s       int32
		t       int32
		a       int32
		b       int32
		apples  []int32
		oranges []int32
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "valid", args: struct {
			s       int32
			t       int32
			a       int32
			b       int32
			apples  []int32
			oranges []int32
		}{s: 7, t: 11, a: 5, b: 15, apples: []int32{-2, 2, 1}, oranges: []int32{5, -6}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			countApplesAndOranges(tt.args.s, tt.args.t, tt.args.a, tt.args.b, tt.args.apples, tt.args.oranges)
		})
	}
}
