package skUtils

import (
	"bufio"
	"github.com/json-iterator/go"
	"github.com/parnurzeal/gorequest"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strings"
)

type HARParser struct {
	headers                           map[string][]string
	method, url, body, postdata, mime string
	postParmas, queryString           map[string]string
	ok                                bool
}



const (
	quickPathWindow = "d:\\tmp\\1.har"
	quickPathMac = "/User/nowind/Document/1.har"
)

func QuickLoad() *HARParser {
	return QuickLoadWithNo(0)
}
func QuickLoadWithNo(no int) *HARParser {
	quickPath:=quickPathWindow
	switch runtime.GOOS {
		case "darwin":
			quickPath=quickPathMac
			break
	}
	return NewHARParserWithNo(quickPath,no)
}
func NewHARParserWithNo(path string, no int) *HARParser {
	f, err := os.Open(path)
	ret := new(HARParser)
	ret.ok = false
	if err != nil {
		return ret
	}
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
		headers := ent.Get("headers").GetInterface().([]map[string]string)
		ret.headers= map[string][]string{}
		for _, header := range headers {
			name, ok1 := header["name"]
			value, ok2 := header["value"]
			if ok1 && ok2 {
				if v, ok := ret.headers[name]; ok {
					ret.headers[name] = append(v, value)
				} else {
					ret.headers[name] = []string{value}
				}
			}
		}
		//处理url查询
		if ent.Get("queryString").ValueType() != jsoniter.InvalidValue {
			queryString := ent.Get("queryString").GetInterface().([]map[string]string)
			ret.queryString = map[string]string{}
			for _, queryItem := range queryString {
				name, ok1 := queryItem["name"]
				value, ok2 := queryItem["value"]
				if ok1 && ok2 {
					ret.queryString[name] = value
				}
			}
		}

		if ent.Get("postData","mimeType").ValueType() != jsoniter.InvalidValue {
			ret.mime=ent.Get("postData","mimeType").ToString()
			if strings.Contains(ret.mime,"x-www-form-urlencoded"){
				pParmas:=ent.Get("postData","params").GetInterface().([]map[string]string)
				ret.postParmas= map[string]string{}
				for _,v:=range pParmas{
					name, ok1 := v["name"]
					value, ok2 := v["value"]
					if ok1 && ok2 {
						ret.postParmas[name] = value
					}
				}
				ret.postdata=ret.ReGenPayload(ret.postParmas)
			}else{
				ret.postdata=ent.Get("postData","text").ToString()
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
	if v,ok=self.headers[e];ok{
		return v,ok
	}
	if v,ok=self.headers[strings.ToLower(e)];ok{
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
func (self *HARParser) Repeat(request *gorequest.SuperAgent,para map[string]string) * http.Response {
	if request==nil {
		request=self.GenEnv(nil,nil)
	}else{
		request.Method=strings.ToUpper(self.method)
		request.Url=self.url
		request.Errors=nil
	}
	var resp * http.Response
	var err []error
	switch strings.ToLower(self.method) {
		case "post":
			pd:=self.postdata
			if para!=nil {
				pd=self.ReGenPayload(para)
			}
			resp,_,err=request.Send(pd).End()
			break
		case "get":
			url:=self.url
			if para!=nil{
				url=self.ReGenUrl(para)
			}
			request.Url=url
			resp,_,err=request.End()
			break
	}
	if err!=nil{
		return nil
	}
	return resp
}