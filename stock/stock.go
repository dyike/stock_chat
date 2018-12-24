package stock

import (
	"github.com/dyike/httpClient"
	"strings"
	"fmt"
	"encoding/json"
)

type StockItem struct {
	Volume  float64   `json:"volume"`
	Type    string   `json:"type"`
	Name    string    `json:"name"`
	Updown  float64    `json:"updown"`
	Price   float64    `json:"price"`
	Percent float64    `json:"percent"`
	Yestclose float64   `json:"yestclose"`
	High    float64     `json:"high"`
	Low   float64    `json:"low"`
	Open   float64    `json:"open"`
	Time  string  `json:"time"`
}

func GetFeed(list []string) (map[string]StockItem, error) {
	var stocks = map[string]StockItem{}
	feedString := strings.Join(list, ",")
	result, err := fetchData(feedString)
	if err != nil {
		return stocks, err
	}
	fmt.Println(result)
	err = json.Unmarshal([]byte(result), &stocks)
	return stocks, err
}

func fetchData(feedString string) (string, error) {
	url := fmt.Sprintf("http://api.money.netease.com/data/feed/%s,money.api", feedString)
	result, err := httpClient.DoRequest(httpClient.Request{
		Method: "GET",
		URL:    url,
	})

	if err != nil {
		return "", err
	}
	reusltString := strings.Trim(string(result), "_ntes_quote_callback();")
	return reusltString, nil
}