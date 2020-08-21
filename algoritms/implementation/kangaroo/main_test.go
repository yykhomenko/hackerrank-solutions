package main

import "testing"

func Test_kangaroo(t *testing.T) {
	type args struct {
		x1 int32
		v1 int32
		x2 int32
		v2 int32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "valid", args: struct {
			x1 int32
			v1 int32
			x2 int32
			v2 int32
		}{x1: 0, v1: 3, x2: 4, v2: 2}, want: "YES"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kangaroo(tt.args.x1, tt.args.v1, tt.args.x2, tt.args.v2); got != tt.want {
				t.Errorf("kangaroo() = %v, want %v", got, tt.want)
			}
			if got := kangarooBetter(tt.args.x1, tt.args.v1, tt.args.x2, tt.args.v2); got != tt.want {
				t.Errorf("kangarooBetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
