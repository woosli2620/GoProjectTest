package main

import (
	"fmt"
)

func main () {
	//打印空心金字塔

	var totalLevel = 10

	for i := 1;i<=totalLevel;i++ {

		for j := 1;j<=i;j++ {
			fmt.Printf("%d * %d = %d\t",i,j,i*j)
		}

		fmt.Print("\n")
	}
}