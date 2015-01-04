package main

import "fmt"

func main() {
	c := make(chan int, 2) // 第二个参数作为缓冲长度
	fmt.Println(c)
	c <- 1
	c <- 2
//	c <- 3  //缓冲区满的时候运行时会报错
	fmt.Println(<-c)
	fmt.Println(<-c)
}
