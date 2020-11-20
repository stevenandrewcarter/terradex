package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProject(t *testing.T) {
	tests := map[string]struct {
		wantedStatusCode int
	}{
		"projectDoesNotExist": {wantedStatusCode: http.StatusBadRequest},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if req, err := http.NewRequest("GET", "/project/test", nil); err != nil {
				t.Fatal(err)
			} else {
				rr := httptest.NewRecorder()
				handler := http.HandlerFunc(GetProject)
				handler.ServeHTTP(rr, req)
				if status := rr.Code; status != tc.wantedStatusCode {
					t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
				}
			}
		})
	}
}
