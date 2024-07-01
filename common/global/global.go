package global

type Response struct {
	Code    int64       `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}
