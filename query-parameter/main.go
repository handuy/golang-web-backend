package main

import (
	"fmt"
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
	// Trả về trang giới thiệu với dòng chữ: Đây là trang giới thiệu
	r.GET("/about", func(c *gin.Context) {
		c.String(http.StatusOK, "Đây là trang giới thiệu")
	})

	// Truy cập vào đường dẫn: /intro?name=gin
	// In giá trị của tham số name ra màn hình
	// name=gin --> in ra gin
	// name=golang --> in ra golang
	// Nếu truy cập vào /intro (không truyền tham số name) --> name là chuỗi rỗng
	r.GET("/intro", func(c *gin.Context) {
		paramValue := c.Query("name")
		result := fmt.Sprintf("Giá trị tham số name là %s", paramValue)
		c.String(http.StatusOK, result)
	})

	// Chạy web server ở cổng 8080
	r.Run(":8080")
}