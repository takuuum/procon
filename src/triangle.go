package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([][]int, n)
	for i:=0; i<n; i++ {
		a[i] = make([]int, 2)
		fmt.Scan(&a[i][0], &a[i][1])
	}

	b := make([][][]int, n)
	for i:=0; i<n; i++ {
		b[i] = make([][]int, n)
		for j:=0; j<n; j++ {
			b[i][j] = make([]int, n)
		}
	}
	count := 0

	for i:=0; i<n; i++ {
		for j:=0; j<n; j++ {
			for k:=0; k<n; k++ {
				//if (a[j][0]-a[i][0])*(a[k][1]-a[i][1])-(a[k][0]-a[i][0])*(a[j][1]-a[i][1]) != 0 {count++}
				if math.Abs(float64(a[i][1]-a[j][1])/float64(a[i][0]-a[j][0])) != math.Abs(float64(a[i][1]-a[k][1])/float64(a[i][0]-a[k][0])) {
					if i<j && j<k {
						if b[i][j][k] == 1 {continue}
						b[i][j][k] = 1
						count++
					} else if j<k && k<i {
						if b[j][k][i] == 1 {continue}
						b[j][k][i] = 1
						count++
					} else if k<i && i<j {
						if b[k][i][j] == 1 {continue}
						b[k][i][j] = 1
						count++
					} else if i<k && k<j {
						if b[i][k][j] == 1 {continue}
						b[i][k][j] = 1
						count++
					} else if j<i && i<k {
						if b[j][i][k] == 1 {continue}
						b[j][i][k] = 1
						count++
					} else if k<j && j<i {
						if b[k][j][i] == 1 {continue}
						b[k][j][i] = 1
						count++
					}
				}
			}
		}
	}
	fmt.Println(count)
}
