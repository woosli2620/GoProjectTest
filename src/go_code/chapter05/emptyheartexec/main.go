package main

import (
	"fmt"
)

func main () {
	//打印空心金字塔

	var totalLevel = 10

	for i := 1;i<=totalLevel;i++ {

		for k := 1;k<totalLevel - i;k++ {
			fmt.Printf(" ")
		}
		for j := 1;j<=2*i -1;j++ {
			fmt.Print("*")
		}

		fmt.Print("\n")
	}
}