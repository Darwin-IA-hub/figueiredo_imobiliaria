package usecases

import (
	"back-end/models"
	"back-end/repository"
)

type ImovelUseCases struct {
	repository repository.ImovelRepository
}

func NewImovelUseCases(repo repository.ImovelRepository) ImovelUseCases {
	return ImovelUseCases{
		repository: repo,
	}
}

func (usecase ImovelUseCases) GetAllImoveis() ([]models.Imovel, error) {
	return usecase.repository.GetAllImoveis()
}

func (usecase ImovelUseCases) GetImovelById(idImovel int) (models.Imovel, error) {
	return usecase.repository.GetImovelById(idImovel)
}

func (usecase ImovelUseCases) CreateImovel(imovel models.Imovel) (int, error) {
	return usecase.repository.CreateImovel(imovel)
}

func (usecase ImovelUseCases) UpdateImovel(imovel models.Imovel) (models.Imovel, error) {
	return usecase.repository.UpdateImovel(imovel)
}

func (usecase ImovelUseCases) DeleteImovel(idImovel int) error {
	return usecase.repository.DeleteImovel(idImovel)
}

func (usecase ImovelUseCases) GetAllImoveisVenda() ([]models.ImovelVenda, error) {
	return usecase.repository.GetAllImoveisVenda()
}

func (usecase ImovelUseCases) GetImovelVendaById(idImovelVenda int) (models.ImovelVenda, error) {
	return usecase.repository.GetImovelVendaById(idImovelVenda)
}

func (usecase ImovelUseCases) CreateImovelVenda(imovelVenda models.ImovelVenda) (int, error) {
	return usecase.repository.CreateImovelVenda(imovelVenda)
}

func (usecase ImovelUseCases) UpdateImovelVenda(imovelVenda models.ImovelVenda) (models.ImovelVenda, error) {
	return usecase.repository.UpdateImovelVenda(imovelVenda)
}

func (usecase ImovelUseCases) DeleteImovelVenda(idImovelVenda int) error {
	return usecase.repository.DeleteImovelVenda(idImovelVenda)
}
