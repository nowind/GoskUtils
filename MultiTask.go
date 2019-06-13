package skUtils

type MultiTask struct {
	pool []*TimeInterval
	data []interface{}
	real RunnerFunc
	SHour,SMin,SSec int
	IntMicroSec,IntMicroSecBefore int
	Max int
	isStop,isFin bool
	main *TimeInterval
}
func realrun(obj interface{}) (nextint int,err error){
	self:=obj.(*MultiTask)
	for _,i:=range self.pool{
		if !i.IsStop(){
			return 0,nil
		}
	}
	self.isStop=true
	return -1,nil
}
func NewMultiTask(real RunnerFunc) *MultiTask{
	ret:=new(MultiTask)
	ret.SHour=10
	ret.SMin=0
	ret.SSec=0
	ret.Max=9999
	ret.IntMicroSec=1000
	ret.IntMicroSecBefore=1000
	ret.isStop=true
	ret.isFin=false
	ret.real=real
	ret.main=NewTimeInterval(realrun)
	return ret
}
func (self *MultiTask)RunWith(datas []interface{}) *MultiTask {
	self.pool=make([]*TimeInterval,0,10)
	self.data=make([]interface{},0,10)
	for _,i:=range datas{
		newone:=NewTimeInterval(self.real)
		newone.SHour=self.SHour
		newone.SMin=self.SMin
		newone.SSec=self.SSec
		newone.IntMicroSec=self.IntMicroSec
		self.pool=append(self.pool,newone)
		newone.RunWith(i)
	}
	self.main.SHour=self.SHour
	self.main.SMin=self.SMin
	self.main.SSec=self.SSec
	self.main.IntMicroSec=self.IntMicroSec
	self.main.IntMicroSecBefore=self.IntMicroSecBefore
	self.main.RunWith(self)
	return self
}
func (self *MultiTask)SetBefRun(run RunnerFunc) *MultiTask {
	self.main.SetBefRun(run)
	return self
}
func  (self *MultiTask)Stop() *MultiTask {
	for _,i:=range self.pool{
		i.Stop()
	}
	self.main.Stop()
	self.isStop=true
	return self
}