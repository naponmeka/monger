package src

import (
	"github.com/naponmeka/robone/connectdb"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
	"go.mongodb.org/mongo-driver/bson"
)

func NewMainLayout(mongoURI string, keyboardBinder *KeyboardBinder) *widgets.QWidget {
	var widget = widgets.NewQWidget(nil, 0)
	var loader = uitools.NewQUiLoader(nil)
	var file = core.NewQFile2(":/qml/main.ui")

	file.Open(core.QIODevice__ReadOnly)
	mainWidget := loader.Load(file, widget)
	file.Close()
	databaseComboBox := widgets.NewQFontComboBoxFromPointer(mainWidget.FindChild("databaseComboBox", core.Qt__FindChildrenRecursively).Pointer())
	collectionComboBox := widgets.NewQFontComboBoxFromPointer(mainWidget.FindChild("collectionComboBox", core.Qt__FindChildrenRecursively).Pointer())
	queryPlainTextEdit := widgets.NewQPlainTextEditFromPointer(mainWidget.FindChild("queryPlainTextEdit", core.Qt__FindChildrenRecursively).Pointer())
	beginSpinBox := widgets.NewQSpinBoxFromPointer(mainWidget.FindChild("beginSpinBox", core.Qt__FindChildrenRecursively).Pointer())
	endSpinBox := widgets.NewQSpinBoxFromPointer(mainWidget.FindChild("endSpinBox", core.Qt__FindChildrenRecursively).Pointer())

	currentDB := ""
	currentCollection := ""

	dbs := connectdb.ListDB(mongoURI)
	collections := []string{}
	documents := []bson.M{}

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
			queryPlainTextEdit.SetPlainText(".find({})")
		}
	})

	resultTreeview := widgets.NewQTreeViewFromPointer(widget.FindChild("resultTreeView", core.Qt__FindChildrenRecursively).Pointer())
	model := NewCustomTreeModel(nil)
	resultTreeview.SetModel(model)
	maxPossibleDocCount := &[]int{50}[0]
	keyboardBinder.queryPlainTextEdit = queryPlainTextEdit
	keyboardBinder.mongoURI = &mongoURI
	keyboardBinder.currentDB = &currentDB
	keyboardBinder.currentCollection = &currentCollection
	keyboardBinder.model = model
	keyboardBinder.documents = &documents
	keyboardBinder.maxPossibleDocCount = maxPossibleDocCount

	registerTemplateBtn(mainWidget, queryPlainTextEdit)
	registerActionBtn(
		mainWidget,
		queryPlainTextEdit,
		endSpinBox,
		beginSpinBox,
		&mongoURI,
		&currentDB,
		&currentCollection,
		model,
		&documents,
		maxPossibleDocCount,
	)
	registerDocOperationBtn(
		mainWidget,
		resultTreeview,
		&documents,
		&mongoURI,
		&currentDB,
		&currentCollection,
		keyboardBinder,
	)

	return mainWidget
}
