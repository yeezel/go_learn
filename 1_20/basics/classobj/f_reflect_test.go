package classobj

import (
	"fmt"
	"reflect"
	"testing"
)

type NotknownType struct {
	s1, s2, S3 string
}

func (n NotknownType) String() string {
	return n.s1 + " - " + n.s2 + " - " + n.S3
}

func TestReflectDemo(t *testing.T) {
	// variable to investigate:
	var secret interface{} = NotknownType{"Ada", "Go", "Oberon"}
	// 获取实例类型
	fmt.Println("type:", reflect.TypeOf(secret)) // type: senior.NotknownType

	v := reflect.ValueOf(secret)
	// 获取所有值
	fmt.Println(v) // Ada - Go - Oberon
	// 获取实例类型
	fmt.Println("type:", v.Type()) // type: senior.NotknownType
	// 获取数据类型
	fmt.Println("kind:", v.Kind()) // kind: struct
	// 获取接口
	fmt.Printf("value is %5.2e\n", v.Interface()) // value is {%!e(string=   Ad) %!e(string=   Go) %!e(string=   Ob)}
	// 判断接口类型
	_, ok := v.Interface().(NotknownType)
	fmt.Println("type is NotknownType: ", ok) // type is NotknownType:  true
	// 获取所有字段
	for i := 0; i < v.NumField(); i++ {
		// Field 0: s1
		// Field 1: s2
		// Field 2: S3
		fmt.Printf("Field %d: %v\n", i, v.Type().Field(i).Name)
	}

	// 因为接口没有方法，所以获取到内置的String函数并调用
	results := v.Method(0).Call(nil)
	fmt.Println(results) // [Ada - Go - Oberon]

	// 修改字段值，只有被导出字段（首字母大写）才是可设置的
	s := NotknownType{"Ada", "Go", "Oberon"}
	v = reflect.ValueOf(&s).Elem()
	if v.CanSet() {
		// v.Field(0).SetString("aaa") //报错，只有暴露到包
		v.Field(2).SetString("aaa")
	}
}
