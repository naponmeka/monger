package src

import (
	"fmt"
	"strings"

	"github.com/naponmeka/robone/connectdb"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

func registerTemplateBtn(mainWidget *widgets.QWidget, queryPlainTextEdit *widgets.QPlainTextEdit) {
	findBtn := widgets.NewQPushButtonFromPointer(mainWidget.FindChild("findBtn", core.Qt__FindChildrenRecursively).Pointer())
	insertBtn := widgets.NewQPushButtonFromPointer(mainWidget.FindChild("insertBtn", core.Qt__FindChildrenRecursively).Pointer())
	updateBtn := widgets.NewQPushButtonFromPointer(mainWidget.FindChild("updateBtn", core.Qt__FindChildrenRecursively).Pointer())
	deleteBtn := widgets.NewQPushButtonFromPointer(mainWidget.FindChild("deleteBtn", core.Qt__FindChildrenRecursively).Pointer())
	aggBtn := widgets.NewQPushButtonFromPointer(mainWidget.FindChild("aggBtn", core.Qt__FindChildrenRecursively).Pointer())
	findBtn.ConnectClicked(func(bool) {
		queryPlainTextEdit.SetPlainText(".find({})")
	})
	insertBtn.ConnectClicked(func(bool) {
		queryPlainTextEdit.SetPlainText(`.insert(
  [
    {}
  ],
  {
    "ordered": true
  }
])`)
	})
	updateBtn.ConnectClicked(func(bool) {
		txt := `.update(
  {query},
  {update},
  {
    "upsert": false,
    "multi": false
  }
)`
		queryPlainTextEdit.SetPlainText(txt)
	})
	deleteBtn.ConnectClicked(func(bool) {
		txt := `.remove(
  {query},
  {
    justOne: true
  }
)`
		queryPlainTextEdit.SetPlainText(txt)
	})
	aggBtn.ConnectClicked(func(bool) {
		txt := `.aggregate([
  {},
])`
		queryPlainTextEdit.SetPlainText(txt)
	})
}

func registerActionBtn(mainWidget *widgets.QWidget, queryPlainTextEdit *widgets.QPlainTextEdit) {
	countBtn := widgets.NewQPushButtonFromPointer(mainWidget.FindChild("countBtn", core.Qt__FindChildrenRecursively).Pointer())
	explainBtn := widgets.NewQPushButtonFromPointer(mainWidget.FindChild("explainBtn", core.Qt__FindChildrenRecursively).Pointer())

	countBtn.ConnectClicked(func(bool) {
		currentText := queryPlainTextEdit.ToPlainText()
		currentText = strings.TrimSuffix(currentText, ".count()")
		currentText = strings.TrimSuffix(currentText, ".explain()")
		queryPlainTextEdit.SetPlainText(fmt.Sprintf("%s.count()", currentText))
	})
	explainBtn.ConnectClicked(func(bool) {
		currentText := queryPlainTextEdit.ToPlainText()
		currentText = strings.TrimSuffix(currentText, ".count()")
		currentText = strings.TrimSuffix(currentText, ".explain()")
		queryPlainTextEdit.SetPlainText(fmt.Sprintf("%s.explain()", currentText))
	})
}

func NewMainLayout(mongoURI string) *widgets.QWidget {
	var widget = widgets.NewQWidget(nil, 0)
	var loader = uitools.NewQUiLoader(nil)
	var file = core.NewQFile2(":/qml/main.ui")

	file.Open(core.QIODevice__ReadOnly)
	mainWidget := loader.Load(file, widget)
	file.Close()
	databaseComboBox := widgets.NewQFontComboBoxFromPointer(widget.FindChild("databaseComboBox", core.Qt__FindChildrenRecursively).Pointer())
	collectionComboBox := widgets.NewQFontComboBoxFromPointer(widget.FindChild("collectionComboBox", core.Qt__FindChildrenRecursively).Pointer())
	queryPlainTextEdit := widgets.NewQPlainTextEditFromPointer(widget.FindChild("queryPlainTextEdit", core.Qt__FindChildrenRecursively).Pointer())
	beginSpinBox := widgets.NewQSpinBoxFromPointer(widget.FindChild("beginSpinBox", core.Qt__FindChildrenRecursively).Pointer())
	endSpinBox := widgets.NewQSpinBoxFromPointer(widget.FindChild("endSpinBox", core.Qt__FindChildrenRecursively).Pointer())

	currentDB := ""
	currentCollection := ""
	registerTemplateBtn(mainWidget, queryPlainTextEdit)
	registerActionBtn(mainWidget, queryPlainTextEdit)
	dbs := connectdb.ListDB(mongoURI)
	collections := []string{}
	databaseComboBox.AddItems(dbs)
	databaseComboBox.ConnectCurrentIndexChanged(func(idx int) {
		currentDB = dbs[idx]
		collections = connectdb.ListCollection(mongoURI, currentDB)
		collectionComboBox.Clear()
		collectionComboBox.AddItems(collections)
	})
	if len(dbs) > 0 {
		currentDB = dbs[0]
		collections = connectdb.ListCollection(mongoURI, currentDB)
		currentCollection = collections[0]
		collectionComboBox.Clear()
		collectionComboBox.AddItems(collections)
		queryPlainTextEdit.SetPlainText(".find({})")
	}

	collectionComboBox.ConnectCurrentIndexChanged(func(idx int) {
		if idx >= 0 && idx < len(collections) {
			currentCollection = collections[idx]
			// queryPlainTextEdit.SetPlainText(fmt.Sprintf("db.getCollection('%s').find({})", currentCollection))
			queryPlainTextEdit.SetPlainText(".find({})")
		}
	})

	resultTreeview := widgets.NewQTreeViewFromPointer(widget.FindChild("resultTreeView", core.Qt__FindChildrenRecursively).Pointer())
	model := NewCustomTreeModel(nil)
	resultTreeview.SetModel(model)
	executeBtn := widgets.NewQPushButtonFromPointer(widget.FindChild("executeBtn", core.Qt__FindChildrenRecursively).Pointer())
	executeBtn.ConnectClicked(func(bool) {
		maxPossibleDocCount := endSpinBox.Value() - beginSpinBox.Value()
		for i := 0; i < maxPossibleDocCount; i++ {
			model.Remove()
		}
		currentQuery := queryPlainTextEdit.ToPlainText()
		items, err := createItems(mongoURI, currentDB, currentCollection, currentQuery)
		if err != nil {
			widgets.QMessageBox_Information(nil, "Error", "wrong query", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			return
		}
		for _, item := range items {
			model.Add(item)
		}
	})
	editDocBtn := widgets.NewQPushButtonFromPointer(widget.FindChild("editDocBtn", core.Qt__FindChildrenRecursively).Pointer())
	editDocBtn.ConnectClicked(func(bool) {
		selected := findRow(resultTreeview, resultTreeview.CurrentIndex())
		subwin := widgets.NewQDialog(nil, 0)
		subwin.SetWindowTitle(fmt.Sprintf("Edit: %d", selected))
		subwin.SetLayout(widgets.NewQHBoxLayout())
		editLayout := NewEditLayout()
		subwin.Layout().AddWidget(editLayout)
		RegisterEditLayoutBtn(editLayout, subwin)
		subwin.SetModal(true)
		subwin.SetMinimumSize2(640, 480)
		subwin.Exec()
	})
	return mainWidget
}

func findRow(treeView *widgets.QTreeView, current *core.QModelIndex) int {
	// parent := current.Parent()
	// if parent != current {
	// 	return findRow(treeView, parent)
	// }
	return current.Row()

}
