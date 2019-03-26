package src

import (
	"fmt"
	"strings"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"go.mongodb.org/mongo-driver/bson"
)

func registerActionBtn(
	mainWidget *widgets.QWidget,
	queryPlainTextEdit *widgets.QPlainTextEdit,
	endSpinBox *widgets.QSpinBox,
	beginSpinBox *widgets.QSpinBox,
	mongoURI *string,
	currentDB *string,
	currentCollection *string,
	model *CustomTreeModel,
	documents *[]bson.M,
) {
	countBtn := widgets.NewQPushButtonFromPointer(mainWidget.FindChild("countBtn", core.Qt__FindChildrenRecursively).Pointer())
	explainBtn := widgets.NewQPushButtonFromPointer(mainWidget.FindChild("explainBtn", core.Qt__FindChildrenRecursively).Pointer())

	countBtn.ConnectClicked(func(bool) {
		currentText := queryPlainTextEdit.ToPlainText()
		currentText = strings.TrimSuffix(currentText, ".count()")
		currentText = strings.TrimSuffix(currentText, ".explain()")
		queryPlainTextEdit.SetPlainText(fmt.Sprintf("%s.count()", currentText))
	})
	explainBtn.ConnectClicked(func(bool) {
		currentText := queryPlainTextEdit.ToPlainText()
		currentText = strings.TrimSuffix(currentText, ".count()")
		currentText = strings.TrimSuffix(currentText, ".explain()")
		queryPlainTextEdit.SetPlainText(fmt.Sprintf("%s.explain()", currentText))
	})

	executeBtn := widgets.NewQPushButtonFromPointer(mainWidget.FindChild("executeBtn", core.Qt__FindChildrenRecursively).Pointer())
	executeBtn.ConnectClicked(func(bool) {
		maxPossibleDocCount := endSpinBox.Value() - beginSpinBox.Value()
		for i := 0; i < maxPossibleDocCount; i++ {
			model.Remove()
		}
		currentQuery := queryPlainTextEdit.ToPlainText()
		items, docs, err := createItems(*mongoURI, *currentDB, *currentCollection, currentQuery)
		*documents = docs
		if err != nil {
			widgets.QMessageBox_Information(nil, "Error", "wrong query", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			return
		}
		for _, item := range items {
			model.Add(item)
		}
	})
}
