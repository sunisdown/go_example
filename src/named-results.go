package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	z := 19
	fmt.Println(z)
	return
}

func main() {
	fmt.Println(split(17))
}
