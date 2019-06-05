package main

import "time"

type RunnerFunc func (interface{}) int
type AOPFunc func ()
type TimeInterval struct {
	SHour,SMin,SSec int
	IntMicroSec,IntMicroSecBefore int
	Max int
	rf,bf RunnerFunc
	beg,end AOPFunc
	isStop,isFin bool
}

func NewTimeInterval(r RunnerFunc){
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
}
func realRun()
func (self *TimeInterval) runWith(data *interface{}) {
	time.Sleep()
}