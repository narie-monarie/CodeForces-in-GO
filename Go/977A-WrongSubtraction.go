package main

import "fmt"

func WrongSubtraction(m, n int) int {
	for i := 0; i < n; i++ {
		if m%10 != 0 {
			m--
		} else {
			m /= 10
		}
	}

	return m
}

func main() {
	var x, m int
	fmt.Scanln(&x, &m)
	fmt.Println(WrongSubtraction(x, m))
}
