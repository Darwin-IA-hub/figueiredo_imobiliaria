package usecases

import (
	"back-end/models"
	"back-end/repository"
)

type InteresseUseCases struct {
	repository repository.InteresseRepository
}

func NewInteresseUseCases(repo repository.InteresseRepository) InteresseUseCases {
	return InteresseUseCases{
		repository: repo,
	}
}
func (usecase InteresseUseCases) GetAllInteresses() ([]models.Interesse, error) {
	return usecase.repository.GetAllInteresses()
}
func (usecase InteresseUseCases) GetInteressesById(idInteresse int) (models.Interesse, error) {
	return usecase.repository.GetInteressesById(idInteresse)
}
func (usecase InteresseUseCases) CreateInteresse(interesse models.Interesse) (int, error) {
	return usecase.repository.CreateInteresse(interesse)
}
func (usecase InteresseUseCases) UpdateInteresse(interesse models.Interesse) (models.Interesse, error) {
	return usecase.repository.UpdateInteresse(interesse)
}
func (usecase InteresseUseCases) DeleteInteresse(interesseId int) error {
	return usecase.repository.DeleteInteresse(interesseId)
}
