package api

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"aaronbarratt.dev/go/cryptomaster/datatypes"
)

const apiURL = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*datatypes.Rate, error) {
	currency = strings.ToUpper(currency)
	result, err := http.Get(fmt.Sprintf(apiURL, currency))
	if err != nil {
		return nil, err
	}
	if result.StatusCode == http.StatusOK {
		httpBody, err := io.ReadAll(result.Body)
		if err != nil {
			return nil, err
		}
		json := string(httpBody)
		fmt.Println(json)
	} else {
		return nil, fmt.Errorf("status code: %v", result.StatusCode)
	}

	rate := datatypes.Rate{Currency: currency, Price: 20}
	return &rate, nil
}
