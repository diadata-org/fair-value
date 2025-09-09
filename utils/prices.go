package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetDiaQuotationPrice(blockchain string, address string) (float64, error) {
	url := fmt.Sprintf("https://api.diadata.org/v1/assetQuotation/%s/%s", blockchain, address)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var result struct {
		Price float64 `json:"Price"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}
	return result.Price, nil
}
