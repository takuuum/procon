package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([][]int, n-1)
	for i:=0; i<n-1; i++ {
		a[i] = make([]int, 2)
		fmt.Scan(&a[i][0], &a[i][1])
	}

	center := 0
	for i:=0; i<n-1; i++ {
		if center == 0 {
			if a[i][0] == a[i+1][0] || a[i][0] == a[i+1][1] {
				center = a[i][0]
			} else if a[i][1] == a[i+1][0] || a[i][1] == a[i+1][1] {
				center = a[i][1]
			} else {
				fmt.Println("No")
				return
			}
		} else if center == a[i][0] || center == a[i][1] {
			continue
		} else {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
