package tui

import "github.com/abdessamad-zgor/lazyman/event"

type TableOriantation string

const (
    Horizontal  TableOriantation = "horizontal"
    Vertical    TableOriantation = "vertical"
)

type Table struct {
    Widget      *WidgetSrc
    Rows        int
    Columns     int
    Headers     []Text
}

func CreateTable() *Table {
    table := Table{
        Widget: CreateWidget(rune(0)),
    }

    return &table
}

func (table *Table) SetBody(layout Layout) {
    table.Widget.Body = layout
}

func (table *Table) GetBody() Layout {
    return table.Widget.Body
}

func (table *Table) GetMarker() rune {
    return 0
}

func (table *Table) GetPosition() *WidgetPosition {
    return nil
}

func (table *Table) GetStyle() *WidgetStyle {
    return nil
}

func (table *Table) GetTitle() *Text {
    return nil
}

func (table *Table) GetContent() *Text {
    return nil
}

func (table *Table) GetParent() Widget {
    return nil
}

func (table *Table) GetChildren() []Widget {
    return nil
}

func (table *Table) GetEventMap() event.EventMap {
    return table.Widget.EventMap
}


func (table *Table) SetStyle(style WidgetStyle) {
}

func (table *Table) SetWidget(w Widget, col, row, colSpan, rowSpan int) error  {
    return table.Widget.SetWidget(w, col, row, colSpan, rowSpan)
}

func (table *Table) SetRows(n int) error {
    return table.Widget.SetRows(n)
}

func (table *Table) SetPosition(position WidgetPosition) {
}

func (table *Table) SetColumns(n int) error {
    return table.Widget.SetColumns(n)
}

func (table *Table) SetTitle(title string) {
}

func (table *Table) SetContent(contents string) {
}

func (table *Table) SetParent(widget Widget) {
}

func (table *Table) GetCoordinates() (int, int) {
    return table.Widget.GetCoordinates()
}

func (table *Table) GetWidth() int {
    return table.Widget.GetWidth()
}

func (table *Table) GetHeight() int {
    return table.Widget.GetHeight()
}

func (table *Table) SetupEventMap()  {
}
