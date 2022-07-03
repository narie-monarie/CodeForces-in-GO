package main

import (
	"fmt"
	"math"
)

func main() {
	var l, w float64
	fmt.Scanln(&l, &w)

	x := (l * w) / 2
	fmt.Println(math.Floor(x))
}
