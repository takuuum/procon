package main

import (
	"fmt"
	"sort"
)

func main() {
	var S string
	fmt.Scan(&S)
	s := make([]string, len(S))

	for i:=0; i<len(S); i++ {
		s[i] = S[i:] + S[0:i]
	}
	sort.Strings(s)
	fmt.Println(s[0])
	fmt.Println(s[len(S)-1])
}
