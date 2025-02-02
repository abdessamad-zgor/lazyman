package tui

import (
	"errors"

	"github.com/abdessamad-zgor/lazyman/event"
	lcontext "github.com/abdessamad-zgor/lazyman/context"
	"github.com/gdamore/tcell/v2"
)

type DropdownState struct {
    Open bool
}

type Menu struct {
    Widget          *WidgetSrc
    Items           []Text
    Dropdown        *DropdownState
    SelectedIndex   int
    SelectedItem    *Text
}

func CreateMenu(marker rune) *Menu {
    menu := Menu{
        Widget: CreateWidget(marker),
    }

    return &menu
}

func (menu *Menu) SetBody(layout Layout) {
    menu.Widget.SetBody(layout)
}

func (menu *Menu) GetBody() Layout {
    return menu.Widget.GetBody()
}

func (menu *Menu) GetMarker() rune {
    return menu.Widget.GetMarker()
}

func (menu *Menu) GetPosition() *WidgetPosition {
    return menu.Widget.GetPosition()
}

func (menu *Menu) GetStyle() *WidgetStyle {
    return menu.Widget.GetStyle()
}

func (menu *Menu) GetTitle() *Text {
    return menu.Widget.GetTitle()
}

func (menu *Menu) GetContent() *Text {
    return menu.Widget.GetContent()
}

func (menu *Menu) GetParent() Widget {
    return menu.Widget.GetParent()
}

func (menu *Menu) GetChildren() []Widget {
    return menu.Widget.GetChildren()
}

func (menu *Menu) GetEventMap() event.EventMap {
    return menu.Widget.GetEventMap()
}

func (menu *Menu) SetStyle(style WidgetStyle) {
    menu.Widget.SetStyle(style)
}

func (menu *Menu) SetWidget(w Widget, col, row, colSpan, rowSpan int) error {
    return menu.Widget.SetWidget(w, col, row, colSpan, rowSpan)
}

func (menu *Menu) SetRows(n int) error {
    return errors.New("Menu is not a layout widget.")
}

func (menu *Menu) SetPosition(position WidgetPosition) {
    menu.Widget.SetPosition(position)
}

func (menu *Menu) SetColumns(n int) error {
    return menu.Widget.SetColumns(n)
}

func (menu *Menu) SetTitle(title string) {
    menu.Widget.SetTitle(title)
}

func (menu *Menu) SetContent(contents string) {
    menu.Widget.SetContent(contents)
}

func (menu *Menu) SetParent(widget Widget) {
    menu.Widget.SetParent(widget)
}

func (menu *Menu) GetCoordinates() (int, int) {
    return menu.Widget.GetCoordinates()
}

func (menu *Menu) GetWidth() int {
    return menu.Widget.GetWidth()
}

func (menu *Menu) GetHeight() int {
    return menu.Widget.GetHeight()
}

func (menu *Menu) SetupEventMap() {

}

func (menu *Menu) GetItems() []Text {
    return menu.Items
}

func (menu *Menu) SetItems(items []Text) {
    menu.Items = append(menu.Items, items...)
}

func (menu *Menu) IsDropdown() bool {
    return (menu.Dropdown != nil)
}

func (menu *Menu) SetDropDown(dropdown *DropdownState) {
    menu.Dropdown = dropdown
} 

func (menu *Menu) SetSelectedIndex(i int) {
    menu.SelectedIndex = i
    menu.SelectedItem = &menu.Items[i]
}

func (menu *Menu) GetSelectedIndex() int {
    return menu.SelectedIndex
}

func (menu *Menu) GetSelectedItem() *Text {
    return menu.SelectedItem
}

func (menu *Menu) GetItemsCount() int {
    return len(menu.GetItems())
}

// Menu.Render should be queued
func (menu *Menu) Render(screen tcell.Screen, context lcontext.Context) {
    renderFn := func(){
        x, y := menu.GetCoordinates()
        width := menu.GetWidth()
        height := menu.GetHeight()

        box := Box{x, y, width, height, tcell.StyleDefault}
        box.Render(screen)
        title := menu.GetTitle()
        content := menu.GetContent()
        title.Render(box, screen)
        content.Render(box, screen)
    }

    if menu.IsDropdown() {
        lcontext.DispatchEvent(lcontext.QueueRender, renderFn)
    } else {
        renderFn()
    }
}
