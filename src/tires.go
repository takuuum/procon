package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	a := len(s)
	b := strings.Split(s, "")
	if b[a-1] == "r" && b[a-2] == "e" {
		fmt.Println("er")
	} else if b[a-1] == "t" && b[a-2] == "s" && b[a-3] == "i" {
		fmt.Println("ist")
	}
}
