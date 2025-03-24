package config

var httpLogs []interface{}

func SetHttpLog(log interface{}) {
	httpLogs = append(httpLogs, log)
}

func GetHttpLogs() []interface{} {
	return httpLogs
}

func ClearHttpLogs() {
	httpLogs = nil
}
