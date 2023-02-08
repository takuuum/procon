package main

import "fmt"

func main() {
	var h,w int
	fmt.Scan(&h,&w)
	b := make([][]int, w)
	for i:=0; i<w; i++ {
		b[i] = make([]int, h)
	}
	for i:=0; i<h; i++ {
		for j:=0; j<w; j++ {
			fmt.Scan(&b[j][i])
		}
	}
	for i:=0; i<w; i++ {
		for j:=0; j<h; j++ {
			fmt.Print(b[i][j], " ")
		}
		fmt.Print("\n")
	}
}
