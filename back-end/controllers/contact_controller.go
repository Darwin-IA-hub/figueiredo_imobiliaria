package controllers

import (
	"back-end/config"
	"back-end/models"
	"back-end/usecases"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ContactController struct {
	useCase usecases.ContactUseCases
}

func NewContactController(usecase usecases.ContactUseCases) ContactController {
	return ContactController{
		useCase: usecase,
	}
}

func (controller ContactController) InitializeContact(c *gin.Context) {
	var contact models.Contact
	//var verCont models.Contact

	err := c.BindJSON(&contact)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Formato inválido.", "error": err.Error()})
		return
	}

	if contact.Telefone == "status" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Telefone obrigatório."})
		return
	}

	count, err := controller.useCase.GetCountContatosAtivos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao contar contatos ativos", "error": err.Error()})
		return
	}

	verCont, err := controller.useCase.GetContatoByTelefone(contact.Telefone)
	if err != nil {
		if err == sql.ErrNoRows {
			if count >= config.GetPlanoAtual() {
				c.JSON(http.StatusForbidden, gin.H{"message": "Erro limite de plano excedido", "error": err.Error()})
				return
			}
			err = controller.useCase.CreateContato(contact.Nome, contact.Telefone)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.String(http.StatusCreated, "novo")
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conversationID := "NULL"
	if verCont.ConversationID.Valid {
		conversationID = verCont.ConversationID.String
	}
	c.String(http.StatusOK, "veterano "+conversationID)

}

func (controller ContactController) SetConversation(c *gin.Context) {
	var mold models.ContactSetMold
	if err := c.BindJSON(&mold); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Formato inválido.", "error": err.Error()})
		return
	}

	err := controller.useCase.SetConversationID(mold.ConversationID, mold.Telefone)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Sucesso!")
}
