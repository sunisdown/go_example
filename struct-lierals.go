package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1} // Y:0 被省略
	v3 = Vertex{} //X:0 Y:0 被省略
	p = &Vertex{1, 2} //类型为 *Vertex ,struct 指针
)

func main() {
	fmt.Println(v1, v2, v3, p)
}
