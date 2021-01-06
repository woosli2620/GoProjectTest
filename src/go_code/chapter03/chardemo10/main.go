package main

import (
    "fmt"
)


//演示golang字符类型的使用
//golang 没有专门的字符类型，用字节byte，字符串用多个字节标识
//英文1个字节，汉字3个字节
func main() {
    var c1 byte = 'a'
    var c2 byte = 'e'

    fmt.Println("c1 = ",c1)
    fmt.Println("c2 = ",c2)
    fmt.Printf("c1 = %c c2 = %c",c1,c2)

    var c3  int = '北'
    fmt.Printf("\nc3 = %c 对应的码字：",c3,c3)

    var c4 int = 22269
    fmt.Printf("c4=%c",c4)

    var n1 = 10 + 'a'
    fmt.Printf("\nn1 = ",n1)
}
