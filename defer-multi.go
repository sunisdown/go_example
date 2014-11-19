package main

import "fmt"

func main() {
	fmt.Print("counting\n")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Print("done\n")
}
