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
	"fmt"
	"os"
	"testing"
)

func setup() {
	fmt.Println("Before all tests")
}

func teardown() {
	fmt.Println("After all tests")
}

func Test1(t *testing.T) {
	fmt.Println("I'm test1")
}

func Test2(t *testing.T) {
	fmt.Println("I'm test2")
}

// 如果测试文件中包含函数 TestMain，那么生成的测试将调用 TestMain(m)，而不是直接运行测试。
// go test c_testmain_test.go  -v
func TestMain(m *testing.M) {
	setup()
	// 调用 m.Run() 触发所有测试用例的执行，并使用 os.Exit() 处理返回的状态码，如果不为0，说明有用例失败。
	code := m.Run()
	teardown()
	os.Exit(code)
}
