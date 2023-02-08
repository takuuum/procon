package main

import (
	"fmt"
	"math"
)

func main() {
	var n, m int
	var min, mm int64
	fmt.Scan(&n, &m)
	b := make([][]int64, n)
	for i:=0; i<n; i++ {
		b[i] = make([]int64, m)
	}
	for i:=0; i<n; i++ {
		for j:=0; j<m; j++ {
			fmt.Scan(&b[i][j])
			if i==0 && j==0 {
				if (b[i][j]+7)%7+6 < int64(m) {
					fmt.Println("No")
					return
				}
				mm = b[i][j]
				min = mm
			} else if j==0 && b[i][j] == mm {
				min = mm
			} else if j==0 {
				fmt.Println("No")
				return
			} else if b[i][j] == min+1 {
				min = b[i][j]
			} else {
				fmt.Println("No")
				return
			}
		}
		mm += 7
	}
	if (math.Pow(10,100)-1)*7+7 < float64(min) {
		fmt.Println("No")
		return
	}
	fmt.Println("Yes")
}
