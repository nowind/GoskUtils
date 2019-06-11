package skUtils

import (
	"github.com/andlabs/ui"
	"testing"
)
func mainUi(){
	w:=ui.NewWindow("",600,600,false)
	w.SetMargined(true)
	v:=ui.NewVerticalBox()
	v.SetPadded(true)

	asy:=NewAsyLogBox()
	v.Append(asy,true)
	b:=ui.NewButton("OK")
	b.OnClicked(func(button *ui.Button) {
		asy.Println("hi,world")
	})
	v.Append(b,false)
	w.SetChild(v)
	w.Show()

}
func TestNewAsyLogBox(t *testing.T) {
	ui.Main(mainUi)
}