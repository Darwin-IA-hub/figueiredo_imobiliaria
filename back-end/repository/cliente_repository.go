package repository

import (
	"back-end/models"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

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

func (repo ClienteRepository) ClienteExiste(telefoneCliente string) (bool, error) {
	query := `SELECT * FROM cliente WHERE telefone = $1`

	var cliente models.Cliente
	err := repo.connection.Get(&cliente, query, telefoneCliente)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
	}
	if len(cliente.Telefone) > 0 {
		return true, nil
	}
	return false, nil
}

func (repo ClienteRepository) CreateCliente(telefone, nomeCliente string) error {
	query := `INSERT INTO cliente(telefone, nomecliente) VALUES ($1,$2);`

	_, err := repo.connection.Exec(query, telefone, nomeCliente)

	if err != nil {
		return err
	}
	return nil
}

func (repo ClienteRepository) UpdateCliente(cliente models.Cliente) (models.Cliente, error) {
	query := `UPDATE cliente SET 
				nomecliente = :nomecliente,
				datanascimentocliente = :datanascimentocliente,
				rendabrutacliente = :rendabrutacliente,
				quantidadefilhos = :quantidadefilhos,
				anoscarteiraassinada = :anoscarteiraassinada,
				tevesubsidio = :tevesubsidio,
				vaiusarfgts = :vaiusarfgts,
				possuifinanciamento = :possuifinanciamento
			  WHERE telefone = :telefone
			  RETURNING *`

	var updatedCliente models.Cliente
	rows, err := repo.connection.NamedQuery(query, &cliente)
	if err != nil {
		return updatedCliente, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.StructScan(&updatedCliente)
		if err != nil {
			return updatedCliente, err
		}
	}
	return updatedCliente, nil
}

func (repo ClienteRepository) DeleteCliente(telefone string) error {
	query := `DELETE FROM cliente WHERE telefone = $1;`

	_, err := repo.connection.Exec(query, telefone)
	if err != nil {
		return err
	}
	return nil
}

func (repo ClienteRepository) GetClienteByTelefone(telefone string) (models.Cliente, error) {
	query := `SELECT * FROM cliente WHERE telefone = $1`

	var cliente models.Cliente
	err := repo.connection.Get(&cliente, query, telefone)
	if err != nil {
		return cliente, err
	}
	return cliente, nil
}
