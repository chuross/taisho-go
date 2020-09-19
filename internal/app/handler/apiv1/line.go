package apiv1

import "github.com/gin-gonic/gin"

func PostLineCallback(c *gin.Context) {
	c.Status(200)
}
