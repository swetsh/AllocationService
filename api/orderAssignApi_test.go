package api

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPutOrder(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPut {
			http.Error(w, "expected PUT method", http.StatusBadRequest)
			return
		}
		expectedURL := fmt.Sprintf("/api/v1/delivery-persons/%d", 4)
		if r.URL.Path != expectedURL {
			http.Error(w, fmt.Sprintf("expected URL: %s, got: %s", expectedURL, r.URL.Path), http.StatusBadRequest)
			return
		}

		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		orderID := buf.String()
		if orderID != "123" {
			http.Error(w, fmt.Sprintf("expected orderID: 123, got: %s", orderID), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusAccepted)
	}))
	defer server.Close()

	url := server.URL

	err := PutOrder("invalid-url", 4, 123)
	if err == nil {
		t.Error("expected error but got nil")
	}

	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	err = PutOrder(url, 4, 123)
	if err == nil {
		t.Error("expected error but got nil")
	}
}
