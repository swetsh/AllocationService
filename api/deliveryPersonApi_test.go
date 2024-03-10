package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetDeliveryPersons(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `[{"id":1,"name":"John"},{"id":2,"name":"Alice"}]`)
	}))
	defer server.Close()

	url := server.URL

	deliveryPersons, err := GetDeliveryPersons(url)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(deliveryPersons) != 2 {
		t.Errorf("unexpected number of delivery persons, want 2, got %d", len(deliveryPersons))
	}

	deliveryPersons, err = GetDeliveryPersons("invalid-url")
	if err == nil {
		t.Error("expected error but got nil")
	}
	if deliveryPersons != nil {
		t.Errorf("expected nil deliveryPersons but got %v", deliveryPersons)
	}

	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	url = server.URL

	deliveryPersons, err = GetDeliveryPersons(url)
	if err == nil {
		t.Error("expected error but got nil")
	}
	if deliveryPersons != nil {
		t.Errorf("expected nil deliveryPersons but got %v", deliveryPersons)
	}

	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `invalid JSON`)
	}))
	defer server.Close()

	url = server.URL

	deliveryPersons, err = GetDeliveryPersons(url)
	if err == nil {
		t.Error("expected error but got nil")
	}
	if deliveryPersons != nil {
		t.Errorf("expected nil deliveryPersons but got %v", deliveryPersons)
	}
}
