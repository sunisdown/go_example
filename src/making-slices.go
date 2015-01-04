package main

import "fmt"

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d, cap=%d %v\n", s, len(x), cap(x), x) // cap 返回 capacity，容量
}

func main() {
	a := make([]int, 5)
	fmt.Println("a = ", a)
	printSlice("a", a)
	b := make([]int, 0, 5)
	fmt.Println("b = ", b)
	printSlice("b", b)
	c := b[:2]
	fmt.Println("c = ", c)
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)
}
