package checker

import (
	"errors"
	"stock-checker/salesforce"
	"stock-checker/shopify"
)

type Product struct {
	Name string `json:"name"`
	Url  string `json:"url"`
	Type string `json:"type"`
}

type StockChecker interface {
	IsInStock(respBody string) bool
	IsOutOfStock(respBody string) bool
}

func (p *Product) StockChecker() (StockChecker, error) {
	switch p.Type {
	case "shopify":
		return new(shopify.Shopify), nil
	case "salesforceAvailableSoon":
		return new(salesforce.AvailableSoon), nil
	}

	return nil, errors.New("unknown product type - " + p.Type)
}
