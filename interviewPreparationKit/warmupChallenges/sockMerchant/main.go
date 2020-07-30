package main

func sockMerchant(n int32, ar []int32) int32 {
	colors := make(map[int32]int32)
	for _, v := range ar {
		colors[v]++
	}

	pairs := int32(0)
	for _, v := range colors {
		pairs += v / 2
	}

	return pairs
}
