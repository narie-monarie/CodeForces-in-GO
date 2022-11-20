package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
func main() {
	var x, y, z, sum int
	fmt.Scanf("%d %d", &x, &y)
	for x != 0 {
		fmt.Scanf("%d", &z)
		sum += factorial(z)
		x--
	}
	// fmt.Println(sum)
	//fmt.Println(factorial(y))
	//fmt.Println(factorial(y)%sum)

	if sum%factorial(y) == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
