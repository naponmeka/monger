package src

import (
	"fmt"
	"strings"

	"github.com/naponmeka/bsonparser"
	"github.com/naponmeka/robone/connectdb"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"go.mongodb.org/mongo-driver/bson"
)

func getSelectedDoc(
	gs *GlobalState,
	resultTreeview *widgets.QTreeView,
	documents *[]bson.M,
) (string, int) {
	docStr := ""
	selected := 0
	if gs.resultTextEdit.IsHidden() {
		selected = findRow(resultTreeview, resultTreeview.CurrentIndex())
		if selected < 0 {
			selected = 0
		}
	} else {
		text := gs.resultTextEdit.ToPlainText()
		pos := gs.resultTextEdit.TextCursor().Position()
		beginOfDoc := "},\n\n{"
		starter := `"_id": `
		var objIDpos int
		beginOfDocIdx := strings.LastIndex(text[:pos], beginOfDoc)
		lastestIDIdx := strings.LastIndex(text[:pos], starter)
		widgets.QMessageBox_Critical(nil, "Error", fmt.Sprintln(beginOfDocIdx, ", ", lastestIDIdx), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
		if beginOfDocIdx > lastestIDIdx {
			objIDpos = strings.Index(text[pos:], starter) + pos
		} else {
			objIDpos = strings.LastIndex(text[:pos], starter)
		}
		if objIDpos == -1 {
			objIDpos = strings.Index(text, starter)
		}
		closer := strings.Index(text[objIDpos:], "\n")
		id := text[objIDpos+len(starter) : closer+objIDpos]
		id = strings.TrimRight(id, ",")
		id = strings.Replace(id, "ObjectId", "ObjectID", -1)
		selected = 0
		for idx, doc := range *documents {
			docID := fmt.Sprint(doc["_id"])
			if docID == id {
				selected = idx
				break
			}
		}
	}
	if selected < len(*documents) && selected >= 0 {
		docByte, _ := bson.MarshalExtJSON((*documents)[selected], false, true)
		docStr, _ = bsonparser.JsonToBson(string(docByte[:]))
	}
	return docStr, selected
}

func registerDocOperationBtn(
	mainWidget *widgets.QWidget,
	resultTreeview *widgets.QTreeView,
	documents *[]bson.M,
	mongoURI *string,
	currentDB *string,
	currentCollection *string,
	globalState *GlobalState,
) {
	editDocBtn := widgets.NewQPushButtonFromPointer(mainWidget.FindChild("editDocBtn", core.Qt__FindChildrenRecursively).Pointer())
	editDocBtn.ConnectClicked(func(bool) {
		docStr, selected := getSelectedDoc(globalState, resultTreeview, documents)
		docID := fmt.Sprint((*documents)[selected]["_id"])
		subwin := widgets.NewQDialog(nil, 0)
		subwin.SetWindowTitle(fmt.Sprintf("Edit: %s", docID))
		subwin.SetLayout(widgets.NewQHBoxLayout())
		editLayout := NewEditLayout(docStr)
		subwin.Layout().AddWidget(editLayout)
		dbCollection := connectdb.GetCollection(*mongoURI, *currentDB, *currentCollection)
		RegisterEditLayoutBtn(editLayout, subwin, dbCollection, globalState)
		subwin.SetModal(true)
		subwin.SetMinimumSize2(640, 480)
		subwin.Exec()
	})
	viewDocBtn := widgets.NewQPushButtonFromPointer(mainWidget.FindChild("viewDocBtn", core.Qt__FindChildrenRecursively).Pointer())
	viewDocBtn.ConnectClicked(func(bool) {
		docStr, selected := getSelectedDoc(globalState, resultTreeview, documents)
		docID := fmt.Sprint((*documents)[selected]["_id"])
		subwin := widgets.NewQDialog(nil, 0)
		subwin.SetWindowTitle(fmt.Sprintf("View: %s", docID))
		subwin.SetLayout(widgets.NewQHBoxLayout())
		viewLayout := NewViewLayout(docStr)
		subwin.Layout().AddWidget(viewLayout)
		RegisterViewLayoutBtn(viewLayout, subwin)
		subwin.SetModal(true)
		subwin.SetMinimumSize2(640, 480)
		subwin.Exec()
	})
	insertDocBtn := widgets.NewQPushButtonFromPointer(mainWidget.FindChild("insertDocBtn", core.Qt__FindChildrenRecursively).Pointer())
	insertDocBtn.ConnectClicked(func(bool) {
		subwin := widgets.NewQDialog(nil, 0)
		subwin.SetWindowTitle("Insert documents")
		subwin.SetLayout(widgets.NewQHBoxLayout())
		insertLayout := NewInsertLayout()
		subwin.Layout().AddWidget(insertLayout)
		dbCollection := connectdb.GetCollection(*mongoURI, *currentDB, *currentCollection)
		RegisterInsertLayoutBtn(insertLayout, subwin, dbCollection, globalState)
		subwin.SetModal(true)
		subwin.SetMinimumSize2(640, 480)
		subwin.Exec()
	})

	deleteDocBtn := widgets.NewQPushButtonFromPointer(mainWidget.FindChild("deleteDocBtn", core.Qt__FindChildrenRecursively).Pointer())
	deleteDocBtn.ConnectClicked(func(bool) {
		_, selected := getSelectedDoc(globalState, resultTreeview, documents)
		subwin := widgets.NewQDialog(nil, 0)
		docID := fmt.Sprint((*documents)[selected]["_id"])
		subwin.SetWindowTitle(fmt.Sprintf("Delete: %s", docID))
		subwin.SetLayout(widgets.NewQHBoxLayout())
		deleteConfirmLayout := NewConfirmLayout("Confirm delete?\n" + docID)
		subwin.Layout().AddWidget(deleteConfirmLayout)
		dbCollection := connectdb.GetCollection(*mongoURI, *currentDB, *currentCollection)
		RegisterConfirmDeleteLayoutBtn(
			deleteConfirmLayout,
			subwin,
			dbCollection,
			(*documents)[selected],
			globalState,
		)
		subwin.SetModal(true)
		subwin.SetMinimumSize2(100, 100)
		subwin.Exec()
	})
}
