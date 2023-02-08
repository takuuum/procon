package main

import (
	"fmt"
)

func main() {
	var n, q int
	fmt.Scan(&n, &q)
	a := make(map[int][]int)
	for i:=0; i<n; i++ {
		var key int
		fmt.Scan(&key)
		a[key] = append(a[key], i+1)
	}
	x := make([]int, q)
	k := make([]int, q)

	for i:=0; i<q; i++ {
		fmt.Scan(&x[i], &k[i])

		amount := len(a[x[i]])
		if k[i] <= amount {
			fmt.Println(a[x[i]][k[i]-1])
		} else {
			fmt.Println(-1)
		}
	}
}
