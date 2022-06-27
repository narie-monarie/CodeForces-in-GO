package main

import "fmt"

func main() {
	var a int64
	fmt.Scanln(&a)
	if a%2 == 0 {
		fmt.Println(a / 2)
	} else {
		fmt.Println(-(a/2 + 1))
	}

}
