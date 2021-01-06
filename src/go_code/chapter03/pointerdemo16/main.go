package main

import (
    "fmt"
)


//演示golang string转基本数据类型
func main() {
    var i int = 10

    fmt.Println("i的地址 = ",&i)

    var ptr *int = &i
    fmt.Printf("ptr = %v\n",ptr)
    fmt.Println("ptr的地址 = ",&ptr)
    fmt.Println("ptr的值是",*ptr)


    var num int = 80
    fmt.Println("num的地址是 :",&num)

    var ptr2 *int = &num
    *ptr2 = 20
    fmt.Println(num)
}

