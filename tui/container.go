package tui

import "github.com/abdessamad-zgor/lazyman/event"

type Container struct {
    Layout *LayoutSrc
}

func CreateRootContainer(w, h int) *Container {
    container := Container{
        Layout: CreateLayout(w, h),
    }

    return &container
}

func CreateChildContainer() *Container {
    container:=  Container{
        Layout: CreateLayout(0, 0),
    }

    return &container;
}

func (container *Container) SetRows(n int) error {
    return container.Layout.SetRows(n)
}

func (container *Container) SetColumns(n int) error {
    return container.Layout.SetColumns(n)
}

func (container *Container) SetWidget(w Widget, col, row, colSpan, rowSpan int) error  {
    return container.Layout.SetWidget(w, col, row, colSpan, rowSpan)
}

func (container *Container) GetColumnsWidth(start, end int) (int, error) {
    return container.Layout.GetColumnsWidth(start, end)
}

func (container *Container) GetRowsHeight(start, end int) (int, error) {
    return container.Layout.GetRowsHeight(start, end)
}

func (container *Container) GetRowPosition(i int) (int, int, error) {
    return container.Layout.GetRowPosition(i)
}

func (container *Container) GetColumnPosition(i int) (int, int, error) {
    return container.Layout.GetColumnPosition(i)
}

func (container *Container) SetOverflowX(overflow bool) {
    container.Layout.SetOverflowX(overflow)
}

func (container *Container) SetOverflowY(overflow bool) {
    container.Layout.SetOverflowY(overflow)
}

func (container *Container) IsOverflowX() bool {
    return container.Layout.IsOverflowX()
}

func (container *Container) IsOverflowY() bool {
    return container.Layout.IsOverflowY()
}

func (container *Container) SetBody(layout Layout) {
}

func (container *Container) GetBody() Layout {
    return container.Layout
}

func (container *Container) GetMarker() rune {
    return rune(0)
}

func (container *Container) GetPosition() *WidgetPosition {
    return nil
}

func (container *Container) GetStyle() *WidgetStyle {
    return nil
}

func (container *Container) GetTitle() *Text {
    return nil
}

func (container *Container) GetContent() *Text {
    return nil
}

func (container *Container) GetParent() Widget {
    var parent *WidgetSrc = nil
    return parent
}

func (container *Container) GetChildren() []Widget {
    return nil
}

func (container *Container) GetEventMap() event.EventMap {
    return make(event.EventMap)
}

func (container *Container) SetStyle(style WidgetStyle) {

}

func (container *Container) SetPosition(position WidgetPosition) {
}

func (container *Container) SetTitle(title string) {
}

func (container *Container) SetContent(contents string) {
}

func (container *Container) SetParent(widget Widget) {
}

func (container *Container) SetupEventMap() {

}

func (container *Container) GetCoordinates() (int, int) {
    return 0, 0
}

func (container *Container) GetWidth() int {
    return 0
}

func (container *Container) GetHeight() int {
    return 0
}

func (container *Container) GetWidgets() []Widget {
    containerWidgets := container.Layout.Widgets
    widgets := []Widget{}
    for _, widget := range containerWidgets {
        container, ok := widget.(*Container)
        if ok {
            widgets = append(widgets, container.GetWidgets()...)
        } else {
            widgets = append(widgets, widget)
        }
    }
    return widgets
}
