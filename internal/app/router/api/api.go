package router

import (
	"github.com/gin-gonic/gin"
)

func SetUpAPI(r *gin.Engine) {
	api := r.Group("api")
	SetUpAPIV1(api)
}
