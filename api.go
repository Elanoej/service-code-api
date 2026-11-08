package main

import (
	"net/http"
)

func novaApi() http.Handler {
	mux := http.NewServeMux()

	servicos := MockServicos()

	buildHealthHandler().Registrar(mux)
	buildServicoHandler(servicos).Registrar(mux)

	return mux
}
