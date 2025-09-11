package usecases

import (
	"back-end/repository"
)

type ClienteUseCases struct{
	repository repository.ClienteRepository
}

func NewClienteUseCases(repo repository.ClienteRepository) ClienteUseCases{
	return ClienteUseCases{
		repository: repo,
	}
}

func (usecase ClienteUseCases) GetClienteBloqueadoById(telefoneCliente string) (error){
	err := usecase.repository.GetClienteBloqueadoById(telefoneCliente)
	if err != nil {
		return err
	}
	return nil
}

func (usecase ClienteUseCases)SetClienteBloqueado(idCliente string) (error){
	err := usecase.repository.SetClienteBloqueado(idCliente)
	if err != nil {
		return err
	}
	return nil
}

func (usecase ClienteUseCases) DeleteClienteBloqueadoByID(idCliente string) (error){
	err := usecase.repository.DeleteClienteBloqueadoByID(idCliente)
	if err != nil {
		return err
	}
	return nil
}