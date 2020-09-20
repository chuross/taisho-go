package router

import (
	router_api "github.com/chuross/taisho/internal/app/router/api"
	"github.com/gin-gonic/gin"
)

func SetUpAPI(r *gin.Engine) {
	api := r.Group("api")
	
	v1 := api.Group("v1")
	{
		v1.POST("line/callback", apiv1.PostLineCallback)
	}
}
