package main

import (
	"fmt"
	"os"
)

var flag bool

func main() {
	flag = false
	var n,x int
	fmt.Scan(&n,&x)
	a := make([]int, n)
	b := make([]int, n)
	nn := 0
	for i:=0; i<n; i++ {
		fmt.Scan(&a[i],&b[i])
	}
	ma := a[0]
	mb := b[0]
	go aa(x,n-1,nn,ma,a,b)
	aa(x,n-1,nn,mb,a,b)
	fmt.Println("No")
}

func aa(x, n, nn, mm int, a, b []int) {
	if x == mm && n == nn {
		fmt.Println("Yes")
		flag = true
		os.Exit(0)
	} else if x < mm || n == nn {
		return
	} else {
		nn++
		mm += a[nn]
		go aa(x, n, nn, mm, a, b)
		mm = mm-a[nn]+b[nn]
		aa(x, n, nn, mm, a, b)
	}
}
