package main

import (
	"io/ioutil"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("template/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", "")
	})

	r.GET("/fail", func(c *gin.Context) {
		c.HTML(http.StatusInternalServerError, "fail.html", "")
	})

	// Trả về file audio cho client
	r.GET("/public/:file", func(c *gin.Context) {
		fileName := c.Param("file")
		filePath := path.Join("./public", fileName)

		contentType := "audio/mpeg"
		data, err := ioutil.ReadFile(filePath)

		if err != nil {
			c.String(http.StatusInternalServerError, "Lỗi server")
			return
		}

		w := c.Writer
		w.Header().Add("Content Type", contentType)
		w.Write(data)
	})

	r.POST("/upload-audio", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, "Bad request")
			return
		}
		files := form.File["files"]

		var fileListSlice []string
		for _, file := range files {
			fileListSlice = append(fileListSlice, file.Filename)
			filename := path.Join("./public", file.Filename)
			if err := c.SaveUploadedFile(file, filename); err != nil {
				c.Request.URL.Path = "/fail"
				c.Request.Method = "GET"
				r.HandleContext(c)
				return
			}
		}

		c.JSON(http.StatusOK, fileListSlice)
	})

	r.Run(":8080")
}
