package main

import "fmt"

const Pi = 3.14

func main() {
	const world = "世界"
	const hello string = "你好"
	fmt.Println(hello, world)
	fmt.Println("Hello", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
