package controllers

import (
	"back-end/models"
	"back-end/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type InteresseController struct {
	useCases usecases.InteresseUseCases
}

func NewInteresseController(usecases usecases.InteresseUseCases) InteresseController {
	return InteresseController{
		useCases: usecases,
	}
}

func (controller InteresseController) GetAllInteresses(c *gin.Context) {
	interesses, err := controller.useCases.GetAllInteresses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, interesses)
}
func (controller InteresseController) GetInteressesById(c *gin.Context) {
	interesseId := c.Param("interesseId")
	interesseIdInt, err := strconv.Atoi(interesseId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id deve ser um numero", "error": err.Error()})
		return
	}
	interesse, err := controller.useCases.GetInteressesById(interesseIdInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, interesse)
}
func (controller InteresseController) CreateInteresse(c *gin.Context) {
	var interesse models.Interesse
	err := c.BindJSON(&interesse)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "inputs invalidos", "error": err.Error()})
		return
	}
	id, err := controller.useCases.CreateInteresse(interesse)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, id)
}
func (controller InteresseController) UpdateInteresse(c *gin.Context) {
	interesseId := c.Param("interesseId")
	interesseIdInt, err := strconv.Atoi(interesseId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id deve ser um numero", "error": err.Error()})
		return
	}
	var interesse models.Interesse
	err = c.BindJSON(&interesse)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "inputs invalidos", "error": err.Error()})
		return
	}
	interesse.IdInteresse = interesseIdInt
	updatedInteresse, err := controller.useCases.UpdateInteresse(interesse)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedInteresse)
}
func (controller InteresseController) DeleteInteresse(c *gin.Context) {
	interesseId := c.Param("interesseId")
	interesseIdInt, err := strconv.Atoi(interesseId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id deve ser um numero", "error": err.Error()})
		return
	}
	err = controller.useCases.DeleteInteresse(interesseIdInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "interesse deletado com sucesso: "+interesseId)
}
