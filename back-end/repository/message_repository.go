package repository

import (
	"back-end/models"

	"github.com/jmoiron/sqlx"
)

type MessageRepository struct {
	connection *sqlx.DB
}

func NewMessageRepository(conn *sqlx.DB) MessageRepository {
	return MessageRepository{
		connection: conn,
	}
}

func (repo MessageRepository) GetClienteBloqueadoById(telefoneCliente string) error {
	query := `SELECT idcliente FROM clientesbloqueados WHERE idcliente = $1`

	var bloqueado string
	err := repo.connection.Get(&bloqueado, query, telefoneCliente)
	if err != nil {
		return err
	}
	return nil
}

func (repo MessageRepository) GetCountContatosAtivos() (int, error) {
	query := `SELECT COUNT(*) FROM contatos WHERE ativo = true`

	var count int
	err := repo.connection.Get(&count, query)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (repo MessageRepository) GetCountContatosAtivosByTelefone(telefone string) (int, error) {
	query := `SELECT COUNT(*) FROM contatos WHERE telefone = $1 AND ativo = true`

	var existingContact int
	err := repo.connection.Get(&existingContact, query, telefone)
	if err != nil {
		return 0, err
	}
	return existingContact, nil
}

func (repo MessageRepository) CreateMessage(message models.Message) error {
	query := `INSERT INTO mensagens (telefone, conteudo) VALUES ($1, $2)`

	_, err := repo.connection.Exec(query, message.Telefone, message.Conteudo)
	if err != nil {
		return err
	}
	return nil
}

func (repo MessageRepository) GetCountConversas(telefone, data string) (int, error) {
	query := `SELECT COUNT(*) FROM conversas WHERE telefone = $1 AND data = $2`

	var contador int
	err := repo.connection.Get(&contador, query, telefone, data)
	if err != nil {
		return 0, err
	}
	return contador, nil
}

func (repo MessageRepository) CreateConversa(telefone, data string) error {
	query := `INSERT INTO conversas (telefone, data) VALUES ($1, $2)`

	_, err := repo.connection.Exec(query, telefone, data)
	if err != nil {
		return err
	}
	return nil
}

func (repo MessageRepository) GetCountMensagensByTelefone(telefone string) (int, error) {
	query := `SELECT COUNT(*) FROM mensagens WHERE telefone = $1`

	var counter int
	err := repo.connection.Get(&counter, query, telefone)
	if err != nil {
		return 0, err
	}
	return counter, nil
}

func (repo MessageRepository) SetContatoAtivo(telefone string) error {
	query := `UPDATE contatos SET ativo = true WHERE telefone = $1`

	_, err := repo.connection.Exec(query, telefone)
	if err != nil {
		return err
	}
	return nil
}

func (repo MessageRepository) GetMessagesByPhone(telefone string) ([]models.Message, error) {
	query := `SELECT telefone, conteudo FROM mensagens WHERE telefone = $1`

	var mensagens []models.Message
	err := repo.connection.Select(&mensagens, query, telefone)
	if err != nil {
		return nil, err
	}
	return mensagens, nil
}

func (repo MessageRepository) ClearMessagesByPhone(telefone string) error {
	query := `DELETE FROM mensagens WHERE telefone = $1`

	_, err := repo.connection.Exec(query, telefone)
	if err != nil {
		return err
	}
	return nil
}

func (repo MessageRepository) GetConversationIDByTelefone(telefone string) (string, error) {
	query := `SELECT conversation_id FROM contatos WHERE telefone = $1`

	var conversationId string
	err := repo.connection.Get(&conversationId, query, telefone)
	if err != nil {
		return "", err
	}
	return conversationId, nil
}
