package main

import (
	"strconv"
	"syscall/js"
	"time"
)

func main() {
	//demo1: hello world
	alert := js.Global().Get("alert")
	alert.Invoke("Hello World!")

	//创建了信道(chan) done，阻塞主协程(goroutine)。fibFunc 如果在 JavaScript 中被调用，会开启一个新的子协程执行
	done := make(chan int, 0)

	//demo2: 注册函数给js调用
	//使用 js.FuncOf 将函数转换为 Func 类型，只有 Func 类型的函数，才能在 JavaScript 中调用
	js.Global().Set("fibFunc", js.FuncOf(fibFunc))

	// demo3: dom添加点击事件
	btnEle := js.Global().Get("dombtn")
	btnEle.Call("addEventListener", "click", js.FuncOf(domFunc))

	// demo4: 注册回调函数
	js.Global().Set("callbackFunc", js.FuncOf(callbackFunc))
	<-done
}

/*
js.Value 可以将 Js 的值转换为 Go 的值
js.ValueOf，则用来将 Go 的值，转换为 Js 的值
*/
func fibFunc(this js.Value, args []js.Value) interface{} {
	return js.ValueOf(fib(args[0].Int()))
}

// 斐波那契数
func fib(i int) int {
	if i == 0 || i == 1 {
		return 1
	}
	return fib(i-1) + fib(i-2)
}

// DOM操作
func domFunc(this js.Value, args []js.Value) interface{} {
	document := js.Global().Get("document")
	numEle := document.Call("getElementById", "num")
	ansEle := document.Call("getElementById", "ans")
	v := numEle.Get("value")
	if num, err := strconv.Atoi(v.String()); err == nil {
		ansEle.Set("innerHTML", js.ValueOf(fib(num)))
	}
	return nil
}

// 注册回调函数
func callbackFunc(this js.Value, args []js.Value) interface{} {
	callback := args[len(args)-1]
	go func() {
		time.Sleep(3 * time.Second)
		v := fib(args[0].Int())
		callback.Invoke(v)
	}()

	js.Global().Get("ans").Set("innerHTML", "Waiting 3s...")
	return nil
}
