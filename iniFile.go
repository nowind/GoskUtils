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
func (self *IniFile) GetSeq(sec string) (map[string]string,[]string){
	s,err:=self.iniF.GetSection(sec)
	if err==nil{
		k:=make([]string,len(s.Keys()))
		for _,i:=range s.Keys(){
			k=append(k,i.Name())
		}
		return s.KeysHash(),k
	}
	return make(map[string]string),[]string{}
}

func (self *IniFile) Get(sec string) map[string]string {
	r,_:=self.GetSeq(sec)
	return r
}