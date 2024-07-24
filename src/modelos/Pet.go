package modelos

type Pet struct {
	ID        uint64  `json:"id,omitempty"`
	Nome      string  `json:"nome,omitempty"`
	Raca      string  `json:"raca,omitempty"`
	UsuarioID uint64  `json:"usuario,omitempty"`
	Usuario   Usuario `json:"usuario,omitempty"`
}
