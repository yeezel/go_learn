package classobj

import (
	"fmt"
	"strconv"
	"testing"
)

/*
- 方法是函数，所以同样的，不允许方法重载
- 不同结构体可以有相同方法
- 结构体和绑定在它上面的方法的代码可以不放置在一起，它们可以存在在不同的源文件，唯一的要求是：它们必须是同一个包的
- 指针方法和值方法都可以在指针或非指针上被调用
*/
type TwoInts struct {
	a int //不是大写，仅在该文件中使用
	b int
}

type IntVector []int

type NamedPoint struct {
	TwoInts
	name string
}

// 仅绑定，不使用TwoInts类型变量可以使用【 _ 】占位符
func (_ TwoInts) say() string {
	return "hello"
}

// 该方法与TwoInts类型绑定，因为TwoInts是指针，所以这里的实例值改变，源实例的值会改变
func (tn *TwoInts) AddToParam(param int) int {
	return tn.a + tn.b + param
}

// 当调用一个函数或方法时，参数会被复制
// 这里的IntVector是一个实例，从源IntVector实例复制过来的，所以这里的实例值改变不会影响到源实例的值
func (v IntVector) Sum() (s int) {
	for _, x := range v {
		s += x
	}
	return
}

// 类型定义了 String() 方法（重写了string方法），它会被用在 fmt.Printf() 等输出方法中生成默认的输出
func (tn *TwoInts) String() string {
	return "(" + strconv.Itoa(tn.a) + "/" + strconv.Itoa(tn.b) + ")"
}

func TestMethodDemo(t *testing.T) {
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10

	fmt.Printf("two say: %s\n", two1.say())
	fmt.Printf("Add them to the param: %d\n", two1.AddToParam(20))

	//绑定在非结构体类型
	two2 := IntVector{3, 4}
	fmt.Printf("The sum is: %d\n", two2.Sum())
	two3 := &NamedPoint{TwoInts{3, 4}, "Pythagoras"}
	fmt.Printf("two3 say: %s\n", two3.say()) // 打印 5
}
