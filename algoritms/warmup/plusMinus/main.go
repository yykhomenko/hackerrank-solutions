package main

import "fmt"

func plusMinus(arr []int32) {
	var positive, negative, zero int32
	for _, i := range arr {
		switch {
		case i > 0:
			positive++
		case i < 0:
			negative++
		default:
			zero++
		}
	}

	var n = float64(len(arr))
	fmt.Printf("%f\n%f\n%f",
		float64(positive)/n,
		float64(negative)/n,
		float64(zero)/n,
	)
}
