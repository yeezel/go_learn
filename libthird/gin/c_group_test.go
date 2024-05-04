package gin

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGroup(t *testing.T) {
	r := gin.Default() //获取gin引擎

	// 分组路由(Grouping Routes)，配置uri前缀
	// group routes 分组路由
	defaultHandler := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"path": c.FullPath(),
		})
	}
	// group: v1
	v1 := r.Group("/v1")
	{
		v1.GET("/posts", defaultHandler)
		v1.GET("/series", defaultHandler)
	}
	// group: v2
	v2 := r.Group("/v2")
	{
		v2.GET("/posts", defaultHandler)
		v2.GET("/series", defaultHandler)
	}

	// 热加载
	//不填写参数，默认监听端口为8080
	r.Run(":8080")
}
