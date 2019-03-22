package main

import (
	"os"

	"github.com/naponmeka/robone/src"
	"github.com/therecipe/qt/widgets"
)

func newWindow() {
	window := widgets.NewQMainWindow(nil, 0)
	menuBar := widgets.NewQMenuBar(nil)
	fileMenuBar := widgets.NewQMenu(nil)
	fileMenuBar.SetTitle("File")
	newWindowBtn := fileMenuBar.AddAction("New window")
	newWindowBtn.ConnectTriggered(func(bool) {
		newWindow()
	})
	fileMenuBar.AddSeparator()
	exportBtn := fileMenuBar.AddAction("Export")
	// exportBtn.SetDisabled(true)
	exportBtn.ConnectTriggered(func(bool) {
		subwin := widgets.NewQDialog(nil, 0)
		subwin.SetWindowTitle("Export")
		subwin.SetLayout(widgets.NewQHBoxLayout())
		subwin.Layout().AddWidget(src.NewExportLayout())
		subwin.SetModal(true)
		subwin.SetMinimumSize2(640, 480)
		subwin.Show()
	})
	importBtn := fileMenuBar.AddAction("Import")
	importBtn.SetDisabled(true)
	menuBar.AddMenu(fileMenuBar)
	window.SetMenuBar(menuBar)
	window.SetMinimumSize2(720, 480)
	window.SetWindowTitle("Robone")
	widget := src.NewConnectLayout(window)
	widget.SetLayout(widgets.NewQVBoxLayout())
	window.SetCentralWidget(widget)
	window.Show()
}

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)
	newWindow()
	app.Exec()
}
