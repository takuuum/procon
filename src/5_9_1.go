package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	sum := 1
	fmt.Scan(&n)

	mainasu := make([]int, 0)
	mainasu2 := make([]int, 0)
	zero := make([]int, 0)
	a := make([]int, n)
	for i:=0; i<n; i++ {
		fmt.Scan(&a[i])
		if a[i] < 0 {
			mainasu = append(mainasu, i)
			mainasu2 = append(mainasu2, a[i])
		} else if a[i] == 0 {
			zero = append(zero, i)
		}
	}
	fmt.Println(len(mainasu)%2, len(mainasu))
	if len(mainasu)%2 == 0 {
		for i:=0; i<n; i++ {
			var flag bool
			if 0<len(zero) {
				for j:=0; j<len(zero); j++ {
					if zero[j] == i {
						flag = true
					}
				}
			}
			if flag {
				continue
			}
			sum *= a[i]
		}
	} else {
		sort.Slice(mainasu2, func(i, j int) bool { return mainasu2[i] < mainasu2[j] })
		for i:=0; i<n; i++ {
			var flag bool
			var mflag bool
			if 0<len(zero) {
				for j:=0; j<len(zero); j++ {
					if zero[j] == i {
						flag = true
					}
				}
			}
			if 0<len(mainasu) {
				for j:=0; j<len(mainasu); j++ {
					fmt.Println(mainasu[j], a[j], a)
					if mainasu[j] == i && a[i] == mainasu2[0] {
						mflag = true
					}
				}
			}
			if flag || mflag {
				continue
			}
			sum *= a[i]
		}
	}
	fmt.Println(sum, mainasu2, mainasu)
}
