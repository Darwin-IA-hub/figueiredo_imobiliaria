package models

import "database/sql"

type Contact struct {
	ConversationID sql.NullString `db:"conversation_id" json:"conversation_id"`
	Nome           string         `db:"nome" json:"nome"`
	Telefone       string         `db:"telefone" json:"telefone"`
	Ativo          bool           `db:"ativo" json:"ativo"`
}