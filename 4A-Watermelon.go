package main

import "fmt"

func watermelon() {
	var w int
	fmt.Scanln(&w)
	if w%2 == 0 && w != 2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func main() {
	watermelon()
}
