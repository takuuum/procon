package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i:=0; i<n; i++ {
		fmt.Scan(&a[i])
	}

	for i:=0; i<n-1; i++ {
		if 60 < float64(a[i]+a[i+1])/2 {
			fmt.Println("No")
			return
		}
		if i < n-2 {
			if 50 < float64(a[i]+a[i+1]+a[i+2])/3 {
				fmt.Println("No")
				return
			}
		}
		if i< n-3 {
			if 40 < float64(a[i]+a[i+1]+a[i+2]+a[i+3])/4 {
				fmt.Println("No")
				return
			}
		}
	}
	fmt.Println("Yes")
}
