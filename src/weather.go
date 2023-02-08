package main

import (
	"fmt"
	"strings"
)

func main() {
	var N int
	var S string

	fmt.Scan(&N,&S)

	s := strings.Split(S,"")
	if s[N-1] == "x" {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
