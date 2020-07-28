package main

func birthdayCakeCandles(ar []int32) int32 {
	max := ar[0]
	for _, v := range ar {
		if max < v {
			max = v
		}
	}

	cnt := int32(0)
	for _, v := range ar {
		if max == v {
			cnt++
		}
	}

	return cnt
}
