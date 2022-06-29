package main

import "fmt"

func EasyProblem(n int) {
	var x, total int
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &x)
		total += x
	}
	if total >= 1 {
		fmt.Println("HARD")
	} else {
		fmt.Println("EASY")
	}
}

func main() {
	var x int
	fmt.Scanln(&x)
	EasyProblem(x)
}
