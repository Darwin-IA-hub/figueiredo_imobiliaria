package controllers

import (
	"back-end/models"
	"back-end/usecases"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ClienteController struct {
	useCase usecases.ClienteUseCases
}

func NewClienteController(usecase usecases.ClienteUseCases) ClienteController {
	return ClienteController{
		useCase: usecase,
	}
}

func (controller ClienteController) DesligaRobo(c *gin.Context) {
	telefone := c.Param("telefone")
	err := controller.useCase.SetClienteBloqueado(telefone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"valido": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"valido": true, "message": fmt.Sprintf("Robo desligado para o número: %s.", telefone)})
}

func (controller ClienteController) LigaRobo(c *gin.Context) {
	telefone := c.Param("telefone")
	err := controller.useCase.DeleteClienteBloqueadoByID(telefone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"valido": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"valido": true, "message": fmt.Sprintf("Robo ligado para o número: %s.", telefone)})
}

func (controller ClienteController) IsRoboLigado(c *gin.Context) {
	telefone := c.Param("telefone")

	err := controller.useCase.GetClienteBloqueadoById(telefone)
	if err != nil {
		if err == sql.ErrNoRows {
			c.String(http.StatusOK, "true") // Robô está ligado
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao verificar status", "details": err.Error()})
		return
	}
	c.String(http.StatusOK, "false") // Robô está desligado
}

func (controller ClienteController) GetAllClientes(c *gin.Context) {
	clientes, err := controller.useCase.GetAllClientes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, clientes)
}

func (controller ClienteController) CreateCliente(c *gin.Context) {
	var cliente models.Cliente
	err := c.BindJSON(&cliente)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "nome ou telefone invalidos", "error": err.Error()})
		fmt.Println(err.Error())
		return
	}
	err = controller.useCase.CreateCliente(cliente.Telefone, cliente.NomeCliente)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "cliente criado com sucesso", "cliente": cliente})
}

func (controller ClienteController) UpdateCliente(c *gin.Context) {
	telefone := c.Param("telefone")
	var cliente models.Cliente
	err := c.BindJSON(&cliente)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "inputs invalidos", "error": err.Error()})
		fmt.Println(err.Error())
		return
	}

	cliente.Telefone = telefone

	updatedCliente, err := controller.useCase.UpdateCliente(cliente)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, updatedCliente)
}

func (controller ClienteController) DeleteCliente(c *gin.Context) {
	telefone := c.Param("telefone")
	err := controller.useCase.DeleteCliente(telefone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, "cliente deletado com sucesso")
}

func (controller ClienteController) GetClienteByTelefone(c *gin.Context) {
	telefone := c.Param("telefone")
	cliente, err := controller.useCase.GetClienteByTelefone(telefone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, cliente)
}

func (controller ClienteController) ClienteExiste(c *gin.Context) {
	telefone := c.Param("telefone")
	existe, err := controller.useCase.ClienteExiste(telefone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, existe)
}
