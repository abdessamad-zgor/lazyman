package app

import (
	"fmt"
	"os"
	"github.com/gdamore/tcell/v2"
)

type Lazyman struct {
  Screen tcell.Screen
}

func NewLazyman() *Lazyman {
    screen, e := tcell.NewScreen();
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
    for xi := range w {
        for yj := range h {
            if xi == 0 {
                app.Screen.SetContent(xi, yj, tcell.RuneVLine, []rune{}, tcell.StyleDefault)
            }
            if yj == 0 {
                app.Screen.SetContent(xi, yj, tcell.RuneHLine, []rune{}, tcell.StyleDefault)
            }
            if xi == w-1 {
                app.Screen.SetContent(xi, yj, tcell.RuneVLine, []rune{}, tcell.StyleDefault)
            }
            if yj == h-1 {
                app.Screen.SetContent(xi, yj, tcell.RuneHLine, []rune{}, tcell.StyleDefault)
            }
        }
    }
    app.Screen.SetContent(0, 0, tcell.RuneULCorner, []rune{}, tcell.StyleDefault)
    app.Screen.SetContent(w-1, 0, tcell.RuneURCorner, []rune{}, tcell.StyleDefault)
    app.Screen.SetContent(0, h-1, tcell.RuneLLCorner, []rune{}, tcell.StyleDefault)
    app.Screen.SetContent(w-1, h-1, tcell.RuneLRCorner, []rune{}, tcell.StyleDefault)
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
