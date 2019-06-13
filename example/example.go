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
	c:=0
	l:=skUtils.NewIni(`e:\tmp\conf.ini`)
	p:=l.Get("act")
	j:=skUtils.NewMulCheckedList(p)
	b.OnClicked(func(button *ui.Button) {
		asy.Println("hi,world")
		skUtils.NewTimeInterval(func(d interface{}) (i int, e error) {
			asy.Println("dong..")
			return -1,nil
		}).SetTime(0,48,0).SetBeg(func() {
			asy.Println("begin")
		}).SetEnd(func() {
			asy.Println("end")
		}).SetBefRun(func(d interface{}) (i int, e error) {
			if c<10{
				asy.Printf("%d \n",c)
				c++
			}
			return 0,nil
		})
		asy.Println(j.SelList())
		har:=skUtils.NewHARParser(`d:\tmp\1.har`)
		k:=har.Repeat(nil,nil)
		d:=make([]byte,5000,5000)
		i,_:=k.Body.Read(d)
		asy.Println(string(d[:i]))
		asy.Println(har.)
	})
	cc:=[]interface{}{"1","8","4"}
	skUtils.NewMultiTask(func(d interface{}) (i int, e error) {
		asy.Println(d)
		return -1,nil
	}).SetTime(0,50,0).SetTimeInt(1000,1000).RunWith(cc)
	v.Append(b,false)
	h.Append(v,false)


	h.Append(j,true)
	w.SetChild(h)
	w.Show()

}
func main() {
	ui.Main(mainUi)
}