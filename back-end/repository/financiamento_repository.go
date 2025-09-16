package repository

import (
	"back-end/models"

	"github.com/jmoiron/sqlx"
)

type FinanciamentoRepository struct {
	connection *sqlx.DB
}

func NewFinanciamentoRepository(conn *sqlx.DB) FinanciamentoRepository {
	return FinanciamentoRepository{
		connection: conn,
	}
}

func (repo FinanciamentoRepository) GetAllFinanciamentos() ([]models.Financiamento, error) {
	query := `SELECT * FROM financiamento;`

	var financiamentos []models.Financiamento
	err := repo.connection.Select(&financiamentos, query)
	if err != nil {
		return nil, err
	}
	return financiamentos, nil
}

func (repo FinanciamentoRepository) GetFinanciamentoById(id int) (models.Financiamento, error) {
	query := `SELECT * FROM financiamento WHERE idfinanciamento = $1;`

	var financiamento models.Financiamento
	err := repo.connection.Get(&financiamento, query, id)
	if err != nil {
		return financiamento, err
	}
	return financiamento, nil
}

func (repo FinanciamentoRepository) CreateFinanciamento(telefoneCLiente, descricao string) (int, error) {
	query := `INSERT INTO financiamento(descricaofinanciamento,fk_cliente_telefone) VALUES($1, $2) 
						RETURNING idfinanciamento;`

	var id int
	err := repo.connection.QueryRow(query, descricao, telefoneCLiente).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (repo FinanciamentoRepository) UpdateFinanciamento(financiamento models.Financiamento) (models.Financiamento, error) {
	query := `UPDATE financiamento SET	
				descricaofinanciamento = :descricaofinanciamento,
				fk_cliente_telefone = :fk_cliente_telefone
			WHERE idfinanciamento = :idfinanciamento
			RETURNING *;`

	var updatedFinanciamento models.Financiamento
	rows, err := repo.connection.NamedQuery(query, &financiamento)
	if err != nil {
		return updatedFinanciamento, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.StructScan(&updatedFinanciamento)
		if err != nil {
			return updatedFinanciamento, err
		}
	}
	return updatedFinanciamento, nil
}

func (repo FinanciamentoRepository) DeleteFinanciamento(idFinanciamento int) error {
	query := `DELETE FROM financiamento WHERE idfianciamento = $1;`

	_, err := repo.connection.Exec(query, idFinanciamento)
	if err != nil {
		return err
	}
	return nil
}
