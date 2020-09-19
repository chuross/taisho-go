package router

import (
	router "github.com/chuross/taisho/internal/app/router/api"
	"github.com/gin-gonic/gin"
)

func SetUp(r *gin.Engine) {
	router.SetUpAPI(r)
}
