package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Scanln(&a, &b)
	x := math.Pow(a, b)
	y := math.Pow(b, a)
	fmt.Println(int(x - y))
}
