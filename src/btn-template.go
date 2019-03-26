package src

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func registerTemplateBtn(mainWidget *widgets.QWidget, queryPlainTextEdit *widgets.QPlainTextEdit) {
	findBtn := widgets.NewQPushButtonFromPointer(mainWidget.FindChild("findBtn", core.Qt__FindChildrenRecursively).Pointer())
	insertBtn := widgets.NewQPushButtonFromPointer(mainWidget.FindChild("insertBtn", core.Qt__FindChildrenRecursively).Pointer())
	updateBtn := widgets.NewQPushButtonFromPointer(mainWidget.FindChild("updateBtn", core.Qt__FindChildrenRecursively).Pointer())
	deleteBtn := widgets.NewQPushButtonFromPointer(mainWidget.FindChild("deleteBtn", core.Qt__FindChildrenRecursively).Pointer())
	aggBtn := widgets.NewQPushButtonFromPointer(mainWidget.FindChild("aggBtn", core.Qt__FindChildrenRecursively).Pointer())
	findBtn.ConnectClicked(func(bool) {
		queryPlainTextEdit.SetPlainText(".find({})")
	})
	insertBtn.ConnectClicked(func(bool) {
		queryPlainTextEdit.SetPlainText(`.insert(
  [
    {}
  ],
  {
    "ordered": true
  }
])`)
	})
	updateBtn.ConnectClicked(func(bool) {
		txt := `.update(
  {query},
  {update},
  {
    "upsert": false,
    "multi": false
  }
)`
		queryPlainTextEdit.SetPlainText(txt)
	})
	deleteBtn.ConnectClicked(func(bool) {
		txt := `.remove(
  {query},
  {
    justOne: true
  }
)`
		queryPlainTextEdit.SetPlainText(txt)
	})
	aggBtn.ConnectClicked(func(bool) {
		txt := `.aggregate([
  {},
])`
		queryPlainTextEdit.SetPlainText(txt)
	})
}
