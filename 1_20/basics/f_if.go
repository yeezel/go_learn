package basics

import "fmt"

func IfDemo() {
	var age int = 23
	if age == 25 {
		fmt.Println("true")
	} else if age < 25 {
		fmt.Println("too small")
	} else {
		fmt.Println("too big")
	}

	//初始化
	if val := 10; val > 5 {
		fmt.Println("yes")
	}
}
