package main

import (
	"fmt"
	"strings"
)

func main() {
	var S,T,a string
	var order int
	fmt.Scan(&S,&T)
	s:= strings.Split(S, "")
	t:= strings.Split(T, "")

	if len(S) > len(T) {
		order = len(T)
		a = "No"
	} else {
		order = len(S)
		a = "Yes"
	}

	for i:=0; i<order; i++ {
		com := strings.Compare(s[i], t[i])
		if com == 0 {
			continue
		} else if com < 0 {
			fmt.Println("Yes")
			return
		} else {
			fmt.Println("No")
			return
		}
	}

	fmt.Println(a)
}
