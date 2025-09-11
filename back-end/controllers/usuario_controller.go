package controllers

import (
	"back-end/models"
	"back-end/usecases"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type UsuarioController struct {
	useCase usecases.UsuarioUseCases
}

func NewUsuarioController(usecase usecases.UsuarioUseCases) UsuarioController {
	return UsuarioController{
		useCase: usecase,
	}
}

func (controller UsuarioController) Login(c *gin.Context) {
	var usuario models.Usuario
	err := c.BindJSON(&usuario)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Formato inválido.", "error": err.Error()})
		return
	}

	//tratamento do celular
	if !strings.HasPrefix(usuario.Celular, "55") {
		usuario.Celular = "55" + usuario.Celular
	}

	id, _, ehRoot, err := controller.useCase.GetUsuarioLogin(usuario)
	if err != nil {
		if err == fmt.Errorf("CREDENCIAIS INVALIDAS") {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Credenciais inválidas."})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Login sucesso.", "expiration_time": time.Now().Add(24 * time.Hour).Unix(), "id": id, "root": ehRoot})

}
func (controller UsuarioController) Greetings(c *gin.Context) {
	dir, err := os.Getwd()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//pasta := filepath.Base(filepath.Base(filepath.Base(dir)))
	SO := runtime.GOOS
	var pasta []string
	if SO == "windows" {
		pasta = strings.Split(dir, `\`)
	} else {
		pasta = strings.Split(dir, `/`)
	}

	c.String(http.StatusOK, pasta[len(pasta)-3])
}
