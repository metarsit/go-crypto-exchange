package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/metarsit/exchange"
)

func main() {
	market := exchange.NewMarketAPI()
	resp, err := market.TickerPrice()
	if err != nil {
		log.Fatalf("Unable to retrieve Market Ticker: %s", err.Error())
	}

	if resp.Code != "0" {
		log.Fatalf("[%s] API Error %s", resp.Code, resp.Message)
	}

	var data exchange.TickerPrice
	json.Unmarshal(*resp.Data, &data)

	for ticker, price := range data {
		fmt.Printf("%s: %v\n", strings.ToUpper(ticker), price)
	}
}
