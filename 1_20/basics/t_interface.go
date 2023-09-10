package basics

import (
	"fmt"
)

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

func InterfaceDemo() {
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
