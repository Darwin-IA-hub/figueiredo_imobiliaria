package routes

import (
	"back-end/controllers"

	"github.com/gin-gonic/gin"
)

func SetupConjugeRoutes(router *gin.Engine, conjugeController controllers.ConjugeController){
	router.GET("/conjuges", conjugeController.GetAllConjuges)
	router.POST("/conjuges", conjugeController.CreateConjuge)
	router.GET("/conjuges/:idConjuge", conjugeController.GetConjugeById)
	router.PUT("/conjuges/:idConjuge", conjugeController.UpdateConjuge)
	router.DELETE("/conjuges/:idConjuge", conjugeController.DeleteConjuge)
}