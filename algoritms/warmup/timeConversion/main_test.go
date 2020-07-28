package main

import "testing"

func Test_timeConversion(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{name: "midnight", args: "12:00:00AM", want: "00:00:00"},
		{name: "noon", args: "12:00:00PM", want: "12:00:00"},
		{name: "am", args: "07:05:45AM", want: "07:05:45"},
		{name: "pm", args: "07:05:45PM", want: "19:05:45"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := timeConversion(tt.args); got != tt.want {
				t.Errorf("timeConversion() = %v, want %v", got, tt.want)
			}
		})
	}
}
