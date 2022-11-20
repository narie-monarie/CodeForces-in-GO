package main

import (
	"fmt"
	"strings"
)

//	func stringTask(x string)string{
//		for i := 0; i <
//	}
func main() {
	var z string
	fmt.Scanf("%s", &z)
	x := strings.ToUpper(z)
	chCh := []byte{}
	var zd int
	for i := 0; i < len(x); i++ {
		if x[i] == 'A' || x[i] == 'O' || x[i] == 'Y' || x[i] == 'E' ||
			x[i] == 'U' || x[i] == 'I' {
			zd++
		} else {
			chCh = append(chCh, x[i])
		}
	}

	var stack []byte
	for _, v := range chCh {
		stack = append(stack, '.')
		stack = append(stack, v)
	}

	for i := 0; i < len(stack); i++ {
		fmt.Print(strings.ToLower(string(stack[i])))
	}
}
