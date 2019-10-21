package controllers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLockProjectWithOutProjectID(t *testing.T) {
	req, err := http.NewRequest("PUT", "/lock", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(LockProject)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestLockProjectWithOutUsername(t *testing.T) {
	req, err := http.NewRequest("PUT", "/lock", nil)
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.WithValue(req.Context(), "projectID", "1234")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(LockProject)
	handler.ServeHTTP(rr, req.WithContext(ctx))
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestLockProject(t *testing.T) {
	req, err := http.NewRequest("PUT", "/lock", nil)
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.WithValue(req.Context(), "projectID", "1234")
	ctx2 := context.WithValue(ctx, "username", "test")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(LockProject)
	handler.ServeHTTP(rr, req.WithContext(ctx2))
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
