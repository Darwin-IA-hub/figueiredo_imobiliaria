package controllers

import (
	"back-end/models"
	"back-end/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ImovelController struct {
	useCases usecases.ImovelUseCases
}

func NewImovelController(usecases usecases.ImovelUseCases) ImovelController {
	return ImovelController{
		useCases: usecases,
	}
}

func (controller ImovelController) GetAllImoveis(c *gin.Context) {
	imoveis, err := controller.useCases.GetAllImoveis()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, imoveis)
}

func (controller ImovelController) GetImovelById(c *gin.Context) {
	idImovel := c.Param("idImovel")
	idImovelInt, err := strconv.Atoi(idImovel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id invalido", "error": err.Error()})
		return
	}
	imovel, err := controller.useCases.GetImovelById(idImovelInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, imovel)
}

func (controller ImovelController) CreateImovel(c *gin.Context) {
	var imovel models.Imovel
	err := c.BindJSON(&imovel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "inputs invalidos", "error": err.Error()})
		return
	}
	id, err := controller.useCases.CreateImovel(imovel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	imovel.IdImovel = id
	c.JSON(http.StatusOK, imovel)
}

func (controller ImovelController) UpdateImovel(c *gin.Context) {
	idImovel := c.Param("idImovel")
	idImovelInt, err := strconv.Atoi(idImovel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id invalido", "error": err.Error()})
		return
	}
	var imovel models.Imovel
	err = c.BindJSON(&imovel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "inputs invalidos", "error": err.Error()})
		return
	}
	imovel.IdImovel = idImovelInt
	updatedImovel, err := controller.useCases.UpdateImovel(imovel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedImovel)
}

func (controller ImovelController) DeleteImovel(c *gin.Context) {
	idImovel := c.Param("idImovel")
	idImovelInt, err := strconv.Atoi(idImovel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id invalido", "error": err.Error()})
		return
	}
	err = controller.useCases.DeleteImovel(idImovelInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "imovel deletado com sucesso")
}

func (controller ImovelController) GetAllImoveisVenda(c *gin.Context) {
	imoveisVenda, err := controller.useCases.GetAllImoveisVenda()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, imoveisVenda)
}

func (controller ImovelController) GetImovelVendaById(c *gin.Context) {
	idImovelVenda := c.Param("idImovelVenda")
	idImovelVendaInt, err := strconv.Atoi(idImovelVenda)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id invalido", "error": err.Error()})
		return
	}
	imovel, err := controller.useCases.GetImovelById(idImovelVendaInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, imovel)
}

func (controller ImovelController) CreateImovelVenda(c *gin.Context) {
	var imovelVenda models.ImovelVenda
	err := c.BindJSON(&imovelVenda)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "inputs invalidos", "error": err.Error()})
		return
	}
	id, err := controller.useCases.CreateImovelVenda(imovelVenda)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	imovelVenda.Fk_Imovel_idImovel = id
	c.JSON(http.StatusOK, id)
}

func (controller ImovelController) UpdateImovelVenda(c *gin.Context) {
	idImovelVenda := c.Param("idImovelVenda")
	idImovelVendaInt, err := strconv.Atoi(idImovelVenda)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id invalido", "error": err.Error()})
		return
	}
	var imovelVenda models.ImovelVenda
	err = c.BindJSON(&imovelVenda)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "inputs invalidos", "error": err.Error()})
		return
	}
	imovelVenda.Fk_Imovel_idImovel = idImovelVendaInt
	updatedImovelVenda, err := controller.useCases.UpdateImovelVenda(imovelVenda)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedImovelVenda)
}

func (controller ImovelController) DeleteImovelVenda(c *gin.Context) {
	idImovelVenda := c.Param("idImovelVenda")
	idImovelVendaInt, err := strconv.Atoi(idImovelVenda)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id invalido", "error": err.Error()})
		return
	}
	err = controller.useCases.DeleteImovelVenda(idImovelVendaInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "imovelVenda deletado com sucesso")
}
