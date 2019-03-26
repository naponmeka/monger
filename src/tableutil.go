package src

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func findRow(treeView *widgets.QTreeView, current *core.QModelIndex) int {
	parent := current.Parent()
	if parent.IsValid() {
		return findRow(treeView, parent)
	}
	return current.Row()
}
