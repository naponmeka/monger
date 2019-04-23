package table

import "github.com/therecipe/qt/core"

type CustomTableModel struct {
	core.QAbstractTableModel

	_ func() `constructor:"init"`

	_ func()                                         `signal:"remove,auto"`
	_ func(item TableItem)                           `signal:"add,auto"`
	_ func(firstName string, lastName string)        `signal:"edit,auto"`
	_ func(i int)                                    `signal:"removeAt,auto"`
	_ func(i int, firstName string, lastName string) `signal:"editAt,auto"`

	modelData []TableItem
}

func (m *CustomTableModel) init() {
	m.modelData = []TableItem{}
	m.ConnectHeaderData(m.headerData)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectData(m.data)
}

func (m *CustomTableModel) headerData(section int, orientation core.Qt__Orientation, role int) *core.QVariant {
	if role != int(core.Qt__DisplayRole) || orientation == core.Qt__Vertical {
		return m.HeaderDataDefault(section, orientation, role)
	}

	switch section {
	case 0:
		return core.NewQVariant14("CSV Field")
	case 1:
		return core.NewQVariant14("From Field")
	}
	return core.NewQVariant()
}

func (m *CustomTableModel) rowCount(*core.QModelIndex) int {
	return len(m.modelData)
}

func (m *CustomTableModel) columnCount(*core.QModelIndex) int {
	return 2
}

func (m *CustomTableModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if role != int(core.Qt__DisplayRole) {
		return core.NewQVariant()
	}

	item := m.modelData[index.Row()]
	switch m.HeaderData(index.Column(), core.Qt__Horizontal, role).ToString() {
	case "CSV Field":
		return core.NewQVariant14(item.CSVField)
	case "From Field":
		return core.NewQVariant14(item.FromField)
	}
	return core.NewQVariant()
}

func (m *CustomTableModel) remove() {
	if len(m.modelData) == 0 {
		return
	}
	m.BeginRemoveRows(core.NewQModelIndex(), len(m.modelData)-1, len(m.modelData)-1)
	m.modelData = m.modelData[:len(m.modelData)-1]
	m.EndRemoveRows()
}

func (m *CustomTableModel) removeAt(i int) {
	if len(m.modelData) == 0 || i > len(m.modelData) {
		return
	}
	m.BeginRemoveRows(core.NewQModelIndex(), i, i)
	m.modelData = append(m.modelData[:i], m.modelData[i+1:]...)
	m.EndRemoveRows()
}

func (m *CustomTableModel) add(item TableItem) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.modelData), len(m.modelData))
	m.modelData = append(m.modelData, item)
	m.EndInsertRows()
}

func (m *CustomTableModel) edit(firstName string, lastName string) {
	if len(m.modelData) == 0 {
		return
	}
	m.modelData[len(m.modelData)-1] = TableItem{firstName, lastName}
	m.DataChanged(m.Index(len(m.modelData)-1, 0, core.NewQModelIndex()), m.Index(len(m.modelData)-1, 1, core.NewQModelIndex()), []int{int(core.Qt__DisplayRole)})
}

func (m *CustomTableModel) editAt(idx int, firstName string, lastName string) {
	if len(m.modelData) == 0 {
		return
	}
	m.modelData[idx] = TableItem{firstName, lastName}
	m.DataChanged(m.Index(idx, 0, core.NewQModelIndex()), m.Index(idx, 1, core.NewQModelIndex()), []int{int(core.Qt__DisplayRole)})
}
