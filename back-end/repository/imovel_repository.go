package repository

import (
	"back-end/models"

	"github.com/jmoiron/sqlx"
)

type ImovelRepository struct {
	connection *sqlx.DB
}

func NewImovelRepository(conn *sqlx.DB) ImovelRepository {
	return ImovelRepository{
		connection: conn,
	}
}

func (repo ImovelRepository) GetAllImoveis() ([]models.Imovel, error) {
	query := `SELECT * FROM imovel`

	var imoveis []models.Imovel
	err := repo.connection.Select(&imoveis, query)
	if err != nil {
		return nil, err
	}
	return imoveis, nil
}

func (repo ImovelRepository) GetImovelById(idImovel int) (models.Imovel, error) {
	query := `SELECT * FROM imovel WHERE idimovel = $1`

	var imovel models.Imovel
	err := repo.connection.Get(&imovel, query, idImovel)
	if err != nil {
		return imovel, err
	}
	return imovel, nil
}

func (repo ImovelRepository) CreateImovel(imovel models.Imovel) (int, error) {
	query := `INSERT INTO imovel(tipoimovel, linkiptu, cidadeimovel) 
			VALUES(:tipoimovel, :linkiptu, :cidadeimovel)
			RETURNING idimovel;`

	rows, err := repo.connection.NamedQuery(query, &imovel)
	if err != nil {
		return 0, err
	}
	var id int
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return 0, err
		}
	}
	return id, nil
}

func (repo ImovelRepository) UpdateImovel(imovel models.Imovel) (models.Imovel, error) {
	query := `UPDATE imovel SET
				tipoimovel = :tipoimovel,
				cidadeimovel= :cidadeimovel, 
				linkiptu= :linkiptu
			WHERE idimovel = :idimovel
			RETURNING *;`
	var updatedImovel models.Imovel
	rows, err := repo.connection.NamedQuery(query, &imovel)
	if err != nil {
		return updatedImovel, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.StructScan(&updatedImovel)
		if err != nil {
			return updatedImovel, err
		}
	}
	return updatedImovel, nil
}

func (repo ImovelRepository) DeleteImovel(idImovel int) error {
	query := `DELETE FROM imovel WHERE idimovel = $1;`
	_, err := repo.connection.Exec(query, idImovel)
	if err != nil {
		return err
	}
	return nil
}

func (repo ImovelRepository) GetAllImoveisVenda() ([]models.ImovelVenda, error) {
	query := `SELECT * FROM imovelVenda;`
	var imovelVenda []models.ImovelVenda
	err := repo.connection.Select(&imovelVenda, query)
	if err != nil {
		return nil, err
	}
	return imovelVenda, nil
}

func (repo ImovelRepository) GetImovelVendaById(idImovelVenda int) (models.ImovelVenda, error) {
	query := `SELECT * FROM imovelVenda WHERE fk_imovel_idimovel = $1`
	var imovelVenda models.ImovelVenda
	err := repo.connection.QueryRow(query, idImovelVenda).Scan(&imovelVenda)
	if err != nil {
		return imovelVenda, err
	}
	return imovelVenda, nil
}

func (repo ImovelRepository) CreateImovelVenda(imovelVenda models.ImovelVenda) (int, error) {
	query := `INSERT INTO imovelvenda(fk_imovel_idimovel, financiadoquitado, docemdia, estahabitado) 
				VALUES (:fk_imovel_idimovel, :financiadoquitado, :docemdia, :estahabitado)
				RETURNING fk_imovel_idimovel`

	rows, err := repo.connection.NamedQuery(query, &imovelVenda)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var id int
	if rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return 0, err
		}
	}
	return id, nil
}

func (repo ImovelRepository) UpdateImovelVenda(imovelVenda models.ImovelVenda) (models.ImovelVenda, error) {
	query := `UPDATE imovelVenda SET
					financiadoquitado = :financiadoquitado,
					docemdia = :docemdia,
					estahabitado = :estahabitado
				WHERE fk_imovel_idimovel = :fk_imovel_idimovel
				RETURNING *;`
	var updatedImovelVenda models.ImovelVenda
	rows, err := repo.connection.NamedQuery(query, &imovelVenda)
	if err != nil {
		return updatedImovelVenda, err
	}
	defer rows.Close()
	if rows.Next() {
		err := rows.StructScan(&updatedImovelVenda)
		if err != nil {
			return updatedImovelVenda, err
		}
	}
	return updatedImovelVenda, nil
}

func (repo ImovelRepository) DeleteImovelVenda(idImovelVenda int) error {
	query := `DELETE FROM imovelvenda WHERE fk_imovel_idimovel = $1`

	_, err := repo.connection.Exec(query, idImovelVenda)
	if err != nil {
		return err
	}
	return nil
}
