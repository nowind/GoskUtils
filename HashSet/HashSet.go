package HashSet

import (
	"fmt"
	"reflect"
)

type HashSet struct {
	m map[interface{}]bool
}

func  New() *HashSet{
	return &HashSet{make(map[interface{}]bool)}
}
func (self *HashSet) Contains(a interface{}) bool{
	_,e:=self.m[a]
	return e
}
func (self *HashSet) Add(a interface{}){
	self.m[a]=true
}
func (self *HashSet) AddAll(a interface{}) error{
	if reflect.ValueOf(a).Kind()==reflect.Slice {
		b:=a.([]interface{})
		for _, c := range b {
			self.m[c] = true
		}
		return nil
	} else{
		return fmt.Errorf("is slice?")
	}
}
func (self *HashSet) AddAllInt(a []int){
		for _, c := range a {
			self.m[c] = true
		}
}
func (self *HashSet) AddAllString(a []string){
	for _, c := range a {
		self.m[c] = true
	}
}