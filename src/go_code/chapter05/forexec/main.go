package main

import (
	"fmt"
)

func main () {
	var max int = 100
	var j int = 1
	var sum int = 0

	for i := 1;i<max;i++ {
		if i % 9 == 0 {
			fmt.Println(i)
			j++
			sum += i
		}
	}
	fmt.Printf("个数为：%d,总和为 :%d\n",j,sum)

	for i := 0;i<100;i++ {
		fmt.Printf("%v + %v = %v \n",i,100-i,100)
	}
}