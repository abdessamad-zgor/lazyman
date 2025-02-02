package app

import "testing"

func TestSetupAppLayout(t *testing.T) {
    container := SetupAppLayout(120, 40)

    widgets := container.GetWidgets() 

    if len(widgets) != 4 {
        t.Errorf("len(widgets) expected to be %d got %d", 4, len(widgets))
    }
}
