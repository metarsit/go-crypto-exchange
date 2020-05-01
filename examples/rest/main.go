package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/metarsit/exchange"
)

func main() {
	var key, secret string

	flag.StringVar(&key, "api", "", "API Key Generated by Crypto Exchange")
	flag.StringVar(&secret, "secret", "", "Secret Key Generated by Crypto Exchange")
	time := time.Now().UnixNano() / 1000000
	flag.Parse()

	api, err := exchange.NewUserAPI(key, secret, time)
	if err != nil {
		log.Fatalf("Unable to create UserAPI Instance: %s", err.Error())
	}
	resp, err := api.Account()
	if err != nil {
		log.Fatalf("Unable to retrieve Account Balance: %s", err.Error())
	}

	if resp.Code != "0" {
		log.Fatalf("[%s] API Error %s", resp.Code, resp.Message)
	}

	fmt.Printf("Total Asset: %s", resp.Data.TotalAsset)
	for _, coin := range resp.Data.CoinList {
		fmt.Printf(
			`
	Coin Name   : %s
	Normal      : %s
	Locked      : %s
	BTCValuation: %s
`,
			coin.Name, coin.Normal, coin.Locked, coin.BTCValuation,
		)
	}
}