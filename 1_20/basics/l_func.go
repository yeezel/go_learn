package basics

import "fmt"

/*
func：函数由 func 开始声明
function_name：函数名称，参数列表和返回值类型构成了函数签名。
parameter list：参数列表，参数就像一个占位符，当函数被调用时，你可以将值传递给参数，这个值被称为实际参数，另外还可以将引用传递给参数，类似指针。参数列表指定的是参数类型、顺序、及参数个数。参数是可选的，也就是说函数也可以不包含参数。
return_types：返回类型，函数返回一列值。return_types 是该列值的数据类型，可以返回多个值。有些功能不需要返回值，这种情况下 return_types 不是必须的。
函数体：函数定义的代码集合。

func function_name( [parameter list] ) [return_types 或者 (return_types1,return_types2)] {
	函数体
	[return num1,str2...]
}
*/

var num int = 10
var numx2, numx3 int

func FuncDemo() {
	numx2, numx3 = getX2AndX3(num)
	PrintValues()
	numx2, numx3 = getX2AndX3_2(num)
	PrintValues()

	//匿名函数
	func() { println("匿名函数直接调用") }()
	pri := func() { println("匿名函数") }
	pri()

	//可变参数测试
	x := min(1, 3, 2, 0)
	fmt.Printf("The minimum is: %d\n", x)
	slice := []int{7, 9, 3, 5, 1}
	x = min(slice...)
	fmt.Printf("The minimum in the slice is: %d", x)
	println()

	// 调用printAll函数，传入不同类型和个数的值
	printAll(1, "hello", true, 3.14, []int{1, 2, 3})

	// 函数作为参数传递
	add := func(a, b int) {
		fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
	}
	callback(1, 2, add)

	//函数作为返回值，这里Adder函数内的参数x是一个共享变量
	var f = Adder()
	fmt.Print(f(1), " - ")  //1
	fmt.Print(f(20), " - ") //21
	fmt.Print(f(300))       //321
}

func PrintValues() {
	fmt.Printf("num = %d, 2x num = %d, 3x num = %d\n", num, numx2, numx3)
}

func getX2AndX3(input int) (int, int) {
	return 2 * input, 3 * input
}

// 尽量使用命名返回值：会使代码更清晰、更简短，同时更加容易读懂。
func getX2AndX3_2(input int) (x2 int, x3 int) {
	x2 = 2 * input
	x3 = 3 * input
	// return x2, x3
	//因为返回值已被赋值，所以直接return也可以
	return
}

// 直接改变外部reply参数的变量值
func Multiply(a, b int, reply *int) {
	*reply = a * b
}

// 可变参数
func min(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}

// 空接口，使用默认的空接口 interface{}，可接受任意长度任意类型的参数
// 定义一个printAll函数，参数是一个空接口类型的可变参数
func printAll(args ...interface{}) {
	// 遍历所有的参数
	for _, arg := range args {
		// 打印出参数的类型和值
		fmt.Printf("type: %T, value: %v\n", arg, arg)
	}
}

func callback(y int, z int, f func(int, int)) {
	f(y, z) // this becomes Add(1, 2)
}

func Adder() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}
