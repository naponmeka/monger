package src

import (
	"github.com/naponmeka/robone/tree"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"go.mongodb.org/mongo-driver/bson"
)

type KeyboardBinder struct {
	maxPossibleDocCount *int
	queryPlainTextEdit  *widgets.QPlainTextEdit
	mongoURI            *string
	currentDB           *string
	currentCollection   *string
	model               *tree.CustomTreeModel
	documents           *[]bson.M
}

func (kbBinder *KeyboardBinder) BindExecuteQuery(event *gui.QKeyEvent) {
	if event.Modifiers() == core.Qt__ControlModifier {
		if event.Key() == 16777220 {
			if kbBinder.mongoURI != nil {
				executeQuery(
					kbBinder.maxPossibleDocCount,
					kbBinder.queryPlainTextEdit,
					kbBinder.mongoURI,
					kbBinder.currentDB,
					kbBinder.currentCollection,
					kbBinder.model,
					kbBinder.documents,
				)
			}
		}
	}
}

func (kbBinder *KeyboardBinder) ExecuteQuery() {
	if kbBinder.mongoURI != nil {
		executeQuery(
			kbBinder.maxPossibleDocCount,
			kbBinder.queryPlainTextEdit,
			kbBinder.mongoURI,
			kbBinder.currentDB,
			kbBinder.currentCollection,
			kbBinder.model,
			kbBinder.documents,
		)
	}
}

func (kbBinder *KeyboardBinder) BindKeyboardTabControl(event *gui.QKeyEvent, tabsHolder *widgets.QTabWidget, window *widgets.QMainWindow) {
	if event.Modifiers() == core.Qt__ControlModifier {
		if event.Key() == 84 { // T
			tab := NewConnectLayout(tabsHolder, kbBinder)
			tabsHolder.AddTab(tab, "Connect")
			tabsHolder.SetCurrentIndex(tabsHolder.Count() - 1)
		} else if event.Key() == 87 { // W
			CurrentIndex := tabsHolder.CurrentIndex()
			tabsHolder.SetCurrentIndex(CurrentIndex - 1)
			tabsHolder.RemoveTab(CurrentIndex)
			if tabsHolder.Count() < 1 {
				window.Close()
			}
		}
	} else if event.Modifiers() == (core.Qt__ControlModifier | core.Qt__ShiftModifier) {
		if event.Key() == 91 { // [
			if tabsHolder.CurrentIndex() > 0 {
				tabsHolder.SetCurrentIndex(tabsHolder.CurrentIndex() - 1)
			} else {
				tabsHolder.SetCurrentIndex(tabsHolder.Count() - 1)
			}
		} else if event.Key() == 93 { //]
			if tabsHolder.CurrentIndex() < tabsHolder.Count()-1 {
				tabsHolder.SetCurrentIndex(tabsHolder.CurrentIndex() + 1)
			} else {
				tabsHolder.SetCurrentIndex(0)
			}
		}
	}
}
