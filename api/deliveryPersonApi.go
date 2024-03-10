package api

import (
	res "allocation-service/responses"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetDeliveryPersons(url string) ([]*res.DeliveryPerson, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch delivery persons: %s", err)
	}
	defer resp.Body.Close()

	var deliveryPersons []*res.DeliveryPerson
	if err := json.NewDecoder(resp.Body).Decode(&deliveryPersons); err != nil {
		return nil, fmt.Errorf("failed to decode response: %s", err)
	}

	return deliveryPersons, nil
}
