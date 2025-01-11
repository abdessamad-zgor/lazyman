package gui

import (
	"github.com/abdessamad-zgor/lazyman/state"
	"github.com/gdamore/tcell/v2"
)

type Widget struct {
	Box      Box
	EventMap state.EventMap
	Marker   rune
	Title    *Text
	Content  *Text
	Children []*Widget
	Parent   *Widget
}

func (widget Widget) Render(screen tcell.Screen, context state.Context) {
    widget.Box.Render(screen)
    if widget.Title != nil {
        widget.Title.Render(&widget.Box, screen)
    }
    if widget.Content != nil {
        widget.Content.Render(&widget.Box, screen)
    }
}


func CreateWidget(marker rune, x int, y int, w int, h int, boxStyle *BoxStyle) Widget {
    return Widget{
        Box: NewBox(x, y, w, h, boxStyle),
        EventMap: make(state.EventMap),
    }
}

func (pWidget *Widget) AppendChild(cWidget Widget) {
    cWidget.Parent = pWidget
    pWidget.Children = append(pWidget.Children, &cWidget)
}

func (widget *Widget) SetTitle(title string) {
    widget.Title = &Text{
        X: 3,
        Y: -1,
        Contents: title,
    }
}

func (widget *Widget) SetMarker(mark rune) {
    widget.Marker = mark
}

func (widget *Widget) SetContent(content string, x int, y int) {
    widget.Content = &Text{
        Contents: content,
        X: x,
        Y: y,
    }
}

func (widget *Widget) SetEventListner(event state.EventName, cb state.Callback) {
    widget.EventMap[event] = cb
}

func (widget Widget) GetEventListner(event state.EventName) (state.Callback, bool) {
    listner, ok := widget.EventMap[event]
    return listner, ok
}
