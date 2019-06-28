package skUtils

import "github.com/go-ini/ini"

type IniFile struct {
	iniF *ini.File
}

func  NewIni(path string)  *IniFile{
	f,err:=  ini.Load(path)
	if err!=nil{
		panic("load file error")
	}
	ret:=new(IniFile)
	ret.iniF=f
	return ret
}
func (self *IniFile) ReGetSeq(sec string) (map[string]string,[]string){
	if self.iniF.Reload()!=nil{
		return make(map[string]string),[]string{}
	}
	m,s:=self.GetSeq(sec)
	return m,s
}
func (self *IniFile) GetSeq(sec string) (map[string]string,[]string){
	s,err:=self.iniF.GetSection(sec)
	if err==nil{
		k:=make([]string,len(s.Keys()))
		for i,v:=range s.Keys(){
			k[i]=v.Name()
		}
		return s.KeysHash(),k
	}
	return make(map[string]string),[]string{}
}

func (self *IniFile) Get(sec string) map[string]string {
	r,_:=self.GetSeq(sec)
	return r
}