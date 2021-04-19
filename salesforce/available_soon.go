package salesforce

import (
	"encoding/json"
	"log"
	"strings"
)

type AvailableSoon struct{}

func productData(respBody string) map[string]interface{} {
	startStr := "window.productJSON = "
	endStr := "window.priceAvailabilityJSON = "
	jsonStart := strings.Index(respBody, startStr) + len(startStr)
	jsonEnd := strings.Index(respBody, endStr)
	jsonStr := respBody[jsonStart : jsonEnd-2]

	var obj map[string]interface{}

	err := json.Unmarshal([]byte(jsonStr), &obj)
	if err != nil {
		log.Println(err)
	}

	return obj
}

func (a *AvailableSoon) IsInStock(respBody string) bool {
	availableSoon := productData(respBody)["product"].(map[string]interface{})["badgesInformation"].(map[string]interface{})["availableSoonBadgeInformation"]
	return availableSoon.(map[string]interface{})["displayAvailableSoonBadge"] == false
}

func (a *AvailableSoon) IsOutOfStock(respBody string) bool {
	availableSoon := productData(respBody)["product"].(map[string]interface{})["badgesInformation"].(map[string]interface{})["availableSoonBadgeInformation"]
	return availableSoon.(map[string]interface{})["displayAvailableSoonBadge"] == "OOS_Preview_Message_PDP"
}
