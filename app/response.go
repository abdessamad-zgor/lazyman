package app

import "github.com/abdessamad-zgor/lazyman/gui"

func InitResponseWidget(w int, h int) gui.Widget {
	responseWidget := gui.CreateWidget(
		2*int(w/5),
		0,
		w-2*int(w/5),
		h,
		&gui.BoxStyle{
			Border: nil,
		},
	)

    responseWidget.SetTitle("Response")
    responseWidget.SetMarker('3')
    
    return responseWidget
}
