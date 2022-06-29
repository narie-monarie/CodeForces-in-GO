package main

import "fmt"

func main() {
	var n, petya, vasya, tonya int
	number := 0
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&petya, &vasya, &tonya)

		if petya+vasya+tonya >= 2 {
			number += 1
		}
	}
	fmt.Println(number, "\n")
}
