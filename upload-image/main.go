package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("template/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", "")
	})

	r.GET("/success", func(c *gin.Context) {
		filename := c.Query("file")
		src := fmt.Sprintf(`/public/%s`, filename)

		// Trả về file success.html
		// Truyền tham số src chứa đường dẫn ảnh vào thuộc tính src của thẻ img
		// Khi đó client sẽ gửi request lên /public/:file để lấy ảnh
		c.HTML(http.StatusOK, "success.html", gin.H{
			"src": src,
		})
	})

	r.GET("/fail", func(c *gin.Context) {
		c.HTML(http.StatusInternalServerError, "fail.html", "")
	})

	// Trả về file ảnh cho client
	r.GET("/public/:file", func(c *gin.Context) {
		fileName := c.Param("file")
		filePath := path.Join("./public", fileName)

		var contentType string
		if strings.HasSuffix(fileName, ".png") {
			contentType = "image/png"
		} else {
			contentType = "image/jpeg"
		}

		data, err := ioutil.ReadFile(filePath)

		if err != nil {
			c.String(http.StatusInternalServerError, "Lỗi server")
			return
		}

		w := c.Writer
		w.Header().Add("Content Type", contentType)
		w.Write(data)
	})

	r.POST("/upload-image", func(c *gin.Context) {
		// Đọc dữ liệu file gửi lên từ thuộc tính name của thẻ input
		// Ở đây thuộc tính name có giá trị = file
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, "Bad request")
		}

		filename := path.Join("./public", file.Filename)
		// Lưu file vừa đọc vào đường dẫn filename
		// Nếu có lỗi trong quá trình lưu thì redirect sang đường dẫn /fail và return
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.Request.URL.Path = "/fail"
			c.Request.Method = "GET"
			r.HandleContext(c)
			return
		}

		// Sau khi lưu thành công, redirect sang đường dẫn /success
		redirectUrl := fmt.Sprintf(`/success`)
		c.Request.URL.Path = redirectUrl

		// Truyền vào query: ?file=[tên file vừa được lưu]
		// Lúc này redirect path sẽ có dạng: /success?file=[tên file vừa được lưu]
		query := fmt.Sprintf(`file=%s`, file.Filename)
		c.Request.URL.RawQuery = query
		c.Request.Method = "GET"
		r.HandleContext(c)
	})

	r.Run(":8080")
}
