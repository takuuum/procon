package main

import "fmt"

func main() {
	var N,count int
	fmt.Scan(&N)

	if N<=125 {
		count = 4
	} else if N<=211 {
		count = 6
	} else {
		count = 8
	}

	fmt.Println(count)
}
