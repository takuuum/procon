package main

import "fmt"

func main() {
	var a,b int
	fmt.Scan(&a,&b)

	if a==1 && b==10 || a==10 && b==1 || a-b == 1 || a-b==-1 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
