package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	h := make([]int, n)
	for i:=0; i<n; i++ {
		fmt.Scan(&h[i])
	}

	for i:=0; i<n; i++ {
		if h[n-1] == h[i] {
			fmt.Println(h[i])
			return
		} else if h[i] < h[i+1] {
			continue
		} else {
			fmt.Println(h[i])
			return
		}
	}
}
