package controllers

import (
	"back-end/models"
	"back-end/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LancamentoController struct {
	useCases usecases.LancamentoUseCases
}

func NewLancamentoController(usecases usecases.LancamentoUseCases) LancamentoController {
	return LancamentoController{
		useCases: usecases,
	}
}

func (controller LancamentoController) GetAllLancamentos(c *gin.Context) {
	lancamentos, err := controller.useCases.GetAllLancamentos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, lancamentos)
}
func (controller LancamentoController) GetLancamentoById(c *gin.Context) {
	lancamentoId := c.Param("lancamentoId")
	lancamentoIdInt, err := strconv.Atoi(lancamentoId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id invalido, deve ser um número inteiro", "error": err.Error()})
		return
	}
	lancamento, err := controller.useCases.GetLancamentoById(lancamentoIdInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, lancamento)
}
func (controller LancamentoController) CreateLancamento(c *gin.Context) {
	var lancamento models.Lancamento
	err := c.BindJSON(&lancamento)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "inputs invalidos", "error": err.Error()})
		return
	}
	id, err := controller.useCases.CreateLancamento(lancamento)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, id)
}
func (controller LancamentoController) UpdateLancamento(c *gin.Context) {
	lancamentoId := c.Param("lancamentoId")
	lancamentoIdInt, err := strconv.Atoi(lancamentoId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id invalido, deve ser um número inteiro", "error": err.Error()})
		return
	}
	var lancamento models.Lancamento
	err = c.BindJSON(&lancamento)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "inputs invalidos", "error": err.Error()})
		return
	}
	lancamento.IdLancamento = lancamentoIdInt
	updatedLancamento, err := controller.useCases.UpdateLancamento(lancamento)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedLancamento)
}
func (controller LancamentoController) DeleteLancamento(c *gin.Context) {
	lancamentoId := c.Param("lancamentoId")
	lancamentoIdInt, err := strconv.Atoi(lancamentoId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id invalido, deve ser um número inteiro", "error": err.Error()})
		return
	}
	err = controller.useCases.DeleteLancamento(lancamentoIdInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "lancamento excluido com sucesso: "+lancamentoId)
}
