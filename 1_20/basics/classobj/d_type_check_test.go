package classobj

import (
	"fmt"
	"testing"
)

type human interface {
	say()
}

type man struct {
}
type woman struct {
}

func (m man) say() {
	fmt.Println("me man")
}

func (w woman) say() {
	fmt.Println("me womain")
}

func TestTypeCheck(t *testing.T) {
	var hm human = new(man)
	// 接口才能判定类型
	// 接口判定类型的hm必须是一个接口变量
	if v, ok := hm.(*man); ok {
		v.say()
	}

	switch t := hm.(type) {
	case *man:
		println("he is man")
	case *woman:
		println("she is woman")
	case nil:
		println("that is not human")
	default:
		println("this is %s human", t)
	}
}
