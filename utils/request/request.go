package request

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type ReqOptions struct {
	Timeout int
	Proxy   string
	Headers map[string]string
}

// 定义选项函数类型
type Option func(*ReqOptions)

type Request struct {
	Api     string
	Method  string
	Data    map[string]interface{}
	Options *ReqOptions
}

type Response struct {
	Code int
	Msg  string
	Resp []byte
}

func NewReqOptions(timeout int, opts ...Option) *ReqOptions {
	options := &ReqOptions{
		Timeout: timeout,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
	for _, v := range opts {
		v(options)
	}

	return options
}

func WithTimeout(timeout int) Option {
	return func(o *ReqOptions) {
		o.Timeout = timeout
	}
}

func WithProxy(proxy string) Option {
	return func(o *ReqOptions) {
		o.Proxy = proxy
	}
}

func WithHeaders(headers map[string]string) Option {
	return func(o *ReqOptions) {
		o.Headers = headers
	}
}

func NewRequest(api, method string, data map[string]interface{}, opts *ReqOptions) *Request {
	return &Request{
		Api:     api,
		Method:  method,
		Data:    data,
		Options: opts,
	}
}

func (r *Request) Send() (data []byte, err error) {
	ts := &http.Transport{}
	if r.Options.Proxy != "" {
		// 设置代理
		ts.Proxy = func(_ *http.Request) (*url.URL, error) {
			return url.Parse(r.Options.Proxy)
		}
	}

	client := &http.Client{
		Transport: ts,
		Timeout:   5 * time.Second,
	}

	if r.Options.Timeout > 0 {
		client.Timeout = time.Duration(r.Options.Timeout) * time.Second
	}

	var dataReader io.Reader
	contentType := r.Options.Headers["Content-Type"]
	switch contentType {
	// case "application/x-www-form-urlencoded":
	case "application/json":
		// json 数据用 application/json
		jsonData, err := json.Marshal(r.Data)
		if err != nil {
			return nil, err
		}
		dataReader = bytes.NewReader(jsonData)
	default:
		// 默认为 application/x-www-form-urlencoded
		// 表单数据用 x-www-form-urlencoded
		url := url.Values{}
		for k, v := range r.Data {
			url.Add(k, v.(string))
		}
		dataReader = strings.NewReader(url.Encode())
		r.Options.Headers["Content-Type"] = "application/x-www-form-urlencoded"
	}

	req, err := http.NewRequest(r.Method, r.Api, dataReader)
	if err != nil {
		return nil, err
	}

	if len(r.Options.Headers) > 0 {
		for k, v := range r.Options.Headers {
			req.Header.Set(k, v)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	resp.Body.Read(data)
	return
}
