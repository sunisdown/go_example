package main

import "fmt"

func main() {
	var z []int
	fmt.Println(z, len(z), cap(z)) //nil
	fmt.Printf("a is type of %T\n", z)
	if z == nil {
		fmt.Print("nil\n")
	}

	a := make([]int, 0, 5)
	fmt.Println(a, len(a), cap(a))
	fmt.Printf("a is type of %T\n", a)
	if a == nil {
		fmt.Print("a is nil")
	}
}
