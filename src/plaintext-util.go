package src

import (
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

func registerAutoCloseBracket(plainTextEdit *widgets.QPlainTextEdit, debounce *bool, textLength *int) {
	plainTextEdit.ConnectTextChanged(func() {
		text := plainTextEdit.ToPlainText()
		if *debounce && len(text) > *textLength {
			pos := plainTextEdit.TextCursor().Position()
			char := text[pos-1]
			if char == '"' && text[pos] != '"' {
				*debounce = false
				plainTextEdit.InsertPlainText(`"`)
				plainTextEdit.MoveCursor(gui.QTextCursor__Left, gui.QTextCursor__MoveAnchor)
			} else if char == '{' && text[pos] != '}' {
				*debounce = false
				plainTextEdit.InsertPlainText(`}`)
				plainTextEdit.MoveCursor(gui.QTextCursor__Left, gui.QTextCursor__MoveAnchor)
			} else if char == '[' && text[pos] != ']' {
				*debounce = false
				plainTextEdit.InsertPlainText(`]`)
				plainTextEdit.MoveCursor(gui.QTextCursor__Left, gui.QTextCursor__MoveAnchor)
			} else if char == '(' && text[pos] != ')' {
				*debounce = false
				plainTextEdit.InsertPlainText(`(`)
				plainTextEdit.MoveCursor(gui.QTextCursor__Left, gui.QTextCursor__MoveAnchor)
			}
		} else {
			*debounce = true
		}
		*textLength = len(text)
	})
}
