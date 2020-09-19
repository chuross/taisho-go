package router

import (
	apiv1 "github.com/chuross/taisho/internal/app/handler/api/v1"
	"github.com/gin-gonic/gin"
)

func SetUpAPIV1(r *gin.Engine) {
	v1 := r.Group("v1")
	{
		v1.POST("line/callback", apiv1.PostLineCallback)
	}
}
