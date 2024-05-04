package basics

import "fmt"

// 声明一个myFunc类型，是一个函数类型，接受一个int类型的参数，返回一个int类型的值
type myFunc func(int) int

// 定义一个double函数，实现了myFunc类型
// 只有在参数和返回值类型完全匹配的情况下，函数才会被视为实现了 myFunc 类型
func double(x int) int {
	return x * 2
}

// 定义一个apply函数，接受一个myFunc类型的参数和一个int类型的参数
func apply(f myFunc, x int) int {
	// 调用f对x进行处理
	return f(x)
}

func TypeFuncDemo() {
	// 创建一个myFunc类型的变量f，并赋值为double
	f := myFunc(double) // 等同于f := double
	// 调用apply函数，传入f和10
	result := apply(f, 10)
	// 打印结果
	fmt.Println(result) // 输出20
}
