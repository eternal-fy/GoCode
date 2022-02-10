package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Fetch(url string) string {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("获取失败")
	}

	byte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("获取错误")
	}
	str := string(byte)
	return str

}
