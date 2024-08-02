package api

import (
	"encoding/json"
	"fmt"
	"io"
	"leo/cryptomasters/models"
	"net/http"
	"strings"
)

const apiUrl = "https://cex.io/api/ticket/%s/USD"

func GetRate(currency string) (*models.Rate, error) {
	upCurrency := strings.ToUpper(currency)
	response, err := http.Get(fmt.Sprintf(apiUrl, upCurrency))
	if err != nil {
		return nil, err
	}
	if response.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		var cryptoRate models.Rate
		err = json.Unmarshal(bodyBytes, &cryptoRate)
	} else {
		return nil, fmt.Errorf("status code received: %d", response.StatusCode)
	}
	rate := models.Rate{Currency: currency, Price: 20}
	return &rate, nil
}
