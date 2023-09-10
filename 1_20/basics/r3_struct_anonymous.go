package basics

import "fmt"

type innerS struct {
	in1 int
	in2 int
	b   float32
}

// 同一级别出现了两次，如果这个名字被程序使用了，将会引发一个错误（不使用没关系）
type A struct{ a int }
type B struct{ a int }

type outerS struct {
	b      int //外层名字会覆盖内层名字（但是两者的内存空间都保留）
	c      float32
	int    //匿名字段，类型就是字段的名字
	innerS //内嵌结构体
	A
	B
}

func AnonymousDemo() {
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60
	outer.in1 = 5
	outer.in2 = 10

	fmt.Printf("outer.b is: %d\n", outer.b)
	fmt.Printf("outer.c is: %f\n", outer.c)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.in1)
	fmt.Printf("outer.in2 is: %d\n", outer.in2)

	// 使用结构体字面量
	outer2 := outerS{6, 7.5, 60, innerS{5, 10, 3}, A{1}, B{2}}
	fmt.Println("outer2 is:", outer2)
}
