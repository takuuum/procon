package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	a := make([]float64, N)
	b := make([]float64, N)
	c := make([]int, N)
	sum := 0
	for i:=0; i<N; i++ {
		fmt.Scan(&a[i])
		fmt.Scan(&b[i])
		c[i] = int(a[i] / b[i] * 100000)
		sum += c[i]
	}
	//N = N*10000
	var count int
	A := 0.00000
	sum = sum/2
	for i:=0; i<N; i++ {
		sum -= c[i]
		if sum<0 {
			sum += c[i]
			count = i
			break
		}
		A += a[i]
	}
	cc := b[count]*float64(sum)/100000
	A += cc
	fmt.Println(A)
}
