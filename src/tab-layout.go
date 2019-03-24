package src

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

func NewTabLayout() *widgets.QWidget {
	var widget = widgets.NewQWidget(nil, 0)
	var loader = uitools.NewQUiLoader(nil)
	var file = core.NewQFile2(":/qml/tabs.ui")
	file.Open(core.QIODevice__ReadOnly)
	tabsWidget := loader.Load(file, widget)
	file.Close()
	return tabsWidget
}

func ReplaceTabContent(tabsHolder *widgets.QTabWidget, replacement *widgets.QWidget, tabName string) {
	currentTabIdx := tabsHolder.CurrentIndex()
	tabsHolder.RemoveTab(currentTabIdx)
	tabsHolder.InsertTab(currentTabIdx, replacement, tabName)
	tabsHolder.SetCurrentIndex(currentTabIdx)
}
