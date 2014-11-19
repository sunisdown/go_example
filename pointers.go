package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i //point to i
	fmt.Println(*p) //read i through pointer
	*p = 21// 通过指针 p 设置 i
	fmt.Println(i)

	p = &j
	*p = *p
	fmt.Println(j)
}
