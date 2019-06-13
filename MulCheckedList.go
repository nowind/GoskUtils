package skUtils

import "github.com/andlabs/ui"



type tableitem struct {
	name string
	value string
	_checked bool
}
type tablehandler struct {
	items []*tableitem
	datas map[string]*tableitem
}

func (self *tablehandler)  ColumnTypes(m *ui.TableModel) []ui.TableValue{
		return []ui.TableValue{
			ui.TableTrue,
			ui.TableString(""),
		}
}


func (self *tablehandler)  NumRows(m *ui.TableModel) int{
	return len(self.items)
}

func (self *tablehandler) CellValue(m *ui.TableModel, row, column int) ui.TableValue {
	if len(self.items) <= row{
		panic("unreach")
	}
	if column ==0{
		if self.items[row]._checked{
			return ui.TableTrue
		}
		return ui.TableFalse
	}
	return ui.TableString(self.items[row].name)
}

func (self *tablehandler) SetCellValue(m *ui.TableModel, row, column int, value ui.TableValue) {
	if len(self.items) <= row{
		panic("unreach")
	}
	self.items[row]._checked= bool(value.(ui.TableInt)==ui.TableTrue)
	m.RowChanged(row)
}

type MulCheckedList  struct{
	*ui.Table
	mh *tablehandler
	m *ui.TableModel
}
func NewMulCheckedList(d map[string]string) *MulCheckedList{
	mh := new(tablehandler)
	mh.datas= map[string]*tableitem{}
	for k,v:=range d{
		it:=&tableitem{name:k,value:v,_checked:false}
		mh.items = append(mh.items, it)
		mh.datas[k]=it
	}
	model := ui.NewTableModel(mh)
	m:=new(MulCheckedList)
	table := ui.NewTable(&ui.TableParams{
		Model:                         model,
	})
	table.AppendCheckboxColumn("",
		0, ui.TableModelColumnAlwaysEditable)
	table.AppendTextColumn("",1,ui.TableModelColumnNeverEditable, nil)
	m.Table=table
	m.mh=mh
	m.m=model
	return m
}
func (self *MulCheckedList)SelList()  map[string]string{
	ret:=make(map[string]string)
	for _,v := range self.mh.items{
		if v._checked {
			ret[v.name]=v.value
		}
	}
	return ret
}
func (self *MulCheckedList) SelAll(){
	for i,d:=range self.mh.items{
		d._checked=true
		self.m.RowChanged(i)
	}

}