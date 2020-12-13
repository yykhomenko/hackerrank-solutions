package main

import "strconv"

func superDigit(n string, k int32) int32 {
	if k == 1 {
		return int32(n[0] - '0')
	}
	var i, num int32
	for ; i < k; i++ {
		num += int32(n[i]) - '0'
	}
	str := strconv.Itoa(int(num))
	return superDigit(str, int32(len(str)))
}
