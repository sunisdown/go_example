package main

import "fmt"

func fibonacci() func() int {
	j, k := 0, 1;
	return func() int {
		j, k = k, j+k
		return j
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
