package rest_api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

func GetRestLoggingString(param gin.LogFormatterParams) string {
	return GetLogString(param.ErrorMessage, &Request{
		Method:     param.Method,
		Url:        param.Path,
		Path:       param.Path,
		Parameters: param.Request.URL.Query(),
		ReqId:      param.Request.Header.Get("x-request-id"),
	}, &Response{
		StatusCode:   param.StatusCode,
		ResponseTime: float64(param.Latency.Microseconds()) / 1000,
	})
}

func GetLogString(message string, request *Request, response *Response) string {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "not_defined"
	}
	b, _ := json.Marshal(LogRecord{
		Level:    30,
		Time:     time.Now().UnixMilli(),
		Pid:      os.Getpid(),
		Hostname: hostname,
		Req:      request,
		Res:      response,
		Message:  message,
	})
	return string(b) + "\n"
}
