package router

import (
	router_api "github.com/chuross/taisho/internal/app/router/api"
	"github.com/gin-gonic/gin"
)

func SetUpAPI(r *gin.Engine) {
	api := r.Group("api")
	router_api.SetUpV1API(api)
}
