package main

func compareTriplets(a []int32, b []int32) []int32 {
	var alice, bob int32
	for i := 0; i < 3; i++ {
		if a[i] > b[i] {
			alice++
		} else if a[i] < b[i] {
			bob++
		}
	}

	return []int32{alice, bob}
}
