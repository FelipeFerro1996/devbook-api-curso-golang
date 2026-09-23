package modelos

import (
	"errors"
	"strings"
	"time"
)

type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}

func (usuario *Usuario) Preparar() error {

	erro := usuario.validar()
	if erro != nil {
		return erro
	}

	usuario.formatar()

	return nil
}

func (usuario *Usuario) validar() error {
	if usuario.Nome == "" {
		return errors.New("O campo Nome é obrigatório e não pode estar vazio")
	}
	if usuario.Nick == "" {
		return errors.New("O campo Nick é obrigatório e não pode estar vazio")
	}
	if usuario.Email == "" {
		return errors.New("O campo Email é obrigatório e não pode estar vazio")
	}
	if usuario.Senha == "" {
		return errors.New("O campo Senha é obrigatório e não pode estar vazio")
	}

	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}
