package event

import (
	"github.com/gdamore/tcell/v2"
)

type EventName = string

const (
	InsertMode EventName = "insert-mode"
	VisualMode EventName = "visual-mode"
	NormalMode EventName = "normal-mode"
	Escape     EventName = "escape"
	Quit       EventName = "quit"
	Confirm    EventName = "confirm"
	Key        EventName = "key"
	Help       EventName = "help"
	Save       EventName = "save"
	Left       EventName = "left"
	Right      EventName = "right"
	Top        EventName = "top"
	Bottom     EventName = "bottom"
)

type Event struct {
	Name EventName
	Key  tcell.Key
}

type Callback = func(context Context) func(event Event)

type EventMap = map[EventName]Callback
type Keybindings = map[tcell.Key]EventName

var AppEventMap EventMap

func init() {

}

func SetDefaultEventMap() EventMap {
	appEventMap := make(EventName)
	appEventMap[Edit] = func(appcontext lcontext.Context, event Event) {
		// Get Selected Widget

	}
}
