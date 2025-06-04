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

type Request struct {
	Api        string
	Method     string
	Data       map[string]interface{}
	Timeout    int
	Headers    map[string]string
	ConnectOut int
	Proxy      string
	Code       int
	Msg        string
	Resp       []byte
}

func (r *Request) Send() error {
	ts := &http.Transport{}
	if r.Proxy != "" {
		// 设置代理
		ts.Proxy = func(_ *http.Request) (*url.URL, error) {
			return url.Parse(r.Proxy)
		}
	}
	client := &http.Client{
		Transport: ts,
		Timeout:   5 * time.Second,
	}

	var dataReader io.Reader
	contentType := r.Headers["Content-Type"]
	switch contentType {
	case "application/x-www-form-urlencoded":
		// 表单数据用 x-www-form-urlencoded
		url := url.Values{}
		for k, v := range r.Data {
			url.Add(k, v.(string))
		}
		dataReader = strings.NewReader(url.Encode())
	case "application/json":
		// json 数据用 application/json
		jsonData, err := json.Marshal(r.Data)
		if err != nil {
			r.Code = 500
			r.Msg = "Failed to encode JSON: " + err.Error()
			return err
		}
		dataReader = bytes.NewReader(jsonData)
	default:
		r.Headers["Content-Type"] = "application/x-www-form-urlencoded"
	}

	url := url.Values{}
	for k, v := range r.Data {
		url.Add(k, v.(string))
	}

	req, err := http.NewRequest(r.Method, r.Api, dataReader)
	if err != nil {
		r.Code = 500
		r.Msg = err.Error()
		return err
	}

	if len(r.Headers) > 0 {
		for k, v := range r.Headers {
			req.Header.Set(k, v)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		r.Code = 500
		r.Msg = err.Error()
		return err
	}
	defer resp.Body.Close()
	r.Code = resp.StatusCode
	r.Msg = resp.Status
	r.Resp, _ = io.ReadAll(resp.Body)
	return nil
}
