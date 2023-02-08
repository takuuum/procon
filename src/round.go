package main

import "fmt"

func main() {
	var x float64
	fmt.Scan(&x)

	y := int(x)
	z := int(10*x)-10*y
	if z<5 {
		fmt.Println(y)
	} else {
		fmt.Println(y+1)
	}
}
