package main

import "fmt"

func miniMaxSum(arr []int32) {
	var sum int64
	min := arr[0]
	max := arr[0]

	for _, i := range arr {
		sum += int64(i)
		if min > i {
			min = i
		}
		if max < i {
			max = i
		}
	}

	fmt.Println(
		sum-int64(max),
		sum-int64(min),
	)
}
