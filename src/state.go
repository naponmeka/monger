package src

import (
	"fmt"

	"github.com/naponmeka/monger/tree"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"go.mongodb.org/mongo-driver/bson"
)

type TabState struct {
	documents *[]bson.M
}

type GlobalState struct {
	tabsHolder          *widgets.QTabWidget
	tabStates           map[string]TabState
	window              *widgets.QMainWindow
	maxPossibleDocCount *int
	queryPlainTextEdit  *widgets.QPlainTextEdit
	currentName         *string
	mongoURI            *string
	currentDB           *string
	currentCollection   *string
	currentQuery        *string
	model               *tree.CustomTreeModel
	documents           *[]bson.M
	exportMenuBar       *widgets.QAction
	skip                *int
	limit               *int
	timeLabel           *widgets.QLabel
	resultTextEdit      *widgets.QPlainTextEdit
	resultTextEditJson  *widgets.QPlainTextEdit
	defaultView         *int
	timeout             *int
	indice              *[]bson.M
}

func (gs *GlobalState) BindExecuteQuery(event *gui.QKeyEvent) {
	if event.Modifiers() == core.Qt__ControlModifier {
		if event.Key() == 16777220 {
			if gs.mongoURI != nil {
				gs.ExecuteQuery()
			}
		}
	}
}

func (gs *GlobalState) ExecuteQuery() {
	if gs.mongoURI != nil {
		executeQuery(gs)
	}
}

func (gs *GlobalState) BindKeyboardTabControl(event *gui.QKeyEvent) {
	if event.Modifiers() == core.Qt__ControlModifier {
		if event.Key() == 78 { // N
			tab := NewConnectLayout(gs.tabsHolder, gs)
			gs.tabsHolder.AddTab(tab, "Connect")
			gs.tabsHolder.SetCurrentIndex(gs.tabsHolder.Count() - 1)
		} else if event.Key() == 84 { // T
			mainQueryWidget := NewMainLayout(*gs.mongoURI, gs, *gs.currentName)
			gs.tabsHolder.AddTab(mainQueryWidget, *gs.currentName)
			gs.tabsHolder.SetCurrentIndex(gs.tabsHolder.Count() - 1)
		} else if event.Key() == 87 { // W
			CurrentIndex := gs.tabsHolder.CurrentIndex()
			gs.tabsHolder.SetCurrentIndex(CurrentIndex - 1)
			gs.tabsHolder.RemoveTab(CurrentIndex)
			if gs.tabsHolder.Count() < 1 {
				gs.window.Close()
			}
		}
	} else if event.Modifiers() == (core.Qt__ControlModifier | core.Qt__ShiftModifier) {
		if event.Key() == 91 { // [
			if gs.tabsHolder.CurrentIndex() > 0 {
				gs.tabsHolder.SetCurrentIndex(gs.tabsHolder.CurrentIndex() - 1)
			} else {
				gs.tabsHolder.SetCurrentIndex(gs.tabsHolder.Count() - 1)
			}
		} else if event.Key() == 93 { // ]
			if gs.tabsHolder.CurrentIndex() < gs.tabsHolder.Count()-1 {
				gs.tabsHolder.SetCurrentIndex(gs.tabsHolder.CurrentIndex() + 1)
			} else {
				gs.tabsHolder.SetCurrentIndex(0)
			}
		}
	}
}

func (gs *GlobalState) TabChangeUpdateState() {
	mainLayout := gs.tabsHolder.CurrentWidget()
	hiddenNameLabel := mainLayout.FindChild("name", core.Qt__FindChildrenRecursively)
	if hiddenNameLabel.IsWidgetType() {
		name := widgets.NewQLabelFromPointer(hiddenNameLabel.Pointer()).Text()
		mongoURI := widgets.NewQLabelFromPointer(mainLayout.FindChild("URI", core.Qt__FindChildrenRecursively).Pointer()).Text()
		database := widgets.NewQFontComboBoxFromPointer(mainLayout.FindChild("databaseComboBox", core.Qt__FindChildrenRecursively).Pointer()).CurrentText()
		collection := widgets.NewQFontComboBoxFromPointer(mainLayout.FindChild("collectionComboBox", core.Qt__FindChildrenRecursively).Pointer()).CurrentText()
		skip := widgets.NewQSpinBoxFromPointer(mainLayout.FindChild("skipSpinBox", core.Qt__FindChildrenRecursively).Pointer()).Value()
		limit := widgets.NewQSpinBoxFromPointer(mainLayout.FindChild("limitSpinBox", core.Qt__FindChildrenRecursively).Pointer()).Value()
		timeLabel := widgets.NewQLabelFromPointer(mainLayout.FindChild("timeLabel", core.Qt__FindChildrenRecursively).Pointer())
		gs.currentName = &name
		gs.mongoURI = &mongoURI
		gs.currentDB = &database
		gs.currentCollection = &collection
		gs.skip = &skip
		gs.limit = &limit
		gs.timeLabel = timeLabel
		gs.queryPlainTextEdit = widgets.NewQPlainTextEditFromPointer(mainLayout.FindChild("queryPlainTextEdit", core.Qt__FindChildrenRecursively).Pointer())
		resultTreeview := widgets.NewQTreeViewFromPointer(mainLayout.FindChild("resultTreeView", core.Qt__FindChildrenRecursively).Pointer())
		oldModel := resultTreeview.Model()
		n := new(tree.CustomTreeModel)
		n = &tree.CustomTreeModel{QAbstractItemModel: *oldModel}
		gs.model = n
		if val, ok := gs.tabStates[fmt.Sprintf("%v", mainLayout)]; ok {
			gs.documents = val.documents
		}
		gs.resultTextEdit = widgets.NewQPlainTextEditFromPointer(mainLayout.FindChild("resultTextView", core.Qt__FindChildrenRecursively).Pointer())
	}

}
