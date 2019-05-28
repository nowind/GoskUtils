package skUtils

import "github.com/andlabs/ui"

//go:generate echo 123


type tableitem struct {
	name string
	value interface{}
}
type tablehandler struct {
	items []tableitem
	datas map[string]*tableitem
}

func (self *tablehandler)  ColumnTypes(m *ui.TableModel) []ui.TableValue{
		return []ui.TableValue{
			ui.TableInt(ui.TableTrue),
			ui.TableString(""),
		}
}


func (self *tablehandler)  NumRows(m *ui.TableModel) int{
	return len(self.items)
}

func (self *tablehandler) CellValue(m *ui.TableModel, row, column int) ui.TableValue {
	if len(self.items) >= row{
		panic("unreach")
	}
	return ui.TableString(self.items[row].name)
}

func (self *tablehandler) SetCellValue(m *ui.TableModel, row, column int, value ui.TableValue) {
	if len(self.items) >= row{
		panic("unreach")
	}
	self.items[row].value= bool(value.(ui.TableInt)!=0)
}

type MulCheckedList  struct{
	m *ui.Table
}
func NewMulCheckedList() *MulCheckedList{
	m:=new(MulCheckedList)
	mh := new(tablehandler)

	model := ui.NewTableModel(mh)

	table := ui.NewTable(&ui.TableParams{
		Model:                         model,
		RowBackgroundColorModelColumn: 3,
	})
	m.m=table
	return m
}
func (self *MulCheckedList) getControl() *ui.Table{
	return self.m
}
func (self *MulCheckedList)