package main
import (
	"fmt"
)


//流程控制
func main() {
	//输出10句  "你好，装备！"

	for i := 1;i<=10;i++ {
		fmt.Println("你好，CREC装备！",i)
	}

	j  := 1 
	for ;j<=10; {
		fmt.Println("你好，CREC装备！",j)
		j++
	}

	k := 1
	for{
		if k <= 10 {
			fmt.Println("ok~~~",k)
		} else {
			break
		}

		k++
	}


	var str string = "hello world!北京"
	str2 := []rune(str) //把str 转成[] rune
	for i := 0;i<len(str2);i++ {
		fmt.Printf("%c",str2[i])
	}

	fmt.Println("\n")

	// str = "abc~ok"
	// for index,val := range str {
	// 	fmt.Printf("index = %d,val = %c\n",index,val)
	// }


}