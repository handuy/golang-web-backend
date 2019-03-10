package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var userList []User

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("template/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", "")
	})

	r.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.html", "")
	})

	// Sau khi đăng kí thành công, server sẽ redirect sang đường dẫn này
	r.GET("/signup-success", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup-success.html", "")
	})

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", "")
	})

	// Sau khi đăng nhập thành công, server sẽ redirect sang đường dẫn này
	r.GET("/login-success", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login-success.html", "")
	})

	// Nếu đăng nhập thất bại, server sẽ redirect sang đường dẫn này
	r.GET("/login-fail", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login-fail.html", "")
	})

	r.POST("/signup", func(c *gin.Context) {
		// Đọc dữ liệu từ thuộc tính name của thẻ <input> trong form gửi lên
		username := c.PostForm("username")
		password := c.PostForm("password")

		// Tạo user mới và lưu vào slice userList
		newUser := User{
			Username: username,
			Password: password,
		}
		userList = append(userList, newUser)

		// Sau khi đăng kí thành công, server redirect sang /signup-success
		c.Request.URL.Path = "/signup-success"
		c.Request.Method = "GET"
		r.HandleContext(c)
	})

	r.POST("/login", func(c *gin.Context) {
		// Đọc dữ liệu từ thuộc tính name của thẻ <input> trong form gửi lên
		username := c.PostForm("username")
		password := c.PostForm("password")

		checkResult := false

		for _, user := range userList {
			// Kiểm tra trong userList xem có cặp username và password hay ko
			if user.Username == username && user.Password == password {
				checkResult = true
			}
		}

		if checkResult {
			// Nếu có --> redirect sang /login-success
			c.Request.URL.Path = "/login-success"
			c.Request.Method = "GET"
			r.HandleContext(c)
		} else {
			// Nếu không có --> redirect sang /login-fail
			c.Request.URL.Path = "/login-fail"
			c.Request.Method = "GET"
			r.HandleContext(c)
		}
		
	})

	r.Run(":8080")
}
