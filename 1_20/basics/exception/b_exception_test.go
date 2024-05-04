package exception

import (
	"fmt"
	"os"
	"testing"
)

func TestExceptionDemo(t *testing.T) {

	//运行时异常处理
	// 最佳实践：
	// 	1）在包内部，总是应该从 panic 中 recover：不允许显式的超出包范围的 panic()
	// 	2）向包的调用者返回错误值。

	// 类似java的finally语句块，defer 定义在函数结束或return的时候才执行
	// 必须在panic之前定义defer，类似catch异常
	defer func() {
		//recover() 用于从 panic 或错误场景中恢复，停止终止过程进而恢复正常执行
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s\r\n", e)
		}
	}()
	//当有多个 defer 行为被注册时，会以逆序执行（类似栈，即后进先出）
	fmt.Printf("In function1 at the top\n")
	defer fmt.Println("Function2: Deferred until the end of the calling function!")
	fmt.Printf("In function1 at the bottom!\n")
	for i := 0; i < 5; i++ {
		defer println(i)
	}

	//panic() 在程序死亡时被调用。
	var user = os.Getenv("USER")
	if user == "" {
		//抛出异常
		panic("Unknown user: no value for $USER")
	}
	fmt.Printf("After bad call\r\n") // <-- would not reach
}
