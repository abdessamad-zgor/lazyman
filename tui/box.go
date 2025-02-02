package tui

import (
	"github.com/gdamore/tcell/v2"
)

type Box struct {
	X      int
	Y      int
	Width  int
	Height int
	Style  tcell.Style
}

func (box Box) Render(screen tcell.Screen) {
	boxStyle := tcell.StyleDefault
	for xi := range box.Width {
		for yi := range box.Height {
			borderX, borderY := xi+box.X, yi+box.Y
			if borderX == box.X || borderX == box.X+box.Width-1 {
				screen.SetContent(borderX, borderY, tcell.RuneVLine, nil, boxStyle)
			}
			if borderY == box.Y || borderY == box.Y+box.Height-1 {
				screen.SetContent(borderX, borderY, tcell.RuneHLine, nil, boxStyle)
			}
		}
	}

	screen.SetContent(box.X, box.Y, tcell.RuneULCorner, nil, boxStyle)
	screen.SetContent(box.X+box.Width-1, box.Y, tcell.RuneURCorner, nil, boxStyle)
	screen.SetContent(box.X, box.Y+box.Height-1, tcell.RuneLLCorner, nil, boxStyle)
	screen.SetContent(box.X+box.Width-1, box.Y+box.Height-1, tcell.RuneLRCorner, nil, boxStyle)
}
