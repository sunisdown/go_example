package main

import "fmt"

func main() {
	p := []int{2, 3, 4, 6, 8, 10}
	fmt.Println("p ==", p)
	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] is %d\n", i, p[i])
	}
}
