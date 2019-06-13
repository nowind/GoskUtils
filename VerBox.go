package skUtils

import (
	"github.com/andlabs/ui"
	"image"
)

type VerBox struct {
	*ui.Area
	img *image.Image
}
type areaHandler struct {
}
func (areaHandler) Draw(a *ui.Area, dp *ui.AreaDrawParams){
		ui.DrawNewPath()
}
func (areaHandler) MouseEvent(a *ui.Area, me *ui.AreaMouseEvent){

}
func (areaHandler) MouseCrossed(a *ui.Area, left bool) {
	// do nothing
}

func (areaHandler) DragBroken(a *ui.Area) {
	// do nothing
}

func (areaHandler) KeyEvent(a *ui.Area, ke *ui.AreaKeyEvent) (handled bool) {
	// reject all keys
	return false
}
func NewVerBox() *VerBox{
	imgbox:=ui.NewArea()
	ret:=new(VerBox)
	ret.Area=imgbox

}