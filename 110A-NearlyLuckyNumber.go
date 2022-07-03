package main

import "fmt"

func main() {
	var n string
	fmt.Scanln(&n)
	count := 0
	for i := 0; i < len(n); i++ {
		if n[i] == '4' || n[i] == '7' {
			count++
		}
	}
	if count == 4 || count == 7 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
