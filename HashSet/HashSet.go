package HashSet

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
func (self *HashSet) AddAll(a []interface{}){
	for _,b:=range a{
		self.m[b]=true
	}
}
func (self *HashSet) AddAllString(a []string){
	for _,b:=range a{
		self.m[b]=true
	}
}
func (self *HashSet) AddAllInt(a []int){
	for _,b:=range a{
		self.m[b]=true
	}
}