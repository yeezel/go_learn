package classobj

import (
	"fmt"
	"testing"
)

type EA struct {
	a  int
	EB //使用匿名方式，相当于EA继承了EB的属性和方法
}

type EB struct {
	b int
}

func TestExtend(t *testing.T) {
	ea := &EA{}
	ea.b = 123
	fmt.Println(ea.b)
	ea.say()
}

func (eb *EB) say() {
	fmt.Println("this is EB ", eb.b, eb)
}
