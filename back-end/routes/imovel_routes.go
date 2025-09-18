package routes

import (
	"back-end/controllers"

	"github.com/gin-gonic/gin"
)

func SetupImovelRoutes(router *gin.Engine, imovelController controllers.ImovelController) {
	router.GET("/imoveis", imovelController.GetAllImoveis)
	router.POST("/imoveis", imovelController.CreateImovel)
	router.GET("/imoveis/:idImovel", imovelController.GetImovelById)
	router.PUT("/imoveis/:idImovel", imovelController.UpdateImovel)
	router.DELETE("/imoveis/:idImovel", imovelController.DeleteImovel)

	router.GET("/imoveis-venda", imovelController.GetAllImoveisVenda)
	router.POST("/imoveis-venda", imovelController.CreateImovelVenda)
	router.GET("/imoveis-venda/:idImovelVenda", imovelController.GetImovelVendaById)
	router.PUT("/imoveis-venda/:idImovelVenda", imovelController.UpdateImovelVenda)
	router.DELETE("/imoveis-venda/:idImovelVenda", imovelController.DeleteImovelVenda)
}
