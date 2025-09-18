package usecases

import (
	"back-end/config"
	"back-end/models"
	"back-end/repository"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type MessageUseCases struct {
	repository repository.MessageRepository
}

func NewMessageUseCases(repo repository.MessageRepository) MessageUseCases {
	return MessageUseCases{
		repository: repo,
	}
}

func (usecase MessageUseCases) Responder(telefone, conversation_id, mensagem string) error {
	config.Counter++
	if config.Counter > 1 {
		config.Counter = 0
		return nil
	}
	if !config.Respondendo || mensagem == "" {
		return nil
	}
	baseURL := "http://localhost:5678/webhook/74fac712-b9a5-458e-8652-c8eedeee7837Figueiredo"
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	// Use a função auxiliar
	telefonePadronizado := config.PadronizaTelefone(telefone)

	data := url.Values{}
	data.Set("user", telefonePadronizado)
	data.Set("data", mensagem)
	data.Set("telefone", telefonePadronizado)
	data.Set("conversation_id", conversation_id)

	req, err := http.NewRequest("POST", baseURL, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("erro ao enviar requisição: %s", resp.Status)
	}
	return nil
}

func EnviarMensagem(telefone, mensagem string) error {
	baseURL := "http://localhost:5678/webhook/04b06825-2c4d-471b-86cc-9b32ad6f9d8eFigueiredo"
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	data := url.Values{}
	data.Set("telefone", telefone)
	data.Set("mensagem", mensagem)

	req, err := http.NewRequest("POST", baseURL, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("erro ao enviar requisição: %s", resp.Status)
	}
	return nil
}

func (usecase MessageUseCases) GetClienteBloqueadoById(telefoneCliente string) error {
	err := usecase.repository.GetClienteBloqueadoById(telefoneCliente)
	if err != nil {
		return err
	}
	return nil
}

func (usecase MessageUseCases) GetCountContatosAtivos() (int, error) {
	count, err := usecase.repository.GetCountContatosAtivos()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (usecase MessageUseCases) GetCountContatosAtivosByTelefone(telefone string) (int, error) {
	existingContact, err := usecase.repository.GetCountContatosAtivosByTelefone(telefone)
	if err != nil {
		return 0, err
	}
	return existingContact, nil
}

func (usecase MessageUseCases) CreateMessage(message models.Message) error {
	err := usecase.repository.CreateMessage(message)
	if err != nil {
		return err
	}
	return nil
}

func (usecase MessageUseCases) GetCountConversas(telefone, data string) (int, error) {
	contador, err := usecase.repository.GetCountConversas(telefone, data)
	if err != nil {
		return 0, err
	}
	return contador, err
}

func (usecase MessageUseCases) CreateConversa(telefone, data string) error {
	err := usecase.repository.CreateConversa(telefone, data)
	if err != nil {
		return err
	}
	return nil
}

func (usecase MessageUseCases) GetCountMensagensByTelefone(telefone string) (int, error) {
	counter, err := usecase.repository.GetCountMensagensByTelefone(telefone)
	if err != nil {
		return 0, err
	}
	return counter, nil
}

func (usecase MessageUseCases) SetContatoAtivo(telefone string) error {
	err := usecase.repository.SetContatoAtivo(telefone)
	if err != nil {
		return err
	}
	return nil
}

func (usecase MessageUseCases) GetMessagesByPhone(telefone string) (string, error) {
	mensagens, err := usecase.repository.GetMessagesByPhone(telefone)
	if err != nil {
		return "", err
	}
	var finalMessage string
	for i := 0; i < len(mensagens); i++ {
		finalMessage += mensagens[i].Conteudo + " "
	}

	return finalMessage, nil
}

func (usecase MessageUseCases) ClearMessagesByPhone(telefone string) error {
	err := usecase.repository.ClearMessagesByPhone(telefone)
	if err != nil {
		return err
	}
	return nil
}

func (usecase MessageUseCases) GetConversationIDByTelefone(telefone string) (string, error) {
	conversationId, err := usecase.repository.GetConversationIDByTelefone(telefone)
	if err != nil {
		return "", err
	}
	return conversationId, nil
}
