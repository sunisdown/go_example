package main

import (
	"fmt"
	"sort"
)

func binary_search(data []int, x int) int {
	start_index, end_index := 0, len(data)

	for start_index <= end_index {
		h := start_index + (end_index-start_index)/2
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
	i := sort.Search(len(data), func(i int) bool { return data[i] >= x }) //使用官方的 sort.Search查找库
	fmt.Println(i)
	fmt.Println(binary_search(data, x))
}
