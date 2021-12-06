package main

import "strings"

func repeatedString(s string, n int64) int64 {
	l := int64(len(s))
	part := n / l
	rest := n % l

	numAPart := int64(strings.Count(s, "a"))
	numARest := int64(strings.Count(s[:rest], "a"))

	return part*numAPart + numARest
}
