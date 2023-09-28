package main

import (
	"fmt"
	"sync"

	"github.com/banaaron/cryptomaster/api"
)

func main() {
	currencies := [3]string{"BTC", "ETH", "BCH"}
	var waitGroup sync.WaitGroup

	for _, currency := range currencies {
		waitGroup.Add(1)
		// create a copy of currency so that we access it in the anon function
		currency := currency
		// create an anon function
		go func() {
			getCurrencyData(currency)
			// go supports closures, so we can access waitGroup from the outer scope in here
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()
}

func getCurrencyData(currency string) {
	rate, err := api.GetRate(currency)
	if err != nil {
		panic(err)
	}
	fmt.Printf("The rate for %v is 1:%.2f \n", rate.Currency, rate.Price)
}
