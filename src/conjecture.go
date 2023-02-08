package main

import "fmt"

func main() {
	var n,i,count int64
	fmt.Scan(&n)
	count = 0

	for i=1; i<=n; i++ {
		for j:=i; j<=n; j++ {
			if n/(i*j)-j+1 > 0 {
				fmt.Println(i,j,count,n/(i*j)-j+1)
				count += n/(i*j)-j+1
			} else {
				if i==j {
					fmt.Println(count)
					return
				}
				break
			}
		}
	}
}
