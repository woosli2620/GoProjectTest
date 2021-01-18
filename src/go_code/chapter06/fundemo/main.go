package main

import "fmt"

func main()  {
	//输入两个数，计算+，-，*，/
	var n1 float64
	var n2 float64

	var operator byte
	var res float64

	fmt.Println("请输入两个数和运算符，如：1+2")
	fmt.Scanf("%f %c %f",&n1,&operator,&n2)

	switch operator {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Println("操作符有误。。。")
	}

	fmt.Println("res = ",res)
}