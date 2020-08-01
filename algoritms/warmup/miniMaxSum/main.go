package main

import "fmt"

func miniMaxSum(arr []int32) {
	min, max, sum := arr[0], arr[0], int64(0)

	for _, i := range arr {
		sum += int64(i)
		if min > i {
			min = i
		}
		if max < i {
			max = i
		}
	}

	fmt.Println(sum-int64(max), sum-int64(min))
}
