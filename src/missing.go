package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, 4*n-1)
	b := make([]int, n)
	for i:=0; i<4*n-1; i++ {
		fmt.Scan(&a[i])
		b[a[i]-1] += 1
	}
	for i:=0; i<n; i++ {
		if b[i] != 4 {
			fmt.Println(i+1)
			return
		}
	}
}
