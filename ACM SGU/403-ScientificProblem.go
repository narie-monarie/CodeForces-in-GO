package main

import "fmt"

func ScientificProblem(n int) int {
	v := 2*n + 1
	return v
}

func main() {
	var a int
	fmt.Scanf("%d", &a)
	fmt.Println(ScientificProblem(a))
}
