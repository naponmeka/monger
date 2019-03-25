package src

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

func NewEditLayout(docStr string) *widgets.QWidget {
	var widget = widgets.NewQWidget(nil, 0)
	var loader = uitools.NewQUiLoader(nil)
	var file = core.NewQFile2(":/qml/edit.ui")
	file.Open(core.QIODevice__ReadOnly)
	exportDialogWidget := loader.Load(file, widget)
	file.Close()
	plainTextEdit := widgets.NewQPlainTextEditFromPointer(widget.FindChild("plainTextEdit", core.Qt__FindChildrenRecursively).Pointer())
	plainTextEdit.SetPlainText(docStr)
	return exportDialogWidget
}

func RegisterEditLayoutBtn(widget *widgets.QWidget, subwin *widgets.QDialog) {
	cancelBtn := widgets.NewQPushButtonFromPointer(widget.FindChild("cancelBtn", core.Qt__FindChildrenRecursively).Pointer())
	cancelBtn.ConnectClicked(func(bool) {
		subwin.Close()
	})
}
