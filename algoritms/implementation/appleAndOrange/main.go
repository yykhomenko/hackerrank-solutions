package main

import "fmt"

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	fmt.Println(count(s, t, a, apples))
	fmt.Println(count(s, t, b, oranges))
}

func count(s int32, t int32, tree int32, fruits []int32) int32 {
	counter := int32(0)
	for _, v := range fruits {
		x := tree + v
		if s <= x && x <= t {
			counter++
		}
	}

	return counter
}
