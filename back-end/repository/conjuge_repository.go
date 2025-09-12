package repository

import (
	"back-end/models"

	"github.com/jmoiron/sqlx"
)

type ConjugeRepository struct {
	connection *sqlx.DB
}

func NewConjugeRepository(conn *sqlx.DB) ConjugeRepository {
	return ConjugeRepository{
		connection: conn,
	}
}

func (repo ConjugeRepository) GetAllConjuges() ([]models.Conjuge, error) {
	query := `SELECT * FROM conjuge;`

	var conjuges []models.Conjuge
	err := repo.connection.Select(&conjuges, query)
	if err != nil {
		return nil, err
	}
	return conjuges, nil
}

func (repo ConjugeRepository) GetConjugeById(idConjuge int) (models.Conjuge, error) {
	query := `SELECT * FROM conjuge WHERE idconjuge = $1;`

	var conjuge models.Conjuge
	err := repo.connection.Get(&conjuge, query, idConjuge)
	if err != nil {
		return conjuge, err
	}
	return conjuge, nil
}

func (repo ConjugeRepository) CreateConjuge(conjuge models.Conjuge) (int, error) {
	querry := `INSERT INTO conjuge(rendabrutamensalconjuge,datanascimentoconjuge,fk_cliente_telefone) 
				VALUES(:rendabrutamensalconjuge,:datanascimentoconjuge,:fk_cliente_telefone)
				RETURNING idConjuge; `
	var id int
	rows,err:= repo.connection.NamedQuery(querry,&conjuge)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	if rows.Next(){
		err = rows.Scan(&id)
		if err != nil {
			return 0, nil
		}
	}
	return id, nil
}

func (repo ConjugeRepository) UpdateConjuge(conjuge models.Conjuge) (models.Conjuge, error){
	query:=`UPDATE conjuge SET
					rendabrutamensalconjuge = :rendabrutamensalconjuge,
					datanascimentoconjuge = :datanascimentoconjuge,
					fk_cliente_telefone = :fk_cliente_telefone
				WHERE idconjuge = :idconjuge
				RETURNING *;`
	var updatedConjuge models.Conjuge
	rows, err := repo.connection.NamedQuery(query, &conjuge)
	if err != nil {
		return updatedConjuge, err
	}
	defer rows.Close()

	if rows.Next(){
		err = rows.StructScan(&updatedConjuge)
		if err != nil {
			return updatedConjuge, err
		}
	}
	return updatedConjuge,nil
}

func (repo ConjugeRepository) DeleteConjuge(idConjuge int)(error) {
	query:= `DELETE FROM conjuge WHERE idconjuge = $1;`
	_, err := repo.connection.Exec(query, idConjuge)
	if err != nil {
		return err
	}
	return nil
}