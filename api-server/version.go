package api_server

import "github.com/gin-gonic/gin"

//	{
//		"version": "saikan.1"
//	}
type Version struct {
	Versions string `json:"versions"`
}

func GetVersionHandler(c *gin.Context) {
	c.JSON(200, Version{
		Versions: "saikan.1",
	})
}
