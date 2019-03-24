package main

import (
	"os"

	"github.com/naponmeka/robone/src"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func newWindow() {
	window := widgets.NewQMainWindow(nil, 0)
	tabsLayout := src.NewTabLayout()
	tabsHolder := widgets.NewQTabWidgetFromPointer(tabsLayout.FindChild("tabWidget", core.Qt__FindChildrenRecursively).Pointer())
	tabsHolder.ConnectTabCloseRequested(func(idx int) {
		tabsHolder.RemoveTab(idx)
	})
	tabsHolder.ConnectTabBarDoubleClicked(func(idx int) {
		if idx == -1 {
			tab := src.NewConnectLayout(tabsHolder)
			tabsHolder.AddTab(tab, "Connect")
			tabsHolder.SetCurrentIndex(tabsHolder.Count() - 1)
		}
	})
	menuBar := widgets.NewQMenuBar(nil)
	fileMenuBar := widgets.NewQMenu(nil)
	fileMenuBar.SetTitle("File")
	newWindowBtn := fileMenuBar.AddAction("New window")
	newWindowBtn.ConnectTriggered(func(bool) {
		newWindow()
	})
	newTabBtn := fileMenuBar.AddAction("New tab")
	newTabBtn.ConnectTriggered(func(bool) {
		tab := src.NewConnectLayout(tabsHolder)
		tabsHolder.AddTab(tab, "Connect")
	})
	fileMenuBar.AddSeparator()
	exportBtn := fileMenuBar.AddAction("Export")
	// exportBtn.SetDisabled(true)
	exportBtn.ConnectTriggered(func(bool) {
		subwin := widgets.NewQDialog(nil, 0)
		subwin.SetWindowTitle("Export")
		subwin.SetLayout(widgets.NewQHBoxLayout())
		exportLayout := src.NewExportLayout()
		subwin.Layout().AddWidget(exportLayout)
		subwin.SetModal(true)
		subwin.SetMinimumSize2(640, 480)
		src.RegisterExportLayoutBtn(exportLayout, subwin)
		subwin.Show()
	})
	importBtn := fileMenuBar.AddAction("Import")
	importBtn.SetDisabled(true)
	menuBar.AddMenu(fileMenuBar)
	window.SetMenuBar(menuBar)
	window.SetMinimumSize2(720, 480)
	window.SetWindowTitle("Robone")
	tab := src.NewConnectLayout(tabsHolder)
	tabsHolder.AddTab(tab, "Connect")
	window.SetCentralWidget(tabsLayout)
	window.Show()
}

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)
	newWindow()
	app.Exec()
}
