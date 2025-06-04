package main

import (
	"fmt"
	"net/http"

	"39.com/utils/request"
)

func main() {
	fmt.Println("Hello World")

	req := &request.Request{
		Api:    "http://127.0.0.1:8080/user/add",
		Method: http.MethodPost,
		Data: map[string]interface{}{
			"password": "123456",
			"username": "wuh",
		},
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	err := req.Send()
	if err != nil {
		fmt.Println("err : ", err)
	}
	fmt.Println("resp:", string(req.Resp))
}
