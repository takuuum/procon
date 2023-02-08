package main

import (
	"fmt"
	"reflect"
)

func main() {
	var n,l int
	fmt.Scan(&n)
	a := make([][]int, n)
	b := make([][]int, 1)
	for i:=0; i<n; i++ {
		fmt.Scan(&l)
		a[i] = make([]int, l)
		for j:=0; j<l; j++ {
			fmt.Scan(&a[i][j])
		}
	}
	count := 0
	for i:=0; i<n; i++ {
		if i==0 {
			b[0] = a[0]
		}
		for j:=0; j<len(b); j++ {
			if reflect.DeepEqual(a[i], b[j]) {
				count++
			}
		}
		if count == 0 {
			b = append(b, a[i])
		}
		count = 0
	}
	fmt.Println(len(b))
}
