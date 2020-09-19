package router_api

import (
	"github.com/chuross/taisho/internal/app/handler/api_v1"
	"github.com/gin-gonic/gin"
)

func SetUpV1API(r *gin.RouterGroup) {
	v1 := r.Group("v1")
	{
		v1.POST("line/callback", api_v1.PostLineCallback)
	}
}
