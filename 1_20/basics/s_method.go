package basics

import "fmt"

type TwoInts struct {
	a int //不是大写，仅在该文件中使用
	b int
}

type IntVector []int

type NamedPoint struct {
	TwoInts
	name string
}

// 仅绑定，不使用TwoInts类型变量可以使用_
func (_ TwoInts) say() string {
	return "hello"
}

// 该方法与TwoInts类型绑定，因为TwoInts是指针，所以
func (tn *TwoInts) AddToParam(param int) int {
	return tn.a + tn.b + param
}

func (v IntVector) Sum() (s int) {
	for _, x := range v {
		s += x
	}
	return
}

func MethodDemo() {
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10

	fmt.Printf("two say: %s\n", two1.say())
	fmt.Printf("Add them to the param: %d\n", two1.AddToParam(20))

	//绑定在非结构体类型
	two2 := IntVector{3, 4}
	fmt.Printf("The sum is: %d\n", two2.Sum())
	two3 := &NamedPoint{TwoInts{3, 4}, "Pythagoras"}
	fmt.Println("two3 say: %s\n", two3.say()) // 打印 5
}
