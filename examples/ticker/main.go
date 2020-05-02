package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/metarsit/exchange"
)

func main() {
	var symbol string

	flag.StringVar(&symbol, "symbol", "", "Market mark, ethbtc, See below for details")
	flag.Parse()

	market := exchange.NewMarketAPI()

	symbol = strings.ToLower(symbol)
	resp, err := market.Ticker(symbol)
	if err != nil {
		log.Fatalf("Unable to retrieve Market Ticker: %s", err.Error())
	}

	if resp.Code != "0" {
		log.Fatalf("[%s] API Error %s", resp.Code, resp.Message)
	}

	// If symbol flag is empty, we will print all tickers
	if symbol == "" {
		var data exchange.TickerAll
		json.Unmarshal(*resp.Data, &data)

		fmt.Printf("Server Time: %s\n", time.Unix(data.Date/1000, 0).Format(time.RFC822Z))
		for _, ticker := range data.List {
			prettyPrint(ticker.Name, ticker)
		}
	} else {
		var data exchange.Symbol
		json.Unmarshal(*resp.Data, &data)
		data.Change = "Not found" // Is not return by server
		prettyPrint(symbol, data)
	}
}

func prettyPrint(name string, data exchange.Symbol) {
	fmt.Printf(`%s:
	High  : %s
	Low   : %s
	Volume: %s
	Last  : %s
	Buy   : %s
	Sell  : %s
	Change: %s
	Rose  : %s
`, strings.ToUpper(name), data.High, data.Low, data.Volume,
		data.Last, data.Buy, data.Sell,
		data.Change, data.Rose)
}
