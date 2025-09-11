package routes

import (
	"back-end/controllers"

	"github.com/gin-gonic/gin"
)

func SetMessageRoutes(router *gin.Engine, messageController controllers.MessageController) {
	router.POST("/addMessage", messageController.AddMessage)
	router.GET("/addMessage", messageController.AddMessage)
	router.GET("/getMessages/:telefone", messageController.GetMessagesByPhone)
	router.DELETE("/clearMessages/:telefone", messageController.ClearMessagesByPhone)
}
