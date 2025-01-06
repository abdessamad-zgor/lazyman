package app

import (
	"fmt"
	"os"

	"github.com/abdessamad-zgor/lazyman/gui"
	"github.com/abdessamad-zgor/lazyman/logger"
	"github.com/abdessamad-zgor/lazyman/state"
	"github.com/gdamore/tcell/v2"
)

type App struct {
	Screen      tcell.Screen
	Context     state.Context
	Widgets     []gui.Widget
	EventMap    state.EventMap
	Keybindings state.Keybindings
    EventChannel chan state.Event
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

func (app *App) StartEventListner() {
	for {
		ev := app.Screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			key := ev.Key()
            event, ok := app.Keybindings[key]
            if ok {
                app.EventChannel <- event
            }
		case *tcell.EventResize:
			app.Screen.Sync()
		}
	}
}

func (app *App) StartEventLoop() {
	for {
        select {
        case event :=<- app.EventChannel:
            callback, ok := app.EventMap[event]
            if ok {
                callback(app.Context)
            }
        }
	}
}

func Init() {
	screen, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	context := state.InitContext()
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
	go app.StartEventListner()
    go app.StartEventLoop()
}
