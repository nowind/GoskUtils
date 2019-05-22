package skUtils

import "github.com/andlabs/ui"

type tableitem struct {
	name string
	value interface{}
}
type tablehandler struct {
	Items []tableitem
}

func (self *tablehandler)  ColumnTypes(m *ui.TableModel) []ui.TableValue{
		return []ui.TableValue{

		}
}


func (self *tablehandler)  NumRows(m *ui.TableModel) int{

}
