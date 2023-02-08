package main

import "fmt"

func main() {
	var n,x int
	fmt.Scan(&n,&x)
	a := make([]int, n+1)
	a[0] = 0
	for i:=1; i<n+1; i++ {
		fmt.Scan(&a[i])
	}
	b := make([]bool, n+1)
	c := 0

	secret(x, a, b)
	for i:=0; i<n+1; i++ {
		if b[i] {
			c++
		}
	}
 	fmt.Println(c)
}

func secret(fre int, a []int, b []bool) {
	if b[fre] {
		return
	} else {
		b[fre] = true
		secret(a[fre], a, b)
	}


}
