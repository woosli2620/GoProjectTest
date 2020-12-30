
package main //每个文件必须属于一个包

import "fmt" //引入一个包，引入改包后就可以使用fmt包的函数，比如fmp.Println

func main() { //func是一个关键字，标识函数。main是函数名，主函数及函数入口
	fmt.Println("Hello World!")  //标识调用fmt的函数，输出Hello World！
	fmt.Println("Hello World!")  //标识调用fmt的函数，输出Hello World！
	fmt.Println("Hello World!")  //标识调用fmt的函数，输出Hello World！

	var num = 10 //变量定义完必须使用，否则会报错
	fmt.Println(num)

	fmt.Println("天龙八部\t张飞")
	fmt.Println("天龙八部\n张飞")
	fmt.Println("天龙八部\\张飞")
	fmt.Println("天龙八部\r张飞")
	fmt.Println("天龙八部\"张飞")
	
}
