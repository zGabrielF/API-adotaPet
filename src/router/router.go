package router

import "github.com/gorilla/mux"

//Retorna um router com as rotas configuradas
func Gerar() *mux.Router {
	return mux.NewRouter()
}
