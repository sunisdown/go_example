package main

import "fmt"

func main() {
	p := []int{1,3,4,7,8}
	fmt.Println("p is ", p)
	fmt.Println("p[1:4] is ", p[1:4])

	fmt.Println("p[:3] is ", p[:3]) // 类似于 Python 的写法
	fmt.Println("p[4:] is ", p[4:]) // 类似于 Python 的写法
}
