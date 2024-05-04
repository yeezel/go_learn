package gin

import (
	"log"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestUpload(t *testing.T) {
	r := gin.Default() //获取gin引擎

	// 上传单个文件
	r.POST("/upload1", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		// c.SaveUploadedFile(file, dst)
		c.String(http.StatusOK, "%s uploaded!", file.Filename)
	})

	//上传多个文件
	r.POST("/upload2", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)
			// c.SaveUploadedFile(file, dst)
		}
		c.String(http.StatusOK, "%d files uploaded!", len(files))
	})

	// 热加载
	//不填写参数，默认监听端口为8080
	r.Run(":8080")
}
