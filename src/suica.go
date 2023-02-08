package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	p := make([][]int, N)
	q := 100
	var s float64
	s = 0

	for i:=0; i<N; i++ {
		p[i] = make([]int, 2)
		fmt.Scan(&p[i][0], &p[i][1])
		s += float64(p[i][0])*float64(q)/10000
		q = 100-p[i][0]-p[i][1]
		fmt.Println(q)
	}

	fmt.Println(s)
}
