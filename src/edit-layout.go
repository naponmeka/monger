package src

import (
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
	exportDialogWidget := loader.Load(file, widget)
	file.Close()
	plainTextEdit := widgets.NewQPlainTextEditFromPointer(widget.FindChild("plainTextEdit", core.Qt__FindChildrenRecursively).Pointer())
	plainTextEdit.SetPlainText(docStr)
	return exportDialogWidget
}

func RegisterEditLayoutBtn(widget *widgets.QWidget, subwin *widgets.QDialog, collection *mongo.Collection) {
	cancelBtn := widgets.NewQPushButtonFromPointer(widget.FindChild("cancelBtn", core.Qt__FindChildrenRecursively).Pointer())
	cancelBtn.ConnectClicked(func(bool) {
		subwin.Close()
	})

	saveBtn := widgets.NewQPushButtonFromPointer(widget.FindChild("saveBtn", core.Qt__FindChildrenRecursively).Pointer())
	saveBtn.ConnectClicked(func(bool) {
		plainTextEdit := widgets.NewQPlainTextEditFromPointer(widget.FindChild("plainTextEdit", core.Qt__FindChildrenRecursively).Pointer())
		raw := plainTextEdit.ToPlainText()
		originalDoc := bson.M{}
		bson.UnmarshalExtJSON([]byte(raw), false, &originalDoc)

		objID := originalDoc["_id"].(primitive.ObjectID)
		objIDStr := objID.Hex()
		oid, _ := primitive.ObjectIDFromHex(objIDStr)
		filter := bson.D{
			{"_id", oid},
		}
		connectdb.Replace(collection, filter, originalDoc, nil)
		subwin.Close()
	})
}
