package state

import (
	"github.com/gdamore/tcell/v2"
)

type Event = string

const (
	Edit        Event = "edit"
	Escape      Event = "escape"
	EscapeEdit  Event = "escape-edit"
	Highlight   Event = "highlight"
	Loading     Event = "loading"
	Press       Event = "press"
	TextSelect  Event = "text-select"
	ToggleMenu  Event = "toggle-menu"
	ToggleFloat Event = "toggle-float"
)

type Callback = func(context Context)

type EventMap = map[Event]Callback
type Keybindings = map[tcell.Key]Event
