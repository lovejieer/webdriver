package webdriver

import (
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
)

const (
	POST = "POST"
	GET = "GET"
)
type Webdriver interface {
	Start() error
	Stop() error
	NewSession() (*Session, error)


}

type Capabilities map[string]interface{}
type params map[string]interface{}

type Session struct {
	Id           string
	Capabilities Capabilities
	RequestUrl string
}

func (s *Session) Request(p params, method string, url string) (string ,[]error){
	jsonParams, _ := json.Marshal(p)
	request := gorequest.New()
	if method == POST{
		_, body, errs := request.Post(s.RequestUrl).
		Send(string(jsonParams)).
		End()
		return body, errs
	}else {
		_, body, errs := request.Get(s.RequestUrl).
			End()
		return body, errs
	}
}


func (s *Session) Url(url string) string{
	ru := fmt.Sprintf("http://127.0.0.1:1234/session/%s/url",s.Id)
	p := params{"url": url}
	body,_ := s.Request(p,POST,ru)
	return body
}

func (s *Session) Source() string{
	ru := fmt.Sprintf("http://127.0.0.1:1234/session/%s/source",s.Id)
	body,_ := s.Request(nil,GET,ru)
	return body
}

func (s *Session) GetUrl() string{
	ru := fmt.Sprintf("http://127.0.0.1:1234/session/%s/url",s.Id)
	body,_ := s.Request(nil,GET,ru)
	return body
}
