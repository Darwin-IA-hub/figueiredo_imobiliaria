package models

type ContactSetMold struct {
	ConversationID string `db:"conversation_id" json:"conversation_id"`
	Telefone       string `db:"telefone" json:"telefone"`
}
