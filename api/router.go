package api

import (
	"auth/api/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(h *handler.Handler) *gin.Engine {
	router := gin.Default()

	user := router.Group("/user")
	{
		user.POST("/createregister", h.CreateRegister)
		user.POST("/addimage", h.AddImage)
		user.GET("/getprofile", h.GetRegister)
		user.GET("/login", h.Login)
		user.PUT("/update", h.Update)
		user.PUT("/updatepassword", h.UpdatePassword)
	}

	return router
}
