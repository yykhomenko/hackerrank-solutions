package main

import (
	"reflect"
	"testing"
)

func Test_gradingStudents(t *testing.T) {
	tests := []struct {
		name string
		args []int32
		want []int32
	}{
		{name: "valid", args: []int32{73, 67, 38, 33}, want: []int32{75, 67, 40, 33}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gradingStudents(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("gradingStudents() = %v, want %v", got, tt.want)
			}
		})
	}
}
