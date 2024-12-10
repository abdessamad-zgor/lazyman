package gui

import (
	"github.com/abdessamad-zgor/lazyman/src/app"
	"github.com/gdamore/tcell/v2"
)

type Callback func (context app.AppContext) 

type Listener struct {
    Event string
    Callback Callback
}

type Box struct {
    X int
    Y int
    W int
    H int
    Title string
    Content string
    Style BoxStyle
    Editable bool
    Float bool
    Listeners []Listener 
    Children []Box
}

type BoxStyle struct {
    Border tcell.Style
    OnHighlight tcell.Style
    Default tcell.Style
    TitleStyle TextStyle
    ContentStyle TextStyle
}

type TextStyle struct {
    // these are relative to the containing box
    X int
    Y int
    Color tcell.Color
    Background tcell.Color
    Bold bool
    Italic bool
}
