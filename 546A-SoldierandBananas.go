package main

import (
	"fmt"
)

func soldierAndBananas(k, n, w int) int {
	var total int
	for i := 1; i <= w; i++ {
		total += i * k
	}
	borrow := total - n
	if borrow < 0 {
		borrow = 0
	}
	return borrow
}

//1 5 6 16
func main() {
	var a, b, c int
	fmt.Scanln(&a, &b, &c)
	fmt.Println(soldierAndBananas(a, b, c))
}
