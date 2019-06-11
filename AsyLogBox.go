package skUtils

import (
	"github.com/andlabs/ui"
	"log"
)

type AsyLogBox struct{
	*ui.MultilineEntry
	*log.Logger
}
func (self *AsyLogBox)Write(p []byte) (n int, err error){
	str:=string(p)
	ui.QueueMain(func(){
		self.Append(str)
	})
	return len(p),nil
}
func NewAsyLogBox() *AsyLogBox {
	ret:=new(AsyLogBox)
	ret.MultilineEntry=ui.NewMultilineEntry()
	ret.Logger=log.New(ret,"",log.Ltime)
	return ret
}