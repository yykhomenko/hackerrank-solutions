package main

import (
	"strconv"
	"strings"
)

func superDigit(n string, k int32) int32 {
	return calc(strings.Repeat(n, int(k)))
}

func calc(str string) int32 {
	if len(str) == 1 {
		return int32(str[0] - '0')
	}
	var sum int32
	for _, c := range str {
		sum += c - '0'
	}
	str = strconv.Itoa(int(sum))
	return calc(str)
}
