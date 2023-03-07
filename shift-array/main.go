package main

import "fmt"

func shiftArray(array []int, i int) []int {
	shiftedArray := make([]int, 9)
	shiftPattern := [][]int{
		{3, 0, 1, 6, 4, 2, 7, 8, 5},
		{6, 3, 0, 7, 4, 1, 8, 5, 2},
		{7, 6, 3, 8, 4, 0, 5, 2, 1},
		{8, 7, 6, 5, 4, 3, 2, 1, 0},
		{5, 8, 7, 2, 4, 6, 1, 0, 3},
		{2, 5, 8, 1, 4, 7, 0, 3, 6},
		{1, 2, 5, 0, 4, 8, 3, 6, 7},
		{0, 1, 2, 3, 4, 5, 6, 7, 8},
	}

	for j := 0; j < 9; j++ {
		shiftedArray[j] = array[shiftPattern[i-1][j]]
	}

	return shiftedArray
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	i := 2
	fmt.Println("Original array:", array)
	shiftedArray := shiftArray(array, i)
	fmt.Println("Shifted array:", shiftedArray)
}
