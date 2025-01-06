package app

import (
	"github.com/abdessamad-zgor/lazyman/state"
)

func GetAppEventMap() state.EventMap {
	eventMap := make(state.EventMap)

	eventMap[state.Escape] = func(context state.Context) {
        if _, ok:=context.State[state.SelectedWidget]; ok {

		}
	}
	return eventMap
}
