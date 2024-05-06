package rest_api

type LogRecord struct {
	Level    int       `json:"level"`
	Time     int64     `json:"time"`
	Pid      int       `json:"pid"`
	Hostname string    `json:"hostname"`
	Message  string    `json:"message"`
	Req      *Request  `json:"req"`
	Res      *Response `json:"res"`
}

type Request struct {
	Method     string `json:"method"`
	Url        string `json:"url"`
	Path       string `json:"path"`
	ReqId      string `json:"reqId"`
	Parameters interface {
	} `json:"parameters"`
}

type Response struct {
	StatusCode   int     `json:"statusCode"`
	ResponseTime float64 `json:"responseTime"`
}
