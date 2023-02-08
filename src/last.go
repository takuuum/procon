package main

import "fmt"

func main() {
	var n,k,a int
	fmt.Scan(&n,&k,&a)

	for ;; {
		k--
		if k==0 {
			fmt.Println(a)
			return
		}
		a++
		if a>n {
			a = 1
		}
	}
}
