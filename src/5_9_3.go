package main

import "fmt"

func main() {
	var n,d int
	fmt.Scan(&n, &d)

	a := make([]int, n)
	for i:=0; i<n; i++ {
		fmt.Scan(&a[i])
	}
}
