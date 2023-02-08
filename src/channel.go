package main

import (
	"fmt"
	"math"
)

func primeNumber() chan int {
	result := make(chan int)
	go func() {
		result <- 2
		for i:=3; i<100; i+=2 {
			l := int(math.Sqrt(float64(i)))
			found := false
			for j:=3; i<l; j+=2 {
				if i%j == 0 {
					found = true
					break
				}
			}
			if !found {
				result <- i
				fmt.Println("b")
			}
		}
		close(result)
	}()
	return result
}

func main() {
	pn := primeNumber()
	fmt.Println("a")
	for n := range pn {
		fmt.Println(n)
	}
}
