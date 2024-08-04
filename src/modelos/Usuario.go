package modelos

import (
	"errors"
	"strings"
)

type Usuario struct {
	ID       uint64 `json:"id,omitempty"`
	Nome     string `json:"nome,omitempty"`
	Nick     string `json:"nick,omitempty"`
	Senha    string `json:"senha,omitempty"`
	Telefone string `json:"telefone,omitempty"`
	Pets     []Pet  `json:"pets,omitempty"`
}

func (usuario *Usuario) validar() error {
	if usuario.Nome == "" {
		return errors.New("o nome é obrigatório, digite novamente")
	}

	if usuario.Nick == "" {
		return errors.New("o nick é obrigatório, digite novamente")
	}
	if usuario.Senha == "" {
		return errors.New("a senha é obrigatório, digite novamente")
	}
	if usuario.Telefone == "" {
		return errors.New("o telefone é obrigatório, digite novamente")
	}
	return nil
}

func (usuario *Usuario) formatarStrings() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Telefone = strings.TrimSpace(usuario.Telefone)

}

func (usuario *Usuario) Preparar() error {
	if erro := usuario.validar(); erro != nil {
		return erro
	}
	usuario.formatarStrings()
	return nil
}
