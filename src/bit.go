package main

import (
	"fmt"
	"strconv"
)

func main() {
	var S,M string
	sum := 0
	fmt.Scan(&S)
	//sum, _ := strconv.Atoi(S)

	for bits:=0; bits<(1<<uint(len(S)-1)); bits++ {
		//fmt.Printf("%03b\n", bits)
		M = S
		count := 0
		for i:=0; i<len(S)-1; i++ {
			if (bits>>uint(i))&1 == 1 {
				fmt.Printf("%d %10b : %10b\n", i, bits,bits>>uint(i))
				j := i-count
				split, _ := strconv.Atoi(M[0:j+1])
				M = M[j+1:]
				count++
				sum += split
				fmt.Println("splitは", split, i)
			}
		}
		split, _ := strconv.Atoi(M)
		sum += split
		fmt.Println("Mは", M ,bits)
		//L := strings.Split(M, "+")
		//for i:=0; i<len(L); i++ {
		//	num, _ := strconv.Atoi(L[i])
		//	fmt.Println("num", num, "sp", M)
		//	sum += num
		//}
	}
	fmt.Println(sum)
}
