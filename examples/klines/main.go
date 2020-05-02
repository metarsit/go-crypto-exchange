package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/metarsit/exchange"
)

func main() {
	var symbol string
	var period int

	flag.StringVar(&symbol, "symbol", "", "Market mark, ethbtc, See below for details")
	flag.IntVar(&period, "period", 1, "Given in minutes. Possible values are [1, 5, 15, 30, 60, 1440, 10080, 43200]")
	flag.Parse()

	market := exchange.NewMarketAPI()
	resp, err := market.KLines(symbol, period)
	if err != nil {
		log.Fatalf("Unable to retrieve Market KLines: %s", err.Error())
	}

	if resp.Code != "0" {
		log.Fatalf("[%s] API Error %s", resp.Code, resp.Message)
	}

	var data [][]float64
	json.Unmarshal(*resp.Data, &data)

	for i, line := range data {
		fmt.Printf(`%d:
	Timestamp     : %e
	Opening Price : %v
	Highest       : %v
	Minimum       : %v
	Closing price : %v
	Volume        : %v
`, i+1, line[0], line[1], line[2], line[3], line[4], line[5])
	}
}
