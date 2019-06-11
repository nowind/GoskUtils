package skUtils

import (
	"time"
)

type RunnerFunc func (interface{}) (int,error)
type AOPFunc func ()
type TimeInterval struct {
	SHour,SMin,SSec int
	IntMicroSec,IntMicroSecBefore int
	Max int
	rf,bf RunnerFunc
	beg,end AOPFunc
	isStop,isFin bool
}

func NewTimeInterval(r RunnerFunc) *TimeInterval{
	ret:=new(TimeInterval)
	ret.rf=r
	ret.SHour=10
	ret.SMin=0
	ret.SSec=0
	ret.IntMicroSec=0
	ret.IntMicroSecBefore=0
	ret.Max=99999
	ret.bf=nil
	ret.beg=nil
	ret.end=nil
	ret.isStop=true
	ret.isFin=false
	return ret
}
func (self *TimeInterval) realRun(data interface{}){
	const MAX_ERROR  =  25
	if self.IntMicroSecBefore<1{
		self.IntMicroSecBefore=1000
	}
 	errcount:=0
 	for beftimer:=time.After(time.Duration(self.IntMicroSecBefore)*time.Microsecond);!self.isStop;{  //运行前处理逻辑
 		<-beftimer
 		if self.bf!=nil{ //有设置运行函数则执行
			c,err:=self.bf(data)
			if err !=nil { //运行后错误次数判断
				if errcount >= MAX_ERROR {
					self.isStop = true
					break
				} else {
					errcount += 1
				}
			}else{ //运行后间隔时间修正逻辑
			if c>0 && c!=self.IntMicroSecBefore{
				beftimer=time.After(time.Duration(c)*time.Microsecond)
				}
			}
		}
	}
 	if self.isStop{
 		return
	}
 	if self.beg !=nil {
 		self.beg()
	}
 	errcount=0
 	for runtimer,i:=time.After(time.Duration(self.IntMicroSec)*time.Microsecond),0;i<self.Max;i++{
 		if self.isStop{ //运行前判断是否结束
			break
		}
 		<-runtimer
 		result,err:=self.rf(data)
 		if err!=nil{
			if errcount >= MAX_ERROR {
				self.isStop = true
				break
			} else {
				errcount += 1
			}
		}else {
			if  result<0 {
				self.isFin=true
				self.isStop=true
				break
			} else if result>0 && result!=self.IntMicroSec{
				runtimer=time.After(time.Duration(result)*time.Microsecond)
			}
		}
	}
	if self.end !=nil{
		self.end()
	}
}
func (self *TimeInterval) RunWith(data interface{}) {
	go self.realRun(data)
}
func (self *TimeInterval) Run() {
	self.RunWith(nil)
}
func (self *TimeInterval) Stop(){
	self.isStop=true
}
func (self *TimeInterval) SetBeg(beg AOPFunc)  {
	self.beg=beg
}
func (self *TimeInterval) SetEnd(end AOPFunc)  {
	self.end=end
}
func (self *TimeInterval) SetBefRun(befrun RunnerFunc)  {
	self.bf=befrun
}
func (self *TimeInterval) IsOK() bool{
	return self.isFin
}

func (self *TimeInterval) IsStop() bool{
	return self.isStop
}