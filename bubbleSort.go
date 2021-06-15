package main

import (
	"fmt"
)

func bubbleSort(sl1 []int) []int {
	for i := 0; i < len(sl1); i++ {
		for j := len(sl1) - 1; j > i; j-- {
			if sl1[j-1] > sl1[j] {
				swap(sl1, (j-1))
			}
		}
	}
	fmt.Println("Here is the sorted slice: ")
	fmt.Println(sl1)
	return sl1
}

func swap(x []int, y int) []int {
	if x[y] > x[y+1] {
		x[y], x[y+1] = x[y+1], x[y]
	}
	return x
}	

func createSlice() []int {
	sl := make([]int, 0)
	var idx int
	for i := 0; i < 10; i++ {
		fmt.Println("Please enter an interger: ")
		fmt.Scanln(&idx)
		sl = append(sl, idx)
	}
	fmt.Println("Here is the original slice: ")
	fmt.Println(sl)
	return sl
}

func main() {
	bubbleSort(createSlice())
}
