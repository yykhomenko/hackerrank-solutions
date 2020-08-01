package main

func simpleArraySum(ar []int32) int32 {
	var sum int32
	for _, i := range ar {
		sum += i
	}
	return sum
}
