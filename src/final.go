package main

import (
	"fmt"
	"sort"
)

func main() {
	var k,n int
	fmt.Scan(&n,&k)
	p := make([]int, n)
	for i:=0; i<n; i++ {
		for j:=0; j<3; j++ {
			var c int
			fmt.Scan(&c)
			p[i] += c
		}
	}
	tmp := append([]int{}, p...)
	str := make([]string, n)
	sort.Slice(p, func(i, j int) bool { return p[i] > p[j] })
	fmt.Println(p)
	for i:=0; i<n; i++ {
		if p[k-1] <= p[i]+300 {
			for j:=0; j<n; j++ {
				if p[i] == tmp[j] {
					str[j] = "Yes"
				}
			}
		} else {
			for j:=0; j<n; j++ {
				if p[i] == tmp[j] {
					str[j] = "No"
				}
			}
		}
	}
	for i:=0; i<n; i++ {
		fmt.Println(str[i])
	}
}
