package app

import "github.com/abdessamad-zgor/lazyman/gui"

func InitUrlWidget(w int, h int) gui.Widget {
    urlWidget :=  gui.CreateWidget(
		0,
		0,
		2*int(w/5),
		int(h/9),
		&gui.BoxStyle{
            Border: nil,
        },
	)

    urlWidget.SetTitle("URL")
    urlWidget.SetMarker('1')

    return urlWidget
}
