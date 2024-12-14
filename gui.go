package main

import (
	"github.com/abdessamad-zgor/lazyman/logger"
	"github.com/gdamore/tcell/v2"
)

type Callback func(context AppContext)

type Listener struct {
	Event    string
	Callback Callback
}

type Box struct {
	X         int
	Y         int
	W         int
	H         int
	Title     *Text
	Content   *Text
	Style     *BoxStyle
	Editable  bool
	Float     bool
	Listeners []Listener
	Children  []Box
}

type BoxStyle struct {
	Border      *tcell.Style
	OnHighlight *tcell.Style
	Default     *tcell.Style
}

func NewBox(x int, y int, w int, h int, title *Text, content *Text, style *BoxStyle, editable bool, float bool, listeners []Listener, children []Box) *Box {
	return &Box{
		X:         x,
		Y:         y,
		W:         w,
		H:         h,
		Title:     title,
		Content:   content,
		Style:     style,
		Editable:  editable,
		Float:     float,
		Listeners: listeners,
		Children:  children,
	}
}

type RenderFn func(screen tcell.Screen)

// a box should provide a method which provides for it's caller a way to draw the box
func (box *Box) GetDrawF() RenderFn {
	return func(screen tcell.Screen) {
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

		// TODO: a title should be less than width or it'll be truncated
		if box.Title != nil {
			logger.Info(box.Title)
			box.Title.Render(box)(screen)
		}

		// TODO:scroll
		if box.Content != nil {
			logger.Info(box.Content)
			box.Content.Render(box)(screen)
		}
	}
}
