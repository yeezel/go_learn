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

Fail()	标记测试函数为失败，然后继续执行（剩下的测试）。
FailNow()  标记测试函数为失败并中止执行；文件中别的测试也被略过，继续执行下一个文件。
Log(args ...interface{})  args 被用默认的格式格式化并打印到错误日志中。
Fatal(args ...interface{}) 结合 先执行 Log，然后执行 FailNow 的效果。
*/

import (
	"golearn/basics/testing/src"
	"testing"
)

// 运行命令 go test a201_test_src_test.go a200_test_src.go
func TestEven(t *testing.T) {
	if src.Even(10) {
		t.Log(" 10 must be even!")
		// 遇到错误不会停止
		t.Errorf("error")
	}
}

func TestOdd(t *testing.T) {

	//运行某个测试用例的子测试：go test -run TestOdd/pos -v
	t.Run("pos", func(t *testing.T) {
		if !src.Odd(11) {
			t.Log(" 11 must be odd!")
			// t.Fail() //标记测试函数为失败，然后继续执行（剩下的测试）。
			// t.FailNow() //标记测试函数为失败并中止执行；文件中别的测试也被略过，继续执行下一个文件。
			// t.Fatal(" 10 must be even!") //先执行Log，再执行FailNow
			t.Fail()
		}
	})

	// 提取通用判单
	assertError := func(t *testing.T, err bool, msg string) {
		if !err {
			t.Log(msg)
			t.Fail()
		}
	}
	t.Run("pos1", func(t *testing.T) {
		assertError(t, src.Odd(11), " 11 must be odd!")
	})
}

// 表格驱动测试在我们要创建一系列相同测试方式的测试用例时很有用
// 多个子测试的场景，更推荐如下的写法
func TestMul(t *testing.T) {
	cases := []struct {
		Name           string
		A, B, Expected int
	}{
		{"pos", 2, 3, 6},
		{"neg", 2, -3, -6},
		{"zero", 2, 0, 0},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if ans := src.Mul(c.A, c.B); ans != c.Expected {
				t.Fatalf("%d * %d expected %d, but %d got",
					c.A, c.B, c.Expected, ans)
			}
		})
	}
}
