package main

import (
	"log"
	"runtime"
)

// 实现一个 where() 闭包函数来打印函数执行的位置
func main() {
	//使用runtime
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}

	where()
	println("some code")
	where()
	println("some code-----")
	where()

	//使用log
	log.SetFlags(log.Llongfile)
	log.Print("this is log")

	var where1 = log.Print
	where1()
	println("some log code")
	where1()
	println("some log code-----")
	where1()
}
