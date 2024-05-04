package basics

import (
	"fmt"
	"testing"
)

// 必须在全局定义
type myname string
type 你好 string

func TestTypeDemo(t *testing.T) {
	var myname = "hello"
	你好 := "世界"
	fmt.Println(myname, 你好) // hello 世界
}
