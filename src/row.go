package main

import (
	"fmt"
	"strings"
)

func main() {
	var S string
	count := 0
	fmt.Scan(&S)

	s := strings.Split(S, "")
	for i:=0; i<len(s); i++ {
		if s[i] == s[i+1] {
			count++
			if i==0 {
				if s[i] != s[i+2] {
					if s[i] == "0" {
						s[i] = "1"
					} else {
						s[i] = "0"
					}
				} else {
					if s[i+1] == "0" {
						s[i+1] = "1"
					} else {
						s[i+1] = "0"
					}
				}
			} else if i==len(s)-1 {
				if s[i] == s[i-1] {
					if s[i] == "0" {
						s[i] = "1"
					} else {
						s[i] = "0"
					}
				}
			} else {
				if s[i-1] == s[i] {
					if s[i] == "0" {
						s[i] = "1"
					} else {
						s[i] = "0"
					}
				} else {
					if s[i+1] == "0" {
						s[i+1] = "1"
					} else {
						s[i+1] = "0"
					}
				}
			}
		} else {
			continue
		}
	}

	fmt.Println(count)
}
