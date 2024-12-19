package gui

import (
	"strings"

	"github.com/abdessamad-zgor/lazyman/logger"
	"github.com/gdamore/tcell/v2"
)

type Text struct {
	// X, Y are relative to the containing box
	X        int
	Y        int
	Contents string
    Style *TextStyle
}

type TextStyle struct {
	Color      tcell.Color
	Background tcell.Color
	Bold       bool
	Italic     bool
}

//TODO: wrapping
// render a text object
func (t Text) Render(box *Box) RenderFn {
    return func(screen tcell.Screen) {
        var words []string
        splitContents := strings.Split(t.Contents, " ") 
        for iword, word := range splitContents {
            words = append(words, word)
            if iword != len(splitContents) - 1 {
                words = append(words, " ")
            }
        }
        logger.Info(words)
        line:=0
        lineRunesCount:=0
        for _ , word := range(words) {
            for _, wrune := range(word) {
                // we add one to exclude border positions 
                runeposx, runeposy := box.X + lineRunesCount + t.X+1 , box.Y + line + t.Y +1
                lineRunesCount += 1
                textStyle := tcell.StyleDefault
                if t.Style != nil {
                    textStyle = tcell.StyleDefault.
                        Foreground(t.Style.Color).
                        Background(t.Style.Background).
                        Bold(t.Style.Bold).
                        Italic(t.Style.Italic)
                }
                logger.Info(runeposx, runeposy, wrune)
                screen.SetContent(runeposx, runeposy, wrune, nil, textStyle)
            }
        }
    }
}
