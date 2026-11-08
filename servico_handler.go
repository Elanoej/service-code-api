package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/Elanoej/service-code-api/utils"
)

type ServicoHandler struct {
	servicos []Servico
}

func buildServicoHandler(servicos []Servico) *ServicoHandler {
	return &ServicoHandler{servicos: servicos}
}

func (h *ServicoHandler) Registrar(mux *http.ServeMux) {
	mux.HandleFunc("GET /codigos", h.buscar)
}

func (h *ServicoHandler) buscar(w http.ResponseWriter, r *http.Request) {
	search := strings.ToLower(r.URL.Query().Get("search"))

	log.Printf("Parâmetro: %q", search)

	filtradas := []Servico{}

	for _, servico := range h.servicos {

		codigoContemBusca := strings.Contains(
			strings.ToLower(servico.Codigo), search,
		)
		servicoContemBusca := strings.Contains(
			strings.ToLower(servico.Descricao), search,
		)

		if servicoContemBusca || codigoContemBusca {
			filtradas = append(filtradas, servico)
		}
	}

	utils.JsonResponse(w, http.StatusOK, filtradas)
}
