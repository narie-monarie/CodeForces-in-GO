package main

import "fmt"

func BigBearBro(a, b int) int {
	count := 0
	for b >= a {
		a *= 3
		b *= 2
		count++
	}

	return count
}

func main() {
	var a, b int
	fmt.Scanln(&a, &b)
	fmt.Println(BigBearBro(a, b))
}
