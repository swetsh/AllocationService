package api

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPutOrder(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, "/1", r.URL.Path)

		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		assert.Equal(t, "4", buf.String())
		assert.NotEqual(t, "3", buf.String())

		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, `{"id":1,"name":"d1","location":{"geoCoordinate":"x"},"orderId":1}`)
	}))

	defer server.Close()

	url := server.URL

	orderResp, err := PutOrder(url, 1, 4)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	assert.NoError(t, err)

	assert.NotNil(t, orderResp)
	assert.Equal(t, 1, orderResp.ID)
	assert.Equal(t, "d1", orderResp.Name)
	assert.Equal(t, "x", orderResp.Location.GeoCoordinate)
	assert.Equal(t, 1, orderResp.OrderID)
}
