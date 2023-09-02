package main

import (
	"fmt"
	"io"
	"log"
)

func main() {
	fmt.Printf("In function1 at the top\n")
	// 类似java的finally语句块，defer 定义在函数结束或return的时候才执行
	defer fmt.Println("Function2: Deferred until the end of the calling function!")
	fmt.Printf("In function1 at the bottom!\n")

	//当有多个 defer 行为被注册时，会以逆序执行（类似栈，即后进先出）
	for i := 0; i < 5; i++ {
		defer println(i)
	}

	//场景一：实现代码追踪
	b()
	//场景二：释放资源
	doDBOperations()

	//场景三：记录函数的参数与返回值
	func1("Go")
}

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

func connectToDB() {
	fmt.Println("ok, connected to db")
}

func disconnectFromDB() {
	fmt.Println("ok, disconnected from db")
}

func doDBOperations() {
	connectToDB()
	fmt.Println("Defering the database disconnect.")
	defer disconnectFromDB() //function called here with defer
	fmt.Println("Doing some DB operations ...")
	fmt.Println("Oops! some crash or network error ...")
	fmt.Println("Returning from function here!")
	return //terminate the program
	// deferred function executed here just before actually returning, even if
	// there is a return or abnormal termination before
}

func func1(s string) (n int, err error) {
	defer func() {
		log.Printf("func1(%q) = %d, %v", s, n, err)
	}()
	return 7, io.EOF
}
