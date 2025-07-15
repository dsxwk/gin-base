package helper

import (
	"encoding/json"
	"fmt"
	"gin-base/common/global"
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

	userAgent := headers["User-Agent"]
	if userAgent == "" {
		req.SetHeader("User-Agent", userAgent)
	}

	// 设置全局请求头
	for key, value := range headers {
		req.SetHeader(key, value)
	}

	// 发布事件
	e := global.Config.Event
	e.Name = "sendHttp"
	e.Data = map[string]interface{}{
		"method": strings.ToUpper(method),
		"uri":    url,
		"header": headers,
		"agent":  userAgent,
		"body":   body,
	}
	global.Event.Publish(e)

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
