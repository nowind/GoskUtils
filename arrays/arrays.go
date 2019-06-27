package arrays

import (
	"github.com/nowind/GoskUtils/stringmap"
	"reflect"
)

func IsArray(a interface{})bool{
	return reflect.ValueOf(a).Kind()==reflect.Slice
}
func In(m interface{},a interface{}) bool{
	if IsArray(m){
		mVal:=reflect.ValueOf(m)
		b:=mVal.Len()
		for i:=0;i<b;i++{
			if mVal.Index(i).Interface()==a{
				return true
			}
		}
		return false
	}else{
		return false
	}
}
func InString(b []string,a string) bool{
		for _,v:=range b{
			if v==a{
				return true
			}
		}
		return false
}
func InInt(b []int,a int) bool{
	for _,v:=range b{
		if v==a{
			return true
		}
	}
	return false
}
func ToStringMap(a []string,f func (string) string) stringmap.StringMap{
	ret:=make(map[string]string)
	for _,i:=range a{
		ret[i]=f(i)
	}
	return ret
}
