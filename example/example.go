package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"github.com/nowind/skUtils-go"
)
func mainUi(){
	w:=ui.NewWindow("",600,600,false)
	w.SetMargined(true)
	w.OnClosing(func(window *ui.Window) bool {
		ui.Quit()
		return false
	})
	ui.OnShouldQuit(func() bool {
		w.Destroy()
		return true
	})
	v:=ui.NewVerticalBox()
	v.SetPadded(true)

	asy:=skUtils.NewAsyLogBox()
	v.Append(asy,true)
	b:=ui.NewButton("OK")
	b.OnClicked(func(button *ui.Button) {
		asy.Println("hi,world")
		skUtils.NewTimeInterval(func(d interface{}) (i int, e error) {
			asy.Println("doing")
			return -1,nil
		}).SetTime(9,11,0).SetBeg(func() {
			asy.Println("begin")
		}).SetEnd(func() {
			asy.Println("end")
		}).Run()
	})
	v.Append(b,false)
	w.SetChild(v)
	w.Show()

}
func main() {
	ui.Main(mainUi)
}