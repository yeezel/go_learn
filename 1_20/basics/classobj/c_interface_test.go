package classobj

import (
	"fmt"
	"testing"
)

/*
- 类型不需要显式声明它实现了某个接口：接口被隐式地实现。
- 多个类型可以实现同一个接口。
- 一个类型可以实现多个接口。
- 只要类型实现了接口中的方法，它就实现了此接口。
- 一个接口可以包含一个或多个其他的接口
- 任何其他类型都实现了空接口
- 一个接口的值可以赋值给另一个接口变量，只要底层类型实现了必要的方法，若没有实现，则在运行时调用会报错
- 接口方法集的调用规则：
  - 类型 *T 的可调用方法集包含接受者为 *T 或 T 的所有方法集
  - 类型 T 的可调用方法集包含接受者为 T 的所有方法
  - 类型 T 的可调用方法集不包含接受者为 *T 的方法
*/
type Any1 interface{} //空接口

type One interface{ a() }
type Two interface{ b() }

type Phone interface { //定义接口
	// One //可以包含多个接口
	// Two
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() { //接口实现
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone *IPhone) call() { //接口实现
	fmt.Println("I am iPhone, I can call you!")
}

func TestInterfaceDemo(t *testing.T) {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

	// 接口判定类型的phone必须是一个接口变量
	if v, ok := phone.(*IPhone); ok {
		v.call()
	}

	switch t := phone.(type) {
	case *NokiaPhone:
		println("this is NokiaPhone")
	case *IPhone:
		println("this is IPhone")
	case nil:
		println("this is not Phone")
	default:
		println("this is %s Phone", t)
	}

}
