package main

import (
	"fmt"
	"os"
	"github.com/gdamore/tcell/v2"
)

type Lazyman struct {
	Screen tcell.Screen
}

func NewLazyman() *Lazyman {
	screen, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	return &Lazyman{Screen: screen}
}

func (app *Lazyman) Init() {
	if e := app.Screen.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	app.Screen.SetStyle(tcell.StyleDefault)
	app.Screen.Clear()

	app.InitLayout()

	app.Screen.Show()
	app.StartEventLoop()
}

func (app *Lazyman) InitLayout() {
	w, h := app.Screen.Size()
	var boxes [1]Box
	boxes[0] = Box{
		X: 0,
		Y: 0,
		W: int(w / 3),
		H: int(h / 8),
		Title: nil,
		Content: &Text{
			X:        1,
			Y:        1,
			Contents: "https://api.nasa.com/v69/challanger",
		},
		Style: &BoxStyle{
			TitleStyle:   &TextStyle{Bold: true},
			ContentStyle: nil,
			OnHighlight:  nil,
			Border:       nil,
		},
	}

	//boxes[1] = Box{
	//	X: 0,
	//	Y: int(h/8) + 1,
	//	W: int(w / 3),
	//	H: h - int(h/8),
	//	Title: &Text{
	//		X:        2,
	//		Y:        -1,
	//		Contents: "Request",
	//	},
	//	Content: &Text{
	//		X:        1,
	//		Y:        1,
	//		Contents: "https://api.nasa.com/v69/challanger",
	//	},
	//	Style: &BoxStyle{
	//		TitleStyle:   &TextStyle{Bold: true},
	//		ContentStyle: nil,
	//		OnHighlight:  nil,
	//		Border:       nil,
	//	},
	//}

	//boxes[2] = Box{
	//	X: int(w/3) + 1,
	//	Y: 0,
	//	W: w - int(w/3),
	//	H: h,
	//	Title: &Text{
	//		X:        2,
	//		Y:        -1,
	//		Contents: "Response",
	//	},
	//	Content: &Text{
	//		X:        1,
	//		Y:        1,
	//		Contents: "https://api.nasa.com/v69/challanger",
	//	},
	//	Style: &BoxStyle{
	//		TitleStyle:   &TextStyle{Bold: true},
	//		ContentStyle: nil,
	//		OnHighlight:  nil,
	//		Border:       nil,
	//	},
	//}

    for i:=0; len(boxes)>i;i++ {
        box:= boxes[i]
        box.GetDrawF()(app.Screen)
    }
}

func (app *Lazyman) StartEventLoop() {
	for {
		ev := app.Screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape, tcell.KeyEnter:
				app.Screen.Fini()
				os.Exit(0)
			case tcell.KeyCtrlL:
				app.Screen.Sync()
			}
		case *tcell.EventResize:
			app.Screen.Sync()
		}
	}
}
