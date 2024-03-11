package api

import (
	res "allocation-service/responses"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func GetDeliveryPersons(url string) ([]*res.DeliveryPerson, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %s", err)
	}

	auth := os.Getenv("USERNAME") + ":" + os.Getenv("PASSWORD")
	fmt.Println(auth)
	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Set("Authorization", basicAuth)

	resp, err := client.Do(req)
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
