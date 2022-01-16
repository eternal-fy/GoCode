package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
	"strconv"
)

func main() {

	uploadSingMulFile("other")
}
func uploadSingMulFile(root string) {
	r := gin.Default()
	//限制上传最大尺寸
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		pathName := c.PostForm("pathName")
		os.MkdirAll("./"+root+"/"+pathName, 0777)
		for i := 0; i < 9; i++ {
			file, err := c.FormFile("file" + strconv.Itoa(i))
			filename := file.Filename
			ext := path.Ext(filename)
			if flag := HandleEr(c, ext, err); flag {
				continue
			}
			c.SaveUploadedFile(file, "./"+root+"/"+pathName+"/"+strconv.Itoa(i)+".png")
		}
	})
	r.Run()
}
func HandleEr(c *gin.Context, ext string, err error) bool {
	count := 0
	if flag := IsPic(ext); !flag {
		c.String(500, "有些文件不是图片")
		count++
	}
	if err != nil {
		c.String(500, "上传图片出错")
		count++
	}
	if count > 0 {
		return true
	}
	return false
}

func IsPic(ext string) bool {
	picContainer := []string{".png", ".PNG", ".JEG", ".jpg", ".JEPG", ".jepg"}
	for _, con := range picContainer {
		if con == ext {
			return true
		}
	}
	return false
}

func uploadMulFile() {
	r := gin.Default()
	// 限制表单上传大小 8MB，默认为32MB
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
		}
		// 获取所有图片
		files := form.File["files"]
		// 遍历所有图片
		for _, file := range files {
			// 逐个存
			if err := c.SaveUploadedFile(file, "./"+file.Filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
				return
			}
		}
		c.String(200, fmt.Sprintf("upload ok %d files", len(files)))
	})
	//默认端口号是8080
	r.Run(":8000")
}
func uploadSingFile() {
	r := gin.Default()
	//限制上传最大尺寸
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(500, "上传图片出错")
		}
		// c.JSON(200, gin.H{"message": file.Header.Context})
		c.SaveUploadedFile(file, "d:/goupload/"+file.Filename)
		c.String(http.StatusOK, file.Filename)
	})
	r.Run()
}
