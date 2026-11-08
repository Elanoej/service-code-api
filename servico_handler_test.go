package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func servicosDeTeste() []Servico {
	return []Servico{
		{Codigo: "010101", Descricao: "Análise e desenvolvimento de sistemas."},
		{Codigo: "010201", Descricao: "Programação."},
		{Codigo: "010301", Descricao: "Processamento de dados."},
	}
}

func novoServicoMuxDeTeste() http.Handler {
	mux := http.NewServeMux()
	buildServicoHandler(servicosDeTeste()).Registrar(mux)
	return mux
}

func TestServicoHandlerBuscar(t *testing.T) {

	casos := []struct {
		nome     string
		url      string
		esperado int
	}{
		{"sem filtro devolve todos", "/codigos", 3},
		{"filtra por código", "/codigos?search=0103", 1},
		{"filtra por descrição", "/codigos?search=programa", 1},
		{"não diferencia maiúsculas", "/codigos?search=PROGRAMA", 1},
		{"sem resultado devolve lista vazia", "/codigos?search=zzz", 0},
	}

	for _, caso := range casos {
		t.Run(caso.nome, func(t *testing.T) {
			mux := novoServicoMuxDeTeste()

			req := httptest.NewRequest(http.MethodGet, caso.url, nil)
			rec := httptest.NewRecorder()

			mux.ServeHTTP(rec, req)

			if rec.Code != http.StatusOK {
				t.Fatalf("status esperado %d, veio %d", http.StatusOK, rec.Code)
			}
			if ct := rec.Header().Get("Content-Type"); ct != "application/json" {
				t.Fatalf("Content-Type esperado application/json, veio %q", ct)
			}

			var obtidos []Servico

			if err := json.NewDecoder(rec.Body).Decode(&obtidos); err != nil {
				t.Fatalf("não consegui decodificar o JSON: %v", err)
			}
			if obtidos == nil {
				t.Fatal("espera-se [] e não null")
			}
			if len(obtidos) != caso.esperado {
				t.Fatalf("esperado %d serviços, veio %d", caso.esperado, len(obtidos))
			}
		})
	}
}

func TestServicoHandlerMetodoErrado(t *testing.T) {
	mux := novoServicoMuxDeTeste()

	req := httptest.NewRequest(http.MethodPost, "/codigos", nil)
	rec := httptest.NewRecorder()

	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("esperado 405, veio %d", rec.Code)
	}
}
