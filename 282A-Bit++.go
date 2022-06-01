package main

import "fmt"

func main() {
	var n int
	var input string
	fmt.Scanln(&n)
	add := 0
	for i := 0; i < n; i++ {
		fmt.Scanln(&input)
		if input == "++X" || input == "X++" {
			add++
		} else {
			add--
		}
	}

	fmt.Println(add)
}
