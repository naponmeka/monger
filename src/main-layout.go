package src

import (
	"fmt"

	"github.com/naponmeka/robone/connectdb"
	"github.com/naponmeka/robone/tree"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
	"go.mongodb.org/mongo-driver/bson"
)

func NewMainLayout(mongoURI string, globalState *GlobalState, name string) *widgets.QWidget {
	var widget = widgets.NewQWidget(nil, 0)
	var loader = uitools.NewQUiLoader(nil)
	var file = core.NewQFile2(":/qml/main.ui")

	file.Open(core.QIODevice__ReadOnly)
	mainWidget := loader.Load(file, widget)
	file.Close()
	URIText := widgets.NewQLabelFromPointer(mainWidget.FindChild("URI", core.Qt__FindChildrenRecursively).Pointer())
	NameText := widgets.NewQLabelFromPointer(mainWidget.FindChild("name", core.Qt__FindChildrenRecursively).Pointer())
	databaseComboBox := widgets.NewQFontComboBoxFromPointer(mainWidget.FindChild("databaseComboBox", core.Qt__FindChildrenRecursively).Pointer())
	collectionComboBox := widgets.NewQFontComboBoxFromPointer(mainWidget.FindChild("collectionComboBox", core.Qt__FindChildrenRecursively).Pointer())
	queryPlainTextEdit := widgets.NewQPlainTextEditFromPointer(mainWidget.FindChild("queryPlainTextEdit", core.Qt__FindChildrenRecursively).Pointer())

	URIText.SetText(mongoURI)
	NameText.SetText(name)

	timeLabel := widgets.NewQLabelFromPointer(mainWidget.FindChild("timeLabel", core.Qt__FindChildrenRecursively).Pointer())
	skipSpinBox := widgets.NewQSpinBoxFromPointer(mainWidget.FindChild("skipSpinBox", core.Qt__FindChildrenRecursively).Pointer())
	skip := skipSpinBox.Value()
	skipSpinBox.ConnectValueChanged(func(i int) {
		skip = i
		globalState.skip = &skip
	})
	limitSpinBox := widgets.NewQSpinBoxFromPointer(mainWidget.FindChild("limitSpinBox", core.Qt__FindChildrenRecursively).Pointer())
	limit := limitSpinBox.Value()
	limitSpinBox.ConnectValueChanged(func(i int) {
		limit = i
		globalState.limit = &limit
	})
	pageLeftBtn := widgets.NewQPushButtonFromPointer(mainWidget.FindChild("pageLeftBtn", core.Qt__FindChildrenRecursively).Pointer())
	pageLeftBtn.ConnectClicked(func(bool) {
		skip -= limit
		if skip < 0 {
			skip = 0
		}
		globalState.skip = &skip
		skipSpinBox.SetValue(skip)
		globalState.ExecuteQuery()
	})
	pageRightBtn := widgets.NewQPushButtonFromPointer(mainWidget.FindChild("pageRightBtn", core.Qt__FindChildrenRecursively).Pointer())
	pageRightBtn.ConnectClicked(func(bool) {
		skip += limit
		globalState.skip = &skip
		skipSpinBox.SetValue(skip)
		globalState.ExecuteQuery()
	})

	currentDB := ""
	currentCollection := ""

	dbs, _ := connectdb.ListDB(mongoURI)
	collections := []string{}
	documents := []bson.M{}

	databaseComboBox.AddItems(dbs)
	databaseComboBox.ConnectCurrentIndexChanged(func(idx int) {
		currentDB = dbs[idx]
		globalState.currentDB = &currentDB
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
			globalState.currentCollection = &currentCollection
			queryPlainTextEdit.SetPlainText(".find({})")
		}
	})
	currentQuery := queryPlainTextEdit.ToPlainText()
	queryPlainTextEdit.ConnectTextChanged(func() {
		currentQuery = queryPlainTextEdit.ToPlainText()
		globalState.currentQuery = &currentQuery
	})

	resultTreeview := widgets.NewQTreeViewFromPointer(widget.FindChild("resultTreeView", core.Qt__FindChildrenRecursively).Pointer())
	model := tree.NewCustomTreeModel(nil)
	resultTreeview.SetModel(model)
	maxPossibleDocCount := &[]int{50}[0]

	resultTextView := widgets.NewQPlainTextEditFromPointer(widget.FindChild("resultTextView", core.Qt__FindChildrenRecursively).Pointer())
	resultTextView.SetTextInteractionFlags(resultTextView.TextInteractionFlags() | core.Qt__TextSelectableByKeyboard)
	resultTextView.Hide()

	globalState.queryPlainTextEdit = queryPlainTextEdit
	globalState.currentName = &name
	globalState.mongoURI = &mongoURI
	globalState.currentDB = &currentDB
	globalState.currentCollection = &currentCollection
	globalState.currentQuery = &currentQuery
	globalState.model = model
	globalState.documents = &documents
	globalState.tabStates[fmt.Sprintf("%v", mainWidget)] = TabState{
		documents: &documents,
	}
	globalState.maxPossibleDocCount = maxPossibleDocCount
	globalState.skip = &skip
	globalState.limit = &limit
	globalState.timeLabel = timeLabel
	globalState.resultTextEdit = resultTextView

	switchViewComboBox := widgets.NewQComboBoxFromPointer(mainWidget.FindChild("switchViewComboBox", core.Qt__FindChildrenRecursively).Pointer())
	switchViewComboBox.ConnectCurrentIndexChanged(func(idx int) {
		if idx == 0 { // table
			resultTextView.Hide()
			resultTreeview.Show()
		} else if idx == 1 { // text
			resultTreeview.Hide()
			resultTextView.Show()
		}
	})

	registerTemplateBtn(mainWidget, queryPlainTextEdit)
	registerActionBtn(
		mainWidget,
		queryPlainTextEdit,
		&mongoURI,
		&currentDB,
		&currentCollection,
		model,
		&documents,
		globalState,
	)
	registerDocOperationBtn(
		mainWidget,
		resultTreeview,
		&documents,
		&mongoURI,
		&currentDB,
		&currentCollection,
		globalState,
	)

	return mainWidget
}
