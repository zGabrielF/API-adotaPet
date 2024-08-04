package rotas

import (
	"apiAdotaPet/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Representa todas as rotas da API
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Coloca todas as rotas dentro do router
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios

	rotas = append(rotas, rotasPets...)

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)

	}
	return r
}
