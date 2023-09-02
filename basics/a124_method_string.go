package main

import (
	"fmt"
	"strconv"
)

type TwoInts1 struct {
	a int
	b int
}

func main() {
	two1 := new(TwoInts1)
	two1.a = 12
	two1.b = 10
	fmt.Printf("two1 is: %v\n", two1)
	fmt.Println("two1 is:", two1)
	fmt.Printf("two1 is: %T\n", two1)
	fmt.Printf("two1 is: %#v\n", two1)
}

// 类型定义了 String() 方法，它会被用在 fmt.Printf() 等输出方法中生成默认的输出
func (tn *TwoInts1) String() string {
	return "(" + strconv.Itoa(tn.a) + "/" + strconv.Itoa(tn.b) + ")"
}
