package tui

import "testing"

func TestSetRowsAndColumns(t *testing.T) {
    layout := CreateLayout(9, 9)

    layout.SetRows(3)
    layout.SetColumns(3)

    if len(layout.Rows) != 3 {
        t.Errorf("layout.Rows expected to be of length %d got %d", 3, len(layout.Rows))
    }

    if len(layout.Columns) != 3 {
        t.Errorf("layout.Columns expected to be of length %d got %d", 3, len(layout.Columns))
    }

    if layout.Rows[0].Height != 3 {
        t.Errorf("layout.Rows[%d].Height expected to be %d got %d", 0, 3, layout.Rows[0].Height)
    }

    if layout.Columns[0].Width != 3 {
        t.Errorf("layout.Columns[%d].Width expected to be %d got %d", 0, 3, layout.Columns[0].Width)
    }
}

func TestSetWidget(t *testing.T) {
    widget := CreateWidget(rune(0))

    layout := CreateLayout(9, 9)

    layout.SetRows(3)
    layout.SetColumns(3)

    layout.SetWidget(widget, 2, 2, 1, 1)

    if widget.GetWidth() != 3 {
        t.Errorf("widget.GetWidth expected to be %d got %d", 3, widget.GetWidth())
    }

    if widget.GetHeight() != 3 {
        t.Errorf("widget.GetHeight expected to be %d got %d", 3, widget.GetHeight())
    }

    if len(layout.Widgets) != 1 {
        t.Errorf("len(layout.Widgets) expected to be %d got %d", 1, len(layout.Widgets) )
    }

    if layout != widget.GetBody().GetParent() {
        t.Errorf("expected widget.Body.Parent equals layout")
    }
}

