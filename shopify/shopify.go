package shopify

import (
	"encoding/json"
	"log"
)

type Shopify struct{}

func productData(respBody string) map[string]interface{} {
	var obj map[string]interface{}
	err := json.Unmarshal([]byte(respBody), &obj)
	if err != nil {
		log.Println(err)
	}
	return obj
}

func (s *Shopify) IsInStock(respBody string) bool {
	return productData(respBody)["available"] == true
}

func (s *Shopify) IsOutOfStock(respBody string) bool {
	return productData(respBody)["available"] == false
}
