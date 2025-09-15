package repository

import (
	"back-end/models"

	"github.com/jmoiron/sqlx"
)

type FotoRepository struct {
	connection *sqlx.DB
}

func NewFotoRepository(conn *sqlx.DB) FotoRepository {
	return FotoRepository{
		connection: conn,
	}
}

func (repo FotoRepository) GetAllFotos() ([]models.Foto, error) {
	query := `SELECT * FROM foto;	`

	var fotos []models.Foto
	err := repo.connection.Select(&fotos, query)
	if err != nil {
		return nil, err
	}
	return fotos, nil
}

func (repo FotoRepository) GetFotoById(idFoto int) (models.Foto, error) {
	query := `SELECT * FROM foto WHERE idfoto = $1`

	var foto models.Foto
	err := repo.connection.Get(&foto, query, idFoto)
	if err != nil {
		return foto, err
	}
	return foto, nil
}

func (repo FotoRepository) CreateFoto(foto models.Foto) (int, error) {
	query := `INSERT INTO foto(linkfoto, fk_cliente_telefone) VALUES(:linkfoto,:fk_cliente_telefone) 
				RETURNING idfoto;`
	var id int
	rows, err := repo.connection.NamedQuery(query, foto)
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

func (repo FotoRepository) UpdateFoto(foto models.Foto) (models.Foto, error) {
	query := `UPDATE foto SET	
				fk_cliente_telefone = :fk_cliente_telefone,
				linkfoto = :linkfoto
			WHERE idfoto = :idfoto
			RETURNING *;`

	var updatedFoto models.Foto
	rows, err := repo.connection.NamedQuery(query, &foto)
	if err != nil {
		return updatedFoto, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.StructScan(&updatedFoto)
		if err != nil {
			return updatedFoto, err
		}
	}
	return updatedFoto, nil
}

func (repo FotoRepository) DeleteFoto(idFoto int) error {
	query := `DELETE FROM foto WHERE idfoto = $1;`
	_, err := repo.connection.Exec(query, idFoto)
	if err != nil {
		return err
	}
	return nil
}
