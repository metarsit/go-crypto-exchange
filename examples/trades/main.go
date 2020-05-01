package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/metarsit/exchange"
)

func main() {
	var symbol string

	flag.StringVar(&symbol, "symbol", "", "Market mark, ethbtc, See below for details")
	flag.Parse()

	market := exchange.NewMarketAPI()
	resp, err := market.Trades(symbol)
	if err != nil {
		log.Fatalf("Unable to retrieve Market Trade: %s", err.Error())
	}

	if resp.Code != "0" {
		log.Fatalf("[%s] API Error %s", resp.Code, resp.Message)
	}

	var data []exchange.MarketTrade
	json.Unmarshal(*resp.Data, &data)

	fmt.Printf("======== %s ========\n", strings.ToUpper(symbol))
	for _, trade := range data {
		fmt.Printf(`ID: %d
	Amount: %v
	Price : %v
	CTime : %v
	Type  : %s
`, trade.ID, trade.Amount, trade.Price, trade.CTime, trade.Type)
	}
}
