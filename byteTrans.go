package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr"
)

func main() {
	r := gin.Default()
	r.GET("/", Handle)
	r.Run(":8888")
}

func Handle(context *gin.Context) {
	writer := context.Writer
	box := packr.NewBox("")
	content, _ := box.FindString("hello.html")
	/*header := writer.Header()
	header.Set("")*/
	writer.Write([]byte(content))
}
