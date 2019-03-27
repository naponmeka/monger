package src

import (
	"github.com/naponmeka/robone/table"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

func NewExportLayout() *widgets.QWidget {
	var widget = widgets.NewQWidget(nil, 0)
	var loader = uitools.NewQUiLoader(nil)
	var file = core.NewQFile2(":/qml/export.ui")
	file.Open(core.QIODevice__ReadOnly)
	exportDialogWidget := loader.Load(file, widget)
	file.Close()

	tableView := widgets.NewQTableViewFromPointer(widget.FindChild("tableView", core.Qt__FindChildrenRecursively).Pointer())
	model := table.NewCustomTableModel(nil)
	tableView.SetModel(model)
	tableData := []table.TableItem{}

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
		// subwin.SetMinimumSize2(100, 80)
		subwin.Exec()
	})

	// csvExportBtn := widgets.NewQPushButtonFromPointer(widget.FindChild("csvExportBtn", core.Qt__FindChildrenRecursively).Pointer())
	// csvExportBtn.ConnectClicked(func(bool) {
		// go func(widget *widgets.QWidget) {
		// 	time.Sleep(5 * time.Second)
		// 	csvExportBtn.SetText("Done")
		// 	// widgets.QMessageBox_Information(widget, "OK", "Export done", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
		// }(widget)

	// })

	return exportDialogWidget
}

func RegisterExportLayoutBtn(widget *widgets.QWidget, subwin *widgets.QDialog) {
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
	csvCancelBtn := widgets.NewQPushButtonFromPointer(widget.FindChild("csvCancelBtn", core.Qt__FindChildrenRecursively).Pointer())
	csvCancelBtn.ConnectClicked(func(bool) {
		subwin.Close()
	})
}
