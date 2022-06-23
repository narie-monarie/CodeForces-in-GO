package main

import "fmt"

func OpenRooms(n int) int {
	count := 0
	for i := 0; i < n; i++ {
		var p, q int
		fmt.Scanln(&p, &q)

		if q > p && q-p >= 2 {
			count++
		}
	}

	return count
}

func main() {
	var a int
	fmt.Scanln(&a)
	fmt.Println(OpenRooms(a))
}
