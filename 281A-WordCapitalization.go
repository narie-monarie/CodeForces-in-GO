package main

import (
	"fmt"
	"strings"
)

func capitalize(w string) string {

	return strings.Title(w)

}

func main() {
	var capital string
	fmt.Scanln(&capital)
	fmt.Println(capitalize(capital))
}
