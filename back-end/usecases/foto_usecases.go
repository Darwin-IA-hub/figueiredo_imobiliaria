package usecases

import (
	"back-end/models"
	"back-end/repository"
)

type FotoUseCases struct {
	repository repository.FotoRepository
}

func NewFotoUseCases(repo repository.FotoRepository) FotoUseCases {
	return FotoUseCases{
		repository: repo,
	}
}

func (usecase FotoUseCases) GetAllFotos() ([]models.Foto, error) {
	return usecase.repository.GetAllFotos()
}

func (usecase FotoUseCases) GetFotoById(idFoto int) (models.Foto, error) {
	return usecase.repository.GetFotoById(idFoto)
}

func (usecase FotoUseCases) CreateFoto(foto models.Foto) (int, error) {
	return usecase.repository.CreateFoto(foto)
}

func (usecase FotoUseCases) UpdateFoto(foto models.Foto) (models.Foto, error) {
	return usecase.repository.UpdateFoto(foto)
}

func (usecase FotoUseCases) DeleteFoto(idFoto int) error {
	return usecase.repository.DeleteFoto(idFoto)
}

func (usecase FotoUseCases) PostFoto(foto models.Foto) (int, error) {
	return usecase.repository.PostFoto(foto)
}
