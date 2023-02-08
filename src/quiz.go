package main

import "fmt"

func main() {
	var S,T,U string
	fmt.Scan(&S,&T,&U)

	matrix := [][]string{{"ABC",""},{"ARC",""},{"AGC",""},{"AHC",""}}

	for i:=0; i<4; i++ {
		if matrix[i][0] == S {
			matrix[i][1] = "ok"
		} else if matrix[i][0] == T {
			matrix[i][1] = "ok"
		} else if matrix[i][0] == U {
			matrix[i][1] = "ok"
		} else {
			fmt.Println(matrix[i][0])
			return
		}
	}
}
