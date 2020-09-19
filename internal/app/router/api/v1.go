package router

import (
	"github.com/gin-gonic/gin"
)

func SetUpAPIV1(r *gin.Engine) {
	v1 := r.Group("v1")
	{
		v1.POST("line/callback")
	}
}
