package main

import (
	"fmt"
)

func main() {
	a := [3]int{0, 1, 2}

	for i, v := range a {
		if i == 0 {
			a[1], a[2] = 123, 456 //修改 a，
			fmt.Println(a)
		}
		a[i] = v + 100 // 使用 value 修改 a
	}

	fmt.Println(a)
} //
