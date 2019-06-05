package main

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
func (self *IniFile) get(sec string ) map[string]string{
	s,err:=self.iniF.GetSection(sec)
	if err==nil{
		return s.KeysHash()
	}
	return make(map[string]string)
}
