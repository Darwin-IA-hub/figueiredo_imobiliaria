package controllers

import (
	"back-end/models"
	"back-end/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FinanciamentoController struct {
	useCases usecases.FinanciamentoUseCases
}

func NewFinanciamentoController(usecases usecases.FinanciamentoUseCases) FinanciamentoController {
	return FinanciamentoController{
		useCases: usecases,
	}
}

func (controller FinanciamentoController) GetAllFinanciamentos(c *gin.Context) {
	financiamentos, err := controller.useCases.GetAllFinanciamentos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, financiamentos)
}

func (controller FinanciamentoController) GetFinanciamentoById(c *gin.Context) {
	id := c.Param("idFinanciamento")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID-FINANCIAMENTO deve ser um numero", "error": err.Error()})
		return
	}
	financiamento, err := controller.useCases.GetFinanciamentoById(idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, financiamento)
}

func (controller FinanciamentoController) CreateFinanciamento(c *gin.Context) {
	var financiamento models.Financiamento
	err := c.BindJSON(&financiamento)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "telefone ou descricao invalido", "error": err.Error()})
		return
	}
	id, err := controller.useCases.CreateFinanciamento(financiamento.Fk_Cliente_telefone, financiamento.DescricaoFinanciamento)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	financiamento.IdFinanciamento = id
	c.JSON(http.StatusOK, id)
}

func (controller FinanciamentoController) UpdateFinanciamento(c *gin.Context) {
	idFinanciamento := c.Param("idFinanciamento")
	var financiamento models.Financiamento
	err := c.BindJSON(&financiamento)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "inputs invalidos", "error": err.Error()})
		return
	}

	idFinanciamentoInt, err := strconv.Atoi(idFinanciamento)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id invalido", "error": err.Error()})
		return
	}
	financiamento.IdFinanciamento = idFinanciamentoInt

	updatedFinanciamento, err := controller.useCases.UpdateFinanciamento(financiamento)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedFinanciamento)
}

func (controller FinanciamentoController) DeleteFinanciamento(c *gin.Context) {
	idFinanciamento := c.Param("idFinanciamento")
	idFinanciamentoInt, err := strconv.Atoi(idFinanciamento)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id invalido", "error": err.Error()})
		return
	}
	err = controller.useCases.DeleteFinanciamento(idFinanciamentoInt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id invalido", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "financiamento deletado com sucesso")
}
