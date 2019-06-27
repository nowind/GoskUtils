package stringmap

import "reflect"

type StringMap map[string]string

func New(m map[string]string) StringMap{
	return m
}
func NewPoint(m interface{}) StringMap{
	v2:=reflect.ValueOf(m)
	if v2.Kind()!=reflect.Map {
		return nil
	}
	newm:=map[string]string{}
	for i:=v2.MapRange();;i.Next(){
		if i !=nil{
			newm[i.Key().String()]=i.Value().String()
		} else {
			break
		}
	}
	return newm
}
func (self StringMap) Keys() []string {
	ret:=make([]string,len(self))
	i:=0
	for k,_:=range self{
		ret[i]=k
		i++
	}
	return ret
}
func (self StringMap) Values() []string {
	ret:=make([]string,len(self))
	i:=0
	for _,v:=range self{
		ret[i]=v
		i++
	}
	return ret
}
func (self StringMap) Contains(s string) bool {
	_,ok:=self[s]
	return ok
}
func (self StringMap) ContainsKey(s string) bool {
	for _,v:=range self {
		if v==s{
			return true
		}
	}
	return false
}