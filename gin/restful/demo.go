package main

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func main() {
	r := gin.Default()
	r.GET("/:username/*action", func(context *gin.Context) {
		param1 := context.Param("username")
		param2 := context.Param("action")
		//截取/
		param2 = strings.Trim(param2, "/") //把字符串前面和后面的/去掉
		println(param1)
		println(param2)
	})
	r.Run(":8888")
}
