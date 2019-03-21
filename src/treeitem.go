package src

import (
	"github.com/therecipe/qt/core"
)

type TreeItem struct {
	core.QObject

	_ func() `constructor:"init"`

	_itemData   []string
	_parentItem *TreeItem

	_childItems []*TreeItem
}

func (i *TreeItem) init() {
	i.ConnectDestroyTreeItem(i.destroyTreeItem)
}

func (i *TreeItem) initWith(data []string) *TreeItem {
	i._itemData = data
	return i
}

func (i *TreeItem) appendChild(child *TreeItem) {
	child._parentItem = i
	i._childItems = append(i._childItems, child)
}

func (i *TreeItem) child(row int) *TreeItem {
	return i._childItems[row]
}

func (i *TreeItem) childCount() int {
	return len(i._childItems)
}

func (i *TreeItem) columnCount() int {
	return len(i._itemData)
}

func (i *TreeItem) data(column int) string {
	return i._itemData[column]
}

func (i *TreeItem) row() int {
	if i._parentItem != nil {
		for index, item := range i._parentItem._childItems {
			if item.Pointer() == i.Pointer() {
				return index
			}
		}
	}
	return 0
}

func (i *TreeItem) parentItem() *TreeItem {
	return i._parentItem
}

func (i *TreeItem) destroyTreeItem() {
	for _, child := range i._childItems {
		child.DestroyTreeItem()
	}
	i.DestroyTreeItemDefault()
}
