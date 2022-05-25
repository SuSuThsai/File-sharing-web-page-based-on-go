package main

import "C"
import (
	"File-sharing-web-page/ControlCenter/CConfig"
	"File-sharing-web-page/ControlCenter/CRouters"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func main() {
	CConfig.CInits()
	viper.OnConfigChange(func(in fsnotify.Event) {
		CConfig.CInits()
	})
	fsnotify.NewWatcher()
	CRouters.IniteRouter()
}
