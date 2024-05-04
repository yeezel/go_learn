package gin

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRedirect(t *testing.T) {
	r := gin.Default() //获取gin引擎

	//重定向(Redirect)
	// 	$ curl "http://localhost:9999/goindex"
	// Who are you?
	r.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/index")
	})

	r.GET("/goindex", func(c *gin.Context) {
		c.Request.URL.Path = "/"
		r.HandleContext(c)
	})

	// 热加载
	//不填写参数，默认监听端口为8080
	r.Run(":8080")
}
