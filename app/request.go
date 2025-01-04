package app

import "github.com/abdessamad-zgor/lazyman/gui"

func InitRequestWidget(w int, h int) gui.Widget{
    requestWidget := gui.CreateWidget(
		0,
		int(h/9),
		2*int(w/5),
		h-int(h/9),
		&gui.BoxStyle{
            Border: nil,
        },
	)

    requestWidget.SetTitle("Request")
    requestWidget.SetMarker('2')

    return requestWidget
}
