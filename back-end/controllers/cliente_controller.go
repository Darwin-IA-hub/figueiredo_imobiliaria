package controllers

import (
	"back-end/usecases"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ClienteController struct {
	useCase usecases.ClienteUseCases
}

func NewClienteController(usecase usecases.ClienteUseCases) ClienteController{
	return ClienteController{
		useCase: usecase,
	}
}

func (controller ClienteController) DesligaRobo(c *gin.Context){
	telefone := c.Param("telefone")
	err := controller.useCase.SetClienteBloqueado(telefone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"valido": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"valido": true, "message": fmt.Sprintf("Robo desligado para o número: %s.", telefone)})
}

func (controller ClienteController) LigaRobo(c *gin.Context){
	telefone := c.Param("telefone")
	err := controller.useCase.DeleteClienteBloqueadoByID(telefone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"valido": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"valido": true, "message": fmt.Sprintf("Robo ligado para o número: %s.", telefone)})	
}

func (controller ClienteController) IsRoboLigado(c *gin.Context){
	telefone := c.Param("telefone")

	err := controller.useCase.GetClienteBloqueadoById(telefone)
	if err != nil {
		if err == sql.ErrNoRows{
			c.String(http.StatusOK, "true") // Robô está ligado
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao verificar status", "details": err.Error()})
		return
	}
	c.String(http.StatusOK, "false") // Robô está desligado
}

