package src

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

func NewConnectLayout(tabsHolder *widgets.QTabWidget, keyboardBinder *KeyboardBinder) *widgets.QWidget {
	var widget = widgets.NewQWidget(nil, 0)
	var loader = uitools.NewQUiLoader(nil)
	var file = core.NewQFile2(":/qml/connect.ui")
	file.Open(core.QIODevice__ReadOnly)
	connectWidget := loader.Load(file, widget)
	file.Close()
	testBtn := widgets.NewQPushButtonFromPointer(widget.FindChild("testBtn", core.Qt__FindChildrenRecursively).Pointer())
	testBtn.ConnectClicked(func(bool) {
		widgets.QMessageBox_Information(nil, "OK", "Success", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	})

	URILineEdit := widgets.NewQLineEditFromPointer(widget.FindChild("mylineedit", core.Qt__FindChildrenRecursively).Pointer())
	URILineEdit.SetText("mongodb://root:root@localhost:27017/admin")
	connectBtn := widgets.NewQPushButtonFromPointer(widget.FindChild("connectBtn", core.Qt__FindChildrenRecursively).Pointer())
	connectBtn.ConnectClicked(func(bool) {
		URI := URILineEdit.Text()
		mainQueryWidget := NewMainLayout(URI, keyboardBinder)
		ReplaceTabContent(tabsHolder, mainQueryWidget, "Query")
	})
	return connectWidget
}