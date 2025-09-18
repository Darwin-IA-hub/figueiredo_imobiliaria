package controllers

import (
	"back-end/models"
	"back-end/usecases"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ConjugeController struct {
	useCases usecases.ConjugeUseCases
}

func NewConjugeController(usecases usecases.ConjugeUseCases) ConjugeController {
	return ConjugeController{
		useCases: usecases,
	}
}

func (controller ConjugeController) GetAllConjuges(c *gin.Context) {
	conjuges, err := controller.useCases.GetAllConjuges()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, conjuges)
}

func (controller ConjugeController) GetConjugeById(c *gin.Context) {
	idConjuge := c.Param("idConjuge")
	idConjugeInt, err := strconv.Atoi(idConjuge)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id invalido", "error": err.Error()})
		fmt.Println(err.Error())
		return
	}
	conjuge, err := controller.useCases.GetConjugeById(idConjugeInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, conjuge)
}

func (controller ConjugeController) CreateConjuge(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	fmt.Println("JSON recebido bruto:", string(body))

	// como o body só pode ser lido uma vez, recria o reader:
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	var conjuge models.Conjuge
	err := c.BindJSON(&conjuge)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "inputs invalidos", "error": err.Error()})
		fmt.Println(err.Error())
		return
	}
	id, err := controller.useCases.CreateConjuge(conjuge)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err.Error())
		return
	}
	conjuge.IdConjuge = id
	c.JSON(http.StatusOK, id)
}

func (controller ConjugeController) UpdateConjuge(c *gin.Context) {
	idConjuge := c.Param("idConjuge")
	idConjugeInt, err := strconv.Atoi(idConjuge)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id invalido", "error": err.Error()})
		fmt.Println(err.Error())
		return
	}
	var conjuge models.Conjuge
	err = c.BindJSON(&conjuge)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "inputs invalidos", "error": err.Error()})
		fmt.Println(err.Error())
		return
	}
	conjuge.IdConjuge = idConjugeInt
	updatedConjuge, err := controller.useCases.UpdateConjuge(conjuge)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, updatedConjuge)
}

func (controller ConjugeController) DeleteConjuge(c *gin.Context) {
	idConjuge := c.Param("idConjuge")
	idConjugeInt, err := strconv.Atoi(idConjuge)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id invalido", "error": err.Error()})
		fmt.Println(err.Error())
		return
	}
	err = controller.useCases.DeleteConjuge(idConjugeInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, "conjuge deletado com sucesso")
}
