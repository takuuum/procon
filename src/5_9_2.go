package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	sum := 0
	min := 1000000
	var minn int

	a := make([]int, n)
	for i:=0; i<n; i++ {
		fmt.Scan(&a[i])
		if a[i] < min {
			min = a[i]
			minn = i
		}
	}

	if n%2==0 {
		for i:=0; i<n; i++ {
			sum += a[i]
		}
	} else if minn%2 == 1 {
		if minn == 0 || minn==n-1 {
			for i:=0; i<n; i++ {
				if i==minn {
					continue
				}
				sum += a[i]
			}
		}
		if a[minn-1]<a[minn+1] {
			for i:=0; i<n; i++ {
				if i==minn-1 {
					continue
				}
				sum += a[i]
			}
		} else {
			for i:=0; i<n; i++ {
				if i==minn+1 {
					continue
				}
				sum += a[i]
			}
		}
	} else {
		for i:=0; i<n; i++ {
			if i==minn {
				continue
			}
			sum += a[i]
		}
	}
	fmt.Println(sum)
}
