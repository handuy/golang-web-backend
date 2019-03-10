package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
	"fmt"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("template/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", "")
	})

	r.POST("/upload-form", func(c *gin.Context){
		// Đọc dữ liệu từ thuộc tính name của thẻ <input> trong form gửi lên
		name := c.PostForm("name")
		height := c.PostForm("height")
		weight := c.PostForm("weight")

		heightInt, _ := strconv.Atoi(height)
		weightInt, _ := strconv.Atoi(weight)
		bmi := float64(weightInt) / (float64(heightInt) * float64(heightInt))

		result := fmt.Sprintf("Xin chào %s. Chỉ số BMI của bạn là %f", name, bmi)
		c.String(http.StatusOK, result)
	})

	r.Run(":8080")
}