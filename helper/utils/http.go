package utils

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"strings"
)

// HttpRequest 发送HTTP请求
func HttpRequest(method, url string, headers map[string]string, body interface{}) (interface{}, error) {
	var (
		resp   *resty.Response
		err    error
		client = resty.New()
		req    = client.R()
		data   = make(map[string]interface{})
	)

	// 设置全局请求头
	for key, value := range headers {
		req.SetHeader(key, value)
	}

	switch strings.ToUpper(method) {
	case "GET":
		resp, err = req.Get(url)
	case "POST":
		resp, err = req.SetBody(body).Post(url)
	case "PUT":
		resp, err = req.SetBody(body).Put(url)
	case "DELETE":
		resp, err = req.Delete(url)
	default:
		return "", fmt.Errorf("unsupported HTTP method: %s", method)
	}

	if err != nil {
		return "", err
	}

	if resp.StatusCode() != 200 {
		return "", fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode())
	}

	err = json.Unmarshal(resp.Body(), &data)
	if err != nil {
		return "", err
	}

	return data, nil
}
