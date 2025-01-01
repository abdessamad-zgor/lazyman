package app

import (
	"fmt"
	"os"

	"github.com/abdessamad-zgor/lazyman/gui"
	"github.com/abdessamad-zgor/lazyman/lcontext"
	"github.com/abdessamad-zgor/lazyman/logger"
	"github.com/gdamore/tcell/v2"
)

type Layout struct {
	Screen   tcell.Screen
	Context  lcontext.Context
	Widgets  []gui.Widget
	EventMap gui.EventMap
}

func Exit() {
	logger.Close()
	os.Exit(0)
}

func (app *Layout) Render() {
    for _, widget := range app.Widgets {
        widget.Render(app.Screen, app.Context)
    }
}

func (app *Layout) StartEventLoop() {
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

func Init() {
	screen, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	context := lcontext.InitContext()
	w, h := screen.Size()
	urlInput := gui.CreateWidget(
		0,
		0,
		2*int(w/5),
		int(h/9),
		&gui.BoxStyle{Border: nil},
	)
	requestWidget := gui.CreateWidget(
		0,
		int(h/9),
		2*int(w/5),
		h-int(h/9),
		&gui.BoxStyle{Border: nil},
	)
	responseWidget := gui.CreateWidget(
		2*int(w/5),
		0,
		w-2*int(w/5),
		h,
		&gui.BoxStyle{
			Border: nil,
		},
	)

	layout := Layout{
		Screen:  screen,
		Widgets: []gui.Widget{urlInput, requestWidget, responseWidget},
		Context: context,
	}
	if e := layout.Screen.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	layout.Screen.SetStyle(tcell.StyleDefault)
	layout.Screen.Clear()
    layout.Render()
    layout.Screen.Show()
    layout.StartEventLoop()
}
