package exception

import (
	"errors"
	"fmt"
	"testing"
)

// 定义一个 DivideError 结构
type DivideError struct {
	dividee int
	divider int
}

// go内置error接口定义, 只需实现 `Error` 接口即可
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

func TestErrorDemo(t *testing.T) {

	// 1. 单独定义错误，也可以直接用于返回值
	errNotFound := errors.New("math - square root of negative number")
	fmt.Printf("error: %v\n", errNotFound)
	err := fmt.Errorf("math: square root of negative number %g", 1.25)
	fmt.Printf("error: %v\n", err)

	// 2. 实现Error接口中的Error方法
	// 当除数为零的时候会返回错误信息
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}

}
