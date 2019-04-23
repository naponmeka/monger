package src

import (
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

func registerAutoCloseBracket(plainTextEdit *widgets.QPlainTextEdit, debounce *bool, textLength *int) {
	plainTextEdit.ConnectTextChanged(func() {
		text := plainTextEdit.ToPlainText()
		if *debounce && len(text) > *textLength {
			textCursor := plainTextEdit.TextCursor()
			pos := textCursor.Position()
			var char uint8
			if len(text) > pos-1 && pos-1 > 0 {
				char = text[pos-1]
			}
			var nextChar uint8
			if len(text) > pos {
				nextChar = text[pos]
			}
			if char == '"' && nextChar != '"' {
				*debounce = false
				plainTextEdit.InsertPlainText(`"`)
				plainTextEdit.MoveCursor(gui.QTextCursor__Left, gui.QTextCursor__MoveAnchor)
			} else if char == '{' {
				*debounce = false
				plainTextEdit.InsertPlainText(`}`)
				plainTextEdit.MoveCursor(gui.QTextCursor__Left, gui.QTextCursor__MoveAnchor)
			} else if char == '[' && nextChar != ']' {
				*debounce = false
				plainTextEdit.InsertPlainText(`]`)
				plainTextEdit.MoveCursor(gui.QTextCursor__Left, gui.QTextCursor__MoveAnchor)
			} else if char == '(' && nextChar != ')' {
				*debounce = false
				plainTextEdit.InsertPlainText(`)`)
				plainTextEdit.MoveCursor(gui.QTextCursor__Left, gui.QTextCursor__MoveAnchor)
			} else if char == '"' && nextChar == '"' {
				*debounce = false
				plainTextEdit.SetPlainText(text[:pos] + text[pos+1:])
				for i := 0; i < pos; i++ {
					plainTextEdit.MoveCursor(gui.QTextCursor__Right, gui.QTextCursor__MoveAnchor)
				}
			} else if char == '}' && nextChar == '}' {
				*debounce = false
				plainTextEdit.SetPlainText(text[:pos] + text[pos+1:])
				for i := 0; i < pos; i++ {
					plainTextEdit.MoveCursor(gui.QTextCursor__Right, gui.QTextCursor__MoveAnchor)
				}
			} else if char == ']' && nextChar == ']' {
				*debounce = false
				plainTextEdit.SetPlainText(text[:pos] + text[pos+1:])
				for i := 0; i < pos; i++ {
					plainTextEdit.MoveCursor(gui.QTextCursor__Right, gui.QTextCursor__MoveAnchor)
				}
			} else if char == ')' && nextChar == ')' {
				*debounce = false
				plainTextEdit.SetPlainText(text[:pos] + text[pos+1:])
				for i := 0; i < pos; i++ {
					plainTextEdit.MoveCursor(gui.QTextCursor__Right, gui.QTextCursor__MoveAnchor)
				}
			}
		} else {
			*debounce = true
		}
		*textLength = len(text)
	})
}
