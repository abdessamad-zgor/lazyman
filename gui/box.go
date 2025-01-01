package gui

import (
	_ "github.com/abdessamad-zgor/lazyman/logger"
	"github.com/gdamore/tcell/v2"
)

type Box struct {
	X        int
	Y        int
	W        int
	H        int
	Style    *BoxStyle
	Editable bool
	Float    bool
}

type BoxStyle struct {
	Border      *tcell.Style
	Default     *tcell.Style
}

func NewBox(x int, y int, w int, h int, style *BoxStyle) Box {
	return Box{
		X:        x,
		Y:        y,
		W:        w,
		H:        h,
		Style:    style,
	}
}

// a box should provide a method which provides for it's caller a way to draw the box
func (box Box) Render(screen tcell.Screen) {
	boxStyle := tcell.StyleDefault
	if box.Style.Border != nil {
		boxStyle = *box.Style.Border
	}
	// we consider box.X and box.Y as origin for the box
	for xi := range box.W {
		for yj := range box.H {
			borderX, borderY := xi+box.X, yj+box.Y
			if borderX == box.X || borderX == box.X+box.W-1 {
				screen.SetContent(borderX, borderY, tcell.RuneVLine, nil, boxStyle)
			}
			if borderY == box.Y || borderY == box.Y+box.H-1 {
				screen.SetContent(borderX, borderY, tcell.RuneHLine, nil, boxStyle)
			}
		}
	}
	screen.SetContent(box.X, box.Y, tcell.RuneULCorner, nil, boxStyle)
	screen.SetContent(box.X+box.W-1, box.Y, tcell.RuneURCorner, nil, boxStyle)
	screen.SetContent(box.X, box.Y+box.H-1, tcell.RuneLLCorner, nil, boxStyle)
	screen.SetContent(box.X+box.W-1, box.Y+box.H-1, tcell.RuneLRCorner, nil, boxStyle)
}

