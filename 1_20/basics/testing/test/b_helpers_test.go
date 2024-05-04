package test

/*
1. 测试文件的文件名满足这种形式 *_test.go
2. 测试文件中必须导入 "testing" 包
3. 测试函数名字以 TestXxxxxxx开头
4. 测试函数只接受一个参数 t *testing.T
5. go test参数说明：
	-v 显示每个用例的测试结果
	-cover 参数可以查看覆盖率
	-run 指定函数测试，该参数支持通配符 *，和部分正则表达式，例如 ^、$
*/

import (
	"golearn/basics/testing/src"
	"testing"
)

// 帮助函数(helpers)
type calcCase struct{ A, B, Expected int }

func createMulTestCase(t *testing.T, c *calcCase) {
	//引入帮助函数，可以清晰帮我们定位哪个行代码调用该函数导致的报错
	t.Helper()
	if ans := src.Mul(c.A, c.B); ans != c.Expected {
		t.Fatalf("%d * %d expected %d, but %d got",
			c.A, c.B, c.Expected, ans)
	}

}

func TestMul1(t *testing.T) {
	createMulTestCase(t, &calcCase{2, 3, 6})
	createMulTestCase(t, &calcCase{2, -3, -6})
	createMulTestCase(t, &calcCase{2, 0, 1}) // wrong case
}
