package skUtils

import (
	"bufio"
	"github.com/json-iterator/go"
	"github.com/parnurzeal/gorequest"
	"net/url"
	"os"
	"strings"
)

type HARParser struct {
	Headers                           map[string][]string
	method, url, body, Postdata, mime string
	PostParmas, QueryString           map[string]string
	ok                                bool
}



func NewHARParser(path string) *HARParser {
	return NewHARParserWithNo(path,0)
}
func NewHARParserWithNo(path string, no int) *HARParser {
	f, err := os.Open(path)
	ret := new(HARParser)
	ret.ok = false
	if err != nil {
		return ret
	}
	defer f.Close()
	fr := bufio.NewReader(f)
	size := fr.Size()
	data := make([]byte, size)
	data = data[:size]
	fr.Read(data)
	if data[0] != '{' {
		return ret
	}
	ent := jsoniter.Get(data, "log", "entries", no, "request")
	if ent.LastError() == nil {
		ret.method = ent.Get("method").ToString()
		ret.url = ent.Get("url").ToString()
		// 以下要处理头
		headers:=ent.Get("headers")
		headersCount := headers.Size()
		ret.Headers= map[string][]string{}
		for i:=0;i<headersCount;i++ {
			_header:=headers.Get(i)
			name := _header.Get("name").ToString()
			value := _header.Get("value").ToString()
				if v, ok := ret.Headers[name]; ok {
					ret.Headers[name] = append(v, value)
				} else {
					ret.Headers[name] = []string{value}
				}
			}

		//处理url查询
		if ent.Get("queryString").ValueType() != jsoniter.InvalidValue {
			queryString := ent.Get("queryString")
			ret.QueryString = map[string]string{}
			for i:=0;i<queryString.Size();i++ {
				_queryItem:=queryString.Get(i)
				name:= _queryItem.Get("name").ToString()
				value := _queryItem.Get("value").ToString()
				ret.QueryString[name] = value

			}
		}

		if ent.Get("postData","mimeType").ValueType() != jsoniter.InvalidValue {
			ret.mime=ent.Get("postData","mimeType").ToString()
			if strings.Contains(ret.mime,"x-www-form-urlencoded"){
				pParmas:=ent.Get("postData","params")
				ret.PostParmas= map[string]string{}
				for i:=0;i<pParmas.Size();i++{
					_v:=pParmas.Get(i)
					name := _v.Get("name").ToString()
					value:= _v.Get("value").ToString()
					ret.PostParmas[name] = value

				}
				ret.Postdata=ret.ReGenPayload(ret.PostParmas)
			}else{
				ret.Postdata=ent.Get("postData","text").ToString()
			}
		}
	}
	ret.ok=true
	return ret
}

func (self *HARParser)ReGenPayload(m map[string]string) string{
	str:=strings.Builder{}
	for k,v:=range m{
		str.WriteString(url.QueryEscape(k))
		str.WriteString("=")
		str.WriteString(url.QueryEscape(v))
		str.WriteString("&")
	}
	ret:=str.String()
	return ret[:len(ret)-1]
}

func (self *HARParser)ReGenUrl(m map[string]string) string{
	url:=strings.Split(self.url,"?")
	return url[0]+"?"+self.ReGenPayload(m)
}

func (self *HARParser)IsOK() bool  {
	return self.ok
}
func (self *HARParser)getHeaderName(e string) (v []string,ok bool){
	if v,ok=self.Headers[e];ok{
		return v,ok
	}
	if v,ok=self.Headers[strings.ToLower(e)];ok{
		return v,ok
	}
	return nil,false
}

func (self *HARParser) GenEnv(request *gorequest.SuperAgent,genHeader []string) *gorequest.SuperAgent{
	if request==nil{
		request=gorequest.New().CustomMethod(strings.ToUpper(self.method),self.url);
	}
	if genHeader==nil{
		genHeader=[]string{}
	}
	defaultHeader:=[]string{"Cookie","User-Agent","Accept","Referer","Accept-Language","Accept-Encoding","Origin","X-Requested-With"}
	allHeader:=append(genHeader,defaultHeader...)
	for _,h:=range allHeader{
		if v,ok:=self.getHeaderName(h);ok{
			request.Set(h,strings.Join(v,";"))
		}
	}
	return request
}
func (self *HARParser)CopyPost() map[string]string {
	if self.PostParmas==nil{
		return nil
	}
	m:=make(map[string]string)
	for i,k :=range self.PostParmas{
		m[i]=k
	}
	return m
}
func (self *HARParser)CopyQuery() map[string]string {
	if self.QueryString==nil{
		return nil
	}
	m:=make(map[string]string)
	for i,k :=range self.QueryString{
		m[i]=k
	}
	return m
}
func (self *HARParser) Repeat(request *gorequest.SuperAgent,para map[string]string) * gorequest.SuperAgent {
	if request==nil {
		request=self.GenEnv(nil,nil)
	}else{
		request.Method=strings.ToUpper(self.method)
		request.Url=self.url
		request.Errors=nil
	}
	switch strings.ToLower(self.method) {
		case "post":
			pd:=self.Postdata
			if para!=nil {
				pd=self.ReGenPayload(para)
			}
			return request.Send(pd)
		case "get":
			url:=self.url
			if para!=nil{
				url=self.ReGenUrl(para)
			}
			request.Url=url
			return request
	}
	return request
}
func (self *HARParser) GetUrl() string{
	return strings.ToLower(self.url)
}
func (self *HARParser) UrlContains(s string) bool {
	return strings.Contains(self.GetUrl(),s)
}