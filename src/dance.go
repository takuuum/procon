package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := fmt.Sprintf("%b", 1)
	ss,_ := strconv.Atoi(s)
	t := fmt.Sprintf("%b", 5)
	tt,_ := strconv.Atoi(t)
	l := strconv.Itoa(ss^tt)
	ll,_ := strconv.ParseInt(l, 2, 0)
	fmt.Println(ll)
}
