package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Truy cập vào /old. Server sẽ điều hướng sang /new")
	})

	// Truy cập vào /old
	// Màn hình hiển thị dòng chữ: Đây là router /old
	// Sau đó server điều hướng request sang đường dẫn /new
	r.GET("/old", func(c *gin.Context) {
		c.String(200, "Đây là router /old \n")
		c.Request.URL.Path = "/new"
		r.HandleContext(c)
	})

	// Request đến đường dẫn /new
	// Màn hình in dòng chữ: Đây là router /new
	r.GET("/new", func(c *gin.Context) {
		c.String(200, "Đây là router /new")
	})

	// Request đến đường dẫn /google
	// Server điều hướng request sang địa chỉ: https://www.google.com/
	// Hiển thị trang chủ google
	r.GET("/google", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.google.com/")
	})

	r.Run(":8080")
}