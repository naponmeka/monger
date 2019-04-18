package src

import (
	"github.com/naponmeka/bsonparser"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

func NewViewLayout(docStr string) *widgets.QWidget {
	var widget = widgets.NewQWidget(nil, 0)
	var loader = uitools.NewQUiLoader(nil)
	var file = core.NewQFile2(":/qml/view.ui")
	file.Open(core.QIODevice__ReadOnly)
	dialogWidget := loader.Load(file, widget)
	file.Close()
	plainTextEdit := widgets.NewQPlainTextEditFromPointer(dialogWidget.FindChild("plainTextEdit", core.Qt__FindChildrenRecursively).Pointer())
	plainTextEdit.SetPlainText(docStr)
	plainTextEdit.SetTextInteractionFlags(plainTextEdit.TextInteractionFlags() | core.Qt__TextSelectableByKeyboard)
	return dialogWidget
}

func RegisterViewLayoutBtn(widget *widgets.QWidget, subwin *widgets.QDialog) {
	plainTextEdit := widgets.NewQPlainTextEditFromPointer(widget.FindChild("plainTextEdit", core.Qt__FindChildrenRecursively).Pointer())

	bsonRadioButton := widgets.NewQRadioButtonFromPointer(widget.FindChild("bsonRadioButton", core.Qt__FindChildrenRecursively).Pointer())
	bsonRadioButton.ConnectClicked(func(bool) {
		currentStr := plainTextEdit.ToPlainText()
		docStr, err := bsonparser.JsonToBson(currentStr)
		if err == nil {
			plainTextEdit.SetPlainText(docStr)
		}
	})

	jsonRadioButton := widgets.NewQRadioButtonFromPointer(widget.FindChild("jsonRadioButton", core.Qt__FindChildrenRecursively).Pointer())
	jsonRadioButton.ConnectClicked(func(bool) {
		currentStr := plainTextEdit.ToPlainText()
		docStr, err := bsonparser.BsonToJson(currentStr)
		if err == nil {
			plainTextEdit.SetPlainText(docStr)
		} else {
			widgets.QMessageBox_Critical(nil, "Error", "Error:\n"+err.Error(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
		}
	})
	cancelBtn := widgets.NewQPushButtonFromPointer(widget.FindChild("cancelBtn", core.Qt__FindChildrenRecursively).Pointer())
	cancelBtn.ConnectClicked(func(bool) {
		subwin.Close()
	})
	copyBtn := widgets.NewQPushButtonFromPointer(widget.FindChild("copyBtn", core.Qt__FindChildrenRecursively).Pointer())
	copyBtn.ConnectClicked(func(bool) {
		gui.QGuiApplication_Clipboard().SetText(plainTextEdit.ToPlainText(), gui.QClipboard__Clipboard)

	})
}
