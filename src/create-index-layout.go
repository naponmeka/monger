package src

import (
	"github.com/naponmeka/robone/connectdb"
	"github.com/naponmeka/robone/tree"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
	"go.mongodb.org/mongo-driver/bson"
)

func NewCreateIndexLayout(
	subwin *widgets.QDialog,
	model *tree.CustomTreeModel,
	globalState *GlobalState,
	isEdit bool,
) *widgets.QWidget {
	var widget = widgets.NewQWidget(nil, 0)
	var loader = uitools.NewQUiLoader(nil)
	var file = core.NewQFile2(":/qml/create-index.ui")
	file.Open(core.QIODevice__ReadOnly)
	layoutWidget := loader.Load(file, widget)
	file.Close()
	cancelBtn := widgets.NewQPushButtonFromPointer(layoutWidget.FindChild("cancelBtn", core.Qt__FindChildrenRecursively).Pointer())
	cancelBtn.ConnectClicked(func(bool) {
		subwin.Close()
	})
	saveBtn := widgets.NewQPushButtonFromPointer(layoutWidget.FindChild("saveBtn", core.Qt__FindChildrenRecursively).Pointer())
	saveBtn.ConnectClicked(func(bool) {
		nameLineEdit := widgets.NewQLineEditFromPointer(layoutWidget.FindChild("nameLineEdit", core.Qt__FindChildrenRecursively).Pointer())
		indexName := nameLineEdit.Text()

		keysPlainTextEdit := widgets.NewQPlainTextEditFromPointer(layoutWidget.FindChild("keysPlainTextEdit", core.Qt__FindChildrenRecursively).Pointer())
		keysText := keysPlainTextEdit.ToPlainText()

		uniqueCheckBox := widgets.NewQCheckBoxFromPointer(layoutWidget.FindChild("uniqueCheckBox", core.Qt__FindChildrenRecursively).Pointer())
		isUnique := uniqueCheckBox.IsChecked()

		sparseCheckBox := widgets.NewQCheckBoxFromPointer(layoutWidget.FindChild("sparseCheckBox", core.Qt__FindChildrenRecursively).Pointer())
		isSparse := sparseCheckBox.IsChecked()

		backgroundCheckBox := widgets.NewQCheckBoxFromPointer(layoutWidget.FindChild("backgroundCheckBox", core.Qt__FindChildrenRecursively).Pointer())
		isBackground := backgroundCheckBox.IsChecked()

		expireAfterCheckBox := widgets.NewQCheckBoxFromPointer(layoutWidget.FindChild("expireAfterCheckBox", core.Qt__FindChildrenRecursively).Pointer())
		isExpireAfter := expireAfterCheckBox.IsChecked()

		spinBox := widgets.NewQSpinBoxFromPointer(layoutWidget.FindChild("spinBox", core.Qt__FindChildrenRecursively).Pointer())
		expireAfter := spinBox.Value()
		if isEdit {
			err := connectdb.DropIndex(*globalState.mongoURI, *globalState.currentDB, *globalState.currentCollection, indexName)
			if err != nil {
				widgets.QMessageBox_Critical(nil, "Error", "Error:\n"+err.Error(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
				return
			}
		}

		keys := bson.M{}
		err := bson.UnmarshalExtJSON([]byte(keysText), true, &keys)
		if err != nil {
			widgets.QMessageBox_Critical(nil, "Error", "Error parsing:\n"+err.Error(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			return
		}

		if !isExpireAfter {
			expireAfter = -1
		}

		err = connectdb.CreateIndex(*globalState.mongoURI,
			*globalState.currentDB,
			*globalState.currentCollection,
			indexName,
			keys,
			isSparse,
			isUnique,
			isBackground,
			expireAfter,
		)
		if err != nil {
			widgets.QMessageBox_Critical(nil, "Error", "Error:\n"+err.Error(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			return
		}
		RefreshIndexTree(model, globalState)
		subwin.Close()
	})
	return layoutWidget
}
