package main

//
import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Student struct {
	Types, Username, Password string
}

func main() {
	r := gin.Default()
	r.POST("/form", func(c *gin.Context) {
		types := c.DefaultPostForm("type", "post")
		username := c.PostForm("username")
		password := c.PostForm("userpassword")
		// c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))

		marshal, _ := json.Marshal(Student{types, username, password})
		c.String(http.StatusOK, string(marshal))
	})
	r.Run()
}
