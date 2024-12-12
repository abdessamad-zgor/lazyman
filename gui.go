package main

import (
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
	Border       *tcell.Style
	OnHighlight  *tcell.Style
	Default      *tcell.Style
	TitleStyle   *TextStyle
	ContentStyle *TextStyle
}

type TextStyle struct {
	// these are relative to the containing box
	Color      tcell.Color
	Background tcell.Color
	Bold       bool
	Italic     bool
}

type Text struct {
	X        int
	Y        int
	Contents string
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
				if xi == box.X {
					screen.SetContent(xi, yj, tcell.RuneVLine, nil, boxStyle)
				}
				if yj == box.Y {
					screen.SetContent(xi, yj, tcell.RuneHLine, nil, boxStyle)
				}
				if xi == box.X+box.W-1 {
					screen.SetContent(xi, yj, tcell.RuneVLine, nil, boxStyle)
				}
				if yj == box.Y+box.H-1 {
					screen.SetContent(xi, yj, tcell.RuneHLine, nil, boxStyle)
				}
			}
		}
		screen.SetContent(box.X, box.Y, tcell.RuneULCorner, nil, boxStyle)
		screen.SetContent(box.X+box.W-1, box.Y, tcell.RuneURCorner, nil, boxStyle)
		screen.SetContent(box.X, box.Y+box.H-1, tcell.RuneLLCorner, nil, boxStyle)
		screen.SetContent(box.X+box.W-1, box.Y+box.H-1, tcell.RuneLRCorner, nil, boxStyle)

		// TODO: a title should be less than width or it'll be truncated
		if box.Title != nil {
			titleStyle := tcell.StyleDefault
			if box.Style.TitleStyle != nil {
				titleStyle = tcell.StyleDefault.
					Foreground(box.Style.TitleStyle.Color).
					Background(box.Style.TitleStyle.Background).
					Bold(box.Style.TitleStyle.Bold).
					Italic(box.Style.TitleStyle.Italic)
			}
			screen.SetContent(box.X+1+box.Title.X, box.Y+1+box.Title.Y, rune(box.Title.Contents[0]), []rune(box.Title.Contents[1:]), titleStyle)
		}

		// TODO: apply wraping and scroll
		if box.Content != nil {
			titleStyle := tcell.StyleDefault
			if box.Style.ContentStyle != nil {
				titleStyle = tcell.StyleDefault.
					Foreground(box.Style.ContentStyle.Color).
					Background(box.Style.ContentStyle.Background).
					Bold(box.Style.ContentStyle.Bold).
					Italic(box.Style.ContentStyle.Italic)
			}
			screen.SetContent(box.X+1+box.Content.X, box.Y+1+box.Content.Y, rune(box.Content.Contents[0]), []rune(box.Content.Contents[1:]), titleStyle)
		}
	}
}
