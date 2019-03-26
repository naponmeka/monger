package src

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

func NewWindow() {
	window := widgets.NewQMainWindow(nil, 0)
	tabsLayout := NewTabLayout()
	tabsHolder := widgets.NewQTabWidgetFromPointer(tabsLayout.FindChild("tabWidget", core.Qt__FindChildrenRecursively).Pointer())
	tabsHolder.ConnectTabCloseRequested(func(idx int) {
		tabsHolder.RemoveTab(idx)
	})
	tabsHolder.ConnectTabBarDoubleClicked(func(idx int) {
		if idx == -1 {
			tab := NewConnectLayout(tabsHolder)
			tabsHolder.AddTab(tab, "Connect")
			tabsHolder.SetCurrentIndex(tabsHolder.Count() - 1)
		}
	})
	tab := NewConnectLayout(tabsHolder)
	tabsHolder.AddTab(tab, "Connect")
	menuBar := CreateMenuBar(tabsHolder)
	window.SetMenuBar(menuBar)
	window.SetMinimumSize2(720, 480)
	window.SetWindowTitle("Robone")
	window.SetCentralWidget(tabsLayout)
	window.ConnectKeyPressEvent(func(event *gui.QKeyEvent) {
		BindKeyboardTabControl(event, tabsHolder, window)
	})
	window.Show()
}
