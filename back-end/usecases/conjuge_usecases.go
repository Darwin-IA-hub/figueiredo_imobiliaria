package usecases

import (
	"back-end/models"
	"back-end/repository"
)

type ConjugeUseCases struct {
	repository repository.ConjugeRepository
}

func NewConjugeUseCases(repo repository.ConjugeRepository) ConjugeUseCases {
	return ConjugeUseCases{
		repository: repo,
	}
}

func (usecase ConjugeUseCases) GetAllConjuges() ([]models.Conjuge, error) {
	return usecase.repository.GetAllConjuges()
}

func (usecase ConjugeUseCases) GetConjugeById(idConjuge int) (models.Conjuge, error){
	return usecase.repository.GetConjugeById(idConjuge)
}

func (usecases ConjugeUseCases) CreateConjuge(conjuge models.Conjuge) (int, error) {
	return usecases.repository.CreateConjuge(conjuge)
}

func (usecases ConjugeUseCases) UpdateConjuge(conjuge models.Conjuge) (models.Conjuge, error){
	return usecases.repository.UpdateConjuge(conjuge)
}

func (usecases ConjugeUseCases) DeleteConjuge(idConjuge int)(error){
	return usecases.repository.DeleteConjuge(idConjuge)
}
