package main

import "fmt"

func main() {
	var h,w int
	fmt.Scan(&h, &w)

	a := make([][]int, h)
	for i:=0; i<h; i++ {
		a[i] = make([]int, w)
	}

	for i:=0; i<h; i++ {
		for j:=0; j<w; j++ {
			fmt.Scan(&a[i][j])
		}
	}

	for i:=0; i<h; i++ {
		for j:=0; j<w; j++ {
			for k:=0; k<h; k++ {
				for l:=0; l<w; l++ {
					if i>=j || k>=l {continue}
					if a[i][k]+a[j][l] > a[j][k]+a[i][l] {
						fmt.Println("No")
						return
					}
					fmt.Println(a[i][k]+a[j][l], a[j][k]+a[i][l])
				}
			}
		}
	}
	fmt.Println("Yes")
}

