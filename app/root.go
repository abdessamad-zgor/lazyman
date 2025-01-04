package app

import (
	"fmt"
	"os"

	"github.com/abdessamad-zgor/lazyman/gui"
	"github.com/abdessamad-zgor/lazyman/lcontext"
	"github.com/abdessamad-zgor/lazyman/logger"
	"github.com/gdamore/tcell/v2"
)

type App struct {
	Screen   tcell.Screen
	Context  lcontext.Context
	Widgets  []gui.Widget
	EventMap gui.EventMap
}

func Exit() {
	logger.Close()
	os.Exit(0)
}

func (app *App) Render() {
	for _, widget := range app.Widgets {
		widget.Render(app.Screen, app.Context)
	}
}

func (app *App) StartEventLoop() {
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
	if e := screen.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	w, h := screen.Size()
	urlInput := InitUrlWidget(w, h)
    requestWidget := InitRequestWidget(w, h)
    responseWidget := InitResponseWidget(w, h)
	app := App{
		Screen:  screen,
		Widgets: []gui.Widget{urlInput, requestWidget, responseWidget},
		Context: context,
	}

	app.Screen.SetStyle(tcell.StyleDefault)
	app.Screen.Clear()
	app.Render()
	app.Screen.Show()
	app.StartEventLoop()
}
