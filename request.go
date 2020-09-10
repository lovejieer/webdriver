package webdriver

import (
	"encoding/json"
	"github.com/parnurzeal/gorequest"
)

type Request struct {
	url string

}

func (r *Request) Post(url string, p params) string{
	jsonParams, _ := json.Marshal(p)
	request := gorequest.New()
		_, body, _ := request.Post(url).
			Send(string(jsonParams)).
			End()
		return body

}

func (r *Request) Get(url string) string{
	request := gorequest.New()
	_, body, _ := request.Get(url).
		End()
	return body

}

