package models

type StatusRobo struct {
	Nome     string `db:"nome" json:"nome"`
	Telefone string `db:"telefone" json:"telefone"`
	Ligado   bool   `db:"ligado" json:"ligado"`
}