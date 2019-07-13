package skUtils

import (
	"fmt"
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
	size int
}

func (self *tablehandler)  ColumnTypes(m *ui.TableModel) []ui.TableValue{
		return []ui.TableValue{
			ui.TableFalse,
			ui.TableString(""),
			ui.TableString(" "),
		}
}


func (self *tablehandler)  NumRows(m *ui.TableModel) int{
	return self.size
}

func (self *tablehandler) CellValue(m *ui.TableModel, row, column int) ui.TableValue {
	if len(self.items) <= row{
		panic(fmt.Sprintf("unreach row:%d ,but len:%d\n",row,len(self.items)))
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
		panic(fmt.Sprintf("unreach row:%d ,but len:%d\n",row,len(self.items)))
	}
	if self.items[row]._checked{
		self.items[row]._checked= false
	}else{
		self.items[row]._checked= true
	}
	m.RowChanged(row)
}

func (self *tablehandler) changeDiff(d map[string]string,sorted []string){
	for _,i:=range sorted {
		if _,ok:=self.datas[i];ok{
			continue
		}
		if _,ok:=d[i];ok{
			//self.datas[i]=&tableitem{TablePair{}}
		}
	}
}
type MulCheckedList  struct{
	*ui.Table
	mh *tablehandler
	m *ui.TableModel
}
func _loadData(d map[string]string,sorted []string) ([]*tableitem,map[string]*tableitem){
	var retd []*tableitem
	retm:=map[string]*tableitem{}
	if sorted==nil{
		retd=make([]*tableitem,len(d))
		i:=0
		for k,v:=range d{
			it:=&tableitem{_checked:false,}
			it.Name=k
			it.Value=v
			retd[i]=it
			i++
			retm[k]=it
		}
	} else {
		retd=make([]*tableitem,len(sorted))
		lend:=0
		for i,k:=range sorted{
			if v,ok:=d[k];ok{
				it:=&tableitem{_checked:false,}
				it.Name=k
				it.Value=v
				retd[i]=it
				retm[k]=it
				lend++
			}

		}
		retd=retd[:lend]
	}
	return retd,retm
}

func NewMulCheckedListSort(d map[string]string,sorted []string) *MulCheckedList{
	mh := new(tablehandler)
	mh.items,mh.datas= _loadData(d,sorted)
	mh.size=len(mh.items)
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

func (self *MulCheckedList) Change(d map[string]string,sorted []string){
	oldLen:=len(self.mh.items)
	self.mh.items,self.mh.datas=_loadData(d,sorted)
	newLen:=len(self.mh.items)
	self.mh.size=1
	det:=newLen-oldLen
	f:=self.m.RowInserted
	if det<0{
		f=self.m.RowDeleted
		det=-det
	}
	//effect ui change
	for i:=0;i<det;i++{
		f(0)
	}
	self.mh.size=(len(self.mh.items))
	for i:=0;i<len(self.mh.items);i++{
		self.m.RowChanged(i)
	}
	//data refill
}