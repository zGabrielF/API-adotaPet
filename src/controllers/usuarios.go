package controllers

import (
	"net/http"
)

// Criar usuario

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("criando usuario"))
}

// Buscar todos usuarios

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("todos usuarios"))
}

//Buscar usuario um usuario

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("um usuario"))
}

// Atualizar um usuario
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("atualizando usuario"))
}

// Deletar um usuario
func DeletaUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("deletando usuario"))
}
