package basics

import (
	"fmt"
	"testing"
)

// 先调用init函数
func init() {
	fmt.Println("init main")
}

func TestInitDemo(t *testing.T) { // 注意：{ 必须与方法的声明放在同一行，这是编译器的强制规定
	fmt.Println("hello world")
}

func init() {
	fmt.Println("可以包含多个init，从上到下按顺序执行")
}

func Init() {
	fmt.Println("这个不会调用")
}
