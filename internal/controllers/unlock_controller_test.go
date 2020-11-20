package controllers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUnlockProject(t *testing.T) {
	tests := map[string]struct {
		wantedStatusCode int
		projectID        string
		username         string
	}{
		"withoutProjectID": {wantedStatusCode: http.StatusBadRequest, projectID: "", username: ""},
		"withoutUsername":  {wantedStatusCode: http.StatusBadRequest, projectID: "1234", username: ""},
		"successful":       {wantedStatusCode: http.StatusOK, projectID: "1234", username: "test"},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if req, err := http.NewRequest("PUT", "/unlock", nil); err != nil {
				t.Fatal(err)
			} else {
				ctx := context.WithValue(req.Context(), "projectID", tc.projectID)
				ctx = context.WithValue(ctx, "username", tc.username)
				rr := httptest.NewRecorder()
				handler := http.HandlerFunc(LockProject)
				handler.ServeHTTP(rr, req.WithContext(ctx))
				if status := rr.Code; status != tc.wantedStatusCode {
					t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
				}
			}
		})
	}
}
