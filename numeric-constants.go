package main

import "fmt"

const (
	Big     = 1 << 100
	Small   = Big >> 99
)

func needInt(x int64) int64 {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(Small)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Big))
	fmt.Println(needFloat(Small))
}
