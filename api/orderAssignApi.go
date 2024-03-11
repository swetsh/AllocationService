package api

import (
	"allocation-service/responses"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func PutOrder(url string, personID int, orderID int) (*responses.OrderResponse, error) {
	url = fmt.Sprintf("%s/%d", url, personID)

	orderBytes := []byte(fmt.Sprintf("%d", orderID))

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(orderBytes))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %s", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	auth := os.Getenv("USERNAME") + ":" + os.Getenv("PASSWORD")
	fmt.Println(auth)
	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Set("Authorization", basicAuth)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var orderResp responses.OrderResponse
	if err := json.NewDecoder(resp.Body).Decode(&orderResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %s", err)
	}

	return &orderResp, nil
}
