package state

import (
	"runtime"
)

type Config = map[string]any


func initConfig() Config {

	if runtime.GOOS == "windows" {
	}

	return Config{}
}
