package main
import (
	"fmt"
)


//do..while
func main() {

	var i int = 1

	for {
		if i > 10 {
			break
		}

		fmt.Println("hello world!",i)
		i++
	}
	fmt.Println("result:",i)

	fmt.Println("*******************************************")

	i = 1
	for {
		fmt.Println("hello world!",i)
		i++
		if i > 10 {
			break
		}
	}
	fmt.Println("result:",i)
}