package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetProjectDoesNotExist(t *testing.T) {
	req, err := http.NewRequest("GET", "/project/test", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetProject)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
