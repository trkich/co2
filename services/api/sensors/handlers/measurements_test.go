package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleMeasurements(t *testing.T) {

	req, err := http.NewRequest("GET", "/api/v1/sensors/TESTAMER2/measurements", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleStatus)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

