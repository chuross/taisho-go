package main

import (
	"github.com/chuross/taisho/internal/app/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.SetUp(r)
	r.Run()
}
