package api_server

import (
	"github.com/disism/saikan/internal/conf"
	"github.com/gin-gonic/gin"
)

func GetNodeInfoHandler(c *gin.Context) {

	c.JSON(200, gin.H{
		"saikan": conf.GetServiceEndpoint(),
	})
}
