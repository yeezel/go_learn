package basics

import (
	"fmt"
	"testing"
)

func TestConstDemo(t *testing.T) {
	//和变量定义一样，只是关键字`var`变成了`const`
	// - iota，特殊常量，可以被编译器修改的常量
	// 	- const 中每新增一行常量声明将使 iota 计数一次
	// 	- iota 在 const关键字出现时将被重置为 0
	// 	- iota 可以被用作枚举值
	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)
	const xx = iota //重置0
	const yy = iota
	fmt.Println(a, b, c, d, e, f, g, h, i, xx, yy) //0 1 2 ha ha 100 100 7 8 0 0
}
