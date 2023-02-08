package main

import (
	"fmt"
	"sort"
)

func main() {
	var A,B,C,K int
	fmt.Scan(&A,&B,&K)

	if A>B {
		C = A
	} else {
		C = B
	}
	k := []int{0}
	for i:=1; i<=C; i++ {
		if A%i==0 && B%i==0 {
			k = append(k, i)
		}
	}

	sort.Slice(k, func(i, j int) bool {
		return k[i] < k[j]
	})

	fmt.Println(k[K])
}
