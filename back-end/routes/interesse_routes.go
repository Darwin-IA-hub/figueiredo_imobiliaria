package routes

import (
	"back-end/controllers"

	"github.com/gin-gonic/gin"
)

func SetupInteresseRoutes(router *gin.Engine, interesseController controllers.InteresseController) {
	router.GET("/interesses", interesseController.GetAllInteresses)
	router.POST("/interesses", interesseController.CreateInteresse)
	router.GET("/interesses/:interesseId", interesseController.GetInteressesById)
	router.PUT("/interesses/:interesseId", interesseController.UpdateInteresse)
	router.DELETE("/interesses/:interesseId", interesseController.DeleteInteresse)
}
