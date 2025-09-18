package controllers

import (
	"back-end/models"
	"back-end/usecases"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FotoController struct {
	useCases usecases.FotoUseCases
}

func NewFotoController(usecase usecases.FotoUseCases) FotoController {
	return FotoController{
		useCases: usecase,
	}
}

func (controller FotoController) GetAllFotos(c *gin.Context) {
	fotos, err := controller.useCases.GetAllFotos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, fotos)
}

func (controller FotoController) GetFotoById(c *gin.Context) {
	idFoto := c.Param("idFoto")
	idFotoInt, err := strconv.Atoi(idFoto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id invalido", "error": err.Error()})
		fmt.Println(err.Error())
		return
	}
	foto, err := controller.useCases.GetFotoById(idFotoInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, foto)
}

func (controller FotoController) CreateFoto(c *gin.Context) {
	var foto models.Foto
	err := c.BindJSON(&foto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "inputs invalidos", "error": err.Error()})
		fmt.Println(err.Error())
		return
	}
	id, err := controller.useCases.CreateFoto(foto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err.Error())
		return
	}
	foto.IdFoto = id
	c.JSON(http.StatusOK, id)
}

func (controller FotoController) UpdateFoto(c *gin.Context) {
	idFoto := c.Param("idFoto")
	idFotoInt, err := strconv.Atoi(idFoto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id deve ser um inteiro", "error": err.Error()})
		fmt.Println(err.Error())
		return
	}
	var foto models.Foto
	err = c.BindJSON(&foto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "inputs invalidos", "error": err.Error()})
		fmt.Println(err.Error())
		return
	}

	foto.IdFoto = idFotoInt
	updatedFoto, err := controller.useCases.UpdateFoto(foto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, updatedFoto)
}

func (controller FotoController) DeleteFoto(c *gin.Context) {
	idFoto := c.Param("idFoto")
	idFotoInt, err := strconv.Atoi(idFoto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id deve ser um inteiro", "error": err.Error()})
		fmt.Println(err.Error())
		return
	}
	err = controller.useCases.DeleteFoto(idFotoInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, "foto deletada com sucesso")
}

func (controller FotoController) PostFoto(c *gin.Context) {
	telefoneCliente := c.Param("telefoneCliente")
	var foto models.Foto
	err := c.BindJSON(&foto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "inputs invalidos", "error": err.Error()})
		fmt.Println(err.Error())
		return
	}
	foto.Fk_Cliente_telefone = telefoneCliente
	id, err := controller.useCases.PostFoto(foto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, id)
}

func (controller FotoController) EnviarFotosClienteParaVendedor(c *gin.Context) {
	telefoneCliente := c.Param("telefoneCliente")
	err := controller.useCases.EnviarFotosClienteParaVendedor(telefoneCliente)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, "fotos enviadas")
}
