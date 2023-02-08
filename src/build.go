package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	s := make([]int, n)
	nn := n

	for i:=0; i<n; i++ {
		fmt.Scan(&s[i])
		a := 0
		for j:=0;j<1; {
			b := 0
			a++
			for k:=0;k<1; {
				b++
				sum := (4*a*b)+3*(a+b)
				if sum == s[i] {
					nn--
					j = 1
					break
				}
				if sum > s[i] {
					k = 1
					if b==1 {
						j = 1
					}
					break
				}
			}
		}
	}
	fmt.Println(nn)
}
