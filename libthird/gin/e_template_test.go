package gin

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestTemplate(t *testing.T) {
	r := gin.Default() //获取gin引擎

	// HTML模板
	type student struct {
		Name string
		Age  int8
	}

	/*
		<!-- templates/arr.tmpl -->
		<html>
		<body>
			<p>hello, {{.title}}</p>
			{{range $index, $ele := .stuArr }}
			<p>{{ $index }}: {{ $ele.Name }} is {{ $ele.Age }} years old</p>
			{{ end }}
		</body>
		</html>
	*/
	r.LoadHTMLGlob("templates/*")

	stu1 := &student{Name: "Geektutu", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}
	r.GET("/arr", func(c *gin.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gin.H{
			"title":  "Gin",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	// 热加载
	//不填写参数，默认监听端口为8080
	r.Run(":8080")
}
