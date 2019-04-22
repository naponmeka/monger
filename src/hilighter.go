package src

import (
	"strings"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
)

type Highlighter struct {
	*gui.QSyntaxHighlighter
	allWords                []string
	highlightingRules       []HighlightingRule
	regexpQuotesAndComment  *core.QRegExp
	keywordFormat           *gui.QTextCharFormat
	identFormat             *gui.QTextCharFormat
	functionFormat          *gui.QTextCharFormat
	numberFormat            *gui.QTextCharFormat
	quotesFormat            *gui.QTextCharFormat
	singleLineCommentFormat *gui.QTextCharFormat
	multiLineCommentFormat  *gui.QTextCharFormat
}

type HighlightingRule struct {
	pattern *core.QRegExp
	format  *gui.QTextCharFormat
}

const (
	STATE_BACKQUOTES         = 0x04
	STATE_SINGLELINE_COMMENT = 0x08
	STATE_MULTILINE_COMMENT  = 0x10
)

func NewHighlighter(document *gui.QTextDocument) *Highlighter {

	hl := &Highlighter{
		QSyntaxHighlighter:      gui.NewQSyntaxHighlighter2(document),
		keywordFormat:           gui.NewQTextCharFormat(),
		identFormat:             gui.NewQTextCharFormat(),
		functionFormat:          gui.NewQTextCharFormat(),
		numberFormat:            gui.NewQTextCharFormat(),
		quotesFormat:            gui.NewQTextCharFormat(),
		singleLineCommentFormat: gui.NewQTextCharFormat(),
		multiLineCommentFormat:  gui.NewQTextCharFormat(),
	}

	hl.keywordFormat.SetForeground(gui.NewQBrush4(core.Qt__darkBlue, core.Qt__SolidPattern))
	hl.keywordFormat.SetFontWeight(75) /// core.QFont__Bold
	hl.identFormat.SetForeground(gui.NewQBrush4(core.Qt__darkBlue, core.Qt__SolidPattern))
	hl.functionFormat.SetForeground(gui.NewQBrush4(core.Qt__blue, core.Qt__SolidPattern))
	hl.numberFormat.SetForeground(gui.NewQBrush4(core.Qt__darkMagenta, core.Qt__SolidPattern))
	hl.quotesFormat.SetForeground(gui.NewQBrush4(core.Qt__darkGreen, core.Qt__SolidPattern))
	hl.singleLineCommentFormat.SetForeground(gui.NewQBrush4(core.Qt__darkCyan, core.Qt__SolidPattern))
	hl.multiLineCommentFormat.SetForeground(gui.NewQBrush4(core.Qt__darkCyan, core.Qt__SolidPattern))

	var words []string
	rule := HighlightingRule{}
	highlightingRules := []HighlightingRule{}

	//number
	rule.pattern = core.NewQRegExp2(`(\b|\.)([0-9]+|0[xX][0-9a-fA-F]+|0[0-7]+)(\.[0-9]+)?([eE][+-]?[0-9]+i?)?\b`, core.Qt__CaseSensitive, core.QRegExp__RegExp)
	rule.format = hl.numberFormat
	highlightingRules = append(highlightingRules, rule)

	//function
	rule.pattern = core.NewQRegExp2(`\b[a-zA-Z_][a-zA-Z0-9_]+\s*(?=\()`, core.Qt__CaseSensitive, core.QRegExp__RegExp)
	rule.format = hl.functionFormat
	highlightingRules = append(highlightingRules, rule)

	//indent
	indent := `bool|byte|complex64|complex128|float32|float64|int8|int16|int32|int64|string|uint8|uint16|uint32|uint64|` +
		`int|uint|uintptr|true|false|iota|nil|append|cap|close|closed|complex|copy|imag|len|make|new|panic|print|println|` +
		`real|recover`
	rule.pattern = core.NewQRegExp2(`\b(`+indent+`)\b`, core.Qt__CaseSensitive, core.QRegExp__RegExp)
	rule.format = hl.identFormat
	highlightingRules = append(highlightingRules, rule)

	words = strings.Split(indent, `|`)
	hl.allWords = append(hl.allWords, words...)

	//keyword
	keyword := `break|default|func|interface|select|case|defer|go|map|struct|chan|else|goto|package|switch|` +
		`const|fallthrough|if|range|type|continue|for|import|return|var`
	rule.pattern = core.NewQRegExp2(`\b(`+keyword+`)\b`, core.Qt__CaseSensitive, core.QRegExp__RegExp)
	rule.format = hl.keywordFormat
	highlightingRules = append(highlightingRules, rule)

	words = strings.Split(keyword, `|`)
	hl.allWords = append(hl.allWords, words...)

	hl.highlightingRules = highlightingRules

	//quotes and comment
	hl.regexpQuotesAndComment = core.NewQRegExp2("//|\\\"|'|`|/\\*", core.Qt__CaseSensitive, core.QRegExp__RegExp)

	hl.ConnectHighlightBlock(hl.highlightBlock)

	return hl
}

func (hl *Highlighter) highlightBlock(stext string) {

	text := core.NewQStringRef3(stext)
	startPos := 0
	endPos := text.Length()
	hl.SetCurrentBlockState(0)

	startPos, endPos, cont := hl.highlightPreBlock(text, startPos, endPos)

	if cont {
		return
	}

	//keyword and func
	for _, rule := range hl.highlightingRules {
		expression := core.NewQRegExp3(rule.pattern)
		index := expression.IndexIn(text.String(), startPos, core.QRegExp__CaretAtZero)
		for index >= 0 {
			length := expression.MatchedLength()
			hl.SetFormat(index, length, rule.format)
			hl.allWords = append(hl.allWords, text.Mid(index, length).String())
			index = expression.IndexIn(text.String(), startPos+index+length, core.QRegExp__CaretAtZero)
		}
	}

	//quote and comment
	for true {
		startPos = hl.regexpQuotesAndComment.IndexIn(text.String(), startPos, core.QRegExp__CaretAtZero)

		if startPos == -1 {
			break
		}

		cap := hl.regexpQuotesAndComment.Cap(0)
		if (cap == "\"") || (cap == "'") || (cap == "`") {
			endPos = hl.findQuotesEndPos(text, startPos+1, cap)

			if endPos == -1 {
				//multiline
				hl.SetFormat(startPos, text.Length()-startPos, hl.quotesFormat)
				if cap == "`" {
					hl.SetCurrentBlockState(STATE_BACKQUOTES)
				}
				return
			} else {
				endPos += 1
				hl.SetFormat(startPos, endPos-startPos, hl.quotesFormat)
				startPos = endPos
			}
		} else if cap == "//" {
			hl.SetFormat(startPos, text.Length()-startPos, hl.singleLineCommentFormat)
			if text.EndsWith("\\", core.Qt__CaseSensitive) {
				hl.SetCurrentBlockState(STATE_SINGLELINE_COMMENT)
			}
			return
		} else if cap == "/*" {
			endPos = text.IndexOf("*/", startPos+2, core.Qt__CaseSensitive)
			if endPos == -1 {
				//multiline
				hl.SetFormat(startPos, text.Length()-startPos, hl.multiLineCommentFormat)
				hl.SetCurrentBlockState(STATE_MULTILINE_COMMENT)
				return
			} else {
				endPos += 2
				hl.SetFormat(startPos, endPos-startPos, hl.multiLineCommentFormat)
				startPos = endPos
			}
		}
	}
}

func (hl *Highlighter) highlightPreBlock(text *core.QStringRef, startPos int, endPos int) (int, int, bool) {

	state := hl.PreviousBlockState()

	if state == -1 {
		state = 0
	}

	if state == STATE_BACKQUOTES {
		endPos = hl.findQuotesEndPos(text, startPos, "`")

		if endPos == -1 {
			hl.SetFormat(0, text.Length(), hl.quotesFormat)
			hl.SetCurrentBlockState(STATE_BACKQUOTES)
			return startPos, endPos, true
		} else {
			endPos += 1
			hl.SetFormat(0, endPos-startPos, hl.quotesFormat)
			startPos = endPos
		}
	} else if state == STATE_MULTILINE_COMMENT {
		endPos = text.IndexOf("*/", startPos, core.Qt__CaseSensitive)
		if endPos == -1 {
			hl.SetFormat(0, text.Length(), hl.multiLineCommentFormat)
			hl.SetCurrentBlockState(hl.PreviousBlockState())
			return startPos, endPos, true
		} else {
			endPos += 2
			hl.SetFormat(0, endPos-startPos, hl.multiLineCommentFormat)
			startPos = endPos
		}
	} else if state == STATE_SINGLELINE_COMMENT {
		hl.SetFormat(0, text.Length(), hl.singleLineCommentFormat)
		if text.EndsWith("\\", core.Qt__CaseSensitive) {
			hl.SetCurrentBlockState(STATE_SINGLELINE_COMMENT)
		}
		return startPos, endPos, true
	}
	return startPos, endPos, false
}

func (hl *Highlighter) findQuotesEndPos(text *core.QStringRef, startPos int, endChar string) int {

	stext := text.String()

	for pos := startPos; pos < len(stext); pos++ {
		if string(stext[pos]) == endChar {
			return pos
		} else if string(stext[pos]) == `\` && endChar != "`" { /// ?
			pos++
		}
	}
	return -1
}
