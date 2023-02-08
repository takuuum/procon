package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	a := make([]int, N)
	b := make([]int, N)
	as := 0
	bs := 0
	for i:=0; i<N; i++ {
		fmt.Scan(&a[i])
		as += a[i]
	}
	for i:=0; i<N; i++ {
		fmt.Scan(&b[i])
		bs += b[i]
	}

	if as < bs {
		fmt.Println("B")
	} else if bs < as {
		fmt.Println("A")
	} else {
		fmt.Println("DRAW")
	}
}
