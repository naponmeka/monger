package src

import (
	"strconv"

	"github.com/naponmeka/robone/list"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

func NewConnectLayout(tabsHolder *widgets.QTabWidget, keyboardBinder *KeyboardBinder) *widgets.QWidget {
	var widget = widgets.NewQWidget(nil, 0)
	var loader = uitools.NewQUiLoader(nil)
	var file = core.NewQFile2(":/qml/connect.ui")
	file.Open(core.QIODevice__ReadOnly)
	connectWidget := loader.Load(file, widget)
	file.Close()
	config := readConfig()
	// savedConnections := []list.ListItem{
	// 	list.ListItem{Name: "local", URI: "mongodb://root:root@localhost:27017"},
	// 	list.ListItem{Name: "production", URI: "mongodb://...."},
	// }
	listView := widgets.NewQListViewFromPointer(widget.FindChild("listView", core.Qt__FindChildrenRecursively).Pointer())
	model := list.NewCustomListModel(nil)
	listView.SetModel(model)
	for _, savedConnection := range config.SavedConnections {
		model.Add(savedConnection)
	}

	testBtn := widgets.NewQPushButtonFromPointer(widget.FindChild("testBtn", core.Qt__FindChildrenRecursively).Pointer())
	testBtn.ConnectClicked(func(bool) {
		widgets.QMessageBox_Information(nil, "OK", "Success", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	})
	nameLineEdit := widgets.NewQLineEditFromPointer(widget.FindChild("nameLineEdit", core.Qt__FindChildrenRecursively).Pointer())
	URILineEdit := widgets.NewQLineEditFromPointer(widget.FindChild("mylineedit", core.Qt__FindChildrenRecursively).Pointer())
	connectBtn := widgets.NewQPushButtonFromPointer(widget.FindChild("connectBtn", core.Qt__FindChildrenRecursively).Pointer())
	connectBtn.ConnectClicked(func(bool) {
		URI := URILineEdit.Text()
		mainQueryWidget := NewMainLayout(URI, keyboardBinder)
		ReplaceTabContent(tabsHolder, mainQueryWidget, "Query")
	})
	saveBtn := widgets.NewQPushButtonFromPointer(widget.FindChild("saveBtn", core.Qt__FindChildrenRecursively).Pointer())
	saveBtn.ConnectClicked(func(bool) {
		name := nameLineEdit.Text()
		URI := URILineEdit.Text()
		newItem := list.ListItem{Name: name, URI: URI}
		model.Add(newItem)
		config.SavedConnections = append(config.SavedConnections, newItem)
		saveConfig(config)
	})

	deleteBtn := widgets.NewQPushButtonFromPointer(widget.FindChild("deleteBtn", core.Qt__FindChildrenRecursively).Pointer())
	deleteBtn.ConnectClicked(func(bool) {
		idx := listView.CurrentIndex().Row()
		idxStr := strconv.Itoa(idx)
		model.RemoveAt(idxStr)
		config.SavedConnections = append(config.SavedConnections[:idx], config.SavedConnections[idx+1:]...)
		saveConfig(config)
	})
	createBtn := widgets.NewQPushButtonFromPointer(widget.FindChild("createBtn", core.Qt__FindChildrenRecursively).Pointer())
	createBtn.ConnectClicked(func(bool) {
		index := core.NewQModelIndex()
		listView.SetCurrentIndex(index)
		nameLineEdit.SetText("")
		URILineEdit.SetText("")
	})

	listView.ConnectClicked(func(index *core.QModelIndex) {
		idx := index.Row()
		item := config.SavedConnections[idx]
		nameLineEdit.SetText(item.Name)
		URILineEdit.SetText(item.URI)
	})

	return connectWidget
}
