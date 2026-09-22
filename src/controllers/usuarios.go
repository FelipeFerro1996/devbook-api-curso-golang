package controllers

import (
	"net/http"
)

func CadastrarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Cadastrando usuario no banco de dados"))
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuarios no banco de dados"))
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuario no banco de dados"))
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Alterando informações do usuario no banco de dados"))
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuario no banco de dados"))
}
