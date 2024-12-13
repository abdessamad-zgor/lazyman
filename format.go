package main

import (
	"strings"
	"github.com/gdamore/tcell/v2"
)

type Text struct {
	X        int
	Y        int
	Contents string
    Style *TextStyle
}

type TextStyle struct {
	// these are relative to the containing box
	Color      tcell.Color
	Background tcell.Color
	Bold       bool
	Italic     bool
}

// render a text object
func (t Text) Render(box *Box) RenderFn {
    return func(screen tcell.Screen) {
        words := strings.Split(t.Contents, " ")
        line:=1
        for _ , word := range(words) {
            for runei, wrune := range(word) {
                runeposx, runposy := box.X + runei , box.Y + line 
                textStyle := tcell.StyleDefault
                if t.Style != nil {
                    textStyle = tcell.StyleDefault.
                        Foreground(t.Style.Color).
                        Background(t.Style.Background).
                        Bold(t.Style.Bold).
                        Italic(t.Style.Italic)
                }
                screen.SetContent(runeposx+1, runposy+1, rune(wrune), nil, textStyle)
            }
        }
    }
}
