package src

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

func BindKeyboardTabControl(event *gui.QKeyEvent, tabsHolder *widgets.QTabWidget, window *widgets.QMainWindow) {
	if event.Modifiers() == core.Qt__ControlModifier {
		if event.Key() == 84 { // T
			tab := NewConnectLayout(tabsHolder)
			tabsHolder.AddTab(tab, "Connect")
			tabsHolder.SetCurrentIndex(tabsHolder.Count() - 1)
		} else if event.Key() == 87 { // W
			CurrentIndex := tabsHolder.CurrentIndex()
			tabsHolder.SetCurrentIndex(CurrentIndex - 1)
			tabsHolder.RemoveTab(CurrentIndex)
			if tabsHolder.Count() < 1 {
				window.Close()
			}
		} else if event.Key() == 16777220 {
			widgets.QMessageBox_Information(nil, "OK", "going to query", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
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
