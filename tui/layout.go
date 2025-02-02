package tui

import (
	"errors"
	"fmt"
)

type Layout interface {
    SetRows(n int) error
    GetRows() []Row
    SetColumns(n int) error
    GetColumns() []Column
    SetWidget(w Widget, col, row, colSpan, rowSpan int) error 
    GetColumnsWidth(start, end int) (int, error)
    GetRowsHeight(start, end int) (int, error)
    GetRowPosition(i int) (int, int, error)
    GetColumnPosition(i int) (int, int, error)
    GetParent() Layout
    SetOverflowX(overflow bool)
    SetOverflowY(overflow bool)
    IsOverflowX() bool
    IsOverflowY() bool
}

type Row struct {
    Height int
}

type Column struct {
    Width int
}


type Viewport struct {}

type LayoutSrc struct {
    Parent      Layout
    Width       int
    Height      int
    Rows        []Row
    Columns     []Column
    Widgets     []Widget
    Viewport    Viewport
    OverflowX   bool 
    OverflowY   bool
}

func CreateLayout(w int, h int) *LayoutSrc {
    layout := LayoutSrc {
        Width: w,
        Height: h,
    }

    return &layout
}

func (layout *LayoutSrc) SetRows(n int) error {
    layoutHeight := layout.Height
    rowHeight := int(layoutHeight / n)
    remainder := layoutHeight % n

    if rowHeight < 1 {
        return errors.New("Impossible number of rows.")
    }
    layout.Rows = []Row{}
    for _ = range n {
        layout.Rows = append(layout.Rows, Row{Height: rowHeight} )
    }

    for i := range remainder {
        layout.Rows[(i + 2) % (len(layout.Rows) - 1)].Height += 1
    }
    return nil
}

func (layout *LayoutSrc) GetRows() []Row {
    return layout.Rows
}

func (layout *LayoutSrc) SetColumns(n int) error {
    layoutWidth := layout.Width
    columnWidth := int(layoutWidth / n)
    remainder := layoutWidth % n

    if columnWidth < 1 {
        return errors.New("Impossible number of columns.")
    }

    for _ = range n{
        layout.Columns = append(layout.Columns, Column{ Width: columnWidth})
    }

    for i := range remainder {
        layout.Columns[(i + 2) % (len(layout.Columns) - 1)].Width += 1
    }
    return nil
}

func (layout *LayoutSrc) GetColumns() []Column {
    return layout.Columns
}

func (layout *LayoutSrc) SetWidget(widget Widget, col, row, spanCol, spanRow int) error {
    if (row + spanRow > len(layout.Rows) && !layout.OverflowY) || (col+ spanCol > len(layout.Columns) && !layout.OverflowX) {
        return errors.New(fmt.Sprintf("'%c' widget position out off view.", widget.GetMarker()))
    }
    widget.SetPosition(WidgetPosition{Parent: layout, Row: row, Col: col, SpanRow: spanRow, SpanCol: spanCol})
    width, err := layout.GetColumnsWidth(col, col + spanCol)
    if err != nil {
        return err
    }
    height, err := layout.GetRowsHeight(row, row + spanCol)
    if err != nil {
        return err
    }
    switch container := widget.(type) {
    case *Container:
        container.Layout.Height = height
        container.Layout.Width = width
        container.Layout.Parent = layout
    default:
        widgetBody := CreateLayout(width, height)
        widgetBody.Parent = layout
        widget.SetBody(widgetBody)
    }
    layout.Widgets = append(layout.Widgets, widget)
    return nil
}

//  GetRowsHeight accepts an exclusive range
func (layout *LayoutSrc) GetRowsHeight(start, end int) (int, error) {
    height := 0
    if start >= end || start < 0 || end > len(layout.Rows) {
        return 0, errors.New("Invalid rows range.")
    }
    rows := []Row{}
    if end == len(layout.Rows) {
        rows = layout.Rows[start:]
    } else {
        rows = layout.Rows[start:end]
    }
    for _, row := range rows {
        height += row.Height
    }
    return height, nil
}

//  GetColumnsWidth accepts an inclusive range
func (layout *LayoutSrc) GetColumnsWidth(start, end int) (int, error) {
    width := 0
    if start >= end || start < 0 || end > len(layout.Columns) {
        return 0, errors.New("Invalid columns range.")
    }
    columns := []Column{}
    if end == len(layout.Columns) {
        columns = layout.Columns[start:]
    } else {
        columns = layout.Columns[start:end]
    }
    for _, column := range columns {
        width += column.Width
    }
    return width, nil
}

// GetRowPosition returns Y position relative to the parent layout
func (layout *LayoutSrc) GetRowPosition(index int) (int, int, error) {
    startY, endY := 0, -1
    if index < 0 || index >= len(layout.Rows) {
        return 0, 0, errors.New("Row index does not exist.")
    }

    for i, row := range layout.Rows {
        startY += endY+1
        endY += row.Height
        if i == index {
            break;
        }
    }

    return startY, endY, nil
}

// GetColumnPosition returns X position relative to the parent layout
func (layout *LayoutSrc) GetColumnPosition(index int) (int, int, error) {
    startX, endX := 0, -1
    if index < 0 || index >= len(layout.Columns) {
        return 0, 0, errors.New("Column index does not exist.")
    }

    for i, column := range layout.Columns {
        startX += endX+1 
        endX += column.Width
        if i == index {
            break;
        }
    }

    return startX, endX, nil
}

func (layout *LayoutSrc) GetParent() Layout {
    return layout.Parent
}

func (layout *LayoutSrc) IsOverflowX() bool {
    return layout.OverflowX
}

func (layout *LayoutSrc) SetOverflowX(overflow bool) {
    layout.OverflowX = overflow
}

func (layout *LayoutSrc) IsOverflowY() bool {
    return layout.OverflowY
}

func (layout *LayoutSrc) SetOverflowY(overflow bool) {
    layout.OverflowY = overflow
}

