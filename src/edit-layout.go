package src

import (
	"github.com/naponmeka/bsonparser"
	"github.com/naponmeka/robone/connectdb"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewEditLayout(docStr string) *widgets.QWidget {
	var widget = widgets.NewQWidget(nil, 0)
	var loader = uitools.NewQUiLoader(nil)
	var file = core.NewQFile2(":/qml/edit.ui")
	file.Open(core.QIODevice__ReadOnly)
	dialogWidget := loader.Load(file, widget)
	file.Close()
	plainTextEdit := widgets.NewQPlainTextEditFromPointer(dialogWidget.FindChild("plainTextEdit", core.Qt__FindChildrenRecursively).Pointer())
	plainTextEdit.SetPlainText(docStr)
	_ = NewHighlighter(plainTextEdit.Document())
	debounce := false
	textLength := len(plainTextEdit.ToPlainText())
	registerAutoCloseBracket(plainTextEdit, &debounce, &textLength)
	return dialogWidget
}

func RegisterEditLayoutBtn(
	widget *widgets.QWidget,
	subwin *widgets.QDialog,
	collection *mongo.Collection,
	globalState *GlobalState,
) {
	cancelBtn := widgets.NewQPushButtonFromPointer(widget.FindChild("cancelBtn", core.Qt__FindChildrenRecursively).Pointer())
	cancelBtn.ConnectClicked(func(bool) {
		subwin.Close()
	})
	saveBtn := widgets.NewQPushButtonFromPointer(widget.FindChild("saveBtn", core.Qt__FindChildrenRecursively).Pointer())
	saveBtn.ConnectClicked(func(bool) {
		plainTextEdit := widgets.NewQPlainTextEditFromPointer(widget.FindChild("plainTextEdit", core.Qt__FindChildrenRecursively).Pointer())
		raw := plainTextEdit.ToPlainText()
		jsonStr, _ := bsonparser.BsonToJson(raw)
		document := bson.M{}
		err := bson.UnmarshalExtJSON([]byte(jsonStr), true, &document)
		if err != nil {
			widgets.QMessageBox_Critical(nil, "Error", "Error:\n"+err.Error(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			return
		}
		filter := bson.D{}
		switch v := document["_id"].(type) {
		case primitive.ObjectID:
			objIDStr := v.Hex()
			oid, _ := primitive.ObjectIDFromHex(objIDStr)
			filter = bson.D{
				{"_id", oid},
			}
		default:
			filter = bson.D{
				{"_id", v},
			}
		}
		connectdb.Replace(*globalState.timeout, collection, filter, document, nil)
		subwin.Close()
		globalState.ExecuteQuery()
	})
}
