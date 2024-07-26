package modelos

type Pet struct {
	ID        uint64  `json:"id,omitempty"`
	Nome      string  `json:"nome,omitempty"`
	Raca      string  `json:"raca,omitempty"`
	Cor       string  `json:"cor,omitempty"`
	Tamanho   string  `json:"tamanho,omitempty"`
	Sexo      string  `json:"sexo,omitempty"`
	Idade     uint8   `json:"idade,omitempty"`
	Peso      float32 `json:"peso,omitempty"`
	UsuarioID uint64  `json:"usuarioId,omitempty"`
}
