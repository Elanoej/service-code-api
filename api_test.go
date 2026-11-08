package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	handler := novaApi()

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("status esperado %d, veio %d", http.StatusOK, rec.Code)
	}

	contentType := rec.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Content-Type esperado application/json, veio %q", contentType)
	}

	var resposta healthResponse
	if err := json.NewDecoder(rec.Body).Decode(&resposta); err != nil {
		t.Fatalf("não consegui decodificar o JSON: %v", err)
	}

	if resposta.Status != "ok" {
		t.Errorf("status esperado %q, veio %q", "ok", resposta.Status)
	}
}

func TestHealthMetodoErrado(t *testing.T) {
	handler := novaApi()

	req := httptest.NewRequest(http.MethodPost, "/health", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("esperado 405m, veio %d", rec.Code)
	}
}
