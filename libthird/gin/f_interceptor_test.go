package gin

import (
	"log"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestInterceptor(t *testing.T) {
	r := gin.Default() //获取gin引擎

	//中间件(Middleware)：拦截器
	// 作用于全局
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 作用于单个路由
	r.GET("/benchmark", gin.Logger(), gin.ErrorLogger())

	// 作用于某个组
	authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth())
	authorized.Use(Logger())
	{
		authorized.POST("/login", gin.Logger())
		authorized.POST("/submit", gin.Logger())
	}

	//自定义中间件
	// 见Logger函数

	// 热加载
	//不填写参数，默认监听端口为8080
	r.Run(":8080")
}

type MyBenchLogger func() gin.HandlerFunc

// 自定义中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// 给Context实例设置一个值
		c.Set("geektutu", "1111")
		// 请求前
		c.Next()
		// 请求后
		latency := time.Since(t)
		log.Print(latency)
	}
}
