package src

import (
	"github.com/naponmeka/robone/tree"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"go.mongodb.org/mongo-driver/bson"
)

type GlobalState struct {
	tabsHolder          *widgets.QTabWidget
	window              *widgets.QMainWindow
	maxPossibleDocCount *int
	queryPlainTextEdit  *widgets.QPlainTextEdit
	mongoURI            *string
	currentDB           *string
	currentCollection   *string
	currentQuery        *string
	model               *tree.CustomTreeModel
	documents           *[]bson.M
	exportMenuBar       *widgets.QAction
}

func (gs *GlobalState) BindExecuteQuery(event *gui.QKeyEvent) {
	if event.Modifiers() == core.Qt__ControlModifier {
		if event.Key() == 16777220 {
			if gs.mongoURI != nil {
				executeQuery(
					gs.maxPossibleDocCount,
					gs.queryPlainTextEdit,
					gs.mongoURI,
					gs.currentDB,
					gs.currentCollection,
					gs.model,
					gs.documents,
				)
			}
		}
	}
}

func (gs *GlobalState) ExecuteQuery() {
	if gs.mongoURI != nil {
		executeQuery(
			gs.maxPossibleDocCount,
			gs.queryPlainTextEdit,
			gs.mongoURI,
			gs.currentDB,
			gs.currentCollection,
			gs.model,
			gs.documents,
		)
	}
}

func (gs *GlobalState) BindKeyboardTabControl(event *gui.QKeyEvent) {
	if event.Modifiers() == core.Qt__ControlModifier {
		if event.Key() == 84 { // T
			tab := NewConnectLayout(gs.tabsHolder, gs)
			gs.tabsHolder.AddTab(tab, "Connect")
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
		} else if event.Key() == 93 { //]
			if gs.tabsHolder.CurrentIndex() < gs.tabsHolder.Count()-1 {
				gs.tabsHolder.SetCurrentIndex(gs.tabsHolder.CurrentIndex() + 1)
			} else {
				gs.tabsHolder.SetCurrentIndex(0)
			}
		}
	}
}
