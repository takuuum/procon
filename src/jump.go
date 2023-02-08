package main

import (
	"fmt"
	"strings"
)

func main() {
	var N,K int
	var S string
	count := 0
	jump := 0
	xcount := 0
	flag := 0
	fmt.Scan(&N,&K,&S)

	s := strings.Split(S, "")

	for i:=1; i<N; i++ {
		if s[i] == "." {
			count++
			if count == K {
				count = 0
				jump++
			}
			if s[i-1] == "X" {
				if xcount == K {
					jump++
				}
				xcount = 0
			}
			if flag == 1 {
				flag = 0
				jump++
			}
		} else if s[i] == "X" {
			count++
			xcount++
			if xcount == K || count == K {
				flag++
			}
			if flag == 2 {
				fmt.Println(-1)
				return
			}
		}
	}
	fmt.Println(jump+1)
}
