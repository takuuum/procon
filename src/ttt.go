package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	taro:=0
	hana:=0

	a := make([][]string, n)
	for i:=0; i<n; i++ {
		a[i] = make([]string, 2)
	}

	for i:=0; i<n; i++ {
		fmt.Scan(&a[i][0], &a[i][1])
		b := a[i][0]
		sort.Slice(a[i], func(k, j int) bool {
			return a[i][k] < a[i][j]
		})
		if a[i][0] == a[i][1] {
			taro++
			hana++
		} else if b==a[i][0] {
			hana+=3
		} else {
			taro+=3
		}
	}
	fmt.Println(taro, hana)
}
// GCMdz35n
