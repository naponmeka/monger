package src

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func CreateMenuBar(tabsHolder *widgets.QTabWidget, globalState *GlobalState) *widgets.QMenuBar {
	menuBar := widgets.NewQMenuBar(nil)
	fileMenuBar := widgets.NewQMenu(nil)
	fileMenuBar.SetTitle("File")
	newWindowBtn := fileMenuBar.AddAction("New window")
	newWindowBtn.ConnectTriggered(func(bool) {
		NewWindow()
	})
	newTabBtn := fileMenuBar.AddAction("New tab")
	newTabBtn.ConnectTriggered(func(bool) {
		tab := NewConnectLayout(tabsHolder, globalState)
		tabsHolder.AddTab(tab, "Connect")
	})
	fileMenuBar.AddSeparator()
	exportBtn := fileMenuBar.AddAction("Export")
	exportBtn.ConnectTriggered(func(bool) {
		subwin := widgets.NewQDialog(nil, 0)
		subwin.SetWindowTitle("Export")
		subwin.SetLayout(widgets.NewQHBoxLayout())
		exportLayout := NewExportLayout(globalState)
		subwin.Layout().AddWidget(exportLayout)
		// subwin.SetModal(true)
		subwin.SetMinimumSize2(640, 480)
		RegisterExportLayoutBtn(exportLayout, subwin)
		subwin.Show()
	})
	exportBtn.SetDisabled(true)
	globalState.exportMenuBar = exportBtn
	importBtn := fileMenuBar.AddAction("Import")
	importBtn.SetDisabled(true)

	settingMenuBar := widgets.NewQMenu(nil)
	settingMenuBar.SetTitle("Other")
	customizeBtn := settingMenuBar.AddAction("Customize")
	customizeBtn.ConnectTriggered(func(bool) {
		subwin := widgets.NewQDialog(nil, 0)
		subwin.SetWindowTitle("Customize")
		subwin.SetLayout(widgets.NewQHBoxLayout())
		exportLayout := NewSettingsLayout(globalState, subwin)
		subwin.Layout().AddWidget(exportLayout)
		subwin.SetModal(true)
		subwin.SetMinimumSize2(640, 400)
		subwin.Show()
	})

	aboutBtn := settingMenuBar.AddAction("Check for updates")
	aboutBtn.ConnectTriggered(func(bool) {
		subwin := widgets.NewQDialog(nil, 0)
		subwin.SetWindowTitle("About")
		subwin.SetLayout(widgets.NewQHBoxLayout())
		versionLayout := NewVersionLayout()
		closeBtn := widgets.NewQPushButtonFromPointer(versionLayout.FindChild("closeBtn", core.Qt__FindChildrenRecursively).Pointer())
		closeBtn.ConnectClicked(func(bool) {
			subwin.Close()
		})
		checkUpdateBtn := widgets.NewQPushButtonFromPointer(versionLayout.FindChild("checkUpdateBtn", core.Qt__FindChildrenRecursively).Pointer())
		checkUpdateBtn.ConnectClicked(func(bool) {
			widgets.QMessageBox_Information(nil, "Status", "Nothing to do.", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
		})
		subwin.Layout().AddWidget(versionLayout)
		subwin.SetModal(true)
		subwin.Show()
	})

	menuBar.AddMenu(fileMenuBar)
	menuBar.AddMenu(settingMenuBar)
	return menuBar
}
