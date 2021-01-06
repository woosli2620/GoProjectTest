package main

import "fmt"

//变量使用注意事项
func main() {
    var a int = 2
    var b = 3
    i := 10
    i = 30
    i = 50

    fmt.Println("i = ",i)

    //错误做法
    //i = 1.5
    fmt.Println("a = ",a)

    fmt.Println("b = ",b)
}
