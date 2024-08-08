package modelos

import (
	"apiAdotaPet/src/seguranca"
	"errors"
	"strings"
)

type Usuario struct {
	ID       uint64 `json:"id,omitempty"`
	Nome     string `json:"nome,omitempty"`
	Nick     string `json:"nick,omitempty"`
	Senha    string `json:"senha,omitempty"`
	Telefone string `json:"telefone,omitempty"`
	//Pets     []Pet  `json:"pets,omitempty"`
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("o nome é obrigatório, digite novamente")
	}

	if usuario.Nick == "" {
		return errors.New("o nick é obrigatório, digite novamente")
	}

	if usuario.Telefone == "" {
		return errors.New("o telefone é obrigatório, digite novamente")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("a senha é obrigatório, digite novamente")
	}
	return nil
}

func (usuario *Usuario) formatarStrings(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Telefone = strings.TrimSpace(usuario.Telefone)

	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(usuario.Senha)
		if erro != nil {
			return erro
		}
		usuario.Senha = string(senhaComHash)
	}
	return nil
}

// Chama o metodo validar e formatar
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	if erro := usuario.formatarStrings(etapa); erro != nil {
		return erro
	}
	return nil
}
