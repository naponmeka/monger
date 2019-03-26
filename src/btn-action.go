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
	maxPossibleDocCount *int,
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
	beginSpinBox.ConnectValueChanged(func(i int) {
		*maxPossibleDocCount = endSpinBox.Value() - i
	})

	endSpinBox.ConnectValueChanged(func(i int) {
		*maxPossibleDocCount = i - beginSpinBox.Value()
	})
	executeBtn := widgets.NewQPushButtonFromPointer(mainWidget.FindChild("executeBtn", core.Qt__FindChildrenRecursively).Pointer())
	executeBtn.ConnectClicked(func(bool) {
		executeQuery(maxPossibleDocCount, queryPlainTextEdit, mongoURI, currentDB, currentCollection, model, documents)
	})
}

func executeQuery(
	maxPossibleDocCount *int,
	queryPlainTextEdit *widgets.QPlainTextEdit,
	mongoURI *string,
	currentDB *string,
	currentCollection *string,
	model *CustomTreeModel,
	documents *[]bson.M,
) {
	currentQuery := queryPlainTextEdit.ToPlainText()
	for i := 0; i < *maxPossibleDocCount; i++ {
		model.Remove()
	}
	items, docs, err := createItems(*mongoURI, *currentDB, *currentCollection, currentQuery)
	*documents = docs
	if err != nil {
		widgets.QMessageBox_Information(nil, "Error", "wrong query", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
		return
	}
	for _, item := range items {
		model.Add(item)
	}
}
