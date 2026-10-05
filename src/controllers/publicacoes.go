package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func InserirPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := autenticacao.ExtrairusuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	body, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	var publicacao modelos.Publicacao
	if erro := json.Unmarshal(body, &publicacao); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro := publicacao.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	publicacao.AutorID = usuarioID

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositoriosDePublicacoes(db)
	publicacao.ID, erro = repositorio.InserirPublicacao(publicacao)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, publicacao)

}

func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {
	usuarioId, erro := autenticacao.ExtrairusuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositoriosDePublicacoes(db)
	publicacoes, erro := repositorio.BuscarPublicacoes(usuarioId)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, publicacoes)

}

func BuscarPublicacao(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)
	publicacaoId, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositoriosDePublicacoes(db)
	publicacao, erro := repositorio.BuscarPorId(publicacaoId)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, publicacao)
}

func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioId, erro := autenticacao.ExtrairusuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	parametros := mux.Vars(r)
	publicacaoId, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositoriosDePublicacoes(db)
	publicacaoSalva, erro := repositorio.BuscarPorId(publicacaoId)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if usuarioId != publicacaoSalva.AutorID {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é possivel editar uma publicação de outro usuario"))
		return
	}

	corporequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	var publicacao modelos.Publicacao
	if erro := json.Unmarshal(corporequisicao, &publicacao); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro := repositorio.AtualizarPublicacao(publicacaoId, publicacao); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)

}

func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioId, erro := autenticacao.ExtrairusuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	parametros := mux.Vars(r)
	publicacoId, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositoriosDePublicacoes(db)
	publicacaoSalva, erro := repositorio.BuscarPorId(publicacoId)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if usuarioId != publicacaoSalva.AutorID {
		respostas.Erro(w, http.StatusForbidden, errors.New("não é póssivel excluir uma publicação de outro usuario"))
		return
	}

	if erro := repositorio.DeletarPublicacao(publicacoId); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)

}
