package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	count := 0
	ss := strings.Split(s, "")
	lens := len(ss)
	for i:=0;;i++ {
		if lens-1-i < 0 {
			break
		}
		if ss[i] == "a" && ss[lens-1-i] == "a" && count == 0 {
			continue
		} else if ss[lens-1-i] == "a" {
			count++
		} else {
			break
		}
	}
	for i:=0; i<lens/2; i++ {
		if lens-1-i-count < 0 {
			break
		}
		if ss[i] == ss[lens-1-i-count] {
			continue
		} else {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
