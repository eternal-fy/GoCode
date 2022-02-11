package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {

	urlValues := url.Values{}
	urlValues.Add("user_id", "906542557")
	urlValues.Add("message", "hehe")
	form, err := http.PostForm("http://106.12.174.72:8081/send_private_msg", urlValues)
	fmt.Println(err)
	all, _ := ioutil.ReadAll(form.Body)
	fmt.Println(string(all))

}
