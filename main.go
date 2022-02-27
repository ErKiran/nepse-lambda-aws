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

func HandleRequest(ctx context.Context, name MyStock) (string, error) {
	nep, err := nepse.NewNepse()
	if err != nil {
		panic(err)
	}

	currentPrice, err := nep.GetCurrentPrice(name.Ticker)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("The price of %s is %.2f it has changed %.2f percentage and traded with volume %d", currentPrice.Ticker, currentPrice.Lasttradedprice, currentPrice.Percentagechange, currentPrice.Totaltradequantity), nil
}

func main() {
	lambda.Start(HandleRequest)
}
