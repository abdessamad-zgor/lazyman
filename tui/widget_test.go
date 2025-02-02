package tui

import "testing"

func TestGetCoordinates(t *testing.T) {
    widget := CreateWidget(rune(0))

    layout := CreateLayout(9, 9)

    layout.SetRows(3)
    layout.SetColumns(3)

    layout.SetWidget(widget, 2, 2, 1, 1)

    x, y := widget.GetCoordinates()

    if x != 6 {
        t.Errorf("x expected to be %d got %d", 6, x)
    }

    if y != 6 {
        t.Errorf("y expected to be %d got %d", 6, y)
    }
}

func TestNestedGetCoordinates(t *testing.T) {
    widget1 := CreateWidget(rune(0))
    widget2 := CreateWidget(rune(0))


    layout := CreateLayout(12, 12)

    layout.SetRows(3)
    layout.SetColumns(3)

    if err := layout.SetWidget(widget1, 2, 2, 1, 1); err != nil {
        t.Fatalf("layout.SetWidget(widget1, 2, 2, 1, 1) failed with %v", err)
    }

    widget1.SetColumns(2)
    widget1.SetRows(2)

    t.Logf("widget1.Body = %v", widget1.Body)
    if widget1.Body == nil {
        t.Errorf("widget1.Body is nil")
    }

    if len(widget1.Body.GetRows()) != 2 {
        t.Errorf("len(widget1.Body.GetRows()) expected to be %d got %d", 2, len(widget1.Body.GetRows()))
    }

    if len(widget1.Body.GetColumns()) != 2 {
        t.Errorf("len(widget1.Body.GetColumns()) expected to be %d got %d", 2, len(widget1.Body.GetColumns()))
    }

    if err := widget1.SetWidget(widget2, 1, 1, 1, 1); err != nil {
        t.Fatalf("widget1.SetWidget(widget2, 1, 1, 1, 1) failed with %v", err)
    }

    x, y := widget2.GetCoordinates()

    if widget2.Parent != widget1 {
        t.Errorf("expected widget1 to equal widget2.Parent")
    }

    if x != 10 {
        t.Errorf("x expected to be %d got %d", 10, x)
    }

    if y != 10 {
        t.Errorf("y expected to be %d got %d", 10, y)
    }
}
