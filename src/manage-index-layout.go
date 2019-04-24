package src

import (
	"github.com/naponmeka/robone/tree"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

func NewManageIndexLayout(subwin *widgets.QDialog, globalState *GlobalState) *widgets.QWidget {
	var widget = widgets.NewQWidget(nil, 0)
	var loader = uitools.NewQUiLoader(nil)
	var file = core.NewQFile2(":/qml/manage-index.ui")
	file.Open(core.QIODevice__ReadOnly)
	layoutWidget := loader.Load(file, widget)
	file.Close()
	indexTreeview := widgets.NewQTreeViewFromPointer(widget.FindChild("indexTreeView", core.Qt__FindChildrenRecursively).Pointer())
	model := tree.NewCustomTreeModel(nil)
	indexTreeview.SetModel(model)

	items, err := tree.CreateIndexItems(*globalState.mongoURI, *globalState.currentDB, *globalState.currentCollection, *globalState.timeout)
	if err != nil {
		widgets.QMessageBox_Critical(nil, "Error", "Error:\n"+err.Error(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	}

	for _, item := range items {
		model.Add(item)
	}
	indexTreeview.ExpandAll()

	addBtn := widgets.NewQPushButtonFromPointer(layoutWidget.FindChild("addBtn", core.Qt__FindChildrenRecursively).Pointer())
	addBtn.ConnectClicked(func(bool) {
		subwin := widgets.NewQDialog(nil, 0)
		subwin.SetWindowTitle("Create index")
		subwin.SetLayout(widgets.NewQHBoxLayout())
		createIndexLayout := NewCreateIndexLayout(subwin)
		subwin.Layout().AddWidget(createIndexLayout)
		subwin.SetModal(true)
		subwin.Exec()
	})

	closeBtn := widgets.NewQPushButtonFromPointer(layoutWidget.FindChild("closeBtn", core.Qt__FindChildrenRecursively).Pointer())
	closeBtn.ConnectClicked(func(bool) {
		subwin.Close()
	})
	return layoutWidget
}
