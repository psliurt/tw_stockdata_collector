package main

import (
	"tw_stockdata_collector/configure"
	"tw_stockdata_collector/env"
	"tw_stockdata_collector/module"
)

func main() {
	//load configuration
	configure.LoadConfiguration()

	// command
	// env init
	env.Initialize()
	g := &module.DailyApiGetter{}
	g.DoGet()
}
