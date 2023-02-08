package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	a := 0
	aa := 0
	b := 0
	bb := 0
	c := 0
	cc := 0
	var flag, flags bool
	fmt.Scan(&s)
	ss := strings.Split(s, "")
	for i:=0; i<len(ss); i++ {
		if ss[i]=="B" {
			bb++
		} else if ss[i]=="C" {
			cc++
		}
	}
	if bb!=cc || bb==0 {
		fmt.Println("NO")
		return
	}
	num := bb
	for i:=0; i<len(ss); i++ {
		if i==len(ss)-1 {
			if ss[i]!="A" {
				break
			} else if num-1==aa {
				flag = true
				break
			}
		}
		if ss[i]=="A" {
			a++
			if b==c && b!=0 && !flags {
				aa++
				flags=true
			}
		} else if ss[i]=="B" {
			b++
			flags=false
			if i==0 || (a<b) {
				break
			}
		} else {
			c++
			flags=false
			if i==0 || b==0 || (b<c) {
				break
			}
		}
	}
	if flag {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
