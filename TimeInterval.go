package skUtils

import (
	"time"
)

type RunnerFunc func (interface{}) (int,error)
type AOPFunc func ()
type TimeInterval struct {
	SHour,SMin,SSec int
	IntMillSec,IntMillSecBefore int
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
	ret.IntMillSec=0
	ret.IntMillSecBefore=0
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
	if self.IntMillSecBefore<1{
		self.IntMillSecBefore=1000
	}
 	errcount:=0
 	for beftimer:=time.NewTicker(time.Duration(self.IntMillSecBefore)*time.Millisecond);!self.isStop && !NowPass(self.SHour,self.SMin,self.SSec);{  //运行前处理逻辑
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
			if c>0 && c!=self.IntMillSecBefore{
				self.IntMillSecBefore=c
				beftimer.Stop()
				beftimer=time.NewTicker(time.Duration(self.IntMillSecBefore)*time.Millisecond)
				}
			}
		}
		<-beftimer.C
	}
 	if self.isStop{
 		return
	}
 	if self.beg !=nil {
 		self.beg()
	}
 	errcount=0
	if self.IntMillSec<1{
		self.IntMillSec=1000
	}
 	for runtimer,i:=time.NewTicker(time.Duration(self.IntMillSec)*time.Millisecond),0;i<self.Max;i++{
 		if self.isStop{ //运行前判断是否结束
			break
		}

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
			} else if result>0 && result!=self.IntMillSec{
				self.IntMillSec=result
				runtimer.Stop()
				runtimer=time.NewTicker(time.Duration(result)*time.Millisecond)
			}
		}
		<-runtimer.C
	}
	if self.end !=nil{
		self.end()
	}
}
func (self *TimeInterval) RunWith(data interface{}) *TimeInterval {
	self.isStop=false
	self.isFin=false
	go self.realRun(data)
	return self
}
func (self *TimeInterval) Run() *TimeInterval {
	return self.RunWith(nil)
}
func (self *TimeInterval) Stop(){
	self.isStop=true
}
func (self *TimeInterval) SetBeg(beg AOPFunc) *TimeInterval {
	self.beg=beg
	return self
}
func (self *TimeInterval) SetEnd(end AOPFunc) *TimeInterval {
	self.end=end
	return self
}
func (self *TimeInterval) SetBefRun(befrun RunnerFunc) *TimeInterval  {
	self.bf=befrun
	return self
}

func (self *TimeInterval) SetTimeInt(bef,run int) *TimeInterval  {
	self.IntMillSecBefore=bef
	self.IntMillSec=run
	return self
}
func  (self *TimeInterval) SetTime(h,m,s int) *TimeInterval  {
	self.SHour=h
	self.SMin=m
	self.SSec=s
	return self
}
func (self *TimeInterval) IsOK() bool{
	return self.isFin
}

func (self *TimeInterval) IsStop() bool{
	return self.isStop
}