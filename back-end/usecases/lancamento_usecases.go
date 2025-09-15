package usecases

import (
	"back-end/models"
	"back-end/repository"
)

type LancamentoUseCases struct {
	repository repository.LancamentoRepository
}

func NewLancamentoUseCases(repo repository.LancamentoRepository) LancamentoUseCases {
	return LancamentoUseCases{
		repository: repo,
	}
}

func (usecase LancamentoUseCases) GetAllLancamentos() ([]models.Lancamento, error) {
	return usecase.repository.GetAllLancamentos()
}
func (usecase LancamentoUseCases) GetLancamentoById(idLancamento int) (models.Lancamento, error) {
	return usecase.repository.GetLancamentoById(idLancamento)
}
func (usecase LancamentoUseCases) CreateLancamento(lancamento models.Lancamento) (int, error) {
	return usecase.repository.CreateLancamento(lancamento)
}
func (usecase LancamentoUseCases) UpdateLancamento(lancamento models.Lancamento) (models.Lancamento, error) {
	return usecase.repository.UpdateLancamento(lancamento)
}
func (usecase LancamentoUseCases) DeleteLancamento(idLancamento int) error {
	return usecase.repository.DeleteLancamento(idLancamento)
}
