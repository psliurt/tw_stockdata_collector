package structs

type DailyStockDayAll struct {
	Code         string
	Name         string
	TradeVolume  string
	TradeValue   string
	OpeningPrice string
	HighestPrice string
	LowestPrice  string
	ClosingPrice string
	Change       string
	Transaction  string
}

type DailyBwibbuAll struct {
	Code          string
	Name          string
	PEratio       string
	DividendYield string
	PBratio       string
}

type DailyStockDayAvgAll struct {
	Code                string
	Name                string
	ClosingPrice        string
	MonthlyAveragePrice string
}
