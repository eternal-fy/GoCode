package demo

import (
	"fmt"
	"net/rpc"
)

func Client() {
	myrequest := Myrequest{1, 2}
	var response Myresponse
	conn, err := rpc.DialHTTP("tcp", ":8081")
	if err != nil {
		fmt.Printf("clienterr:  %v", err)
	}
	fmt.Println("开始远程调用--------")
	conn.Call("Mymethod.Multiply", myrequest, &response)
	fmt.Println("ans=", response.Ans)

}
