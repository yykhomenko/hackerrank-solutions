package main

/*
 * Complete the 'diagonalDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func diagonalDifference(arr [][]int32) int32 {
	var sumA, sumB int32
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if i == j {
				sumA += arr[i][j]
			}
			if i == len(arr[i])-1-j {
				sumB += arr[i][j]
			}
		}
	}

	sum := sumA - sumB
	if sum < 0 {
		return -sum
	} else {
		return sum
	}
}
