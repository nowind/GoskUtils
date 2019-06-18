package requtils

import (
	"github.com/json-iterator/go"
	"github.com/parnurzeal/gorequest"
)

type Response struct{
	gorequest.Response
}

var	BUF_LEN int64=5000

var buf=make([]byte,BUF_LEN,BUF_LEN)
func ResponseWrap(g gorequest.Response) *Response{
	return  &Response{g}
}
func (self *Response) Json() jsoniter.Any {
	if self.ContentLength>BUF_LEN{
		BUF_LEN=self.ContentLength
	}
	buf=buf[:BUF_LEN]
	if i,err:=self.Body.Read(buf);err!=nil{
		buf=buf[:i]
		return jsoniter.Get(buf)
	}
	return jsoniter.Wrap(nil)
}
