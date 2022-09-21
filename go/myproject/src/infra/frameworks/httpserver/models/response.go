package models

import (
	"encoding/json"
	"fmt"
)

type HttpResponse struct {
	Data         interface{} `json:"data,omitempty"`
	ErrorMessage string      `json:"error,omitempty"`
}

func NewHttpResponse() HttpResponse {
	return HttpResponse{}
}

func (r *HttpResponse) SetData(data interface{}) {
	r.Data = data
	r.ErrorMessage = ""
}

func (r *HttpResponse) To(dest interface{}) {
	s, _ := json.Marshal(r.Data)
	json.Unmarshal(s, dest)
}

func (r *HttpResponse) SetError(err error, messages ...string) {
	var message string
	if err != nil {
		message = err.Error()
	}
	for _, m := range messages {
		message = fmt.Sprintf("%s %s", message, m)
	}
	r.ErrorMessage = message
}

func (r *HttpResponse) String() string {
	if r.Data != nil {
		return fmt.Sprintf("%s", r.Data)
	}
	return "Response data is empty"
}
