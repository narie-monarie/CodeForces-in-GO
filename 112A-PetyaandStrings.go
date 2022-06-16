package main

import (
	"fmt"
	"strings"
)

func PetyaStrings(a, b string) int {
	main := strings.ToLower(a)
	target := strings.ToLower(b)
	if main > target {
		return 1
	} else if target > main {
		return -1
	}
	return 0
}

func main() {
	var a, b string
	fmt.Scanln(&a, &b)
	fmt.Println(PetyaStrings(a, b))
}
