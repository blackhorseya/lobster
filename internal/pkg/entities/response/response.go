package response

import "encoding/json"

// Response declare unite api response format
type Response struct {
	Code int         `json:"code,omitempty"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

// WithMessage append message into response
func (resp *Response) WithMessage(message string) *Response {
	return &Response{
		Code: resp.Code,
		Msg:  message,
		Data: resp.Data,
	}
}

// WithData append data into response
func (resp *Response) WithData(data interface{}) *Response {
	return &Response{
		Code: resp.Code,
		Msg:  resp.Msg,
		Data: data,
	}
}

func (resp *Response) String() string {
	ret, _ := json.Marshal(resp)
	return string(ret)
}
