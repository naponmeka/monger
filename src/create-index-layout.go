package src

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

func NewCreateIndexLayout(subwin *widgets.QDialog) *widgets.QWidget {
	var widget = widgets.NewQWidget(nil, 0)
	var loader = uitools.NewQUiLoader(nil)
	var file = core.NewQFile2(":/qml/create-index.ui")
	file.Open(core.QIODevice__ReadOnly)
	layoutWidget := loader.Load(file, widget)
	file.Close()
	cancelBtn := widgets.NewQPushButtonFromPointer(layoutWidget.FindChild("cancelBtn", core.Qt__FindChildrenRecursively).Pointer())
	cancelBtn.ConnectClicked(func(bool) {
		subwin.Close()
	})
	return layoutWidget
}
