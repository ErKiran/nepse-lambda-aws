package main

import (
	"context"
	"fmt"
	"nepse-lambda/nepse"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyStock struct {
	Ticker string `json:"ticker"`
}

type MyStockResponse struct {
	Message string                     `json:"message"`
	Data    *nepse.LastTradingDayStats `json:"data"`
}

func HandleRequest(ctx context.Context, name MyStock) (*MyStockResponse, error) {
	nep, err := nepse.NewNepse()
	if err != nil {
		return nil, err
	}

	currentPrice, err := nep.GetCurrentPrice(name.Ticker)
	if err != nil {
		return nil, err
	}

	return &MyStockResponse{
		Message: fmt.Sprintf("The price of %s is %.2f it has changed %.2f percentage and traded with %d number of stocks", currentPrice.Ticker, currentPrice.Lasttradedprice, currentPrice.Percentagechange, currentPrice.Totaltradequantity),
		Data:    currentPrice,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
