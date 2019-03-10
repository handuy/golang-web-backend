package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("template/*")

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Truy cập vào /language?name=[tên ngôn ngữ] để render HTML với nội dung tương ứng")
	})

	r.GET("/language", func(c *gin.Context) {
		languageName := c.Query("name")
		if languageName == "" {
			languageName = "go"
		}
		c.HTML(http.StatusOK, "language.html", gin.H{
			"language": languageName,
		})
	})

	r.Run(":8080")
}