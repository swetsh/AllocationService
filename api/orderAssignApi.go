package api

import (
	"bytes"
	"fmt"
	"net/http"
)

func PutOrder(url string, personID int, orderID int) error {
	url = fmt.Sprintf("%s/%d", url, personID)

	orderBytes := []byte(fmt.Sprintf("%d", orderID))

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(orderBytes))
	if err != nil {
		return fmt.Errorf("error creating request: %s", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	fmt.Println(resp)

	return nil
}
