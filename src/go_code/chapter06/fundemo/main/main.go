package main

import (
	"fmt"
	AA "go_code/chapter06/fundemo/utils" //别名
)


func main()  {
	//输入两个数，计算+，-，*，/
	var n1 float64
	var n2 float64

	var operator byte

	fmt.Println("Num = ",AA.Num)

	fmt.Println("请输入两个数和运算符，如：1+2")
	fmt.Scanf("%f%c%f",&n1,&operator,&n2)

	fmt.Println("结果为：",AA.Cal(n1,n2,operator))
	
}