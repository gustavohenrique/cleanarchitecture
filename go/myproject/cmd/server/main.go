package main

import (
	"runtime/debug"

	"myproject/src/application"
)

func init() {
	debug.SetGCPercent(500)
}

func main() {
	application.New().
		ParseCommandLineArgs().
		LoadConfigurationFile().
		CreateServers().
		Start()
}
