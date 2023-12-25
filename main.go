package main

import (
	"encoding/json"
	"fmt"
	"tw_stockdata_collector/configure"
	"tw_stockdata_collector/env"
	"tw_stockdata_collector/openapi"
	"tw_stockdata_collector/structs"
)

func main() {
	//load configuration
	configure.LoadConfiguration()

	// command
	// env init
	env.Initialize()
	bs, err := openapi.Instance().CreateRequest().
		Method("GET").Url("https://openapi.twse.com.tw/v1/exchangeReport/STOCK_DAY_ALL").
		AddHeader("Accept", "application/json").
		AddHeader("If-Modified-Since", "Mon, 26 Jul 1997 05:00:00 GMT").
		AddHeader("Cache-Control", "no-cache").
		AddHeader("Pragma", "no-cache").SendQuery()
	if err != nil {
		panic(err)
	}
	var d []structs.STOCK_DAY_ALL
	err = json.Unmarshal(bs, &d)
	if err != nil {
		panic(err)
	}
	fmt.Println(d[0].Name)
}
