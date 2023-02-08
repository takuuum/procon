package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	b := make([]int, 1)
	for i:=0; i<n; i++ {
		fmt.Scan(&a[i])
		if i==0 {
			b[0] = a[i]
			continue
		}
		for j:=0; j<len(b); j++ {
			if a[i] == b[j] {
				break
			} else if j==len(b)-1 {
				b = append(b, a[i])
			} else {
				continue
			}
		}
	}
	fmt.Println(len(b))
}
