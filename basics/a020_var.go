package main

import "fmt"

func main() {

	//变量声明格式：
	//	明确数据类型：	var 变量名 数据类型 = 值
	//	自动类型判断：	var 变量名 = 值
	//	简短声明：		变量名 := 值

	//类型相同多个变量, 非全局变量
	var a, b, c string
	b, c = "3", "a"
	fmt.Printf("a=%s,b=%s,c=%s", a, b, c)
	println()
	// 不需要显示声明类型，在编译得时候会自动推断
	var d, e = 1, "fgo"
	fmt.Printf("d=%v,e=%s", d, e)

	// 出现在 := 左侧的变量不能是已经被声明过的，否则会导致编译错误
	v1, v2, v3 := a, b, c
	println()
	fmt.Printf("v1=%s,v2=%s,v3=%s", v1, v2, v3)
	println()

	// 这种因式分解关键字的写法一般用于声明全局变量
	var (
		a1       = 15
		b1       = false
		str      = "Go says hello to the world!"
		numShips = 50
		city     string
	)

	fmt.Printf("v1=%v\nb1=%v\nstr=%v\nnum=%v\ncity=%v", a1, b1, str, numShips, city)
	println()
	//`_`为空白标识符，占位用，不能使用该变量，
	var _ = "gogogo"
}
