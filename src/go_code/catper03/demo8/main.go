package main

import (
    "fmt"
    "unsafe"
)


//演示golang整数类型的使用
func main() {
    var i int = 1
    fmt.Println("i = ",i)

    //测试下int8 的范围
    //其他int16,int32,int64
    var j int8 = -128
    fmt.Println("j = ",j)

    // 测试uint8的范围
    var k uint8 = 255
    fmt.Println("k = ",k)


    var a int = 89900
    fmt.Println("a = ",a)

    var b uint = 1
    var c  byte = 255

    fmt.Println("a,b,c= ",a,b,c)


    //如何查看变脸的数据类型
    var n1  = 100
    fmt.Println("n1 = ",n1)
    fmt.Printf("n1的类型是 %T",n1)

    //查看占用大小
    var n2 int64 = 10
    fmt.Printf("n2 的类型是 %T，占用字节是 %d\n",n2,unsafe.Sizeof(n2))

    var age byte = 90
    fmt.Printf("age 的类型是 %T，占用字节是 %d\n",age,unsafe.Sizeof(age))
}
