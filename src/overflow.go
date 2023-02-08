package main

import "fmt"

func main() {
	var n int64
	fmt.Scan(&n)
	if -2147483648 <= n && n < 2147483648 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
