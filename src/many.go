package main

import "fmt"

func main() {
	var S,T,a,b,c,M int
	fmt.Scan(&S,&T)
	count := 0


	if S<=T {
		M = S
	} else {
		M = T
	}

	for a=0; a<=M; a++ {
		for b=0; b<=M; b++ {
			//if a>b {continue}
			for c=0; c<=M; c++ {
				//if b>c{continue}
				if a+b+c <= S && a*b*c <= T {
					count++
				}
			}
		}
	}

	fmt.Println(count*6)
}
