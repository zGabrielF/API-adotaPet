package rotas

import (
	"apiAdotaPet/src/controllers"
	"net/http"
)

// Todas as rotas de pets da API

var rotasPets = []Rota{
	{
		URI:                "/pets",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarPet,
		RequerAutenticacao: false,
	},

	{
		URI:                "/pets",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarPets,
		RequerAutenticacao: false,
	},

	{
		URI:                "/pets/{petId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarPet,
		RequerAutenticacao: false,
	},

	{
		URI:                "/pets/{petId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarPet,
		RequerAutenticacao: false,
	},
	{
		URI:                "/pets/{petId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletandoPet,
		RequerAutenticacao: false,
	},
}
