package main

import (
	"fmt"
	"log"

	"github.com/metarsit/exchange"
)

func main() {
	market := exchange.NewMarketAPI()
	resp, err := market.Symbols()
	if err != nil {
		log.Fatalf("Unable to retrieve Market Symbol: %s", err.Error())
	}

	if resp.Code != "0" {
		log.Fatalf("[%s] API Error %s", resp.Code, resp.Message)
	}

	for _, symbol := range resp.Data {
		fmt.Printf(
			`
Symbol          : %s
Count Coin      : %s
Base Coin       : %s
Amount Precision: %d
Price Precision : %d
`,
			symbol.Name, symbol.CountCoin, symbol.BaseCoin,
			symbol.AmountPrecision, symbol.PricePrecision,
		)
	}
}
