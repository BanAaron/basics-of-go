package api

import (
	"fmt"
	"net/http"

	"aaronbarratt.dev/go/cryptomaster/datatypes"
)

const apiURL = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (datatypes.Rate, error) {
	result, err := http.Get(fmt.Sprintf(apiURL, currency))
	if err != nil {

	}
}
