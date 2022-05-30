package main

import "fmt"

func main() {
	//n is the number of questions
	//m is the pass mark
	var n, m int
	fmt.Scanln(&n, &m)
	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}

	count := 0
	for i := 0; i < n; i++ {
		if a[i] >= a[m-1] && a[i] != 0 {
			count++
		}
	}

	fmt.Println(count)
}
