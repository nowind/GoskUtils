package arrays

import (
	"github.com/modern-go/reflect2"
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
	"reflect"
)

func IsArray(a interface{})bool{
	return reflect2.TypeOf(a).Kind()==reflect.Slice
}
func In(m interface{},a interface{}) bool{
	if IsArray(m){
	//	slType:=reflect2.TypeOf(m).(reflect2.SliceType)
		fmt.Printf("%v %v\n",m,m.([]int))
//		pm:=unsafe.Pointer(&m)
//		x:=(*reflect.SliceHeader)(pm)
//		b:=x.Len
		/*for i:=0;i<b;i++{
			if slType.UnsafeGetIndex(pm,i)==a{
				return true
			}
		}*/
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