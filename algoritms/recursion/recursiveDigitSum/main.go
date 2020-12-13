package main

import (
	"strconv"
)

func superDigit(n string, k int32) int32 {
	sum := calc(n)
	sum *= k
	return calc(strconv.Itoa(int(sum)))
}

func calc(n string) int32 {
	if len(n) == 1 {
		return int32(n[0] - '0')
	}
	sum := 0
	for _, c := range n {
		sum += int(c) - '0'
	}
	return calc(strconv.Itoa(sum))
}
