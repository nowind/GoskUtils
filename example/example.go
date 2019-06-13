package main

import (
	"github.com/andlabs/ui"
	 "github.com/nowind/skUtils-go"
	_ "github.com/andlabs/ui/winmanifest"
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
		ti:=skUtils.NewTimeInterval(func(d interface{}) (i int, e error) {
			asy.Println("end")
			return -1,nil
		})

		ti.SHour=0
		ti.SMin=42
		ti.SetBeg(func() {
			asy.Println("aaa")
		})
		ti.Run()
	})
	v.Append(b,false)
	w.SetChild(v)
	w.Show()

}
func main() {
	ui.Main(mainUi)
}