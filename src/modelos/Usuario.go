package modelos

type Usuario struct {
	ID       uint64 `json:"id,omitempty"`
	Nome     string `json:"nome,omitempty"`
	Nick     string `json:"nick,omitempty"`
	Senha    string `json:"senha,omitempty"`
	Telefone string `json:"telefone,omitempty"`
	Pets     []Pet  `json:"pets,omitempty"`
}
