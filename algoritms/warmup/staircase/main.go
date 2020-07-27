package main

import "fmt"

func staircase(n int32) {
	num := int(n)
	row := make([]byte, n)

	for i := 0; i < num; i++ {
		for j := 0; j < num; j++ {
			if i < num-1-j {
				row[j] = ' '
			} else {
				row[j] = '#'
			}
		}
		fmt.Printf("%s\n", row)
	}
}
