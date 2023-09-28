package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/banaaron/cryptomaster/datatypes"
)

const apiURL = "https://cex.io/api/ticker/%s/USD"

func GetRate(cryptoCurrency string) (*datatypes.Rate, error) {
	cryptoCurrency = strings.ToUpper(cryptoCurrency)
	result, err := http.Get(fmt.Sprintf(apiURL, cryptoCurrency))
	if err != nil {
		return nil, err
	}
	var response CexResponse
	if result.StatusCode == http.StatusOK {
		httpBody, err := io.ReadAll(result.Body)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(httpBody, &response)
		if err != nil {
			return nil, err
		}

	} else {
		return nil, fmt.Errorf("status code: %v", result.StatusCode)
	}

	rate := datatypes.Rate{Currency: cryptoCurrency, Price: response.Bid}
	return &rate, nil
}
