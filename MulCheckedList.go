package skUtils

import (
	"github.com/andlabs/ui"
)


type TablePair struct {
	Name,Value string
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
	return ui.TableString(self.items[row].Name)
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
	table.AppendCheckboxColumn("",
		0, ui.TableModelColumnAlwaysEditable)
	table.AppendTextColumn("",1,ui.TableModelColumnNeverEditable, nil)
	m.Table=table
	m.mh=mh
	m.m=model
	return m
}
func NewMulCheckedList(d map[string]string) *MulCheckedList{
	return NewMulCheckedListSort(d,nil)
}
func (self *MulCheckedList)SelPairList()  []TablePair{
	ret:=make([]TablePair,0,len(self.mh.items))
	for _,v := range self.mh.items{
		if v._checked {
			ret=append(ret,v.TablePair)
		}
	}
	return ret
}
func (self *MulCheckedList)SelList() []string{
	ret:=make([]string,0,len(self.mh.items))
	for _,v := range self.mh.items{
		if v._checked {
			ret=append(ret,v.Name)
		}
	}
	return ret
}
func (self *MulCheckedList)SelValList() []string{
	ret:=make([]string,0,len(self.mh.items))
	for _,v := range self.mh.items{
		if v._checked {
			ret=append(ret,v.Value)
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