package main

import "fmt"

func qsort(data []int) {
	if len(data) <= 1 {
		return
	}
	mid, i := data[0], 1
	head, tail := 0, len(data)-1
	for head < tail {
		if data[i] < mid {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[i], data[head] = data[head], data[i]
			head++
			i++
		}
	}
	data[head] = mid
	qsort(data[:head])
	qsort(data[head+1:])
}

func main() {
	v := []int{2, 8, 7, 1, 3, 5, 6, 4, 0, 9, 10, 12, 11}
	qsort(v)
	fmt.Println(v)
}
