package webdriver

import (

	"fmt"

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
	r           *Request
}



func (s *Session) Url(url string) string{
	ru := fmt.Sprintf("http://127.0.0.1:1234/session/%s/url",s.Id)
	p := params{"url": url}
	body:= s.r.Post(ru,p)
	return body
}

func (s *Session) Source() string{
	ru := fmt.Sprintf("http://127.0.0.1:1234/session/%s/source",s.Id)
	body := s.r.Get(ru)
	return body
}

func (s *Session) GetUrl() string{
	ru := fmt.Sprintf("http://127.0.0.1:1234/session/%s/url",s.Id)
	body := s.r.Get(ru)
	return body
}
