package main

import (
	"fmt"
)

func main() {
	
	
	var i int = 1
	
	for i= 1;i<13;i++{
		
		if i == 10 {
			fmt.Println("Jump",i)
			continue

			}

		
		fmt.Println("i = ",i)
	}
}