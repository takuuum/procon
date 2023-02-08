package main

import "fmt"

func main() {
	var s,t,x int
	fmt.Scan(&s,&t,&x)

	if s<=t {
		if s<=x && x<t {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}

	} else if t<s {
		if x<t || s<=x {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
