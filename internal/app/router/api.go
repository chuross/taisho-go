package router

import (
	"github.com/gin-gonic/gin"
)

func SetUpAPI(r *gin.Engine) {
	api := r.Group("api")

	v1 := api.Group("v1")
	{
		v1.POST("line/callback", apiv1.PostLineCallback)
	}
}
