package models

type Lead struct {
	Nome     string `db:"nome" json:"nome"`
	Telefone string `db:"telefone" json:"telefone"`
}
