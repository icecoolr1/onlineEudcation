package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

func main() {
	r := gin.Default() //启动gin路由，携带基础中间件启动 logger and recovery (crash-free) 中间件
	r.Use(CORSMiddleware())
	r.POST("testUpload", func(c *gin.Context) {
		must := uuid.NewV4().String()
		fmt.Println(must)
		file, _ := c.FormFile("file")
		filename := file.Filename
		fmt.Println(filename, "test")
		fileLocation := "d:/images/" + must + ".jpg"
		err := c.SaveUploadedFile(file, fileLocation)
		if err != nil {
			return
		}
		c.JSON(200, gin.H{
			"msg":          file,
			"fileLocation": fileLocation,
		})
	})
	r.Run(":1010") // listen and serve on 0.0.0.0:8080
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:8080") // 可将将 * 替换为指定的域名
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
