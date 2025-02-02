package tui

import (
	"errors"

	"github.com/abdessamad-zgor/lazyman/event"
	lcontext "github.com/abdessamad-zgor/lazyman/context"
	"github.com/gdamore/tcell/v2"
)

type Widget interface {
    // Getters, Setters
    SetBody(layout Layout)
    GetBody() Layout
    GetMarker() rune
    GetPosition() *WidgetPosition
    GetStyle() *WidgetStyle
    GetTitle() *Text
    GetContent() *Text
    GetParent() Widget
    GetChildren() []Widget
    GetEventMap() event.EventMap

    SetStyle(style WidgetStyle)
    SetWidget(w Widget, col, row, colSpan, rowSpan int) error 
    SetRows(n int) error
    SetPosition(position WidgetPosition)
    SetColumns(n int) error
    SetTitle(title string)
    SetContent(contents string)
    SetParent(widget Widget)
    SetupEventMap() 

    GetCoordinates() (int, int)
    GetWidth() int
    GetHeight() int
}

type TUI interface {
    Render(screen tcell.Screen, context lcontext.Context)
}

type WidgetPosition struct {
    Parent  Layout
    Row     int
    Col     int
    SpanRow int
    SpanCol int
}

type WidgetStyle struct {
    Highlight   tcell.Style
    Select      tcell.Style
	Default     tcell.Style
    Border      tcell.Color
}

func NewWidgetStyle() WidgetStyle {
    selectStyle := tcell.StyleDefault.Foreground(tcell.ColorBlue)
    highlightStyle := tcell.StyleDefault.Foreground(tcell.ColorLime)
    defaultStyle :=tcell.StyleDefault 
    return WidgetStyle{
        Highlight: highlightStyle,
        Default: defaultStyle,
        Select: selectStyle,
    }
}

type WidgetSrc struct {
	Marker      rune

	Title       *Text
	Content     *Text

	Children    []Widget
	Parent      Widget

    EventMap    event.EventMap

    Style       WidgetStyle
    Position    WidgetPosition
    Body        Layout
}

func (widget WidgetSrc) Render(screen tcell.Screen, context lcontext.Context) {
}

func CreateWidget(marker rune) *WidgetSrc {
    widget := WidgetSrc{
        EventMap: make(event.EventMap),
        Marker: marker,
        Style: NewWidgetStyle(),
    }

    return &widget
}

func (widget *WidgetSrc) GetMarker() rune {
    return widget.Marker
}

func (widget *WidgetSrc) SetMarker(marker rune) {
    widget.Marker = marker
}

func (parent *WidgetSrc) SetWidget(child Widget, row, col, spanCol, spanRow int) error {
    child.SetParent(parent)
    for _, cw := range parent.Children {
        if cw.GetMarker() == child.GetMarker() {
            return errors.New("Duplicate marker: cannot add a widget with prexisting marker.")
        }
    }
    parent.Children = append(parent.Children, child)
    if err := parent.Body.SetWidget(child, row, col, spanCol, spanRow); err != nil {
        return err
    }
    return nil
}

func (widget *WidgetSrc) SetTitle(title string) {
    widget.Title = &Text{
        X: 3,
        Y: -1,
        Contents: title,
    }
}

func (widget *WidgetSrc) GetTitle() *Text {
    return widget.Title
}

func (widget *WidgetSrc) SetContent(content string) {
    widget.Content = &Text{
        Contents: content,
        X: 0,
        Y: 0,
    }
}

func (widget *WidgetSrc) GetContent() *Text {
    return widget.Content
}
 
func (widget *WidgetSrc) SetOverflowX(value bool) {
    widget.Body.SetOverflowX(value)
}

func (widget *WidgetSrc) IsOverflowX() bool {
    return widget.Body.IsOverflowX()
}

func (widget *WidgetSrc) SetOverflowY(value bool) {
    widget.Body.SetOverflowY(value)
}

func (widget *WidgetSrc) IsOverflowY() bool {
    return widget.Body.IsOverflowY()
}

func (widget *WidgetSrc) SetRows(n int) error {
    return widget.Body.SetRows(n)
}

func (widget *WidgetSrc) SetColumns(n int) error {
    return widget.Body.SetColumns(n)
}

func (widget *WidgetSrc) GetCoordinates() (int, int) {
    position := widget.Position
    rootContainer := position.Parent
    parentWidget := widget.Parent
    // error is ignored becaused an inserted Widget is guarented to have a valid range
    x, _ := rootContainer.GetColumnsWidth(0, position.Col)
    y, _ := rootContainer.GetRowsHeight(0, position.Row) 
    if parentWidget == nil {
        return  x, y
    }

    dx, dy := parentWidget.GetCoordinates()
    x += dx
    y += dy
    return x, y
}

func (widget *WidgetSrc) GetWidth() int {
    position := widget.Position
    container := position.Parent
    width, _ := container.GetColumnsWidth(position.Col, position.Col + position.SpanCol)
    return width
}

func (widget *WidgetSrc) GetHeight() int {
    position := widget.Position
    container := position.Parent
    height, _ := container.GetRowsHeight(position.Row, position.Row + position.SpanRow)
    return height
}

func (widget *WidgetSrc) SetBody(layout Layout) {
    widget.Body = layout
}

func (widget *WidgetSrc) GetBody() Layout {
    return widget.Body
}

func (widget *WidgetSrc) GetChildren() []Widget {
    return widget.Children
}

func (widget *WidgetSrc) GetEventMap() event.EventMap {
    return widget.EventMap
} 

func (widget *WidgetSrc) SetupEventMap() {
}

func (widget *WidgetSrc) SetParent(parent Widget) {
    widget.Parent = parent
}

func (widget *WidgetSrc) GetParent() Widget{
    return widget.Parent
}

func (widget *WidgetSrc) SetPosition(position WidgetPosition) {
    widget.Position = position
} 

func (widget *WidgetSrc) GetPosition() *WidgetPosition {
    return &widget.Position
}

func (widget *WidgetSrc) SetStyle(style WidgetStyle) {
    widget.Style = style
}

func (widget *WidgetSrc) GetStyle() *WidgetStyle {
    return &widget.Style 
} 

