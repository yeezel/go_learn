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
	"io"
	"net/http/httptest"
	"testing"
)

// httptest, 针对 http 开发的场景
func TestConn(t *testing.T) {
	req := httptest.NewRequest("GET", "http://baidu.com", nil)
	w := httptest.NewRecorder()
	// helloHandler(w, req)
	bytes, _ := io.ReadAll(w.Result().Body)

	if string(bytes) != "hello world" {
		t.Fatal(req.URL, " expected hello world, but got", string(bytes))
	}
}
