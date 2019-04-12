package src

import (
	"strings"

	"github.com/naponmeka/bsonparser"
	"github.com/naponmeka/robone/connectdb"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewInsertLayout() *widgets.QWidget {
	var widget = widgets.NewQWidget(nil, 0)
	var loader = uitools.NewQUiLoader(nil)
	var file = core.NewQFile2(":/qml/insert.ui")
	file.Open(core.QIODevice__ReadOnly)
	dialogWidget := loader.Load(file, widget)
	file.Close()
	plainTextEdit := widgets.NewQPlainTextEditFromPointer(dialogWidget.FindChild("plainTextEdit", core.Qt__FindChildrenRecursively).Pointer())
	plainTextEdit.SetPlainText(`[
  {

  }
]`)
	return dialogWidget
}

func RegisterInsertLayoutBtn(
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
		jsonStr, err := bsonparser.BsonToJson(raw)
		if err != nil {
			widgets.QMessageBox_Critical(nil, "Error", "Error:\n"+err.Error(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			return
		}
		jsonStr = strings.TrimSpace(jsonStr)
		isArray := strings.HasPrefix(jsonStr, "[")
		if isArray {
			var documents []interface{}
			err := bson.UnmarshalExtJSON([]byte(jsonStr), false, &documents)
			if err != nil {
				widgets.QMessageBox_Critical(nil, "Error", "Error:\n"+err.Error(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
				return
			}
			connectdb.Insert(collection, documents, nil)
		} else {
			document := bson.M{}
			err := bson.UnmarshalExtJSON([]byte(jsonStr), false, &document)
			if err != nil {
				widgets.QMessageBox_Critical(nil, "Error", "Error:\n"+err.Error(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
				return
			}
			connectdb.Insert(collection, []interface{}{document}, nil)
		}
		subwin.Close()
		globalState.ExecuteQuery()
	})
}
