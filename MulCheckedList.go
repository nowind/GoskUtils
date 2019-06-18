package skUtils

import (
	"github.com/andlabs/ui"
)


type TablePair struct {
	Name string
	Value interface{}
}
type tableitem struct {
	TablePair
	_checked bool
}
type tablehandler struct {
	items []*tableitem
	datas map[string]*tableitem
}

func (self *tablehandler)  ColumnTypes(m *ui.TableModel) []ui.TableValue{
		return []ui.TableValue{
			ui.TableFalse,
			ui.TableString(""),
			ui.TableString(" "),
		}
}


func (self *tablehandler)  NumRows(m *ui.TableModel) int{
	return len(self.items)
}

func (self *tablehandler) CellValue(m *ui.TableModel, row, column int) ui.TableValue {
	if len(self.items) <= row{
		panic("unreach")
	}
	if column == 2 {
		if self.items[row]._checked{
			return ui.TableString("X")
		}
		return ui.TableString(" ")
	}
	return ui.TableString(self.items[row].Name)
}

func (self *tablehandler) SetCellValue(m *ui.TableModel, row, column int, value ui.TableValue) {
	if len(self.items) <= row{
		panic("unreach")
	}
	if self.items[row]._checked{
		self.items[row]._checked= false
	}else{
		self.items[row]._checked= true
	}
	m.RowChanged(row)
}

type MulCheckedList  struct{
	*ui.Table
	mh *tablehandler
	m *ui.TableModel
}
func NewMulCheckedListSort(d map[string]string,sorted []string) *MulCheckedList{
	mh := new(tablehandler)
	mh.datas= map[string]*tableitem{}
	if sorted==nil{
		for k,v:=range d{
			it:=&tableitem{_checked:false,}
			it.Name=k
			it.Value=v
			mh.items = append(mh.items, it)
			mh.datas[k]=it
		}
	} else {
		for _,k:=range sorted{
			if v,ok:=d[k];ok{
				it:=&tableitem{_checked:false,}
				it.Name=k
				it.Value=v
				mh.items = append(mh.items, it)
				mh.datas[k]=it
			}

		}
	}

	model := ui.NewTableModel(mh)
	m:=new(MulCheckedList)
	table := ui.NewTable(&ui.TableParams{
		Model:                         model,
	})
	table.AppendButtonColumn("",
		2,ui.TableModelColumnAlwaysEditable)
	table.AppendTextColumn("",1,ui.TableModelColumnNeverEditable,nil)
	m.Table=table
	m.mh=mh
	m.m=model
	return m
}
func NewMulCheckedList(d map[string]string) *MulCheckedList{
	return NewMulCheckedListSort(d,nil)
}
func (self *MulCheckedList)SelPairList()  []interface{}{
	ret:=make([]interface{},0,len(self.mh.items))
	for _,v := range self.mh.items{
		if v._checked {
			ret=append(ret,v.TablePair)
		}
	}
	return ret
}
func (self *MulCheckedList)SelList() []interface{}{
	ret:=make([]interface{},0,len(self.mh.items))
	for _,v := range self.mh.items{
		if v._checked {
			ret=append(ret,v.Name)
		}
	}
	return ret
}
func (self *MulCheckedList)SelValList() []interface{}{
	ret:=make([]interface{},0,len(self.mh.items))
	for _,v := range self.mh.items{
		if v._checked {
			ret=append(ret,v.Value)
		}
	}
	return ret
}
func (self *MulCheckedList) sel(b bool){
	for i,d:=range self.mh.items{
		d._checked=b
		self.m.RowChanged(i)
	}
}
func (self *MulCheckedList) SelAll(){
	self.sel(true)
}
func (self *MulCheckedList) UnSelAll(){
	self.sel(false)

}