package skUtils

import "time"

func NowPass(sHour int,sMin int,sSec int) bool{
	if sHour<0{
		return true
	}
	now:=time.Now()
	day:=now.Day()
	if now.Hour()>12 && sHour <5{
		day=day+1
	}
	t:=time.Date(now.Year(),now.Month(),day,now.Hour(),now.Minute(),now.Second(),0,time.Local)
	return now.After(t)
}
