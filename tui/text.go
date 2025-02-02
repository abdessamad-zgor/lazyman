package tui

import (
	"github.com/abdessamad-zgor/lazyman/logger"
	"github.com/gdamore/tcell/v2"
	"strings"
)

type Text struct {
	X        int
	Y        int
	Contents string
	Style    *TextStyle
}

type TextStyle struct {
	Default   tcell.Style
	Highlight tcell.Style
	Select    tcell.Style
}

// TODO: wrapping
// render a text object
func (t *Text) Render(box Box, screen tcell.Screen) {
	if t != nil {
		var words []string
		splitContents := strings.Split(t.Contents, " ")
		for iword, word := range splitContents {
			words = append(words, word)
			if iword != len(splitContents)-1 {
				words = append(words, " ")
			}
		}
		logger.Info(words)
		line := 0
		lineRunesCount := 0
		x, y := box.X, box.Y
		for _, word := range words {
			for _, wrune := range word {
				// we add one to exclude border positions
				runeposx, runeposy := x+lineRunesCount+t.X+1, y+line+t.Y+1
				lineRunesCount += 1
				textStyle := tcell.StyleDefault
				if t.Style != nil {
					textStyle = t.Style.Default
				}
				screen.SetContent(runeposx, runeposy, wrune, nil, textStyle)
			}
		}
	}
}
