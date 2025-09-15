package repository

import (
	"back-end/models"

	"github.com/jmoiron/sqlx"
)

type InteresseRepository struct {
	connection *sqlx.DB
}

func NewInteresseRepository(conn *sqlx.DB) InteresseRepository {
	return InteresseRepository{
		connection: conn,
	}
}

func (repo InteresseRepository) GetAllInteresses() ([]models.Interesse, error) {
	query := `SELECT * FROM interesse;`
	var interesses []models.Interesse
	err := repo.connection.Select(&interesses, query)
	if err != nil {
		return nil, err
	}
	return interesses, nil
}
func (repo InteresseRepository) GetInteressesById(idInteresse int) (models.Interesse, error) {
	query := `SELECT * FROM interesse WHERE idInteresse = $1;`
	var interesse models.Interesse
	err := repo.connection.Get(&interesse, query, idInteresse)
	if err != nil {
		return interesse, err
	}
	return interesse, nil
}
func (repo InteresseRepository) CreateInteresse(interesse models.Interesse) (int, error) {
	query := `INSERT INTO interesse(interesseatual, cidadeinteresse, intervalopreco, observacao, tipoimovelinteresse, fk_cliente_telefone, fk_imovel_idimovel, fk_lancamento_idlancamento)
						VALUES(:interesseatual, :cidadeinteresse, :intervalopreco, :observacao, :tipoimovelinteresse, :fk_cliente_telefone, :fk_imovel_idimovel, :fk_lancamento_idlancamento)
						RETURNING idinteresse;`
	var id int
	rows, err := repo.connection.NamedQuery(query, &interesse)
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
func (repo InteresseRepository) UpdateInteresse(interesse models.Interesse) (models.Interesse, error) {
	query := `UPDATE interesse SET	
							interesseatual = :interesseatual,
							cidadeinteresse = :cidadeinteresse,
							intervalopreco = :intervalopreco,
							observacao = :observacao,
							tipoimovelinteresse = :tipoimovelinteresse,
							fk_cliente_telefone = :fk_cliente_telefone,
							fk_imovel_idimovel = :fk_imovel_idimovel,
							fk_lancamento_idlancamento = :fk_lancamento_idlancamento
						WHERE idinteresse = :idinteresse
						RETURNING *;`
	var updatedInteresse models.Interesse
	rows, err := repo.connection.NamedQuery(query, &interesse)
	if err != nil {
		return updatedInteresse, err
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.StructScan(&updatedInteresse)
		if err != nil {
			return updatedInteresse, err
		}
	}
	return updatedInteresse, nil
}
func (repo InteresseRepository) DeleteInteresse(interesseId int) error {
	query := `DELETE FROM interesse WHERE idinteresse = $1`
	_, err := repo.connection.Exec(query, interesseId)
	if err != nil {
		return err
	}
	return nil
}
