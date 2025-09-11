package routes

import(
	"back-end/controllers"

	"github.com/gin-gonic/gin"
)

func SetupContactRoutes(router *gin.Engine, contactController controllers.ContactController){
	router.POST("/init", contactController.InitializeContact)
	router.PUT("/setConversation", contactController.SetConversation)
}