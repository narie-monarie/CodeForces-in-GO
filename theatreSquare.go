package main

import (
	"fmt"
	"math"
)

func theatreSquare(x, y, z float64) float64 {
	a := math.Ceil(x / z)
	b := math.Ceil(y / z)

	result := a * b
	return result
}

func main() {
	var x, y, z float64
	fmt.Scanln(&x, &y, &z)
	fmt.Println(uint64(theatreSquare(x, y, z)))
}
