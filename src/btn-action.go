package src

import (
	"fmt"
	"strings"

	"github.com/naponmeka/robone/tree"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"go.mongodb.org/mongo-driver/bson"
)

func registerActionBtn(
	mainWidget *widgets.QWidget,
	queryPlainTextEdit *widgets.QPlainTextEdit,
	mongoURI *string,
	currentDB *string,
	currentCollection *string,
	model *tree.CustomTreeModel,
	documents *[]bson.M,
	globalState *GlobalState,
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
		executeQuery(queryPlainTextEdit, mongoURI, currentDB, currentCollection, model, documents, globalState.skip, globalState.limit)
	})
}

func executeQuery(
	queryPlainTextEdit *widgets.QPlainTextEdit,
	mongoURI *string,
	currentDB *string,
	currentCollection *string,
	model *tree.CustomTreeModel,
	documents *[]bson.M,
	skip *int,
	limit *int,
) {
	currentQuery := queryPlainTextEdit.ToPlainText()
	for i := 0; i < 100; i++ {
		model.Remove()
	}
	items, docs, err := tree.CreateItems(*mongoURI, *currentDB, *currentCollection, currentQuery, *skip, *limit)
	*documents = docs
	if err != nil {
		widgets.QMessageBox_Information(nil, "Error", "wrong query", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
		return
	}
	for _, item := range items {
		model.Add(item)
	}
}
