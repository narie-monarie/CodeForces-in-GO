package main

import "fmt"

func trial(values string) string {
	v := []rune(values)
	for i := 0; i < len(values)-2; i++ {
		if v[i] > v[i+2] {
			temp := v[i]
			v[i] = v[i+2]
			v[i+2] = temp

		}
		for i := 0; i < len(values)-2; i++ {
			if v[i] > v[i+2] {
				temp := v[i]
				v[i] = v[i+2]
				v[i+2] = temp

			}
		}
	}
	result := string(v)
	return result
}

func main() {
	var values string
	fmt.Scanln(&values)
	fmt.Println(trial(values))
}
