package routes

import (
	"back-end/controllers"

	"github.com/gin-gonic/gin"
)

func SetupFotoRoutes(router *gin.Engine, fotoController controllers.FotoController){
	router.GET("/fotos", fotoController.GetAllFotos)
	router.POST("/fotos", fotoController.CreateFoto)
	router.GET("/fotos/:idFoto", fotoController.GetFotoById)
	router.PUT("/fotos/:idFoto", fotoController.UpdateFoto)
	router.DELETE("/fotos/:idFoto", fotoController.DeleteFoto)
}