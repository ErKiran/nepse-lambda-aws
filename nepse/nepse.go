package nepse

import (
	"context"
	"fmt"
	"nepse-lambda/utils"
	"net/http"
	"os"
	"time"
)

const (
	Header = "tearsheet/header"
	Health = "overview/topGainers/?count=5"
)

type CurrentPrice struct {
	Response int    `json:"response"`
	Error    string `json:"error"`
	Message  struct {
		Ticker           string    `json:"ticker"`
		Company          string    `json:"company"`
		Latestprice      float64   `json:"latestPrice"`
		Pointchange      float64   `json:"pointChange"`
		Percentagechange float64   `json:"percentageChange"`
		Timestamp        time.Time `json:"timestamp"`
		Wtavgprice       float64   `json:"wtAvgPrice"`
		Sharestraded     int       `json:"sharesTraded"`
		Volume           int       `json:"volume"`
		Mktcap           float64   `json:"mktCap"`
	} `json:"message"`
}

type LastTradingDayStats struct {
	Ticker              string      `json:"ticker"`
	Openprice           float64     `json:"openPrice"`
	Highprice           float64     `json:"highPrice"`
	Lowprice            float64     `json:"lowPrice"`
	PointChanged        float64     `json:"pointChanged"`
	Totaltradequantity  int         `json:"totalTradeQuantity"`
	Lasttradedprice     float64     `json:"lastTradedPrice"`
	Percentagechange    float64     `json:"percentageChange"`
	Lastupdateddatetime string      `json:"lastUpdatedDateTime"`
	Lasttradedvolume    interface{} `json:"lastTradedVolume"`
	Previousclose       float64     `json:"previousClose"`
}

type Nepse struct {
	client *utils.Client
}

func NewNepse() (*Nepse, error) {
	client := utils.NewClient(nil, os.Getenv("NEPSE"))

	_, err := client.NewRequest(http.MethodGet, Health, nil)

	if err != nil {
		return nil, err
	}

	nep := &Nepse{
		client: client,
	}
	return nep, nil
}

func (n Nepse) buildTickerSlug(urlPath, ticker string) string {
	return fmt.Sprintf("%s/?tkr=%s", urlPath, ticker)
}

func (n Nepse) GetCurrentPrice(ticker string) (*LastTradingDayStats, error) {
	url := n.buildTickerSlug(Header, ticker)
	req, err := n.client.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	res := &CurrentPrice{}
	if _, err := n.client.Do(context.Background(), req, res); err != nil {
		return nil, err
	}

	currentPrice := LastTradingDayStats{
		Ticker:             res.Message.Ticker,
		Totaltradequantity: res.Message.Sharestraded,
		PointChanged:       res.Message.Pointchange,
		Lasttradedprice:    res.Message.Latestprice,
		Percentagechange:   res.Message.Percentagechange,
		Lasttradedvolume:   res.Message.Volume,
	}

	return &currentPrice, nil
}
