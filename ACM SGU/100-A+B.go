package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanln(&a, &b)
	c := a + b
	fmt.Println(c)
}
