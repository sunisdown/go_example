package main

import "fmt"

var pow = []int{1, 2, 3, 4, 6, 8, 23, 45}

func main() {
	for i, v := range pow {
		fmt.Printf("pow[%d] = %d\n", i, v)
	}
}
