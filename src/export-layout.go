package src

import (
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
	return exportDialogWidget
}
