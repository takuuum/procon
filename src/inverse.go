package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)

	p := make([][]int, N)

	for i:=0; i<N; i++ {
		p[i] = make([]int, 2)
		fmt.Scan(&p[i][0])
		p[i][1] = i+1
	}

	sort.Slice(p, func(i, j int) bool {
		return p[i][0] < p[j][0]
	})

	for i:=0; i<N; i++ {
		fmt.Printf("%d ",p[i][1])
	}
}

