package usecases

import (
	//"back-end/models"
	"back-end/models"
	"back-end/repository"
	//"fmt"
	//"math"
)

type ContactUseCases struct{
	repository repository.ContactRepository
}

func NewContactUseCases(repo repository.ContactRepository) ContactUseCases{
	return ContactUseCases{
		repository: repo,
	}
}

func (usecase ContactUseCases) GetCountContatosAtivos() (int, error){
	countContatos, err := usecase.repository.GetCountContatosAtivos()
	if err != nil {
		return 0, err
	}
	return countContatos,nil
}

func (usecase ContactUseCases) GetContatoByTelefone(telefone string) (models.Contact, error){
	contato, err := usecase.repository.GetContatoByTelefone(telefone)
	if err != nil {
		return models.Contact{}, err
	}
	return contato,nil
}

func (usecase ContactUseCases) CreateContato(nome, telefone string) (error){
	err := usecase.repository.CreateContato(nome,telefone)
	if err != nil {
		return err
	}
	return nil
}

func (usecase ContactUseCases) SetConversationID(conversationID, telefone string) (error){
	err := usecase.repository.SetConversationID(conversationID,telefone)
	if err != nil {
		return err
	}
	return nil
}

