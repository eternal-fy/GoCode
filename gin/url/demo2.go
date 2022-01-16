package main

import (
	"github.com/gin-gonic/gin"
)

type Data struct {
	Path string `uri:"path" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/:path", func(c *gin.Context) {
		data := new(Data)
		err := c.ShouldBindUri(&data)
		if err != nil {
			c.String(500, "路径错误")
		}
		c.String(200, "你的路径为"+data.Path)
	})
	r.Run(":8080")

}
