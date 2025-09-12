package usecases

import (
	"back-end/models"
	"back-end/repository"
)

type FinanciamentoUseCases struct {
	repository repository.FinanciamentoRepository
}

func NewFinanciamentoUseCases(repo repository.FinanciamentoRepository) FinanciamentoUseCases {
	return FinanciamentoUseCases{
		repository: repo,
	}
}

func (usecase FinanciamentoUseCases) GetAllFinanciamentos() ([]models.Financiamento, error) {
	return usecase.repository.GetAllFinanciamentos()
}

func (usecase FinanciamentoUseCases) GetFinanciamentoById(id int) (models.Financiamento, error){
	return usecase.repository.GetFinanciamentoById(id)
}

func (usecase FinanciamentoUseCases) CreateFinanciamento(telefoneCLiente, descricao string) (int, error){
	return usecase.repository.CreateFinanciamento(telefoneCLiente, descricao)
}

func (usecase FinanciamentoUseCases) UpdateFinanciamento(financiamento models.Financiamento) (models.Financiamento, error){
	return usecase.repository.UpdateFinanciamento(financiamento)
}

func (usecase FinanciamentoUseCases) DeleteFinanciamento(idFinanciamento int) (error){
	return usecase.repository.DeleteFinanciamento(idFinanciamento)
}