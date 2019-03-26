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

func NewConfirmLayout(confirmText string) *widgets.QWidget {
	var widget = widgets.NewQWidget(nil, 0)
	var loader = uitools.NewQUiLoader(nil)
	var file = core.NewQFile2(":/qml/confirm.ui")
	file.Open(core.QIODevice__ReadOnly)
	dialogWidget := loader.Load(file, widget)
	file.Close()
	label := widgets.NewQLabelFromPointer(widget.FindChild("label", core.Qt__FindChildrenRecursively).Pointer())
	label.SetText(confirmText)
	return dialogWidget
}

func RegisterConfirmDeleteLayoutBtn(
	widget *widgets.QWidget,
	subwin *widgets.QDialog,
	collection *mongo.Collection,
	document bson.M,
	keyboardBinder *KeyboardBinder,
) {

	buttonBox := widgets.NewQDialogButtonBoxFromPointer(widget.FindChild("buttonBox", core.Qt__FindChildrenRecursively).Pointer())
	buttonBox.ConnectAccepted(func() {
		objID := document["_id"].(primitive.ObjectID)
		objIDStr := objID.Hex()
		oid, _ := primitive.ObjectIDFromHex(objIDStr)
		filter := bson.D{
			{"_id", oid},
		}
		connectdb.Remove(collection, filter, true)
		subwin.Close()
		keyboardBinder.ExecuteQuery()
	})
	buttonBox.ConnectRejected(func() {
		subwin.Close()
	})
}
