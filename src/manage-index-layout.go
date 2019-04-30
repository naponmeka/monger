package src

import (
	"fmt"

	"github.com/naponmeka/monger/connectdb"
	"github.com/naponmeka/monger/tree"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

func RefreshIndexTree(
	indexTreeview *widgets.QTreeView,
	model *tree.CustomTreeModel,
	globalState *GlobalState,
) {
	model.RemoveAll()
	items, documents, err := tree.CreateIndexItems(*globalState.mongoURI, *globalState.currentDB, *globalState.currentCollection, *globalState.timeout)
	if err != nil {
		widgets.QMessageBox_Critical(nil, "Error", "Error:\n"+err.Error(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	}

	for _, item := range items {
		model.Add(item)
	}
	globalState.indice = &documents
	indexTreeview.ExpandAll()
}

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
	RefreshIndexTree(indexTreeview, model, globalState)

	addBtn := widgets.NewQPushButtonFromPointer(layoutWidget.FindChild("addBtn", core.Qt__FindChildrenRecursively).Pointer())
	addBtn.ConnectClicked(func(bool) {
		subwin := widgets.NewQDialog(nil, 0)
		subwin.SetWindowTitle("Create index")
		subwin.SetLayout(widgets.NewQHBoxLayout())
		createIndexLayout := NewCreateIndexLayout(subwin, indexTreeview, model, globalState, false)
		subwin.Layout().AddWidget(createIndexLayout)
		subwin.SetModal(true)
		subwin.Exec()
	})

	removeBtn := widgets.NewQPushButtonFromPointer(widget.FindChild("removeBtn", core.Qt__FindChildrenRecursively).Pointer())
	removeBtn.ConnectClicked(func(bool) {
		selected := findRow(indexTreeview, indexTreeview.CurrentIndex())
		selectedIndex := (*globalState.indice)[selected]
		indexName := tree.CastToString(selectedIndex["name"])
		subwin := widgets.NewQDialog(nil, 0)
		subwin.SetWindowTitle(fmt.Sprintf("Delete: %s", indexName))
		subwin.SetLayout(widgets.NewQHBoxLayout())
		deleteConfirmLayout := NewConfirmLayout("Confirm delete?\n")
		subwin.Layout().AddWidget(deleteConfirmLayout)

		buttonBox := widgets.NewQDialogButtonBoxFromPointer(deleteConfirmLayout.FindChild("buttonBox", core.Qt__FindChildrenRecursively).Pointer())
		buttonBox.ConnectAccepted(func() {
			err := connectdb.DropIndex(*globalState.mongoURI,
				*globalState.currentDB,
				*globalState.currentCollection,
				indexName,
			)
			if err != nil {
				widgets.QMessageBox_Critical(nil, "Error", "Error:\n"+err.Error(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			} else {
				RefreshIndexTree(indexTreeview, model, globalState)
			}
			subwin.Close()
		})
		buttonBox.ConnectRejected(func() {
			subwin.Close()
		})
		subwin.SetModal(true)
		subwin.SetMinimumSize2(100, 100)
		subwin.Exec()
	})

	closeBtn := widgets.NewQPushButtonFromPointer(layoutWidget.FindChild("closeBtn", core.Qt__FindChildrenRecursively).Pointer())
	closeBtn.ConnectClicked(func(bool) {
		subwin.Close()
	})
	return layoutWidget
}
