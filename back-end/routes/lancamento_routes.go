package routes

import (
	"back-end/controllers"

	"github.com/gin-gonic/gin"
)

func SetupLancamentoRoutes(router *gin.Engine, lancamentoController controllers.LancamentoController) {
	router.GET("/lancamentos", lancamentoController.GetAllLancamentos)
	router.POST("/lancamentos", lancamentoController.CreateLancamento)
	router.GET("/lancamentos/:lancamentoId", lancamentoController.GetLancamentoById)
	router.PUT("/lancamentos/:lancamentoId", lancamentoController.UpdateLancamento)
	router.DELETE("/lancamentos/:lancamentoId", lancamentoController.DeleteLancamento)
}
