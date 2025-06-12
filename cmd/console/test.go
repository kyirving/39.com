package main

import (
	"fmt"
	"net/http"

	"39.com/utils/request"
)

func main() {
	fmt.Println("Hello World")
	//接收异常
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	data := make(map[string]interface{})
	req := request.NewRequest("http://127.0.0.1/user/add", http.MethodPost, data, request.NewReqOptions(5))
	output, err := req.Send()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
}
