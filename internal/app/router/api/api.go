package router

import (
	"github.com/gin-gonic/gin"
)

func SetUpAPI(r *gin.Engine) {
	SetUpAPIV1(r)
}
