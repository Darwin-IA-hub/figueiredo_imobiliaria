package repository

import (
	"back-end/models"

	"github.com/jmoiron/sqlx"
	//"fmt"
)

type ContactRepository struct {
	connection *sqlx.DB
}

func NewContactRepository(connection *sqlx.DB) ContactRepository {
	return ContactRepository{
		connection: connection,
	}
}

func (repo ContactRepository) GetCountContatosAtivos() (int, error) {
	query := `SELECT COUNT(*) FROM contatos WHERE ativo = true`

	var count int
	err := repo.connection.Get(&count, query)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (repo ContactRepository) GetContatoByTelefone(telefone string) (models.Contact, error) {
	query := `SELECT nome, telefone, conversation_id, ativo FROM contatos WHERE telefone = $1`

	var contato models.Contact
	err := repo.connection.Get(&contato, query, telefone)
	if err != nil {
		return models.Contact{}, err
	}

	return contato, nil
}

func (repo ContactRepository) CreateContato(nome, telefone string) error {
	query := `INSERT INTO contatos (nome, telefone, conversation_id, ativo) 
	          VALUES ($1, $2, $3, true)`

	_, err := repo.connection.Exec(query, nome, telefone, nil)
	if err != nil {
		return err
	}
	return nil
}

func (repo ContactRepository) SetConversationID(conversationID, telefone string) error {
	query := `UPDATE contatos SET conversation_id = $1 WHERE telefone = $2`

	_, err := repo.connection.Exec(query, conversationID, telefone)
	if err != nil {
		return err
	}
	return nil
}
