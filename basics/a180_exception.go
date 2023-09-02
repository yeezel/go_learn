package main

import (
	"errors"
	"fmt"
	"os"
)

// 定义一个 DivideError 结构
type DivideError struct {
	dividee int
	divider int
}

// go内置error接口定义, 只需实现 `error` 接口即可
func (de *DivideError) Error() string {
	strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
`
	return fmt.Sprintf(strFormat, de.dividee)
}

// 定义 `int` 类型除法运算的函数
func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}

}

func main() {

	// 单独定义错误，也可以直接用于返回值
	errNotFound := errors.New("math - square root of negative number")
	fmt.Printf("error: %v\n", errNotFound)
	err := fmt.Errorf("math: square root of negative number %g", 1.25)
	fmt.Printf("error: %v\n", err)

	// 当除数为零的时候会返回错误信息
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}

	//运行时异常处理
	// 最佳实践：
	// 	1）在包内部，总是应该从 panic 中 recover：不允许显式的超出包范围的 panic()
	// 	2）向包的调用者返回错误值。

	//recover() 用于从 panic 或错误场景中恢复，停止终止过程进而恢复正常执行, 必须在panic之前defer
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s\r\n", e)
		}
	}()

	//panic() 在程序死亡时被调用。
	var user = os.Getenv("USER")
	if user == "" {
		panic("Unknown user: no value for $USER")
	}
	fmt.Printf("After bad call\r\n") // <-- would not reach
}
