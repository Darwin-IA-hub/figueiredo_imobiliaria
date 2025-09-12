package routes

import (
	"back-end/controllers"

	"github.com/gin-gonic/gin"
)	

func SetupFinanciamentoRoutes(router *gin.Engine, financiamentoController controllers.FinanciamentoController){
	router.GET("/financiamentos", financiamentoController.GetAllFinanciamentos)
	router.POST("/financiamentos", financiamentoController.CreateFinanciamento)
	router.GET("/financiamentos/:idFinanciamento", financiamentoController.GetFinanciamentoById)
	router.PUT("/financiamentos/:idFinanciamento", financiamentoController.UpdateFinanciamento)
	router.DELETE("/financiamentos/:idFinanciamento", financiamentoController.DeleteFinanciamento)
}