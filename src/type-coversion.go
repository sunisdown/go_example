package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 2, 3
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z int = int(f)
	fmt.Println(x, y, f, z)
}
