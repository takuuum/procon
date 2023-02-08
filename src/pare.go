package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n+2)
	b := make([]int, n+2)
	a[0] = -1
	a[n+1] = -1
	for i:=1; i<=n; i++ {
		fmt.Scan(&a[i])
	}
	b = append([]int{}, a...)
	sort.Slice(b, func(i, j int) bool { return b[i] > b[j] })
	count := 0

	for j:=0; j<n; j++ {
		for i:=1; i<=n; i++ {
			if b[j] == a[i] {
				count += a[i]
				a[i] = -1
				if a[i-1] < a[i+1] {
					count += a[i+1]
					a[i+1] = -1
					fmt.Println(a, count)
				} else if a[i-1] > a[i+1] || (a[i-1] == a[i+1] && a[i-1] != -1) {
					count += a[i-1]
					a[i-1] = -1
					fmt.Println(a, count)
				} else {
					break
				}
			}
		}
	}
	fmt.Println(count)
}
