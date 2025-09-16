package usecases

import (
	"back-end/models"
	"back-end/repository"
)

type ClienteUseCases struct {
	repository repository.ClienteRepository
}

func NewClienteUseCases(repo repository.ClienteRepository) ClienteUseCases {
	return ClienteUseCases{
		repository: repo,
	}
}

func (usecase ClienteUseCases) GetClienteBloqueadoById(telefoneCliente string) error {
	err := usecase.repository.GetClienteBloqueadoById(telefoneCliente)
	if err != nil {
		return err
	}
	return nil
}

func (usecase ClienteUseCases) SetClienteBloqueado(idCliente string) error {
	err := usecase.repository.SetClienteBloqueado(idCliente)
	if err != nil {
		return err
	}
	return nil
}

func (usecase ClienteUseCases) DeleteClienteBloqueadoByID(idCliente string) error {
	err := usecase.repository.DeleteClienteBloqueadoByID(idCliente)
	if err != nil {
		return err
	}
	return nil
}

func (usecase ClienteUseCases) GetAllClientes() ([]models.Cliente, error) {
	return usecase.repository.GetAllClientes()
}

func (usecase ClienteUseCases) CreateCliente(telefone, nomeCliente string) error {
	return usecase.repository.CreateCliente(telefone, nomeCliente)
}

func (usecase ClienteUseCases) UpdateCliente(cliente models.Cliente) (models.Cliente, error) {
	return usecase.repository.UpdateCliente(cliente)
}

func (usecase ClienteUseCases) DeleteCliente(telefone string) error {
	return usecase.repository.DeleteCliente(telefone)
}

func (usecase ClienteUseCases) GetClienteByTelefone(telefone string) (models.Cliente, error) {
	return usecase.repository.GetClienteByTelefone(telefone)
}

func (usecase ClienteUseCases) ClienteExiste(telefoneCliente string) (bool, error) {
	return usecase.repository.ClienteExiste(telefoneCliente)
}
