package main

import (
	"os"

	"github.com/naponmeka/monger/src"
	"github.com/therecipe/qt/widgets"
)

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)
	src.NewWindow()
	// app.SetStyleSheet(``)
	app.Exec()
}
