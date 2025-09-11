package routes

import (
	"back-end/controllers"

	"github.com/gin-gonic/gin"
)

func SetupClienteRoutes(router *gin.Engine, clienteController controllers.ClienteController){
	router.POST("/desligar/:telefone", clienteController.DesligaRobo)
	router.DELETE("/ligar/:telefone", clienteController.LigaRobo)
	router.GET("/ligado/:telefone", clienteController.IsRoboLigado)
}