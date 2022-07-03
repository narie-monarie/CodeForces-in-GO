package main

import (
	"fmt"
	"math"
)

func Elephant(n float64) float64 {
	var result float64 = n / 5

	return math.Ceil(result)
}

func main() {
	var n float64
	fmt.Scanln(&n)
	fmt.Println(Elephant(n))
}
