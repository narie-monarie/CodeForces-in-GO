package main

import "fmt"

//Not a must you add the arrqay thats a waste of space i only used it for practice
func fence(m, n int) int {
	var a, count int
	var t []int
	for i := 0; i < m; i++ {
		fmt.Scanf("%d", &a)
		t = append(t, a)
	}

	for i := 0; i < len(t); i++ {
		if t[i] > n {
			count += 2
		} else {
			count++
		}
	}
	return count
}

func main() {
	var m, n int
	fmt.Scanln(&m, &n)
	fmt.Println(fence(m, n))
}
