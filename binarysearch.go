package main

import (
	"fmt"
	"sort"
	"time"
)

func binary_search(data []int, x int) int {
	fmt.Println(data)
	start_index, end_index := 0, len(data)
	fmt.Println(start_index, end_index)

	for start_index <= end_index {
		time.Sleep(1000 * time.Millisecond)
		h := start_index + (end_index-start_index)/2
		fmt.Println(start_index, end_index, h, x)
		if data[h] == x {
			return h
		} else if data[h] < x {
			start_index = h + 1
		} else {
			end_index = h
		}
	}
	return start_index
}

func main() {
	x := 23
	data := []int{1, 2, 3, 5, 7, 8, 11, 23, 25, 43}
	i := sort.Search(len(data), func(i int) bool { return data[i] >= x })
	fmt.Println(i)
	fmt.Println(data)
	fmt.Println(binary_search(data, x))
}
