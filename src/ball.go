package main

import (
	"fmt"
	"strings"
)

func main() {
	var N,m int
	var A,B string
	fmt.Scan(&N)

	n := 2
	s := 0
	count := 0

	for ;; {
		n *= 2
		count++
		B += "B"
		if n>N/2 && s==0 {
			s = n/2
		}
		if n>N {
			m = n/2
			break
		}
	}

	if N/2-s>N-m {
		for i:=0; i<N-m; i++ {
			A += "A"
		}
		if N%2==0 {
			fmt.Printf("A%s%s", B, A)
		} else {
			fmt.Printf("A%s%sA", B, A)
		}
	} else {
		B = strings.Replace(B, "B", "", 1)
		for i:=0; i<N/2-s; i++ {
			A += "A"
		}
		if N%2==0 {
			fmt.Printf("A%s%sB", B, A)
		} else {
			fmt.Printf("A%s%sBA", B, A)
		}
	}
}