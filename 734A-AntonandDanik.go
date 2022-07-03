package main

import (
	"fmt"
)

func AntonDanik() {
	var a, m, n int
	var b string
	fmt.Scanln(&a)
	fmt.Scanf("%s ", &b)
	for i := range b {
		if b[i] == 65 {
			m++
		} else {
			n++
		}
	}

	if m > n {
		fmt.Println("Anton")
	} else if n > m {
		fmt.Println("Danik")
	} else {
		fmt.Println("Friendship")
	}

}

func main() {
	AntonDanik()
}
