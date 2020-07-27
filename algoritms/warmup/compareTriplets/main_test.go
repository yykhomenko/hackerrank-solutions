package main

import (
	"reflect"
	"testing"
)

func Test_compareTriplets(t *testing.T) {
	type args struct {
		a []int32
		b []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{name: "valid", args: struct {
			a []int32
			b []int32
		}{a: []int32{1, 2, 3}, b: []int32{3, 2, 3}}, want: []int32{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareTriplets(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("compareTriplets() = %v, want %v", got, tt.want)
			}
		})
	}
}
