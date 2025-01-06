package state

import (
	"runtime"
)

type Context struct {
	State  State
	Config Config
}

type StateKey = string

const (
    SelectedWidget StateKey = "selected-widget"
)

type State = map[string]any

type Config = map[string]any

func InitContext() Context {
	return Context{}
}

func initConfig() Config {

	if runtime.GOOS == "windows" {
	}

	return Config{}
}
