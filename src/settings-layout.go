package src

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

func NewSettingsLayout(globalState *GlobalState, subwin *widgets.QDialog) *widgets.QWidget {
	var widget = widgets.NewQWidget(nil, 0)
	var loader = uitools.NewQUiLoader(nil)
	var file = core.NewQFile2(":/qml/settings.ui")
	file.Open(core.QIODevice__ReadOnly)
	settingsDialogWidget := loader.Load(file, widget)
	file.Close()

	config := readConfig()
	timeoutSpinBox := widgets.NewQSpinBoxFromPointer(settingsDialogWidget.FindChild("timeoutSpinBox", core.Qt__FindChildrenRecursively).Pointer())
	timeoutSpinBox.SetValue(config.Timeout)

	defaultViewComboBox := widgets.NewQFontComboBoxFromPointer(settingsDialogWidget.FindChild("defaultViewComboBox", core.Qt__FindChildrenRecursively).Pointer())
	defaultViewComboBox.SetCurrentIndex(config.DefaultView)

	cancelBtn := widgets.NewQPushButtonFromPointer(widget.FindChild("cancelBtn", core.Qt__FindChildrenRecursively).Pointer())
	cancelBtn.ConnectClicked(func(bool) {
		subwin.Close()
	})

	saveBtn := widgets.NewQPushButtonFromPointer(widget.FindChild("saveBtn", core.Qt__FindChildrenRecursively).Pointer())
	saveBtn.ConnectClicked(func(bool) {
		config.Timeout = timeoutSpinBox.Value()
		config.DefaultView = defaultViewComboBox.CurrentIndex()
		globalState.defaultView = &config.DefaultView
		globalState.timeout = &config.Timeout
		saveConfig(config)
		subwin.Close()
	})
	return settingsDialogWidget
}
