package main

import (
	"fmt"
	"os"

	"github.com/abdessamad-zgor/lazyman/logger"
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
	var boxes [3]Box

	boxes[0] = Box{
		X:       0,
		Y:       0,
		W:       int(w / 3),
		H:       int(h / 9),
		Title:   &Text{
			X:        3,
			Y:        -1,
			Contents: "[1] - URL",
		},
		Content: &Text{
			X:        1,
			Y:        0,
			Contents: "",
		},
		Style: &BoxStyle{
			OnHighlight: nil,
			Border:      nil,
		},
	}

	boxes[1] = Box{
		X:       0,
		Y:       int(h/9),
		W:       int(w / 3),
		H:       h - int(h / 9),
		Title:   &Text{
			X:        3,
			Y:        -1,
			Contents: "[2] - Request",
		},
		Content: &Text{
			X:        1,
			Y:        0,
			Contents: "",
		},
		Style: &BoxStyle{
			OnHighlight: nil,
			Border:      nil,
		},
	}

	boxes[2] = Box{
		X:       int(w / 3),
		Y:       0,
		W:       w-int(w / 3),
		H:       h,
		Title:   &Text{
			X:        3,
			Y:        -1,
			Contents: "[3] - Response",
		},
		Content: &Text{
			X:        1,
			Y:        0,
			Contents: "",
		},
		Style: &BoxStyle{
			OnHighlight: nil,
			Border:      nil,
		},
	}

	for i := 0; len(boxes) > i; i++ {
		box := boxes[i]
		box.GetDrawF()(app.Screen)
	}
}

func Exit() {
	logger.Close()
	os.Exit(0)
}

func (app *Lazyman) StartEventLoop() {
	for {
		ev := app.Screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape, tcell.KeyEnter:
				app.Screen.Fini()
				Exit()
			case tcell.KeyCtrlL:
				app.Screen.Sync()
			}
		case *tcell.EventResize:
			app.Screen.Sync()
		}
	}
}
