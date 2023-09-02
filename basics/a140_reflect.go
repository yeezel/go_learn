// blog: Laws of Reflection
package main

import (
	"fmt"
	"reflect"
)

type NotknownType struct {
	s1, s2, S3 string
}

func (n NotknownType) String() string {
	return n.s1 + " - " + n.s2 + " - " + n.S3
}

func main() {
	// variable to investigate:
	var secret interface{} = NotknownType{"Ada", "Go", "Oberon"}
	fmt.Println("type:", reflect.TypeOf(secret))
	v := reflect.ValueOf(secret)
	fmt.Println(v)
	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind())
	fmt.Printf("value is %5.2e\n", v.Interface())
	_, ok := v.Interface().(NotknownType)
	fmt.Println("type is NotknownType: ", ok)
	// iterate through the fields of the struct:
	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, v.Type().Field(i).Name)
	}

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
