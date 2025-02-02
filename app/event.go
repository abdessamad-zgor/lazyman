package app

import (
	"github.com/abdessamad-zgor/lazyman/tui"
	_ "github.com/abdessamad-zgor/lazyman/tui"
)

func (app *App) IsInput(marker rune) (bool, error) {
    widget, err := app.GetWidgetByMarker(marker)
    if err != nil {
        return false, err
    }
    _, ok := widget.(*tui.Input)

    return ok, nil
}

func (app *App) GetWidgetByMarker(marker rune) (tui.Widget, error) {

    return nil, nil
}

