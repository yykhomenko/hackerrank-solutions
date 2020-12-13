package main

import (
	"strconv"
)

func superDigit(n string, k int32) int32 {
	sum := calc(n)
	sum *= k
	n = strconv.Itoa(int(sum))
	return calc(n)
}

func calc(n string) int32 {
	if len(n) == 1 {
		return int32(n[0] - '0')
	}
	var sum int32
	for _, c := range n {
		sum += c - '0'
	}
	n = strconv.Itoa(int(sum))
	return calc(n)
}
