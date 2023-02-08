package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	a := strings.Split(s, "")
	var count int
	if a[0] == a[1] && a[1] == a[2] {
		count = 1
	} else if (a[0] == a[1] && a[1] != a[2]) || (a[0] != a[1] && a[1] == a[2]) || (a[0] == a[2] && a[1] != a[2]) {
		count = 3
	} else {
		count = 6
	}

	fmt.Println(count)
}
