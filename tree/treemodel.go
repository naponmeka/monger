package tree

import "github.com/therecipe/qt/core"

type CustomTreeModel struct {
	core.QAbstractItemModel

	_ func() `constructor:"init"`

	_ func()                                  `signal:"remove,auto"`
	_ func(item *TreeItem)                    `signal:"add,auto"`
	_ func(firstName string, lastName string) `signal:"edit,auto"`

	rootItem *TreeItem
}

func (m *CustomTreeModel) init() {
	m.rootItem = NewTreeItem(nil).initWith([]string{"Key", "Value", "Type"})

	m.ConnectIndex(m.index)
	m.ConnectParent(m.parent)
	m.ConnectHeaderData(m.headerData)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectData(m.data)
}

func (m *CustomTreeModel) index(row int, column int, parent *core.QModelIndex) *core.QModelIndex {
	if !m.HasIndex(row, column, parent) {
		return core.NewQModelIndex()
	}

	var parentItem *TreeItem
	if !parent.IsValid() {
		parentItem = m.rootItem
	} else {
		parentItem = NewTreeItemFromPointer(parent.InternalPointer())
	}

	childItem := parentItem.child(row).Pointer()
	if childItem != nil {
		return m.CreateIndex(row, column, childItem)
	}
	return core.NewQModelIndex()
}

func (m *CustomTreeModel) parent(index *core.QModelIndex) *core.QModelIndex {
	if !index.IsValid() {
		return core.NewQModelIndex()
	}

	item := NewTreeItemFromPointer(index.InternalPointer())
	parentItem := item.parentItem()

	if parentItem.Pointer() == m.rootItem.Pointer() {
		return core.NewQModelIndex()
	}

	return m.CreateIndex(parentItem.row(), 0, parentItem.Pointer())
}

func (m *CustomTreeModel) headerData(section int, orientation core.Qt__Orientation, role int) *core.QVariant {
	if role == int(core.Qt__DisplayRole) && orientation == core.Qt__Horizontal {
		return core.NewQVariant14(m.rootItem.data(section))
	}
	return m.HeaderDataDefault(section, orientation, role)
}

func (m *CustomTreeModel) rowCount(parent *core.QModelIndex) int {
	if !parent.IsValid() {
		return m.rootItem.childCount()
	}
	return NewTreeItemFromPointer(parent.InternalPointer()).childCount()
}

func (m *CustomTreeModel) columnCount(parent *core.QModelIndex) int {
	if !parent.IsValid() {
		return m.rootItem.columnCount()
	}
	return NewTreeItemFromPointer(parent.InternalPointer()).columnCount()
}

func (m *CustomTreeModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() || role != int(core.Qt__DisplayRole) {
		return core.NewQVariant()
	}

	item := NewTreeItemFromPointer(index.InternalPointer())
	return core.NewQVariant14(item.data(index.Column()))
}

func (m *CustomTreeModel) remove() {
	if m.rootItem.childCount() == 0 {
		return
	}
	m.BeginRemoveRows(core.NewQModelIndex(), len(m.rootItem._childItems)-1, len(m.rootItem._childItems)-1)
	item := m.rootItem._childItems[len(m.rootItem._childItems)-1]
	m.rootItem._childItems = m.rootItem._childItems[:len(m.rootItem._childItems)-1]
	m.EndRemoveRows()
	item.DestroyTreeItem()
}

func (m *CustomTreeModel) add(item *TreeItem) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.rootItem._childItems), len(m.rootItem._childItems))
	m.rootItem.appendChild(item)
	m.EndInsertRows()
}

func (m *CustomTreeModel) edit(firstName string, lastName string) {
	if m.rootItem.childCount() == 0 {
		return
	}
	item := m.rootItem._childItems[len(m.rootItem._childItems)-1]
	item._itemData = []string{firstName, lastName}
	m.DataChanged(m.Index(item.row(), 0, core.NewQModelIndex()), m.Index(item.row(), 1, core.NewQModelIndex()), []int{int(core.Qt__DisplayRole)})
}
