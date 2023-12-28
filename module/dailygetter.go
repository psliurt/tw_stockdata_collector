package module

import (
	"encoding/json"
	"fmt"
	"tw_stockdata_collector/openapi"
	"tw_stockdata_collector/structs"
)

type DailyApiGetter struct {
}

func (dag *DailyApiGetter) DoGet() {
	dag.getStockDayAll()
	dag.getBwibbuAll()
	dag.getStockDayAvgAll()
}

func (dag *DailyApiGetter) getStockDayAll() {

	bs := dag.sendGetRequest("https://openapi.twse.com.tw/v1/exchangeReport/STOCK_DAY_ALL")

	var d []structs.DailyStockDayAll
	err := json.Unmarshal(bs, &d)
	if err != nil {
		panic(err)
	}
	fmt.Println(d[0].Name)
}

func (dag *DailyApiGetter) getBwibbuAll() {
	bs := dag.sendGetRequest("https://openapi.twse.com.tw/v1/exchangeReport/BWIBBU_ALL")
	var d []structs.DailyBwibbuAll
	err := json.Unmarshal(bs, &d)
	if err != nil {
		panic(err)
	}
	fmt.Println(d[0].Name)
}

func (dag *DailyApiGetter) getStockDayAvgAll() {
	bs := dag.sendGetRequest("https://openapi.twse.com.tw/v1/exchangeReport/STOCK_DAY_AVG_ALL")
	var d []structs.DailyBwibbuAll
	err := json.Unmarshal(bs, &d)
	if err != nil {
		panic(err)
	}
	fmt.Println(d[0].Name)
}

func (dag *DailyApiGetter) sendGetRequest(url string) []byte {
	bs, err := openapi.Instance().CreateRequest().
		Method("GET").Url(url).
		AddHeader("Accept", "application/json").
		AddHeader("If-Modified-Since", "Mon, 26 Jul 1997 05:00:00 GMT").
		AddHeader("Cache-Control", "no-cache").
		AddHeader("Pragma", "no-cache").SendQuery()
	if err != nil {
		panic(err)
	}
	return bs
}
