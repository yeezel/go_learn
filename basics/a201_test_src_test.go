package main

import "testing"

// 运行命令 go test a201_test_src_test.go a200_test_src.go
func TestEven(t *testing.T) {
	if Even(10) {
		t.Log(" 10 must be even!")
		t.Fail() //标记测试函数为失败，然后继续执行（剩下的测试）。
		// t.FailNow() //标记测试函数为失败并中止执行；文件中别的测试也被略过，继续执行下一个文件。
		// t.Fatal(" 10 must be even!") //先执行Log，再执行FailNow
	}
	if Even(7) {
		t.Log(" 7 is not even!")
		t.Fail()
	}

}

func TestOdd(t *testing.T) {
	if !Odd(11) {
		t.Log(" 11 must be odd!")
		t.Fail()
	}
	if Odd(10) {
		t.Log(" 10 is not odd!")
		t.Fail()
	}
}
