package main

import "fmt"

func fibonacci(n int) int {
	if n < 2 {
		return n
	}

	var a, b int

	b = 1

	for n--; n > 0; n-- {
		a += b
		a, b = b, a
	}

	return b
}

func main() {
	var input int
	fmt.Scanln(&input)
	var sum int
	for i := 1; i <= input; i++{
		fmt.Scanln(i)
		v := fibonacci(i)
		sum += v
	}
	fmt.Println(sum)
}
