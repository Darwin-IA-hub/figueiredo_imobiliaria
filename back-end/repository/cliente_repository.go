package repository

import (
	"back-end/models"

	"github.com/jmoiron/sqlx"
)

//"fmt"

type ClienteRepository struct {
	connection *sqlx.DB
}

func NewClienteRepository(connection *sqlx.DB) ClienteRepository {
	return ClienteRepository{
		connection: connection,
	}
}

func (repo ClienteRepository) GetClienteBloqueadoById(telefoneCliente string) error {
	query := `SELECT idcliente FROM clientesbloqueados WHERE idcliente = $1`

	var bloqueado string
	err := repo.connection.Get(&bloqueado, query, telefoneCliente)
	if err != nil {
		return err
	}
	return nil
}

func (repo ClienteRepository) SetClienteBloqueado(idCliente string) error {
	query := `INSERT INTO clientesbloqueados (idcliente) VALUES ($1) ON CONFLICT DO NOTHING`

	_, err := repo.connection.Exec(query, idCliente)
	if err != nil {
		return err
	}
	return nil
}

func (repo ClienteRepository) DeleteClienteBloqueadoByID(idCliente string) error {
	query := `DELETE FROM clientesbloqueados WHERE idcliente = $1`

	_, err := repo.connection.Exec(query, idCliente)
	if err != nil {
		return err
	}
	return nil
}

func (repo ClienteRepository) GetAllClientes() ([]models.Cliente, error) {
	query := `SELECT * FROM cliente;`

	var clientes []models.Cliente
	err := repo.connection.Select(&clientes, query)
	if err != nil {
		return nil, err
	}
	return clientes, nil
}

func (repo ClienteRepository) CreateCliente(telefone, nomeCliente string) error {
	query := `INSERT INTO cliente(telefone, nomecliente) VALUES ($1,$2);`

	_, err := repo.connection.Exec(query, telefone, nomeCliente)

	if err != nil {
		return err
	}
	return nil
}
