package tui

import (
	"errors"
	"fmt"

	"github.com/abdessamad-zgor/lazyman/event"
	lcontext "github.com/abdessamad-zgor/lazyman/context"
	"github.com/gdamore/tcell/v2"
)

type Tab struct {
    Header string
    Window Widget
}

type Tabs struct {
    Widget      *WidgetSrc
    Selected    Widget
    Tabs        []Tab
}

func CreateTabs(marker rune) *Tabs {
    tabs := Tabs{
        Widget: CreateWidget(marker),
    }

    return &tabs
}


func (tabs *Tabs) SetBody(layout Layout) {
    tabs.Widget.Body = layout
}

func (tabs *Tabs) GetBody() Layout {
    return tabs.Widget.Body
}

func (tabs *Tabs) GetMarker() rune {
    return tabs.Widget.Marker
}

func (tabs *Tabs) GetPosition() *WidgetPosition {
    return &tabs.Widget.Position
}

func (tabs *Tabs) GetStyle() *WidgetStyle {
    return &tabs.Widget.Style
}

func (tabs *Tabs) GetTitle() *Text {
    return tabs.Widget.Title
}

func (tabs *Tabs) GetContent() *Text {
    return tabs.Widget.Content
}

func (tabs *Tabs) GetParent() Widget {
    return tabs.Widget.Parent
}

func (tabs *Tabs) GetChildren() []Widget {
    return tabs.Widget.Children
}

func (tabs *Tabs) GetEventMap() event.EventMap {
    return tabs.Widget.EventMap
}

func (tabs *Tabs) SetStyle(style WidgetStyle) {
    tabs.Widget.Style = style
}

func (tabs *Tabs) SetWidget(w Widget, col, row, colSpan, rowSpan int) error  {
    return tabs.Widget.SetWidget(w, col, row, colSpan, rowSpan)
}

func (tabs *Tabs) SetRows(n int) error {
    return tabs.Widget.SetRows(n)
}

func (tabs *Tabs) SetPosition(position WidgetPosition) {
    tabs.Widget.SetPosition(position)
}

func (tabs *Tabs) SetColumns(n int) error {
    return tabs.Widget.SetColumns(n)
}

func (tabs *Tabs) SetTitle(title string) {
    tabs.Widget.SetTitle(title)
}

func (tabs *Tabs) SetContent(contents string) {
    tabs.Widget.SetContent(contents)
}

func (tabs *Tabs) SetParent(widget Widget) {
    tabs.Widget.SetParent(widget)
}

func (tabs *Tabs) GetCoordinates() (int, int) {
    return tabs.Widget.GetCoordinates()
}

func (tabs *Tabs) GetWidth() int {
    return tabs.Widget.GetWidth()
}

func (tabs *Tabs) GetHeight() int {
    return tabs.Widget.GetHeight()
}

func (tabs *Tabs) SetupEventMap()  {

}

func (tabs *Tabs) GetHeaders() []string {
    headers := []string{}
    for _, tab := range tabs.Tabs {
        headers = append(headers, tab.Header)
    }
    return headers
}

func (tabs *Tabs) AddTabs(header string, widget Widget) {
    tabs.Tabs = append(tabs.Tabs, Tab{header, widget})
} 

func (tabs *Tabs) GetTabWindow(header string) (Widget, error) {
    var foundTab *Tab = nil
    for _, tab := range tabs.Tabs {
        if tab.Header == header {
            foundTab = &tab
        }
    }
    if foundTab != nil {
        return foundTab.Window, nil
    } else {
        return nil, errors.New(fmt.Sprintf("Header '%s' not found.", header))
    }
}

func (tabs *Tabs) Render(screen tcell.Screen, context lcontext.Context) {
    x, y := tabs.GetCoordinates()
    width := tabs.GetWidth()
    height := tabs.GetHeight()

    box := Box{ x, y, width, height, tcell.StyleDefault }
    box.Render(screen)
}
