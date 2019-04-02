package list

import (
	"fmt"
	"strconv"

	"github.com/therecipe/qt/core"
)

type CustomListModel struct {
	core.QAbstractListModel

	_ func() `constructor:"init"`

	_ func()                        `signal:"remove,auto"`
	_ func(item ListItem)           `signal:"add,auto"`
	_ func(Name string, URI string) `signal:"edit,auto"`
	_ func(iStr string)             `signal:"removeAt,auto"`

	modelData []ListItem
}

func (m *CustomListModel) init() {
	m.modelData = []ListItem{}

	m.ConnectRowCount(m.rowCount)
	m.ConnectData(m.data)
}

func (m *CustomListModel) rowCount(*core.QModelIndex) int {
	return len(m.modelData)
}

func (m *CustomListModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if role != int(core.Qt__DisplayRole) {
		return core.NewQVariant()
	}

	item := m.modelData[index.Row()]
	return core.NewQVariant14(fmt.Sprintf("%v - %v", item.Name, item.URI))
}

func (m *CustomListModel) remove() {
	if len(m.modelData) == 0 {
		return
	}
	m.BeginRemoveRows(core.NewQModelIndex(), len(m.modelData)-1, len(m.modelData)-1)
	m.modelData = m.modelData[:len(m.modelData)-1]
	m.EndRemoveRows()
}

func (m *CustomListModel) removeAt(iStr string) {
	i, _ := strconv.Atoi(iStr)
	if len(m.modelData) == 0 || i > len(m.modelData) {
		return
	}
	m.BeginRemoveRows(core.NewQModelIndex(), i, i)
	m.modelData = append(m.modelData[:i], m.modelData[i+1:]...)
	m.EndRemoveRows()
}

func (m *CustomListModel) add(item ListItem) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.modelData), len(m.modelData))
	m.modelData = append(m.modelData, item)
	m.EndInsertRows()
}

func (m *CustomListModel) edit(firstName string, lastName string) {
	if len(m.modelData) == 0 {
		return
	}
	m.modelData[len(m.modelData)-1] = ListItem{firstName, lastName}
	m.DataChanged(m.Index(len(m.modelData)-1, 0, core.NewQModelIndex()), m.Index(len(m.modelData)-1, 0, core.NewQModelIndex()), []int{int(core.Qt__DisplayRole)})
}
