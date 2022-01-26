package demo

import (
	"fmt"
	"net/http"
	"net/rpc"
)

type Myrequest struct {
	Param1, Param2 float64
}

type Mymethod struct {
}

func (mymethod *Mymethod) Multiply(request Myrequest, response *Myresponse) error {
	response.Ans = request.Param1 * request.Param2
	return nil
}

func (mymethod *Mymethod) Divide(request Myrequest, response *Myresponse) error {
	response.Ans = request.Param1 / request.Param2
	return nil
}

type Myresponse struct {
	Ans float64
}

func Server() {
	serve := new(Mymethod)
	rpc.Register(serve)
	rpc.HandleHTTP()

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Printf("servererr:  %v", err)
	}
}
