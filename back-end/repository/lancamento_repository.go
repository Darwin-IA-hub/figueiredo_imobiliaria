package repository

import (
	"back-end/models"

	"github.com/jmoiron/sqlx"
)

type LancamentoRepository struct {
	connection *sqlx.DB
}

func NewLancamentoRepository(conn *sqlx.DB) LancamentoRepository {
	return LancamentoRepository{
		connection: conn,
	}
}

func (repo LancamentoRepository) GetAllLancamentos() ([]models.Lancamento, error) {
	query := `SELECT * FROM lancamento;`

	var lancamentos []models.Lancamento
	err := repo.connection.Select(&lancamentos, query)
	if err != nil {
		return nil, err
	}
	return lancamentos, nil
}
func (repo LancamentoRepository) GetLancamentoById(idLancamento int) (models.Lancamento, error) {
	query := `SELECT * FROM lancamento WHERE idlancamento = $1;`
	var lancamento models.Lancamento
	err := repo.connection.Get(&lancamento, query, idLancamento)
	if err != nil {
		return lancamento, err
	}
	return lancamento, nil
}
func (repo LancamentoRepository) CreateLancamento(lancamento models.Lancamento) (int, error) {
	query := `INSERT INTO lancamento(cidadelancamento, nomelancamento, detalhes) VALUES(:cidadelancamento, :nomelancamento, :detalhes) 
						RETURNING idlancamento;`
	var id int
	rows, err := repo.connection.NamedQuery(query, lancamento)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return 0, err
		}
	}
	return id, nil
}
func (repo LancamentoRepository) UpdateLancamento(lancamento models.Lancamento) (models.Lancamento, error) {
	query := `UPDATE lancamento SET	
							cidadelancamento = :cidadelancamento,
							nomelancamento = :nomelancamento,
							detalhes = :detalhes
					 	WHERE idlancamento = :idlancamento
						RETURNING *;`
	var updatedLancamento models.Lancamento
	rows, err := repo.connection.NamedQuery(query, lancamento)
	if err != nil {
		return updatedLancamento, err
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.StructScan(&updatedLancamento)
		if err != nil {
			return updatedLancamento, err
		}
	}
	return updatedLancamento, nil
}
func (repo LancamentoRepository) DeleteLancamento(idLancamento int) error {
	query := `DELETE FROM lancamento WHERE idlancamento = $1;`
	_, err := repo.connection.Exec(query, idLancamento)
	if err != nil {
		return err
	}
	return nil
}
