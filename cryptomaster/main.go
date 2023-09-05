package main

import (
	"fmt"

	"aaronbarratt.dev/go/cryptomaster/api"
)

func main() {
	rate, err := api.GetRate("BTC")
	if err != nil {
		panic(err)
	}

	fmt.Println(rate.Currency, rate.Price)
}
