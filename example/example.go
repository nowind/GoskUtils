package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	skUtils "github.com/nowind/GoskUtils"
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
	h:=ui.NewHorizontalBox()
	h.SetPadded(true)
	j:=skUtils.NewMulCheckedListSort(map[string]string{"a":"b"},[]string{"a"})
	b.OnClicked(func(button *ui.Button) {
		j.Change(map[string]string{"a":"b","b":"c"},[]string{"a","b"})
	})

	v.Append(b,false)
	h.Append(v,false)


	h.Append(j,true)
	w.SetChild(h)
	w.Show()

}
func main() {
	ui.Main(mainUi)
}