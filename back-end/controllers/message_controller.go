package controllers

import (
	"back-end/config"
	"back-end/models"
	"back-end/usecases"
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type MessageController struct {
	useCase usecases.MessageUseCases
}

func NewMessageController(usecase usecases.MessageUseCases) MessageController{
	return MessageController{
		useCase: usecase,
	}
}

func (controller MessageController) AddMessage(c *gin.Context){
	var msg models.Message

	//capturar o body do JSON
	err := c.BindJSON(&msg)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message":"Formato Invalido", "error":err.Error()})
		return
	}

	//verificar se cliente esta bloqueado
	err = controller.useCase.GetClienteBloqueadoById(msg.Telefone)
	if err != sql.ErrNoRows {
		c.String(http.StatusOK, "bloqueado")
		return
	}

	//contagem de contatos ativos
	count, err := controller.useCase.GetCountContatosAtivos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message":"Erro ao encontrar contatos ativos","error":err.Error()})
		return
	}

	//verificar se o contato ja existe
	existingContact, err := controller.useCase.GetCountContatosAtivosByTelefone(msg.Telefone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message":"Erro ao verificar contato existente", "error":err.Error()})	
		return
	}

	//verifica se excedeu o limite do plano
	if count >= config.GetPlanoAtual() && existingContact == 0 {
		c.JSON(http.StatusForbidden, gin.H{"error":"Limite do plano excedido"})
		return
	}

	//inserir a mensagem na tabela de mensagens
	err = controller.useCase.CreateMessage(msg)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
	// 	return
	// }

	//inserir o registro da conversa, caso não exista
	hoje := time.Now()
	data := hoje.Format("02/01/2006")
	contador, err := controller.useCase.GetCountConversas(msg.Telefone, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	if contador < 1{
		err = controller.useCase.CreateConversa(msg.Telefone,data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}
	}

	// Se for a primeira mensagem, começa o contador de resposta
	counter, err := controller.useCase.GetCountMensagensByTelefone(msg.Telefone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	err = controller.useCase.SetContatoAtivo(msg.Telefone)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	if counter > 1 {
		config.Mensagem += msg.Conteudo + " "
		c.String(http.StatusOK, "segunda")
		return
	}
	config.Respondendo = true

	c.String(http.StatusOK, "primeira")
}

func (controller MessageController) GetMessagesByPhone(c *gin.Context){
	tel := c.Param("telefone")
	now := time.Now()
	dataAtual := now.Format("02/01/2006")
	hora := now.Format("15:04")
	diaSemana := []string{"Domingo", "Segunda-feira", "Terça-feira", "Quarta-feira", "Quinta-feira", "Sexta-feira", "Sábado"}[now.Weekday()]

	finalMessage, err := controller.useCase.GetMessagesByPhone(tel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"telefone_cliente": tel, "mensagem": finalMessage, "data_atual": dataAtual, "hora_atual": hora, "dia_semana": diaSemana})

}

func (controller MessageController) ClearMessagesByPhone(c *gin.Context){
	tel := c.Param("telefone")
	err := controller.useCase.ClearMessagesByPhone(tel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	conversationID, err := controller.useCase.GetConversationIDByTelefone(tel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.String(http.StatusOK, "Sucesso!")

	err = controller.useCase.Responder(tel, conversationID, config.Mensagem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	config.Respondendo = false
	config.Mensagem = ""
} 

