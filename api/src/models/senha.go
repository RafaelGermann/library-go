package models

type Senha struct {
	SenhaNova  string `json:"senhaNova,omitempty"`
	SenhaAtual string `json:"senhaAtual,omitempty"`
}
