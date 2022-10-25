package main

import (
	"fmt"
	"math"
)

func candies(n float64) float64 {
	x := n / 2
	newN := n - x + 1
	if n <= 2 {
		return 0
	}
	return n - newN
}

func main() {
	var n, j float64
	fmt.Scan(&n)
	for n != 0 {
		n--
		fmt.Scan(&j)
		x := math.Round(candies(j))
		fmt.Println(int(x))
	}
}
