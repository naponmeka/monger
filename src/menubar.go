package src

import (
	"github.com/therecipe/qt/widgets"
)

func CreateMenuBar(tabsHolder *widgets.QTabWidget, keyboardBinder *KeyboardBinder) *widgets.QMenuBar {
	menuBar := widgets.NewQMenuBar(nil)
	fileMenuBar := widgets.NewQMenu(nil)
	fileMenuBar.SetTitle("File")
	newWindowBtn := fileMenuBar.AddAction("New window")
	newWindowBtn.ConnectTriggered(func(bool) {
		NewWindow()
	})
	newTabBtn := fileMenuBar.AddAction("New tab")
	newTabBtn.ConnectTriggered(func(bool) {
		tab := NewConnectLayout(tabsHolder, keyboardBinder)
		tabsHolder.AddTab(tab, "Connect")
	})
	fileMenuBar.AddSeparator()
	exportBtn := fileMenuBar.AddAction("Export")
	exportBtn.ConnectTriggered(func(bool) {
		subwin := widgets.NewQDialog(nil, 0)
		subwin.SetWindowTitle("Export")
		subwin.SetLayout(widgets.NewQHBoxLayout())
		exportLayout := NewExportLayout()
		subwin.Layout().AddWidget(exportLayout)
		// subwin.SetModal(true)
		subwin.SetMinimumSize2(640, 480)
		RegisterExportLayoutBtn(exportLayout, subwin)
		subwin.Show()
	})
	importBtn := fileMenuBar.AddAction("Import")
	importBtn.SetDisabled(true)

	settingMenuBar := widgets.NewQMenu(nil)
	settingMenuBar.SetTitle("Setting")
	timeoutBtn := settingMenuBar.AddAction("Timeout")
	timeoutBtn.SetDisabled(true)

	menuBar.AddMenu(fileMenuBar)
	menuBar.AddMenu(settingMenuBar)
	return menuBar
}
