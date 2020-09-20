package router

import (
	"github.com/chuross/taisho/internal/app/handler/apiv1"
	"github.com/chuross/taisho/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func SetUpAPI(r *gin.Engine) {
	api := r.Group("api")

	v1 := api.Group("v1")
	{
		line := v1.Group("line").Use(middleware.ValidateLineSignature)
		{
			line.POST("callback", apiv1.PostLineCallback)
		}
	}
}
