package routes

import (
	"back-end/controllers"

	"github.com/gin-gonic/gin"
)

func SetupClienteRoutes(router *gin.Engine, clienteController controllers.ClienteController) {
	router.POST("/desligar/:telefone", clienteController.DesligaRobo)
	router.DELETE("/ligar/:telefone", clienteController.LigaRobo)
	router.GET("/ligado/:telefone", clienteController.IsRoboLigado)

	router.GET("/clientes", clienteController.GetAllClientes)
	router.GET("/clientes/:telefone", clienteController.GetClienteByTelefone)
	router.GET("/cliente/:telefone", clienteController.ClienteExiste)
	router.POST("/clientes", clienteController.CreateCliente)
	router.PUT("/clientes", clienteController.UpdateCliente)
	router.DELETE("/clientes/:telefone", clienteController.DeleteCliente)
}
