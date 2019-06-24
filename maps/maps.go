package maps

import (
	"github.com/modern-go/reflect2"
	"reflect"
)

func IsMap(m interface{}) bool{
	return reflect2.TypeOf(m).Kind()==reflect.Map
}
func IsEmptyMap(m interface{}) bool{
	if IsMap(m){
		return !reflect2.TypeOf(m).(reflect2.MapType).Iterate(m).HasNext()
	}else{
		return false
	}
}
