package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	p := make([]int, N)
	count := 0
	for i:=0; i<N; i++ {
		fmt.Scan(&p[i])
		if i>0 && p[i]>p[i-1] {
			p[i] = p[i-1]
		}
		count += p[i]
	}
	fmt.Println(count)
}
