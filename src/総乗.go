package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	minasu := 0
	miss := 0
	a := make([]int, n)
	b := make([]int, 2)
	for i:=0; i<n; i++ {
		fmt.Scan(&a[i])
		if a[i] < 0 {
			minasu++
		}
	}
	if minasu%2 == 0 {
		for i:=0; i<n; i++ {
			b[0] *= a[i]
		}
	} else {
		for i:=0; i<n; i++ {
			if miss == 0 {
				b[0] *= a[i]
			} else if 0 < miss && miss < minasu  {
				b[0] *= a[i]
				b[1] *= a[i]
			} else if miss == minasu {
				b[1] *= a[i]
			}
			if a[i] < 0 {
				miss++
			}
		}
	}
	if b[0] < b[1] {
		fmt.Println(b[1])
	} else {
		fmt.Println(b[0])
	}
}
