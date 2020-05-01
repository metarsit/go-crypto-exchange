package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/metarsit/exchange"
)

func main() {
	var symbol, option string

	flag.StringVar(&symbol, "symbol", "", "Market mark, ethbtc, See below for details")
	flag.StringVar(&option, "option", "", "The depth type -- options: step0, step1, step2 (Merger depth0-2). step0 has the highest accuracy")
	flag.Parse()

	market := exchange.NewMarketAPI()
	resp, err := market.Depth(symbol, option)
	if err != nil {
		log.Fatalf("Unable to retrieve Market Depth: %s", err.Error())
	}

	if resp.Code != "0" {
		log.Fatalf("[%s] API Error %s", resp.Code, resp.Message)
	}

	var data exchange.MarketDepth
	json.Unmarshal(*resp.Data, &data)

	for _, j := range data.Tick.Asks {
		fmt.Printf(`Asks: %v
`, j)
	}
	for _, j := range data.Tick.Bids {
		fmt.Printf(`Bids: %v
`, j)
	}
	fmt.Printf(`Time: %d`, data.Tick.Time)
}
