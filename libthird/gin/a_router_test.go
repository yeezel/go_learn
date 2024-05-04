package gin

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRouter(t *testing.T) {
	r := gin.Default() //获取gin引擎
	//配置路由
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Geektutu")
	})
	// 动态路由，格式：/user/:name/*role，* 代表可选
	// 请求示例：curl http://localhost:9999/user/geektutu -> Hello geektutu
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})
	// 匹配users?name=xxx&pwd=xxx，pwd可选
	// 	$ curl "http://localhost:9999/users?name=Tom&pwd=student"
	// Tom is a student
	r.GET("/users", func(c *gin.Context) {
		name := c.Query("name")
		role := c.DefaultQuery("pwd", "teacher") //如果没有pwd参数则使用默认值
		c.String(http.StatusOK, "%s is a %s", name, role)
	})

	// POST
	// 	$ curl http://localhost:9999/form  -X POST -d 'username=geektutu&password=1234'
	// {"password":"1234","username":"geektutu"}
	r.POST("/form", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.DefaultPostForm("password", "000000") // 可设置默认值

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})

	// GET 和 POST 混合 请求
	// 	$ curl "http://localhost:9999/posts?id=9876&page=7"  -X POST -d 'username=geektutu&password=1234'
	// {"id":"9876","page":"7","password":"1234","username":"geektutu"}
	r.POST("/posts", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		username := c.PostForm("username")
		password := c.DefaultPostForm("username", "000000") // 可设置默认值

		c.JSON(http.StatusOK, gin.H{
			"id":       id,
			"page":     page,
			"username": username,
			"password": password,
		})
	})

	// map参数
	// 	$ curl -g "http://localhost:9999/post?ids[Jack]=001&ids[Tom]=002" -X POST -d 'names[a]=Sam&names[b]=David'
	// {"ids":{"Jack":"001","Tom":"002"},"names":{"a":"Sam","b":"David"}}
	r.POST("/post", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		c.JSON(http.StatusOK, gin.H{
			"ids":   ids,
			"names": names,
		})
	})
	// 热加载
	//不填写参数，默认监听端口为8080
	r.Run(":8080")
}
