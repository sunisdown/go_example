package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func test(fn func() int) int {
	return fn() + 1
}

type FormatFunc func(s string, x, y int) string //定义函数类型

func format(fn FormatFunc, s string, x, y int) string {
	return fn(s, x, y)
}

func main() {
	fmt.Println(add(3, 4))
	s1 := test(func() int { return 100 }) //将地名函数作为参数

	s2 := format(func(s string, x, y int) string {
		return fmt.Sprintf(s, x, y)
	}, "%d, %d", 10, 20)

	fmt.Println(s1, s2)
}
