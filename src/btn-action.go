package src

import (
	"fmt"
	"strings"
	"time"

	"github.com/naponmeka/bsonparser"
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
		executeQuery(globalState)
	})
}

func executeQuery(gs *GlobalState) {
	gs.timeLabel.SetText("Loading..")
	start := time.Now()

	currentQuery := gs.queryPlainTextEdit.ToPlainText()
	for i := 0; i < 100; i++ {
		gs.model.Remove()
	}
	items, docs, err := tree.CreateItems(*gs.mongoURI, *gs.currentDB, *gs.currentCollection, currentQuery, *gs.skip, *gs.limit)
	if err != nil {
		widgets.QMessageBox_Critical(nil, "Error", "Error:\n"+err.Error(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
		return
	}
	*gs.documents = docs
	for _, item := range items {
		gs.model.Add(item)
	}
	resultText := "[\n"
	resultTextJson := "[\n"

	docTexts := []string{}
	docTextJson := []string{}
	for _, doc := range docs {
		docByte, _ := bson.MarshalExtJSON(doc, false, true)
		docStr, _ := bsonparser.JsonToBson(string(docByte[:]))
		docTexts = append(docTexts, docStr)
		docStrJson, _ := bsonparser.BsonToJson(docStr)
		docTextJson = append(docTextJson, docStrJson)
	}
	resultText += strings.Join(docTexts, ",\n\n")
	resultText += "\n]"

	resultTextJson += strings.Join(docTextJson, ",\n\n")
	resultTextJson += "\n]"

	gs.resultTextEdit.SetPlainText(resultText)
	_ = NewHighlighter(gs.resultTextEdit.Document())
	gs.resultTextEditJson.SetPlainText(resultTextJson)
	_ = NewHighlighter(gs.resultTextEditJson.Document())

	t := time.Now()
	elapsed := t.Sub(start)
	gs.timeLabel.SetText(fmt.Sprintf("Took: %.3f sec.", elapsed.Seconds()))
}
