package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var XY string
	fmt.Scan(&XY)

	xy := strings.Split(XY, ".")

	y, _ := strconv.Atoi(xy[1])
	x := xy[0]

	if y<=2 {
		fmt.Printf("%s-", x)
	} else if y<=6 {
		fmt.Println(x)
	} else {
		fmt.Printf("%s+", x)
	}
}
