package models

type Message struct {
	Telefone string `db:"telefone" json:"telefone"`
	Conteudo string `db:"conteudo" json:"conteudo"`
}