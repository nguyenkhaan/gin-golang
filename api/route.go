package api

import (
	"cloudian/cloudian-restful/internal/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	api := r.Group("/api") 
	{
		user := api.Group("/user") 
		{
			user.GET("/" , controller.GetAllUsers) 
			user.GET("/:userID" , controller.GetDetailUser) 
			user.PUT("/:userID" , controller.UpdateUser) 

		}
	}
}