package pkg

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type ReqOption struct {
	Api            string
	Method         string
	Data           map[string]interface{}
	Header         map[string]string
	Timeout        int
	ProxyHost      string
	ConnectTimeout int
	Response       string
	Msg            string
}

// 发送请求
func (req *ReqOption) sendRequest() {
	client := &http.Client{}

	// 设置超时时间
	if req.Timeout > 0 {
		client.Timeout = time.Duration(req.Timeout) * time.Second
	}

	// 设置代理
	if req.ProxyHost != "" {
		proxyUrl, err := url.Parse(req.ProxyHost)
		if err != nil {
			panic(err)
		}
		ts := &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		}
		client.Transport = ts
	}

	body := url.Values{}
	for k, v := range req.Data {
		switch v.(type) {
		case string, int, float64:
			body.Add(k, v.(string))
		case []string:
			//todo 处理数组
		}
	}
	request, err := http.NewRequest(req.Method, req.Api, strings.NewReader(body.Encode()))
	if err != nil {
		panic("request fail: " + err.Error())
	}

	if req.Header != nil {
		for k, v := range req.Header {
			request.Header.Add(k, v)
		}
	}

	resp, err := client.Do(request)
	if err != nil {
		panic("client send request fail: " + err.Error())
	}

	bodyByte, err := io.ReadAll(resp.Body)
	if err != nil {
		panic("read response fail: " + err.Error())
	}
	req.Response = string(bodyByte)

	defer func() {
		if err := recover(); err != nil {
			req.Msg = "请求失败"
		}

		resp.Body.Close()
	}()
}
