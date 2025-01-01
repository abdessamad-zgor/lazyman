package lcontext

import (
	"runtime"
)

type Context struct {
	State  State
	Config Config
}

type State map[string]any

type Config map[string]any

func InitContext() Context {
	return Context{}
}

func initConfig() Config {

	if runtime.GOOS == "windows" {
	}

	return Config{}
}
