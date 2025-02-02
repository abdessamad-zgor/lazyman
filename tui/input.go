package tui

import (
	"errors"

	"github.com/abdessamad-zgor/lazyman/event"
	lcontext "github.com/abdessamad-zgor/lazyman/context"
	"github.com/gdamore/tcell/v2"
)

type Input struct {
    Widget      *WidgetSrc
    Multiline   bool
}

func CreateInput(marker rune) *Input {
    input := Input{
        Widget: CreateWidget(marker),
    }

    return &input
}

func (input *Input) GetMarker() rune {
    return input.Widget.Marker
}

func (input *Input) SetMarker(marker rune) {
    input.Widget.Marker = marker
}

func (parent *Input) SetWidget(child Widget, row, col, spanCol, spanRow int) error {
    return errors.New("Input element cannot have child elements.")
}

func (input *Input) SetTitle(title string) {
    input.Widget.Title = &Text{
        X: 3,
        Y: -1,
        Contents: title,
    }
}

func (input *Input) GetTitle() *Text {
    return input.Widget.Title
}

func (input *Input) SetContent(content string) {
    input.Widget.Content = &Text{
        Contents: content,
        X: 0,
        Y: 0,
    }
}

func (input *Input) GetContent() *Text {
    return input.Widget.Content
}
 
func (input *Input) SetOverflowX(value bool) {
    input.Widget.Body.SetOverflowX(value)
}

func (input *Input) IsOverflowX() bool {
    return input.Widget.Body.IsOverflowX()
}

func (input *Input) SetOverflowY(value bool) {
    input.Widget.Body.SetOverflowY(value)
}

func (input *Input) IsOverflowY() bool {
    return input.Widget.Body.IsOverflowY()
}

func (input *Input) SetRows(n int) error {
    return input.Widget.Body.SetRows(n)
}

func (input *Input) SetColumns(n int) error {
    return input.Widget.Body.SetColumns(n)
}

func (input *Input) GetCoordinates() (int, int) {
    position := input.Widget.Position
    rootContainer := position.Parent
    parentWidget := input.Widget.Parent
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

func (input *Input) GetWidth() int {
    return input.Widget.GetWidth()
}

func (input *Input) GetHeight() int {
    return input.Widget.GetHeight()
}

func (input *Input) SetBody(layout Layout) {
    input.Widget.Body = layout
}

func (input *Input) GetBody() Layout {
    return input.Widget.Body
}

func (input *Input) GetChildren() []Widget {
    return input.Widget.Children
}

func (input *Input) GetEventMap() event.EventMap {
    return input.Widget.EventMap
} 

func (input *Input) SetupEventMap() {
}

func (input *Input) SetParent(parent Widget) {
    input.Widget.Parent = parent
}

func (input *Input) GetParent() Widget{
    return input.Widget.Parent
}

func (input *Input) SetPosition(position WidgetPosition) {
    input.Widget.Position = position
} 

func (input *Input) GetPosition() *WidgetPosition {
    return &input.Widget.Position
}

func (input *Input) SetStyle(style WidgetStyle) {
    input.Widget.Style = style
}

func (input *Input) GetStyle() *WidgetStyle {
    return &input.Widget.Style 
} 

func (input *Input) AddRune(i int, letter rune) error {
    content := input.Widget.GetContent()
    contentRunes := []rune(content.Contents)
    if i<0 || i > len(content.Contents) {
        return errors.New("Invalid rune index.")
    }
    contentRunes = append(contentRunes[:i+1], contentRunes[i:]...)
    contentRunes[i] = letter
    content.Contents = string(contentRunes)
    return nil
} 

func (input *Input) RemoveRune(i int) error {
    content := input.Widget.GetContent()
    contentRunes := []rune(content.Contents)
    if i <= 0 || i > len(content.Contents) {
        return errors.New("Invalid rune index.")
    }
    contentRunes = append(contentRunes[0:i-1], contentRunes[i:]...)
    return nil
}

func (input *Input) Render(screen tcell.Screen, context lcontext.Context) {
    x, y := input.GetCoordinates()
    width := input.GetWidth()
    height := input.GetHeight()

    box := Box{ x, y, width, height, tcell.StyleDefault }
    box.Render(screen)
}
