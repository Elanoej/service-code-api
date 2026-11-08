package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/Elanoej/service-code-api/utils"
)

type healthResponse struct {
	Status string `json:"status"`
}

func novaApi() http.Handler {
	mux := http.NewServeMux()

	servicos := MockServicos()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		utils.JsonResponse(w, http.StatusOK, healthResponse{Status: "ok"})
	})

	mux.HandleFunc("GET /codigos", func(w http.ResponseWriter, r *http.Request) {
		search := strings.ToLower(r.URL.Query().Get("search"))

		log.Printf("Parâmetro: %q", search)

		filtradas := []Servico{}

		for _, servico := range servicos {

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
	})

	return mux
}
