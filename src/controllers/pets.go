package controllers

import (
	"net/http"
)

// Criar pet
func CriarPet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando pet"))
}

// Buscar todos pets
func BuscarPets(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os pets"))
}

// Buscar um pet
func BuscarPet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um pet"))
}

// Atualizar um pet
func AtualizarPet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando o pet"))
}

// Deletar um pet
func DeletandoPet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando o pet"))
}
