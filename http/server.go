package main

import (
	"fmt"
	"net/http"
)

func main() {
	get()
	http.HandleFunc("/hello", MyHandle)
	http.ListenAndServe("localhost:8888", nil)

}
func MyHandle(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("hehe"))
}

func MyHandle1(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("hehe"))
}

func get() {
	resp, _ := http.Get("https://www.baidu.com")
	body := resp.Body
	buf := make([]byte, 2048)
	body.Read(buf)
	fmt.Println(string(buf))

}
