package src

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/naponmeka/robone/connectdb"
	"github.com/naponmeka/robone/table"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
	"go.mongodb.org/mongo-driver/bson"
	gojsonq "gopkg.in/thedevsaddam/gojsonq.v2"
)

func NewExportLayout(globalState *GlobalState) *widgets.QWidget {
	var widget = widgets.NewQWidget(nil, 0)
	var loader = uitools.NewQUiLoader(nil)
	var file = core.NewQFile2(":/qml/export.ui")
	file.Open(core.QIODevice__ReadOnly)
	exportDialogWidget := loader.Load(file, widget)
	file.Close()

	detailLabel := widgets.NewQLabelFromPointer(widget.FindChild("detailLabel", core.Qt__FindChildrenRecursively).Pointer())
	detailLabel.SetText(fmt.Sprintf("DB:%s Collection: %s Query:%s", *globalState.currentDB, *globalState.currentCollection, *globalState.currentQuery))

	bsonFileLineEdit := widgets.NewQLineEditFromPointer(widget.FindChild("bsonFileLineEdit", core.Qt__FindChildrenRecursively).Pointer())
	csvLineEdit := widgets.NewQLineEditFromPointer(widget.FindChild("csvLineEdit", core.Qt__FindChildrenRecursively).Pointer())

	fileDialog := widgets.NewQFileDialog(nil, 0)
	fileDialog.SetFileMode(widgets.QFileDialog__DirectoryOnly)
	fileDialog.ConnectAccepted(func() {
		paths := fileDialog.SelectedFiles()
		bsonFileLineEdit.SetText(paths[0])
		csvLineEdit.SetText(paths[0])
	})
	bsonBrowseBtn := widgets.NewQPushButtonFromPointer(widget.FindChild("bsonBrowseBtn", core.Qt__FindChildrenRecursively).Pointer())
	bsonBrowseBtn.ConnectClicked(func(bool) {
		fileDialog.Show()
	})
	csvBrowseBtn := widgets.NewQPushButtonFromPointer(widget.FindChild("csvBrowseBtn", core.Qt__FindChildrenRecursively).Pointer())
	csvBrowseBtn.ConnectClicked(func(bool) {
		fileDialog.Show()
	})

	tableView := widgets.NewQTableViewFromPointer(widget.FindChild("tableView", core.Qt__FindChildrenRecursively).Pointer())
	model := table.NewCustomTableModel(nil)
	tableView.SetModel(model)
	tableData := []table.TableItem{}

	// edit
	tableView.ConnectDoubleClicked(func(index *core.QModelIndex) {
		idx := index.Row()
		subwin := widgets.NewQDialog(nil, 0)
		subwin.SetWindowTitle("Add config")
		subwin.SetLayout(widgets.NewQHBoxLayout())
		editTableLayout := NewEditTableLayout()
		csvFieldLineEdit := widgets.NewQLineEditFromPointer(editTableLayout.FindChild("csvFieldLineEdit", core.Qt__FindChildrenRecursively).Pointer())
		csvFieldLineEdit.SetText(tableData[idx].CSVField)
		fromFieldLineEdit := widgets.NewQLineEditFromPointer(editTableLayout.FindChild("fromFieldLineEdit", core.Qt__FindChildrenRecursively).Pointer())
		fromFieldLineEdit.SetText(tableData[idx].FromField)
		cancelBtn := widgets.NewQPushButtonFromPointer(editTableLayout.FindChild("cancelBtn", core.Qt__FindChildrenRecursively).Pointer())
		cancelBtn.ConnectClicked(func(bool) {
			subwin.Close()
		})
		saveBtn := widgets.NewQPushButtonFromPointer(editTableLayout.FindChild("saveBtn", core.Qt__FindChildrenRecursively).Pointer())
		saveBtn.ConnectClicked(func(bool) {
			model.EditAt(idx, csvFieldLineEdit.Text(), fromFieldLineEdit.Text())
			tableData[idx].CSVField = csvFieldLineEdit.Text()
			tableData[idx].FromField = fromFieldLineEdit.Text()
			subwin.Close()
		})
		subwin.Layout().AddWidget(editTableLayout)
		subwin.SetModal(true)
		subwin.Exec()
	})

	addBtn := widgets.NewQPushButtonFromPointer(widget.FindChild("addBtn", core.Qt__FindChildrenRecursively).Pointer())
	addBtn.ConnectClicked(func(bool) {
		subwin := widgets.NewQDialog(nil, 0)
		subwin.SetWindowTitle("Add config")
		subwin.SetLayout(widgets.NewQHBoxLayout())
		editTableLayout := NewEditTableLayout()
		cancelBtn := widgets.NewQPushButtonFromPointer(editTableLayout.FindChild("cancelBtn", core.Qt__FindChildrenRecursively).Pointer())
		cancelBtn.ConnectClicked(func(bool) {
			subwin.Close()
		})
		saveBtn := widgets.NewQPushButtonFromPointer(editTableLayout.FindChild("saveBtn", core.Qt__FindChildrenRecursively).Pointer())
		saveBtn.ConnectClicked(func(bool) {
			csvFieldLineEdit := widgets.NewQLineEditFromPointer(editTableLayout.FindChild("csvFieldLineEdit", core.Qt__FindChildrenRecursively).Pointer())
			fromFieldLineEdit := widgets.NewQLineEditFromPointer(editTableLayout.FindChild("fromFieldLineEdit", core.Qt__FindChildrenRecursively).Pointer())
			newItem := table.TableItem{
				CSVField:  csvFieldLineEdit.Text(),
				FromField: fromFieldLineEdit.Text(),
			}
			model.Add(newItem)
			tableData = append(tableData, newItem)
			subwin.Close()
		})
		subwin.Layout().AddWidget(editTableLayout)
		subwin.SetModal(true)
		subwin.Exec()
	})
	removeBtn := widgets.NewQPushButtonFromPointer(widget.FindChild("removeBtn", core.Qt__FindChildrenRecursively).Pointer())
	removeBtn.ConnectClicked(func(bool) {
		idx := tableView.CurrentIndex().Row()
		if idx >= 0 {
			tableData = append(tableData[:idx], tableData[idx+1:]...)
			model.RemoveAt(idx)
		}
	})

	csvFilenameLineEdit := widgets.NewQLineEditFromPointer(widget.FindChild("csvFilenameLineEdit", core.Qt__FindChildrenRecursively).Pointer())
	csvFilenameLineEdit.SetText("export-result")

	csvExportBtn := widgets.NewQPushButtonFromPointer(widget.FindChild("csvExportBtn", core.Qt__FindChildrenRecursively).Pointer())
	csvExportBtn.ConnectClicked(func(bool) {
		documents, err := connectdb.Query(
			*globalState.mongoURI,
			*globalState.currentDB,
			*globalState.currentCollection,
			*globalState.currentQuery,
			-1,
			-1,
			*globalState.timeout,
		)
		if err != nil {
			widgets.QMessageBox_Information(exportDialogWidget, "Error", "Cannot Query", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			return
		}
		countResult, err := connectdb.Query(
			*globalState.mongoURI,
			*globalState.currentDB,
			*globalState.currentCollection,
			*globalState.currentQuery+".count()",
			-1,
			-1,
			*globalState.timeout,
		)
		total := countResult[0]["count"].(int64)
		totalDocuments := int(total)
		if err != nil {
			widgets.QMessageBox_Information(exportDialogWidget, "Error", "Cannot Count", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			return
		}

		csvData := [][]string{}
		firstLine := []string{}
		for _, colSetting := range tableData {
			firstLine = append(firstLine, colSetting.CSVField)
		}
		csvData = append(csvData, firstLine)
		for _, doc := range documents {
			line := []string{}
			for _, colSetting := range tableData {
				jsonByte, _ := bson.MarshalExtJSON(doc, false, true)
				rawData := gojsonq.New().JSONString(string(jsonByte[:])).Find(colSetting.FromField)
				line = append(line, fmt.Sprint(rawData))
			}
			csvData = append(csvData, line)
		}
		path := fmt.Sprintf("%s/%s.csv", csvLineEdit.Text(), csvFilenameLineEdit.Text())
		file, err := os.Create(path)
		if err != nil {
			widgets.QMessageBox_Information(exportDialogWidget, "Error", "Cannot create file", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			return
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		p := widgets.NewQProgressDialog(nil, core.Qt__Dialog)
		p.SetModal(true)
		p.SetLabelText(fmt.Sprintf("Exporting (%d documents)", totalDocuments))
		p.SetMaximum(totalDocuments)
		for i, value := range csvData {
			err := writer.Write(value)
			p.SetValue(i)
			if err != nil {
				widgets.QMessageBox_Information(exportDialogWidget, "Error", "Cannot write file", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
				return
			}
		}
		p.SetValue(totalDocuments)
		p.Close()
		widgets.QMessageBox_Information(exportDialogWidget, "Status", "Exported", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	})

	return exportDialogWidget
}

func RegisterExportLayoutBtn(widget *widgets.QWidget, subwin *widgets.QDialog) {
	csvCancelBtn := widgets.NewQPushButtonFromPointer(widget.FindChild("csvCancelBtn", core.Qt__FindChildrenRecursively).Pointer())
	csvCancelBtn.ConnectClicked(func(bool) {
		subwin.Close()
	})
}
