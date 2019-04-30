package src

import (
	"fmt"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

func NewWindow() {
	window := widgets.NewQMainWindow(nil, 0)
	tabsLayout := NewTabLayout()
	tabsHolder := widgets.NewQTabWidgetFromPointer(tabsLayout.FindChild("tabWidget", core.Qt__FindChildrenRecursively).Pointer())
	config := readConfig()
	gs := GlobalState{
		tabsHolder:  tabsHolder,
		window:      window,
		tabStates:   make(map[string]TabState),
		timeout:     &config.Timeout,
		defaultView: &config.DefaultView,
	}
	tabsHolder.ConnectCurrentChanged(func(idx int) {
		gs.TabChangeUpdateState()
	})

	tabsHolder.ConnectTabCloseRequested(func(idx int) {
		tabsHolder.RemoveTab(idx)
		mainLayout := gs.tabsHolder.CurrentWidget()
		key := fmt.Sprintf("%v", mainLayout)
		if _, ok := gs.tabStates[key]; ok {
			delete(gs.tabStates, key)
		}
	})
	tabsHolder.ConnectTabBarDoubleClicked(func(idx int) {
		if idx == -1 {
			tab := NewConnectLayout(tabsHolder, &gs)
			tabsHolder.AddTab(tab, "Connect")
			tabsHolder.SetCurrentIndex(tabsHolder.Count() - 1)
		}
	})
	connectLayout := NewConnectLayout(tabsHolder, &gs)
	tabsHolder.AddTab(connectLayout, "Connect")
	menuBar := CreateMenuBar(tabsHolder, &gs)
	window.SetMenuBar(menuBar)
	window.SetMinimumSize2(720, 480)
	window.SetWindowTitle("Monger")
	window.SetCentralWidget(tabsLayout)

	window.ConnectKeyPressEvent(func(event *gui.QKeyEvent) {
		gs.BindKeyboardTabControl(event)
		gs.BindExecuteQuery(event)
	})
	window.Show()
}
