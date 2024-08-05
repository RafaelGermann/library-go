package models

import (
	"api/src/seguranca"
	"errors"
	"github.com/badoux/checkmail"
	"strings"
	"time"
)

type Usuario struct {
	ID        uint64    `json:"id,omitempty"`
	Nome      string    `json:"nome,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Senha     string    `json:"senha,omitempty"`
	CriandoEm time.Time `json:"criandoEm,omitempty"`
}

func (usuario *Usuario) Preparar(etapa string) error {
	usuario.formatar()

	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	if etapa == "cadastro" {
		if erro := usuario.gerarHashSenha(); erro != nil {
			return erro
		}
	}

	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco")
	}
	if usuario.Nick == "" {
		return errors.New("O nick é obrigatório e não pode estar em branco")
	}
	if usuario.Email == "" {
		return errors.New("O email é obrigatório e não pode estar em branco")
	}

	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("O email inserido é inválido")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("A senha é obrigatório e não pode estar em branco")
	}

	return nil
}
func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}

func (usuario *Usuario) gerarHashSenha() error {
	senhaComHash, erro := seguranca.Hash(usuario.Senha)
	if erro != nil {
		return erro
	}
	usuario.Senha = string(senhaComHash)

	return nil
}
