package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Truy cập vào đường dẫn: /
	// Trả về trang chủ với dòng chữ: Đây là trang chủ
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Đây là trang chủ")
	})

	// Truy cập vào đường dẫn: /about
	// Trả về trang chủ với dòng chữ: Đây là trang giới thiệu
	r.GET("/about", func(c *gin.Context) {
		c.String(http.StatusOK, "Đây là trang giới thiệu")
	})

	// Chạy web server ở cổng 8080
	r.Run(":8080")
}