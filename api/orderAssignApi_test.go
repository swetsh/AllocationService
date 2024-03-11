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

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"id":1,"name":"d1","location":{"geoCoordinate":"x"},"orderId":1}`)
	}))
	defer server.Close()

	orderResp, err := PutOrder(server.URL, 1, 4)

	assert.NoError(t, err)

	assert.NotNil(t, orderResp)
	assert.Equal(t, 1, orderResp.ID)
	assert.Equal(t, "d1", orderResp.Name)
	assert.Equal(t, "x", orderResp.Location.GeoCoordinate)
	assert.Equal(t, 1, orderResp.OrderID)
}
