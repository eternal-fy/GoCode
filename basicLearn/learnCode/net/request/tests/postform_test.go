package tests

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNet(t *testing.T) {
	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}
	u := "http://106.12.174.72"
	//values := url.Values{}
	//values.Add("hehe", "meme")
	//form, err := client.PostForm(u, values)
	request, _ := http.NewRequest(http.MethodGet, u, nil)
	do, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	fmt.Println(do)

}

func redirectPolicyFunc(req *http.Request, via []*http.Request) error {
	fmt.Println("--")
	return nil
}
