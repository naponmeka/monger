package src

import (
	"fmt"
	"strings"

	"github.com/naponmeka/robone/connectdb"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

func registerTemplateBtn(widget *widgets.QWidget, queryPlainTextEdit *widgets.QPlainTextEdit, currentCollection *string) {
	findBtn := widgets.NewQPushButtonFromPointer(widget.FindChild("findBtn", core.Qt__FindChildrenRecursively).Pointer())
	insertBtn := widgets.NewQPushButtonFromPointer(widget.FindChild("insertBtn", core.Qt__FindChildrenRecursively).Pointer())
	updateBtn := widgets.NewQPushButtonFromPointer(widget.FindChild("updateBtn", core.Qt__FindChildrenRecursively).Pointer())
	deleteBtn := widgets.NewQPushButtonFromPointer(widget.FindChild("deleteBtn", core.Qt__FindChildrenRecursively).Pointer())
	aggBtn := widgets.NewQPushButtonFromPointer(widget.FindChild("aggBtn", core.Qt__FindChildrenRecursively).Pointer())
	findBtn.ConnectClicked(func(bool) {
		text := ""
		if currentCollection != nil {
			text = *currentCollection
		}
		queryPlainTextEdit.SetPlainText(fmt.Sprintf("db.getCollection('%s').find({})", text))
	})
	insertBtn.ConnectClicked(func(bool) {
		text := ""
		if currentCollection != nil {
			text = *currentCollection
		}
		queryPlainTextEdit.SetPlainText(fmt.Sprintf("db.getCollection('%s').insert([\n{}\n])", text))
	})
	updateBtn.ConnectClicked(func(bool) {
		text := ""
		if currentCollection != nil {
			text = *currentCollection
		}
		template := `db.getCollection('%s').update(
  {query},
  {update}
)`
		queryPlainTextEdit.SetPlainText(fmt.Sprintf(template, text))
	})
	deleteBtn.ConnectClicked(func(bool) {
		text := ""
		if currentCollection != nil {
			text = *currentCollection
		}
		template := `db.getCollection('%s').remove(
  {query},
  {
    justOne: true
   }
)`
		queryPlainTextEdit.SetPlainText(fmt.Sprintf(template, text))
	})
	aggBtn.ConnectClicked(func(bool) {
		text := ""
		if currentCollection != nil {
			text = *currentCollection
		}
		template := `db.getCollection('%s').aggregate([
  {},
])`
		queryPlainTextEdit.SetPlainText(fmt.Sprintf(template, text))
	})
}

func registerActionBtn(widget *widgets.QWidget, queryPlainTextEdit *widgets.QPlainTextEdit, currentCollection *string) {
	countBtn := widgets.NewQPushButtonFromPointer(widget.FindChild("countBtn", core.Qt__FindChildrenRecursively).Pointer())
	explainBtn := widgets.NewQPushButtonFromPointer(widget.FindChild("explainBtn", core.Qt__FindChildrenRecursively).Pointer())

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
	registerTemplateBtn(widget, queryPlainTextEdit, &currentCollection)
	registerActionBtn(widget, queryPlainTextEdit, &currentCollection)
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
	}

	collectionComboBox.ConnectCurrentIndexChanged(func(idx int) {
		if idx >= 0 && idx < len(collections) {
			currentCollection = collections[idx]
			queryPlainTextEdit.SetPlainText(fmt.Sprintf("db.getCollection('%s').find({})", currentCollection))
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
		for _, item := range createItems(mongoURI, currentDB, currentCollection, currentQuery) {
			model.Add(item)
		}
	})

	return mainWidget
}
