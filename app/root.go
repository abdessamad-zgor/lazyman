package app

import (
	"fmt"
	"os"
	"reflect"

	lcontext "github.com/abdessamad-zgor/lazyman/context"
	levent "github.com/abdessamad-zgor/lazyman/event"
	"github.com/abdessamad-zgor/lazyman/logger"
	"github.com/abdessamad-zgor/lazyman/tui"
	"github.com/gdamore/tcell/v2"
)

type App struct {
	Screen       tcell.Screen
	Context      lcontext.Context
	Root         *tui.Container
	EventMap     levent.EventMap
	Keybindings  levent.Keybindings
	EventChannel chan levent.Event
}

func Exit() {
	logger.Close()
	os.Exit(0)
}

func (app *App) Render() {
    widgets := app.Root.GetWidgets()
    logger.Info("widgets: ", widgets)
    for _, widget := range widgets {
        logger.Info("widget type: ", reflect.TypeOf(widget).String())
        tuiElement, ok := widget.(tui.TUI) 
        if ok {
            logger.Info("Tui: ", tuiElement)
            tuiElement.Render(app.Screen, app.Context)
        }
    }
}

func (app *App) StartEventListener() {
	for {
		ev := app.Screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			key := ev.Key()
            logger.Info(key)
			event, ok := app.Keybindings[key]
			if ok {
				app.EventChannel <- levent.Event{event, key}
			}
		case *tcell.EventResize:
			app.Screen.Sync()
		}
	}
}

func (app *App) StartEventLoop() {
	for {
		select {
		case event := <-app.EventChannel:
			callback, ok := app.EventMap[event.Name]
			if ok {
				callback(app.Context, event)
			}
		case dispatchEvent, _:= <-lcontext.DispatchContextChannel:
			switch dispatchEvent.Event {
			case lcontext.HighlightWidget:
			case lcontext.SelectWidget:
            case lcontext.QueueRender:
                value, ok := app.Context.GetValue(lcontext.RenderQueue)
                payload := dispatchEvent.Payload
                if ok {
                    queue, qOk := value.([]func ()) 
                    rendefFn, fOk := payload.(func())

                    if qOk && fOk {
                        queue = append(queue, rendefFn)
                        app.Context.SetValue(lcontext.RenderQueue, queue)
                    } else {
                        panic(fmt.Sprintf("Invalid context value '%s' : %v or invalid cast %v.", lcontext.RenderQueue, value, payload))
                    }
                } else {
                    app.Context.SetValue(lcontext.RenderQueue, [](func()){})
                }
			}
		}
	}
}

func SetupAppLayout(width, height int) *tui.Container {
	root := tui.CreateRootContainer(width, height)
	root.SetColumns(2)
	root.SetRows(1)

    // Request container
	requestContainer := tui.CreateChildContainer()
	err := root.SetWidget(requestContainer, 0, 0, 1, 1)
	if err != nil {
		panic(err)
	}

	err = requestContainer.SetColumns(5)
	if err != nil {
		panic(err)
	}

	err = requestContainer.SetRows(8)
	if err != nil {
		panic(err)
	}

    // Response tabs
	responseTabs := tui.CreateTabs('P')
	err = root.SetWidget(responseTabs, 1, 0, 1, 1)
	if err != nil {
		panic(err)
	}

    // Method menu
	method := tui.CreateMenu('M')
	err = requestContainer.SetWidget(method, 0, 0, 1, 1)
	if err != nil {
		panic(err)
	}

    // Url input
	url := tui.CreateInput('U')
	err = requestContainer.SetWidget(url, 1, 0, 4, 1)
    if err != nil {
        panic(err)
    }

    // Request tabs
    requestTabs := tui.CreateTabs('Q')
    err = requestContainer.SetWidget(requestTabs, 0, 1, 5, 7)
    if err != nil {
        panic(err)
    }
    headersTable := tui.CreateTable()
    requestTabs.AddTabs("Headers", headersTable)

    return root
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
    logger.Info(context)
	w, h := screen.Size()
    logger.Info(w, h)

	app := App{
		Screen:    screen,
		Root: SetupAppLayout(w, h),
		Context:   context,
        Keybindings: make(map[tcell.Key]string),
        EventMap: make(map[string]func(context lcontext.Context, event levent.Event)),
        EventChannel: make(chan levent.Event),
	}
    logger.Info(app)

	app.Screen.SetStyle(tcell.StyleDefault)
	app.Screen.Clear()
	app.Render()
	app.Screen.Show()

	go app.StartEventListener()
	go app.StartEventLoop()
}
